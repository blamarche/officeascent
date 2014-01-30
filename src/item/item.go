/*

    item.go - "class" to handle item specific operations and collection

*/
package item


import (
    "../constants"
)


const (
    KIND_WEAPON = iota   
    KIND_POTION
    KIND_COLLECTIBLE //non-functional paychecks, etc
)

const (
	SUBKIND_NONE = iota
	SUBKIND_POTION_GOOD
	SUBKIND_POTION_BAD
	SUBKIND_WEAPON_RANGED
)

//Item struct
//if you add anything here don't forget to update Clone()!
type Item struct { 
    Name string
	Use_desc string
    Rune string
	Floor_min int
	Floor_max int
    Max_per_floor float32 //per 10000 squares (ie 100x100 room)
	Fgcolor int
    Kind int //type of item (weapon, potion, collectible, etc)
	Subkind int //subtype of item    
}

//TODO: item runes should be randomized?

//Item runes will be lowercase, enemies uppercase.
var ItemList []Item = []Item{
	Item{
		"executive report",
		"present an",
		"e",
		1, //min floor
		20, //max floor
        10, //frequency
		3, //fgcolor
		KIND_WEAPON,
		SUBKIND_NONE,
	},
	Item{
		"bonus check",
		"cash a",
		"b",
		1, //min floor
		20, //max floor
        5, //frequency
		6, //fgcolor
		KIND_POTION,
		SUBKIND_POTION_GOOD,
	},
    Item{
		"stapler",
		"mess up printouts with a",
		"s",
		2, //min floor
		4, //max floor
        50, //frequency
		2, //fgcolor
		KIND_WEAPON,
		SUBKIND_WEAPON_RANGED,
	},
	Item{
		"fake wall",
		"crumble a",
		"#",
		2, //min floor
		999, //max floor
        4, //frequency
		constants.COLOR_WALL, //fgcolor
		KIND_COLLECTIBLE,
		SUBKIND_NONE,
	},
}


//FUNCTIONS
func (i *Item) Clone() *Item {
    //not exactly memory efficient :)
    copy := Item{
        i.Name,
        i.Use_desc, 
        i.Rune ,
        i.Floor_min ,
        i.Floor_max ,
        i.Max_per_floor, 
        i.Fgcolor ,
        i.Kind ,
        i.Subkind,  
    }
    
    return &copy
}




