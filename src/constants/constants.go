/*

    Constants.go - util functions

*/
package constants

//*********************************************
//*********************************************

const VERSION = "0.1 Alpha"
const START_RADIUS = 2
const MAX_MAPS = 42

//ansi colors for game runes
const (
	COLOR_PLAYER = 2
	COLOR_WALL = 5
    COLOR_STAIR = 2
)

//runes
const (
	RUNE_PLAYER = "@"
	RUNE_WALL = "#"
    RUNE_FLOOR = "."
    RUNE_HIDDEN = " "
    RUNE_BLANK = " "
)

//walls
const (
    WALL_NONE = iota
    WALL_NORMAL
    WALL_BLANK
)

//game states
const (
	STATE_NORMAL = iota
	STATE_LOOK
    STATE_INVENTORY
    STATE_INVENTORY_DROP
)

//map gen methods
const (
    MAPGEN_PROC = iota
    MAPGEN_BSP
    MAPGEN_CELL
    MAPGEN_NOISE
)
const (
    MAPGEN_PROC_ITERATIONS = 100
    MAPGEN_PROC_ROOMSIZE_MINY = 1
    MAPGEN_PROC_ROOMSIZE_MAXY = 13
    MAPGEN_PROC_ROOMSIZE_MINX = 1
    MAPGEN_PROC_ROOMSIZE_MAXX = 20
    MAPGEN_PROC_PATH_FREQ = 5 //higher == less frequent
    MAPGEN_PROC_LEFT = 0
    MAPGEN_PROC_RIGHT = 1
    MAPGEN_PROC_UP = 2
    MAPGEN_PROC_DOWN = 3
)

//player states 
const (
    PLAYERSTATE_NORMAL = iota
)

//debug settings
const (
    DEBUG_NOWALL = 0
    DEBUG_REVEAL_ALL = 0
    DEBUG_NOCLEAR = 0
)