package gocolor

import (
	"bytes"
)

func RenderMsgBlack(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;0m"), msg)
}

func RenderMsgRed(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;31m"), msg)
}

func RenderMsgGreen(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;32m"), msg)
}

func RenderMsgYellow(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;33m"), msg)
}

func RenderMsgBlue(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;34m"), msg)
}

func RenderMsgMaganta(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;35m"), msg)
}

func RenderMsgCyan(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;36m"), msg)
}

func RenderMsgWhite(msg []byte) []byte {
	return ColorMsg([]byte("\033[01;37m"), msg)
}

func ColorMsg(color []byte, msg []byte) []byte {
	buf := bytes.NewBuffer(color)

	buf.Write(msg)
	buf.Write([]byte("\033[0m"))

	return buf.Bytes()
}
