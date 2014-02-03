/*

    Ai_logic.go - handle ai logic

*/
package world


import (
	"math/rand"

	"../interfaces"
	"../constants"
)



func (m *Map) StepAIAt(x, y int, p1 interfaces.IPlayer) {
	px, py := p1.GetXY()
	ai := m.Tiles[y][x].Ai 
	
	//TODO: different movement types, attacking player, etc
	if !ai.Dirty {
	
		_=px
		_=py
		
		xx:=x
		yy:=y
		
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
				ai.Dirty=true
				return
		}
		
		if m.InBounds(xx,yy) && m.Tiles[yy][xx].Ai==nil && m.Tiles[yy][xx].Wall==constants.WALL_NONE {
			m.Tiles[yy][xx].Ai = ai
			m.Tiles[y][x].Ai = nil
		}
		
		ai.Dirty = true
		
	}
}