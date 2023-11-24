package go_mediainfo

import (
	"testing"
)

func TestReadFrameCount(t *testing.T) {
	info, err := NewMediaInfo()
	if err != nil {
		panic(err)
	}
	defer func() {
		info.Destroy()
	}()
	err = info.Option("CharSet", "UTF-8")
	if err != nil {
		panic(err)
	}
	err = info.Open("./test.mp4")
	if err != nil {
		panic(err)
	}
	defer info.Close()
	s, err := info.Get(StreamGeneral, 0, "FileSize", InfoDefault, InfoDefault)
	if err != nil {
		panic(err)
	}
	if s != "2199442" {
		t.Failed()
	}
}
