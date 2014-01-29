package main


import (
    "github.com/blamarche/ansiterm"
    "github.com/blamarche/Go-Term/term"
    
	"fmt"
    "os"
    "math/rand"
    "time"

    "./player"
	//"./item"
    "./world"
    "./utils"
)

//TODO: "Look" functionality w/ "cursor"


// GLOBALS
var (
	license = "GPLv2 - See LICENSE file for details"
	tick_count int
	
    worldmap *world.Map
    
	p1 *player.Player
	
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

    showIntro()
    
    //config world and player here    
    worldmap = world.NewMap(100, 60, 1) //width, height, floor
    p1 = player.NewPlayerXY(50,30, 10, 10, 8, 8, 8, 8, player.CLASS_ENGINEER)	
    
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
		
        worldmap.Display(p1.X, p1.Y, wx, wy)
        
        p1.Display(wx,wy)

        ansiterm.SetFGColor(7)
		ansiterm.MoveToXY(0,0)
        fmt.Printf("%s ", utils.GetMessage())
        
        
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
            case "ESC:A", "8", "k":
                p1.MoveUp()
            case "ESC:B", "2", "j":
                p1.MoveDown()
            case "ESC:C", "6", "l":
                p1.MoveRight()
            case "ESC:D", "4", "h":
                p1.MoveLeft() 
            case "7", "y": 
                p1.MoveUpLeft()
            case "9", "u": 
                p1.MoveUpRight()
            case "1", "b": 
                p1.MoveDownLeft()
            case "3", "n": 
                p1.MoveDownRight()
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
    ws := utils.GetWinSize()
    wy = int(ws[0])//rows
    wx = int(ws[1])//cols
    
    if wx<64 || wy<24 {
        fmt.Printf("%d, %d", wx, wy)
        panic("Your terminal is not big enough! (64x24)")
    }
    
    wx = wx/2*2 //dont count odd width 
    
    term.Echo(false)
	term.KeyPress()
    ansiterm.HideCursor()
    ansiterm.ClearPage()
}

func showIntro() {
    ansiterm.ClearPage()
    ansiterm.MoveToXY(0,0)
    fmt.Println("_-= OFFICE ASCENT =-_")
    fmt.Println("Version: "+utils.VERSION)
    fmt.Println("")
    fmt.Println("Movement Controls: Arrows, Numpad (numlock ON), hjklyubn keys")
    fmt.Println("")
    fmt.Println("Press any key to start...")
    getKeypress()
}
