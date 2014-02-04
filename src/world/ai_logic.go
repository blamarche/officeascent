/*

    Ai_logic.go - handle ai logic

*/
package world


import (
	"math/rand"

	"../interfaces"
	"../constants"
	"../ai"
)



func (m *Map) StepAIAt(x, y int, p1 interfaces.IPlayer) {
	a := m.Tiles[y][x].Ai 
	
	//TODO: different movement types, attacking player, etc
	if !a.Dirty {
		px, py := p1.GetXY()
		
		xx:=x
		yy:=y
		
		switch a.Pattern {
			case ai.PATTERN_RANDOM:
				dir := rand.Intn(9)
				switch dir {
					case 0://up
						yy--
					case 1://down	
						yy++
					case 2://left
						xx--
					case 3://right
						xx++
					case 4://upleft
						yy--
						xx--
					case 5://upright
						yy--
						xx++
					case 6://downleft
						yy++
						xx--
					case 7://downright
						yy++
						xx++
					case 8:
						a.Dirty=true
						return
				}

			case ai.PATTERN_WANDER:				
				//pick rand nonwall dest and move towards it
				if a.Path==nil {
					destx, desty := m.GetRandomOpenXY()
					a.PathIndex = 0
					a.Path = m.GetPath(x,y,destx,desty) //might be nil, but thats ok
				}

				if a.Path!=nil {
					wait := rand.Intn(3)
					if wait!=0 {
						//move to next in node
						a.PathIndex++
						xx = a.Path[a.PathIndex].X
						yy = a.Path[a.PathIndex].Y

						//if at last index, set path to nil
						if a.PathIndex >= len(a.Path)-1 {
							a.PathIndex = 0
							a.Path = nil					
						}
					}
				}
			
			case ai.PATTERN_STILL:
				a.Dirty = true 
				return
		}
		
		if m.InBounds(xx,yy) && m.Tiles[yy][xx].Ai==nil && m.Tiles[yy][xx].Wall==constants.WALL_NONE && xx!=px && yy!=py {
			m.Tiles[yy][xx].Ai = a
			m.Tiles[y][x].Ai = nil
		} else {
			if a.Path!=nil && x!=xx && y!=yy {
				a.Path = nil
				a.PathIndex = 0
			}			
		}
		
		a.Dirty = true
		
	}
}

func (m *Map) GetRandomOpenXY() (int, int) {
	for {
		x:= rand.Intn(m.Width)
		y:= rand.Intn(m.Height)

		if m.Tiles[y][x].Wall == constants.WALL_NONE && m.Tiles[y][x].Ai==nil {
			return x, y
		}
	}
}