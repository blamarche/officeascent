/*

    Entity.go - "class" to handle entity specific operations

*/
package entity


import (
    "github.com/blamarche/ansiterm"
    "fmt"
)


const (
    KIND_PLAYER = iota   
    KIND_MONSTER
)


type Entity struct {
    x int
    y int
    rune string
    hp int
    kind int //type of entity (player, monster, pet, item, etc)
}

func NewEntity(rune string, hp int, kind int) *Entity {
    return NewEntityXY(0,0,rune,hp,kind)
}

func NewEntityXY(x int, y int, rune string, hp int, kind int) *Entity {
    var e Entity = Entity{x, y, rune, hp, kind}
    e.Display()
    return &e
}

func (e *Entity) GetXY() (int,int) {
    return e.x, e.y
}

func (e *Entity) SetXY(x,y int) {
    e.x=x
    e.y=y
}

//these will eventually return whether movement was sucessful
func (e *Entity) MoveUp() bool {
    e.Clear()
    e.y--
    e.Display()
    
    return true
}

func (e *Entity) MoveDown() bool {
    e.Clear()
    e.y++
    e.Display()
    return true
}

func (e *Entity) MoveLeft() bool {
    e.Clear()
    e.x--
    e.Display()
    
    return true 
}

func (e *Entity) MoveRight() bool {
    e.Clear()    
    e.x++
    e.Display()
    
    return true
}

//handle show/hide of entity
func (e *Entity) Clear() {
    ansiterm.MoveToXY(e.x,e.y)
    fmt.Print(" ") //eventually will lookup floor tile type below the entity instead of blank space    
}

func (e *Entity) Display() {
    ansiterm.MoveToXY(e.x,e.y)
    fmt.Print(e.rune)    
}


