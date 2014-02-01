package main

/*
TODO:

Inventory & item pickup
Doors & keys
Item randomization (attributes & runes)
Enemies, Basic AI
Player "attacks", weapon equip
Using potions
Ranged AI
Stats log
Disguise
Backstab

*/


import (
    "github.com/blamarche/ansiterm"
    "github.com/blamarche/Go-Term/term"
    
	"fmt"
    "os"
    "math/rand"
    "time"
    "strconv"

    "./player"
    "./world"
    "./utils"
    "./constants"
    "./cursor"
)


// GLOBALS
var (
	license = "GPLv2 - See LICENSE file for details"
	tick_count int
	
    worldmap *world.Map
    
	p1 *player.Player
    cur *cursor.Cursor
	
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
    if constants.DEBUG_NOCLEAR==0 {
        defer ansiterm.ClearPage()
    }

    showIntro()
    
    //config world and player here    
    worldmap = world.NewMap(150, 80, 1) //width, height, floor
    p1 = player.NewPlayerXY(75,40, 10, 10, 8, 8, 8, 8, player.CLASS_ENGINEER)			
		
    //game turn loop
    for {
        tick_count++       
        
        //per-turn logic and actions
        worldmap.AdvanceTurn()
        worldmap.RevealTilesAround(p1.X, p1.Y, p1.LightRadius)
        
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

        if cur!=nil {
            cur.Display()            
        }

        switch worldmap.GameState {
            case constants.STATE_LOOK:
                pos := cur.GetMapXY(p1.X, p1.Y)
                if pos[1]>=0 && pos[1]<worldmap.Height && pos[0]>=0 && pos[0]<worldmap.Width && worldmap.Tiles[pos[1]][pos[0]].Revealed && worldmap.Tiles[pos[1]][pos[0]].Item!=nil {
                    utils.SetMessage("You see: "+worldmap.Tiles[pos[1]][pos[0]].Item.Name)
                } else {
                    utils.SetMessage("Cur: "+strconv.Itoa(pos[0])+", "+strconv.Itoa(pos[1]))
                }
        }

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
        
        if worldmap.GameState == constants.STATE_LOOK {
            
            switch input {
                case "ESC:A", "8", "k":
                    cur.MoveUp()
                case "ESC:B", "2", "j":
                    cur.MoveDown()
                case "ESC:C", "6", "l":
                    cur.MoveRight()
                case "ESC:D", "4", "h":
                    cur.MoveLeft() 
                case "7", "y": 
                    cur.MoveUpLeft()
                case "9", "u": 
                    cur.MoveUpRight()
                case "1", "b": 
                    cur.MoveDownLeft()
                case "3", "n": 
                    cur.MoveDownRight()
                case " ":                    
                    cur = nil
                    worldmap.GameState = constants.STATE_NORMAL
            }
            return false
            
        } else {

            switch input {
                case "L", "s":
                    cur = cursor.NewCursor(wx/2, wy/2, 1, wx, 2, wy-2, 1, 2)
                    worldmap.GameState = constants.STATE_LOOK

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
    }
    
    return false
}

func getKeypress() string {
    var buffer [1]byte
    os.Stdin.Read(buffer[:])
    
    if buffer[0]==27 {
        getKeypress()
        return "ESC:"+getKeypress()
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
    fmt.Println("Version: "+constants.VERSION)
    fmt.Println("")
    fmt.Println("Movement Controls: Arrows, Numpad (numlock ON), hjklyubn keys")
    fmt.Println("")
    fmt.Println("Press any key to start...")
    getKeypress()
}
