package main

/*
TODO:

Weapon/armor equip
Enemies, Basic AI
Player "attacks"
Using potions
Item randomization (attributes & runes)
Doors & keys
Ranged attacks
Ranged AI
Stats log / collectibles
Disguise
Autoexplore
Backstab

*/


import (
    "github.com/blamarche/ansiterm"
    "github.com/blamarche/Go-Term/term"
    
	"fmt"
    "os"
    "math/rand"
    "time"
    //"strconv"

    "./player"
    "./world"
    "./utils"
    "./constants"
    "./cursor"
)


// GLOBALS
var (
	license string = "GPLv2 - See LICENSE file for details"
	tick_count  int
	
    floors      [constants.MAX_MAPS]*world.Map
    worldmap    *world.Map
    
	p1          *player.Player
    cur         *cursor.Cursor
	
	wx int = 64 //screen width
	wy int = 24 //screen height
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
    for i:=1; i<=constants.MAX_MAPS; i++ {
        floors[i-1] = world.NewMap(120, 60, i) //width, height, floor
    }
    worldmap = world.SetCurrentMap(floors[0])
    p1 = player.NewPlayerXY(60, 30, 10, 10, rand.Intn(4)+4, rand.Intn(4)+4, rand.Intn(4)+4, rand.Intn(4)+4, player.CLASS_ENGINEER)			
		
    //game turn loop
    for {
        if worldmap.GameState==constants.STATE_NORMAL {
			tick_count++       
				
  			//per-turn logic and actions
			worldmap.AdvanceTurn(p1)
			worldmap.RevealTilesAround(p1.X, p1.Y, p1.LightRadius)
        }
		
		drawScreen()
        
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


//MAIN DRAW
func drawScreen() {
    //draw screen
    ansiterm.SetFGColor(7)
    ansiterm.SetBGColor(0)
	ansiterm.SetColorBright()
    //ansiterm.ClearPage()       
    
    ansiterm.MoveToXY(0, wy-1)
	fmt.Printf("Turn: %d | HP: %d/%d | Amb/Chr/Spr/Grd: %d/%d/%d/%d", tick_count, p1.Hp, p1.Max_hp, p1.Ambition, p1.Charm, p1.Spirit, p1.Greed)
    
    ansiterm.MoveToXY(0, wy)
    fmt.Printf("Weap | Armor | Status | Floor %d", worldmap.Floor)
    
    worldmap.Display(p1.X, p1.Y, wx, wy)
    
    p1.Display(wx,wy)

    if cur!=nil {
        cur.Display()            
    }

    switch worldmap.GameState {
        case constants.STATE_LOOK:
            pos := cur.GetMapXY(p1.X, p1.Y)
            if pos[1]>=0 && pos[1]<worldmap.Height && pos[0]>=0 && pos[0]<worldmap.Width && worldmap.Tiles[pos[1]][pos[0]].Revealed {
                if worldmap.Tiles[pos[1]][pos[0]].Ai!=nil {
					utils.SetMessage("You see: "+worldmap.Tiles[pos[1]][pos[0]].Ai.Name)
				} else if worldmap.Tiles[pos[1]][pos[0]].Item!=nil {
					utils.SetMessage("You see: "+worldmap.Tiles[pos[1]][pos[0]].Item.Name)
				} else {
					utils.SetMessage("")
				}		
            } else {
				utils.SetMessage("")
			}
            
        case constants.STATE_INVENTORY:
            p1.ListInventory("Use / equip which item?")
        case constants.STATE_INVENTORY_DROP:
            p1.ListInventory("Drop which item?")
    }

    ansiterm.SetFGColor(7)
    ansiterm.MoveToXY(0,0)
	fmt.Printf(utils.Space(wx-1))
	ansiterm.MoveToXY(0,0)
    fmt.Printf("%s ", utils.GetMessage())    
}


//MAIN INPUT
//**********
func doInput(input string) bool {
    if input=="Q" {
        return true
    } else {        
        
        if worldmap.GameState == constants.STATE_INVENTORY_DROP {            
            if worldmap.InBounds(p1.X, p1.Y) && worldmap.Tiles[p1.Y][p1.X].Item==nil {
                ind := p1.StringToInvIndex(input)
                if ind>=0 {
                    item := p1.RemoveInventoryItem(ind)
                    worldmap.Tiles[p1.Y][p1.X].Item = item
                    if item!=nil {
                        utils.SetMessage("You dropped: "+item.Name)                        
                    }                    
                }
            } else {
                utils.SetMessage("You can't drop here!")
            }
            
            worldmap.GameState = constants.STATE_NORMAL  
			ansiterm.ClearPage() 			
            return false
            
        } else if worldmap.GameState == constants.STATE_INVENTORY {            
            
            worldmap.GameState = constants.STATE_NORMAL

			ind := p1.StringToInvIndex(input)
			utils.SetMessage(p1.UseItem(ind))
			
			ansiterm.ClearPage() 			
			
            return false
            
        } else if worldmap.GameState == constants.STATE_LOOK {
            
            cursorInput(input)            
            switch input {                
                case " ":                    
                    cur = nil
                    worldmap.GameState = constants.STATE_NORMAL
            }
            return false
            
        } else {
            
            playerMoveInput(input)
            switch input {
                //grab
                case "g", "5": //TODO: collectibles handled differently (stats)
                    if worldmap.InBounds(p1.X, p1.Y) && worldmap.Tiles[p1.Y][p1.X].Item!=nil {
                        i := p1.AddInventoryItem(worldmap.Tiles[p1.Y][p1.X].Item)
                        if i>=0 {                            
                            utils.SetMessage("You picked up: "+worldmap.Tiles[p1.Y][p1.X].Item.Name)
                            worldmap.Tiles[p1.Y][p1.X].Item = nil
                        } else {
                            utils.SetMessage("Your inventory is full!")
                        }
                    } else {
                        utils.SetMessage("There is nothing here to get!")
                    }
                //drop
                case "d":
                    worldmap.GameState = constants.STATE_INVENTORY_DROP                
                //use
                case "i":
                    worldmap.GameState = constants.STATE_INVENTORY                
                //look
                case "L", "s":
                    cur = cursor.NewCursor(wx/2, wy/2, 1, wx, 2, wy-2, 1, 2)
                    worldmap.GameState = constants.STATE_LOOK
                //downstairs
                case ">": 
                    if worldmap.DownXY[0]==p1.X && worldmap.DownXY[1]==p1.Y && worldmap.Tiles[worldmap.DownXY[1]][worldmap.DownXY[0]].DownStair {
                        ansiterm.ClearPage()
						worldmap = world.SetCurrentMap(floors[worldmap.Floor])
                        p1.SetXY(worldmap.UpXY[0], worldmap.UpXY[1])
                    } else {
                        utils.SetMessage("No up escalator here!")
                    }
                //upstairs
                case "<": 
                    if worldmap.UpXY[0]==p1.X && worldmap.UpXY[1]==p1.Y && worldmap.Tiles[worldmap.UpXY[1]][worldmap.UpXY[0]].UpStair {
                        ansiterm.ClearPage()
						worldmap = world.SetCurrentMap(floors[worldmap.Floor-2])
                        p1.SetXY(worldmap.DownXY[0], worldmap.DownXY[1])
                    } else {
                        utils.SetMessage("No down escalator here!")
                    }
                //?
                case "?":
                    showHelp()
                /*    
                default:
                    fmt.Printf("Command not recognized: %s, press ? for help\n", input)
                */
            }
            
        }  
    }
    
    return false
}

func playerMoveInput(input string) {
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
    }
}

func cursorInput(input string) {
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
    }
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


//DISPLAY
//*******
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
    fmt.Println("Press ? for help in-game")
    fmt.Println("")
    fmt.Println("Press any key to start...")
    getKeypress()
}

func showHelp() {
    ansiterm.ClearPage()
    ansiterm.MoveToXY(0,0)
    fmt.Println("_-= OFFICE ASCENT =-_")
    fmt.Println("Version: "+constants.VERSION)
    fmt.Println("")
    fmt.Println("Player & Cursor Movement: Up/Down/Left/Right, hjklyunm, Numpad (numlock ON)")
    
    fmt.Println("")
    fmt.Println("i - Use/equip item, view inventory")
    fmt.Println("g - Grab item")
    fmt.Println("d - Drop item")
    fmt.Println("s, L - Look around")
    
    fmt.Println("Q - quit")
    fmt.Println("Press any key to continue...")
    getKeypress()
}
