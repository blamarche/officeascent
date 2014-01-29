/*

    World.go - interface to handle worlds

*/
package world


import (
    "github.com/blamarche/ansiterm"
    "math/rand"
    "fmt"
    
    "../utils"    
)

type Map struct {
    Walls [][]int   //[y][x] or [row][col]
    //Monsters
    //Items
    //Doors
    //Other
}

//generate a new floor
func NewMap(width, height, floor int) *Map {
    m := Map{}    
    m.Walls = make([][]int, height)
    
    for i := range m.Walls {
        m.Walls[i] = make([]int, width)
    }
    
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            m.Walls[y][x] = 0
        }
    }
    
    m.GenerateWalls()
    
    return &m
}

func (m *Map) GenerateWalls() {
    for x := 0; x < len(m.Walls); x++ {
        for y := 0; y < len(m.Walls[x]); y++ {
            m.Walls[y][x] = rand.Intn(2)
        }
    }
}

func (m *Map) Display(px, py, wx, wy int) {
    for y := py-wy/2; y < py+wy/2; y++ {
        if y>=0 && y<len(m.Walls) {    
            
            var line string = ""
            
            for x := px-wx/2+1; x < px+wx/2+1; x++ {
                if x>=0 && x<len(m.Walls[y]) && m.Walls[y][x]!=0 {
                    line += utils.RUNE_WALL                    
                } else {
                    line += utils.RUNE_FLOOR
                }
            }
            
            dy := wy/2+y-py
            if dy>1 && dy<wy-1 {
                ansiterm.SetFGColor(utils.COLOR_WALL)
                ansiterm.MoveToXY(0,dy)
                fmt.Print(line)
            }
        }
    }
}

