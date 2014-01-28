/*

    World.go - interface to handle worlds

*/
package world


import (
    "github.com/blamarche/ansiterm"
    "math/rand"
    "fmt"
)

var wallRune string = "#"

type Map struct {
    Walls [][]int    
    //Monsters
    //Items
    //Doors
    //Other
}

//generate a new floor
func NewMap(width, height, floor int) *Map {
    m := Map{}    
    m.Walls = make([][]int, width)
    
    for i := range m.Walls {
        m.Walls[i] = make([]int, height)
    }
    
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            m.Walls[x][y] = 0
        }
    }
    
    m.GenerateWalls()
    
    return &m
}

func (m *Map) GenerateWalls() {
    for x := 0; x < len(m.Walls); x++ {
        for y := 0; y < len(m.Walls[x]); y++ {
            m.Walls[x][y] = rand.Intn(2)
        }
    }
}

func (m *Map) Display(px, py, wx, wy int) {
    //TODO: build whole lines then display, need to reverse x/y indices, woops    
    for x := px-wx/2; x < px+wx/2; x++ {
        if x>=0 && x<len(m.Walls) {        
            for y := py-wy/2; y < py+wy/2; y++ {
                if y>=0 && y<len(m.Walls[x]) && m.Walls[x][y]!=0 {
                    ansiterm.SetFGColor(4)
                    dx := wx/2+x-px
                    dy := wy/2+y-py        
                    if dy>1 && dy<wy-1 {
                        ansiterm.MoveToXY(dx,dy)
                        fmt.Print(wallRune)
                    }
                }
            }
        }        
    }
}

