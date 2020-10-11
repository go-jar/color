package gocolor

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(string(GetBlackMsg([]byte("Black"))))
	fmt.Println(string(GetRedMsg([]byte("Red"))))
	fmt.Println(string(GetGreenMsg([]byte("Green"))))
	fmt.Println(string(GetBlueMsg([]byte("Blue"))))
	fmt.Println(string(GetMagantaMsg([]byte("Maganta"))))
	fmt.Println(string(GetCyanMsg([]byte("Cyan"))))
	fmt.Println(string(GetWhiteMsg([]byte("White"))))
}
