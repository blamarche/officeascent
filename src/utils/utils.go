/*

    Utils.go - util functions

*/
package utils


import (
    "syscall"
    "unsafe"
)


//*********************************************
//*********************************************

const VERSION = "0.1 Alpha"

//ansi colors for game runes
const (
	COLOR_PLAYER = 2
	COLOR_WALL = 5
)

//runes
const (
	RUNE_PLAYER = "@"
	RUNE_WALL = "#"
    RUNE_FLOOR = " "
)


//*********************************************
//*********************************************

//Message for top of screen
var message string = "Welcome, office adventurer!"
func SetMessage(s string) {
    message = s
}

func GetMessage() string {
    return message
}


//*********************************************
//*********************************************

//Gets size of terminal
const TIOCGWINSZ = 0x5413
type Winsize struct {
    ws_row, ws_col uint16
    ws_xpixel, ws_ypixel uint16
}

func GetWinSize() []uint16 {
    ws := Winsize{}
    syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), uintptr(TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
    return []uint16{ws.ws_row, ws.ws_col, ws.ws_xpixel, ws.ws_ypixel}
} 
