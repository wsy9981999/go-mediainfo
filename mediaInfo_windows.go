package go_mediainfo

import (
	"errors"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type MediaInfo struct {
	DllInfo  *syscall.LazyDLL
	Instance uintptr
}

func NewMediaInfo(path ...string) (*MediaInfo, error) {
	DllPath := "MediaInfo.dll"
	if len(path) > 0 {
		DllPath = path[0]
	}
	dll := syscall.NewLazyDLL(DllPath)
	ptr, _, err := dll.NewProc(`MediaInfo_New`).Call()
	if !isError(err) {
		return nil, err
	}
	return &MediaInfo{
		DllInfo:  dll,
		Instance: ptr,
	}, nil
}
func (m *MediaInfo) Option(name string, value string) error {
	np, err := strToPtr(name)
	if err != nil {
		return err
	}
	vp, err := strToPtr(value)
	if err != nil {
		return err
	}
	_, _, err = m.DllInfo.NewProc(`MediaInfo_Option`).Call(m.Instance, np, vp)
	if !isError(err) {
		return err
	}
	return nil
}
func (m *MediaInfo) Destroy() error {
	_, _, err := m.DllInfo.NewProc(`MediaInfo_Delete`).Call(m.Instance)
	if !isError(err) {
		return err
	}
	return nil
}
func (m *MediaInfo) Open(name string) error {
	ptr, err := strToPtr(strings.ReplaceAll(name, "\\", "/"))
	if err != nil {
		return err
	}
	_, _, err = m.DllInfo.NewProc(`MediaInfo_Open`).Call(m.Instance, ptr)
	if !isError(err) {
		return err
	}
	return nil
}
func (m *MediaInfo) Close() error {
	_, _, err := m.DllInfo.NewProc(`MediaInfo_Close`).Call(m.Instance)
	if !isError(err) {
		return err
	}
	return nil
}
func (m *MediaInfo) Get(streamKind Stream, streamNumber uint32, parameter string, infoKind Info, searchKind Info) (string, error) {
	_infoKind := infoKind
	if _infoKind == InfoDefault {
		_infoKind = InfoText
	}

	_searchKind := searchKind
	if _searchKind == InfoDefault {
		_searchKind = InfoName
	}

	streamKindPtr := uint32ToPtr(uint32(streamKind))
	streamNumberPtr := uint32ToPtr(streamNumber)
	infoKindPtr := uint32ToPtr(uint32(_infoKind))
	searchKindPtr := uint32ToPtr(uint32(_searchKind))
	pp, err := strToPtr(parameter)
	if err != nil {
		return "", err
	}
	str, _, err := m.DllInfo.NewProc(`MediaInfo_Get`).Call(m.Instance, streamKindPtr, streamNumberPtr, pp, infoKindPtr, searchKindPtr)
	if !isError(err) {
		return "", nil
	}

	return ptrToStr(str), nil
}
func (m *MediaInfo) GetI(streamKind Stream, streamNumber uint32, parameter uint32, infoKind ...Info) (string, error) {
	_s := InfoText
	if len(infoKind) > 0 {
		_s = infoKind[0]
	}
	streamKindPtr := uint32ToPtr(uint32(streamKind))
	streamNumberPtr := uint32ToPtr(streamNumber)
	infoKindPtr := uint32ToPtr(uint32(_s))
	pp := uint32ToPtr(parameter)
	str, _, err := m.DllInfo.NewProc(`MediaInfo_Get`).Call(m.Instance, streamKindPtr, streamNumberPtr, pp, infoKindPtr)
	if !isError(err) {
		return "", err
	}

	return ptrToStr(str), nil
}
func strToPtr(str string) (uintptr, error) {
	fromString, err := windows.UTF16PtrFromString(str)
	if err != nil {
		return 0, err
	}
	return uintptr(unsafe.Pointer(fromString)), nil
}
func uint32ToPtr(i uint32) uintptr {
	return uintptr(i)
}
func isError(err error) bool {
	return errors.Is(err, windows.ERROR_SUCCESS)
}
func ptrToStr(str uintptr) string {
	u := (*uint16)(unsafe.Pointer(str))
	return windows.UTF16PtrToString(u)
}
