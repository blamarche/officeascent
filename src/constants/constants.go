/*

    Constants.go - util functions

*/
package constants

//*********************************************
//*********************************************

const VERSION = "0.1 Alpha"
const START_RADIUS = 7

//ansi colors for game runes
const (
	COLOR_PLAYER = 2
	COLOR_WALL = 5
)

//runes
const (
	RUNE_PLAYER = "@"
	RUNE_WALL = "#"
    RUNE_FLOOR = " "
    RUNE_HIDDEN = "."
)

//walls
const (
    WALL_NONE = iota
    WALL_NORMAL
)

//game states
const (
	STATE_NORMAL = iota
	STATE_LOOK
)

//map gen methods
const (
    MAPGEN_PROC = iota
    MAPGEN_BSP
    MAPGEN_CELL
    MAPGEN_NOISE
)

//player states 
const (
    PLAYERSTATE_NORMAL = iota
)