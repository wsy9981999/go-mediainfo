package go_mediainfo

type Stream uint32

const (
	StreamGeneral Stream = iota
	StreamVideo
	StreamAudio
	StreamText
	StreamOther
	StreamImage
	StreamMenu
	StreamMax
)

type Info uint32

const (
	InfoName Info = iota
	InfoText
	InfoMeasure
	InfoOptions
	InfoNameText
	InfoMeasureText
	InfoInfo
	InfoHowTo
	InfoMax
	InfoDefault
)
