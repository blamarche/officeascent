/*

    Mapgen.go - logic to generate world maps

*/
package world


import (
    "../constants"
    "math/rand"
)



func (m *Map) genPath(x,y,dir int) [4]int {
    r := [4]int{x-5,y,10,1}
    return r
}

func (m *Map) genRoom(x,y,dir int) [4]int {
    r := [4]int{0,0,1,1}
    r[2] = rand.Intn(constants.MAPGEN_PROC_ROOMSIZE_MAXX - constants.MAPGEN_PROC_ROOMSIZE_MINX) + constants.MAPGEN_PROC_ROOMSIZE_MINX
    r[3] = rand.Intn(constants.MAPGEN_PROC_ROOMSIZE_MAXY - constants.MAPGEN_PROC_ROOMSIZE_MINY) + constants.MAPGEN_PROC_ROOMSIZE_MINY
    
    is_room := rand.Intn(constants.MAPGEN_PROC_PATH_FREQ)
    
    if is_room==0 {
        vert := rand.Intn(2)
        if vert==1 {
            r[3]=1
        } else {
            r[2]=1
        }
    }
    
    switch dir {
        case constants.MAPGEN_PROC_LEFT:
            r[0] = x-r[2]
            r[1] = y-r[3]/2            
        case constants.MAPGEN_PROC_RIGHT:
            r[0] = x
            r[1] = y-r[3]/2            
        case constants.MAPGEN_PROC_UP:
            r[0] = x-r[2]/2
            r[1] = y-r[3]            
        case constants.MAPGEN_PROC_DOWN:
            r[0] = x-r[2]/2
            r[1] = y
    }        
    
    return r
}

func (m *Map) genNext(x,y,dir int) [4]int {    
    return m.genRoom(x,y,dir)
}

func (m *Map) addRoomToWorld(r [4]int) bool {    
    //check for enough space
    for i:=r[0]; i<r[0]+r[2]; i++ {
        for j:=r[1]; j<r[1]+r[3]; j++ {
            if i>=0 && i<m.Width && j>=0 && j<m.Height && m.Tiles[j][i].Wall == constants.WALL_NONE {
                return false
            }
            if i==0 || j==0 || i==m.Width-1 || j==m.Height-1 {
                return false
            }
        }
    }
    
    //now loop through and make the open room
    //TODO: doors, etc.
    for i:=r[0]; i<r[0]+r[2]; i++ {
        for j:=r[1]; j<r[1]+r[3]; j++ {
            if i>=0 && i<m.Width && j>=0 && j<m.Height {
                m.Tiles[j][i].Wall = constants.WALL_NONE
            }
        }
    }
    
    return true
}

func (m *Map) GenerateWalls() {
    if gen_method == constants.MAPGEN_PROC {
        
        //PROC method        
        /*
        Fill the whole map with solid earth
        Dig out a single room in the centre of the map
        Pick a wall of any room
        Decide upon a new feature to build
        See if there is room to add the new feature through the chosen wall
        If yes, continue. If no, go back to step 3
        Add the feature through the chosen wall
        Go back to step 3, until the dungeon is complete
        */
        
        for y := 0; y < len(m.Tiles); y++ {
            for x := 0; x < len(m.Tiles[y]); x++ {
                m.Tiles[y][x].Wall = constants.WALL_NORMAL
            }
        }
        
        //var rooms [constants.MAPGEN_PROC_ITERATIONS][4]int
        ri := 0
        m.rooms[ri] = m.genRoom(m.Width/2, m.Height/2, constants.MAPGEN_PROC_LEFT) //wont really matter which params are here due to override
        m.rooms[ri][0] = m.Width/2 - m.rooms[ri][2]/2
        m.rooms[ri][1] = m.Height/2 - m.rooms[ri][3]/2
        m.addRoomToWorld(m.rooms[ri])
        
        tot:=1
        for i:=1; i<constants.MAPGEN_PROC_ITERATIONS; i++ {
            tot++
            
            //pick random wall of random room & wall-side
            r := m.rooms[rand.Intn(i)]
            dir := rand.Intn(4)
            var x, y int
            
            switch dir {
                case constants.MAPGEN_PROC_LEFT:
                    x = r[0]
                    y = rand.Intn(r[3])+r[1]
                case constants.MAPGEN_PROC_RIGHT:
                    x = r[0]+r[2]
                    y = rand.Intn(r[3])+r[1]
                case constants.MAPGEN_PROC_UP:
                    x = rand.Intn(r[2])+r[0]
                    y = r[1]
                case constants.MAPGEN_PROC_DOWN:
                    x = rand.Intn(r[2])+r[0]
                    y = r[1]+r[3]
            }
            
            next := m.genNext(x,y,dir)
            if m.addRoomToWorld(next) {
                m.rooms[i]=next 
            } else {
                i--
            }
            
            if tot>=constants.MAPGEN_PROC_ITERATIONS*5 {
                break
            }
        }
        

    } else {
    
        //NOISE
        
        for y := 0; y < len(m.Tiles); y++ {
            for x := 0; x < len(m.Tiles[y]); x++ {
                if x==0 || x==len(m.Tiles[y])-1 || y==0 || y==len(m.Tiles)-1 {
                    m.Tiles[y][x].Wall = constants.WALL_NORMAL
                } else {
                    if rand.Intn(4)==1 {
                        m.Tiles[y][x].Wall = constants.WALL_NORMAL
                    } else {
                        m.Tiles[y][x].Wall = constants.WALL_NONE
                    }
                }
            }
        }
    }
}
