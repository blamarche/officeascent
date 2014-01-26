package main


import (
    "github.com/blamarche/ansiterm"
    "github.com/blamarche/Go-Term/term"
    
	"fmt"
    "os"
    "syscall"
    "unsafe"
        
    "./entity"
	"./item"
)


// GLOBALS
var (
	license = "GPLv2 - See LICENSE file for details"
	tick_count int
	
	player *entity.Entity
	
	wx = 64 //screen width
	wy = 24 //screen height
)


// MAIN FUNCTION
func main() {
    input := ""
    initTerm()
    
    defer term.RestoreTerm()    
    defer ansiterm.ResetTerm(ansiterm.NORMAL)
    defer term.Echo(true)
    defer ansiterm.ClearPage()

    //config world and player here    
    player = entity.NewEntityXY(10,10,"X", 10, entity.KIND_PLAYER)
    
	fmt.Println(item.ItemList)
	
    //game turn loop
    for {
        tick_count++       
        
		//draw screen	
        ansiterm.MoveToXY(0,2)
        fmt.Printf("Tick: %d", tick_count)
        
        ansiterm.MoveToXY(0,4)
        fmt.Printf("w/h: %d, %d", wx, wy)
		
		ansiterm.MoveToXY(0,0)
		
		//wait for input
		input = getKeypress()
        
        if input!="" {
            if doInput(input) {
                break
            }
        }
        
        
    }	
}


// CHANNELS 
func inputChannel(in chan string) {
    for {
        s:=getKeypress()
        in <- s
    }
}


// FUNCTIONS
func doInput(input string) bool {
    if input=="q" {
        return true
    } else {        
        cx, cy := player.GetXY()

        switch input {
            case "ESC:A":
                if cy>1 {
                    player.MoveUp()
                }
            case "ESC:B":
                if cy<wy {
                    player.MoveDown()
                }
            case "ESC:C":
                if cx<wx {
                   player.MoveRight()
                }
            case "ESC:D":
                if cx>1 {
                    player.MoveLeft()
                }
            default:
                fmt.Printf("You pressed: %s\n", input)
        }            
    }
    
    return false
}

func getKeypress() string {
    var buffer [1]byte
    os.Stdin.Read(buffer[:])
    
    if buffer[0]==27 {
        getKeypress()
        return "ESC:"+getKeypress();
    }
    
    return string(buffer[:])
}

func initTerm() {
    ws := GetWinSize()
    wy = int(ws.ws_row)
    wx = int(ws.ws_col)
    
    if wx<64 || wy<24 {
        fmt.Printf("%d, %d", wx, wy)
        panic("Your terminal is not big enough! (64x24)")
    }
    
    term.Echo(false)
	term.KeyPress()
    ansiterm.HideCursor()
    ansiterm.ClearPage()
    //ansiterm.SetFGColor(5)
}


//gets size of terminal TODO: move to util module
const TIOCGWINSZ = 0x5413
type winsize struct {
    ws_row, ws_col uint16
    ws_xpixel, ws_ypixel uint16
}

func GetWinSize() winsize {
    ws := winsize{}
    syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), uintptr(TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
    return ws
} 