/*

    Player.go - interface to handle player specific operations

*/
package player


import (
    "github.com/blamarche/ansiterm"
    "fmt"
    
    "../utils"
)


const (
    CLASS_ENGINEER = iota   
    CLASS_OTHER
)

const (
    STATE_NORMAL = iota
    STATE_INVENTORY
    STATE_HELP
)


type Player struct {
    X int
    Y int
    
	Hp int
	Max_hp int
	
	Ambition int
	Charm int
	Spirit int
	Greed int
	
    Class int //classof player
    State int
}

func NewPlayer(hp int, max_hp int, amb int, charm int, spirit int, greed int, class int) *Player {
    return NewPlayerXY(0,0, hp, max_hp, amb, charm, spirit, greed, class)
}

func NewPlayerXY(x int, y int, hp int, max_hp int, amb int, charm int, spirit int, greed int, class int) *Player {
    var e Player = Player{x, y,  hp, max_hp, amb, charm, spirit, greed, class,0}
    return &e
}

func (e *Player) GetXY() (int,int) {
    return e.X, e.Y
}

func (e *Player) SetXY(x,y int) {
    e.X=x
    e.Y=y
}

//these will eventually return whether movement was sucessful
func (e *Player) MoveUp() bool {
    e.Y--
    return true
}

func (e *Player) MoveDown() bool {
    e.Y++
    return true
}

func (e *Player) MoveLeft() bool {
    e.X--
    return true 
}

func (e *Player) MoveRight() bool {
    e.X++
    return true
}

func (e *Player) Display(wx, wy int) {
    ansiterm.SetFGColor(utils.COLOR_PLAYER)
    ansiterm.MoveToXY(wx/2,wy/2)
    fmt.Print(utils.RUNE_PLAYER)
}

