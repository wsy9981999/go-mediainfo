package go_mediainfo

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

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
