package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(string(RenderMsgBlack([]byte("Black"))))
	fmt.Println(string(RenderMsgRed([]byte("Red"))))
	fmt.Println(string(RenderMsgGreen([]byte("Green"))))
	fmt.Println(string(RenderMsgYellow([]byte("Yellow"))))
	fmt.Println(string(RenderMsgBlue([]byte("Blue"))))
	fmt.Println(string(RenderMsgMaganta([]byte("Maganta"))))
	fmt.Println(string(RenderMsgCyan([]byte("Cyan"))))
	fmt.Println(string(RenderMsgWhite([]byte("White"))))
}
