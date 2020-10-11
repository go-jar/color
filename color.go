package gocolor

import (
	"bytes"
)

func GetBlackMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;0m"), msg)
}

func GetRedMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;31m"), msg)
}

func GetGreenMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;32m"), msg)
}

func GetYellowMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;33m"), msg)
}

func GetBlueMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;34m"), msg)
}

func GetMagantaMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;35m"), msg)
}

func GetCyanMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;36m"), msg)
}

func GetWhiteMsg(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;37m"), msg)
}

func ColorMsg(color []byte, msg []byte) []byte {
	buf := bytes.NewBuffer(color)

	buf.Write(msg)
	buf.Write([]byte("\033[0m"))

	return buf.Bytes()
}
