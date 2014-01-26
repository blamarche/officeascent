/*

    item.go - "class" to handle item specific operations and collection

*/
package item


import (

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


type Item struct {
    name string
	use_desc string
    rune string
	floor_min int
	floor_max int
	fgcolor int
    kind int //type of item (weapon, potion, collectible, etc)
	subkind int //subtype of item
}


//Item runes will be lowercase, enemies uppercase.
//ItemList := map[string]Item{
var ItemList []Item = []Item{
	Item{
		"executive report",
		"present an",
		"e",
		0,
		20,
		3, //fgcolor
		KIND_WEAPON,
		SUBKIND_NONE,
	},
	Item{
		"bonus check",
		"cash a",
		"b",
		0,
		20,
		4, //fgcolor
		KIND_POTION,
		SUBKIND_POTION_GOOD,
	},
};
