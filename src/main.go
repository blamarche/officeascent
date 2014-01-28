package main


import (
    "github.com/blamarche/ansiterm"
    "github.com/blamarche/Go-Term/term"
    
	"fmt"
    "os"
    "syscall"
    "unsafe"
    "math/rand"
    "time"

    "./player"
	//"./item"
    "./world"
    //"./utils"
)


// GLOBALS
var (
	license = "GPLv2 - See LICENSE file for details"
	tick_count int
	
    worldmap *world.Map
    
	p1 *player.Player
    message string = "Welcome, office adventurer!"
	
	wx = 64 //screen width
	wy = 24 //screen height
)


// MAIN FUNCTION
func main() {
    rand.Seed( time.Now().UTC().UnixNano())

    input := ""
    initTerm()
    
    defer term.RestoreTerm()    
    defer ansiterm.ResetTerm(ansiterm.NORMAL)
    defer term.Echo(true)
    defer ansiterm.ClearPage()

    //config world and player here    
    worldmap = world.NewMap(150, 150, 1) //width, height, floor
    p1 = player.NewPlayerXY(140,140,"@", 10, 10, 8, 8, 8, 8, player.CLASS_ENGINEER)	
    
	//fmt.Println(item.ItemList[0].Use_desc)
	
    //game turn loop
    for {
        tick_count++       
        
		//draw screen
        ansiterm.SetFGColor(7)
        ansiterm.SetBGColor(0)
        ansiterm.ClearPage()       
        
        ansiterm.MoveToXY(0, wy-1)
        fmt.Printf("Turn: %d | HP: %d/%d | Amb/Chr/Spr/Grd: %d/%d/%d/%d", tick_count, p1.Hp, p1.Max_hp, p1.Ambition, p1.Charm, p1.Spirit, p1.Greed)
        
        ansiterm.MoveToXY(0, wy)
        fmt.Printf("ITEM / ARMOR / STATUS / DEBUG x/y: %d, %d", p1.X, p1.Y)
		
		ansiterm.MoveToXY(0,0)
        fmt.Printf("%s ", message)
        
        worldmap.Display(p1.X, p1.Y, wx, wy)
        
        p1.Display(wx,wy)

		//wait for input
        ansiterm.MoveToXY(0,0)
		input = getKeypress()
        
        if input!="" {
            if doInput(input) {
                break
            }
        }
        
        
    }	
}


// FUNCTIONS
func doInput(input string) bool {
    if input=="Q" {
        return true
    } else {        
        switch input {
            case "ESC:A":
                p1.MoveUp()
            case "ESC:B":
                p1.MoveDown()
            case "ESC:C":
                p1.MoveRight()
            case "ESC:D":
                p1.MoveLeft()                
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