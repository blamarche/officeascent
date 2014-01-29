/*

    Constants.go - util functions

*/
package constants

//*********************************************
//*********************************************

const VERSION = "0.1 Alpha"

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