/*

    World.go - interface to handle worlds

*/
package world


import (
    "github.com/blamarche/ansiterm"
    "github.com/blamarche/astar"

    "math/rand"
    "math"
    "fmt"
    
    "../item"
    "../constants"  
)

type MapTile struct {
    Wall int
    Item *item.Item
    Revealed bool
    //Enemy
    //Door
    //Other
}

type Map struct {  //[y][x] or [row][col]
    Tiles [][]*MapTile
    PathData astar.MapData 
    Floor int
    Width int
    Height int
    GameState int
    
    rooms [constants.MAPGEN_PROC_ITERATIONS][4]int
}

var (
    CurrentMap *Map
    
    gen_method int = constants.MAPGEN_PROC
)


//CREATION
//********
func NewMap(width, height, floor int) *Map {
    m := Map{}    

    m.Floor = floor 
    m.Width = width 
    m.Height = height 
    
    m.PathData = astar.NewMapData(m.Width, m.Height)
    for y:=0; y<m.Height; y++ {
        for x:=0; x<m.Width; x++ {
            m.PathData[x][y] = astar.LAND            
        }
    }

    m.Tiles = make([][]*MapTile, height)
    
    for i := range m.Tiles {
        m.Tiles[i] = make([]*MapTile, width)
    }
    
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            m.Tiles[y][x] = &MapTile{}
            m.Tiles[y][x].Wall = 0
            m.Tiles[y][x].Item = nil
            m.Tiles[y][x].Revealed = false
            if constants.DEBUG_REVEAL_ALL==1 {
                m.Tiles[y][x].Revealed = true
            }
        }
    }
    
    m.GenerateWalls()
    m.GenerateItems()
    
    m.RefreshPathData()

    CurrentMap = &m
    return &m
}

func (m *Map) GenerateItems() {
    tilecount := len(m.Tiles[0])*len(m.Tiles)
    ratio := float32(tilecount) / 10000.0
    
    for i:=0; i<len(item.ItemList); i++ {        
        if m.Floor >= item.ItemList[i].Floor_min && m.Floor <= item.ItemList[i].Floor_max {
            count := item.ItemList[i].Max_per_floor * ratio
            
            for j:=0; j<=int(count)+1; j++ {
                x := rand.Intn(len(m.Tiles[0]))
                y := rand.Intn(len(m.Tiles))
                
                if m.Tiles[y][x].Wall==constants.WALL_NONE {
                    m.Tiles[y][x].Item = item.ItemList[i].Clone()
                }
            }
        }
    }
}


//GAMEPLAY
//********
func (m *Map) InBounds(x, y int) bool {
    return (x>=0 && y>=0 && x<m.Width && y<m.Height)
}

func (m *Map) RefreshPathData() {
    for y:=0; y<m.Height; y++ {
        for x:=0; x<m.Width; x++ {
            if m.Tiles[y][x].Wall!=constants.WALL_NONE {
                m.PathData[x][y] = astar.WALL
            } else {
                m.PathData[x][y] = astar.LAND
            }
        }
    }
}

func (m *Map) GetPath(sx, sy, ex, ey int) []*astar.Node {
    path := astar.Astar(m.PathData, sx, sy, ex, ey, true)
    return path
}

func (m *Map) AdvanceTurn() {
    //enemy ai, etc

    //TODO: add dirty flag to this
    m.RefreshPathData()
}

func (m *Map) RevealTilesAround(x, y, radius int) {
    //not really a radius
    
    for xx:=x-radius; xx<x+radius; xx++ {
        for yy:=y-radius; yy<y+radius; yy++ {
            diff := int(math.Abs(float64(x-xx)) + math.Abs(float64(y-yy)))
            if diff < radius && xx>=0 && xx<m.Width && yy>=0 && yy<m.Height {
                m.Tiles[yy][xx].Revealed = true
            }
        }
    }
}

func (m *Map) IsBlocked(x, y int) bool {
    if constants.DEBUG_NOWALL==1 {
        return false
    }
    
    if y>=0 && y<len(m.Tiles) {
        if x>=0 && x<len(m.Tiles[y]) {
            return m.Tiles[y][x].Wall!=constants.WALL_NONE
        }
    }
    return false
}

func (m *Map) Display(px, py, wx, wy int) {
    //walls
    for y := py-wy/2; y < py+wy/2; y++ {
        if y>=0 && y<len(m.Tiles) {    
            
            var line string = ""
            
            for x := px-wx/2+1; x < px+wx/2+1; x++ {
                if x>=0 && x<len(m.Tiles[y]) && !m.Tiles[y][x].Revealed {
                    line += constants.RUNE_HIDDEN
                } else if x>=0 && x<len(m.Tiles[y]) && m.Tiles[y][x].Wall!=constants.WALL_NONE {
                    line += constants.RUNE_WALL                    
                } else {
                    line += constants.RUNE_FLOOR
                }
            }
            
            dy := wy/2+y-py
            if dy>1 && dy<wy-1 {
                ansiterm.SetFGColor(constants.COLOR_WALL)
                ansiterm.MoveToXY(0,dy)
                fmt.Print(line)
            }
        }
    }
    
    //items
    for y := py-wy/2; y < py+wy/2; y++ {
        if y>=0 && y<len(m.Tiles) {                
            for x := px-wx/2+1; x < px+wx/2+1; x++ {
                if x>=0 && x<len(m.Tiles[y]) && m.Tiles[y][x].Item!=nil {
                    dx := wx/2+x-px
                    dy := wy/2+y-py
                    if dy>1 && dy<wy-1 && m.Tiles[y][x].Revealed {
                        ansiterm.SetFGColor(m.Tiles[y][x].Item.Fgcolor)
                        ansiterm.MoveToXY(dx,dy)
                        fmt.Print(m.Tiles[y][x].Item.Rune)
                    }
                }
            }
        }
    }
}

