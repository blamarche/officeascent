/*

    Ai.go - "class" to handle AIs

*/
package ai


import (
    //"../constants"
)


const (
    KIND_PASSIVE = iota   
    KIND_AGGRESSIVE
    KIND_FRIENDLY
)

const (
	WEAK_PROMOTE = iota 
	WEAK_DEMOTE
)

const (
	PATTERN_RANDOM = iota
	PATTERN_WANDER
)

const (
	STATE_NORMAL = iota
	STATE_ATTACK
	STATE_RUN
)

//Item struct
//if you add anything here don't forget to update Clone()!
type AI struct { 
    Name string
	Hit_desc string
    Rune string
	Floor_min int
	Floor_max int
    Max_per_floor float32 //per 10000 squares (ie 100x100 room)
	Fgcolor int
    Kind int //type of ai (passive, aggresive, etc)
	Pattern int //move pattern
	WeakAgainst int //WEAK_
	Attention int
	Ambition int
	Charm int
	Spirit int
	Greed int
	State int // STATE_
	Dirty bool //used by step logic
}

//TODO: ai runes should be randomized

//Item runes will be lowercase, enemies uppercase.
var AIList []AI = []AI{
	AI{
		"office secretary",
		"puts you on hold",
		"S",
		1,
		999,
		20, //per 10000 squares (ie 100x100 room)
		3,
		KIND_AGGRESSIVE,
		PATTERN_RANDOM,
		WEAK_PROMOTE,
		15, //attention dist
		1, //amb
		10, //charm
		3, //spirit
		1, //greed
		STATE_NORMAL,
		false, //dirty
	},
}


//CLONE
func (a *AI) Clone() *AI {
    //not exactly memory efficient but should allow for interesting effects  :)
    copy := AI{
        a.Name ,
		a.Hit_desc ,
		a.Rune ,
		a.Floor_min ,
		a.Floor_max ,
		a.Max_per_floor ,
		a.Fgcolor ,
		a.Kind ,
		a.Pattern ,
		a.WeakAgainst ,
		a.Attention , 
		a.Ambition ,
		a.Charm ,
		a.Spirit ,
		a.Greed ,
		a.State ,
		a.Dirty, 
    }
    
    return &copy
}
