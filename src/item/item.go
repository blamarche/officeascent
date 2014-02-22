/*

    item.go - "class" to handle item specific operations and collection

*/
package item


import (
    "../constants"
)


const (
    KIND_WEAPON_BOTH = iota   
    KIND_WEAPON_PROMOTE
    KIND_WEAPON_DEMOTE
	KIND_ARMOR_SUIT
	KIND_ARMOR_TIE
    KIND_POTION
    KIND_COLLECTIBLE //non-functional paychecks, etc
)

const (
	SUBKIND_NONE = iota
	SUBKIND_POTION_GOOD
	SUBKIND_POTION_BAD
	SUBKIND_WEAPON_RANGED
	SUBKIND_ARMOR_PROMOTE
	SUBKIND_ARMOR_DEMOTE
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
	Power int
    Kind int //type of item (weapon, potion, collectible, etc)
	Subkind int //subtype of item    
}

//TODO: item runes should be randomized

//Item runes will be lowercase, enemies uppercase.
var ItemList []Item = []Item{
	Item{
		"executive report",
		"present an",
		"e",
		1, //min floor
		20, //max floor
        15, //frequency
		3, //fgcolor
		2, //power
		KIND_WEAPON_DEMOTE,
		SUBKIND_NONE,
	},
	Item{
		"bonus check",
		"cash a",
		"b",
		1, //min floor
		20, //max floor
        10, //frequency
		6, //fgcolor
		5, //power
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
		1, //power
		KIND_WEAPON_PROMOTE,
		SUBKIND_WEAPON_RANGED,
	},
	Item{
		"fake wall",
		"crumble a",
		"#",
		1, //min floor
		999, //max floor
        6, //frequency
		constants.COLOR_WALL, //fgcolor
		0, //power
		KIND_COLLECTIBLE,
		SUBKIND_NONE,
	},
	Item{
		"thrift store suit",
		"wear a",
		"r",
		1, //min floor
		999, //max floor
        50, //frequency
		3, //fgcolor
		1, //power
		KIND_ARMOR_SUIT,
		SUBKIND_ARMOR_PROMOTE,
	},
	Item{
		"coffee-stained tie",
		"wear a",
		"t",
		1, //min floor
		999, //max floor
        50, //frequency
		6, //fgcolor
		1, //power
		KIND_ARMOR_TIE,
		SUBKIND_ARMOR_DEMOTE,
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
		i.Power,
        i.Kind ,
        i.Subkind,  
    }
    
    return &copy
}




