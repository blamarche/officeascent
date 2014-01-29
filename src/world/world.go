/*

    World.go - interface to handle worlds

*/
package world


import (
    "github.com/blamarche/ansiterm"
    "math/rand"
    "fmt"
    
    "../item"
    "../utils"    
)

//TODO: optimize into single array of a MapTile struct
type Map struct {  //[y][x] or [row][col]
    Walls [][]int  
    //Enemies
    Items [][]*item.Item
    //Doors
    //Other
}

var CurrentMap *Map


//CREATION
//********
func NewMap(width, height, floor int) *Map {
    m := Map{}    
    m.Walls = make([][]int, height)
    m.Items = make([][]*item.Item, height)
    
    for i := range m.Walls {
        m.Walls[i] = make([]int, width)
        m.Items[i] = make([]*item.Item, width)
    }
    
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            m.Walls[y][x] = 0
            m.Items[y][x] = nil
        }
    }
    
    m.GenerateWalls(floor)
    m.GenerateItems(floor)
    
    CurrentMap = &m
    return &m
}

func (m *Map) GenerateWalls(floor int) {
    _=floor
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

func (m *Map) GenerateItems(floor int) {
    tilecount := len(m.Walls[0])*len(m.Walls)
    ratio := float32(tilecount) / 10000.0
    
    for i:=0; i<len(item.ItemList); i++ {        
        if floor >= item.ItemList[i].Floor_min && floor <= item.ItemList[i].Floor_max {
            count := item.ItemList[i].Max_per_floor * ratio
            
            for j:=0; j<=int(count)+1; j++ {
                x := rand.Intn(len(m.Walls[0]))
                y := rand.Intn(len(m.Walls))
                
                if m.Walls[y][x]==0 {
                    m.Items[y][x] = item.ItemList[i].Clone()
                }
            }
        }
    }
}


//GAMEPLAY
//********

func (m *Map) IsBlocked(x, y int) bool {
    if y>=0 && y<len(m.Walls) {
        if x>=0 && x<len(m.Walls[y]) {
            return m.Walls[y][x]==1
        }
    }
    return false
}

func (m *Map) Display(px, py, wx, wy int) {
    //walls
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
    
    //items
    for y := py-wy/2; y < py+wy/2; y++ {
        if y>=0 && y<len(m.Walls) {                
            for x := px-wx/2+1; x < px+wx/2+1; x++ {
                if x>=0 && x<len(m.Walls[y]) && m.Items[y][x]!=nil {
                    dx := wx/2+x-px
                    dy := wy/2+y-py
                    if dy>1 && dy<wy-1 {
                        ansiterm.SetFGColor(m.Items[y][x].Fgcolor)
                        ansiterm.MoveToXY(dx,dy)
                        fmt.Print(m.Items[y][x].Rune)
                    }
                }
            }
        }
    }
}

