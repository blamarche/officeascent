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

//Message for top of screen, maybe move this into world.go?
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
