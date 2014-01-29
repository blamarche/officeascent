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

var CurrentMap *Map

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
    
    CurrentMap = &m
    return &m
}

func (m *Map) GenerateWalls() {
    for y := 0; y < len(m.Walls); y++ {
        for x := 0; x < len(m.Walls[y]); x++ {
            if x==0 || x==len(m.Walls[y])-1 || y==0 || y==len(m.Walls)-1 {
                m.Walls[y][x] = 1
            } else {
                if rand.Intn(4)==1 {
                    m.Walls[y][x] = 1
                } else {
                    m.Walls[y][x] = 0
                }
            }
        }
    }
    
}

func (m *Map) IsBlocked(x, y int) bool {
    if y>=0 && y<len(m.Walls) {
        if x>=0 && x<len(m.Walls[y]) {
            return m.Walls[y][x]==1
        }
    }
    return false
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

