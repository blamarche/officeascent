/*

    Player.go - interface to handle player specific operations

*/
package player


import (
    "github.com/blamarche/ansiterm"
    "fmt"
    
    "../constants"
    "../world"
    "../item"
)


const INVENTORY_MAX = 26

const (
    CLASS_ENGINEER = iota   
    CLASS_OTHER
)

const (
    STATE_NORMAL = iota
    STATE_INVENTORY
    STATE_HELP
)


type Player struct {
    X int
    Y int
    
	Hp int
	Max_hp int
	
	Ambition int
	Charm int
	Spirit int
	Greed int
	
    LightRadius int
    
    Class int //classof player
    PlayerState int
    Inventory []*item.Item
	EquippedArmorSuit int 
	EquippedArmorTie int
	EquippedWeaponClose int
	EquippedWeaponRanged int
}


//NEW
//***
func NewPlayer(hp int, max_hp int, amb int, charm int, spirit int, greed int, class int) *Player {
    return NewPlayerXY(0,0, hp, max_hp, amb, charm, spirit, greed, class)
}

func NewPlayerXY(x int, y int, hp int, max_hp int, amb int, charm int, spirit int, greed int, class int) *Player {
    
    e := Player{
        x, 
        y,  
        hp, 
        max_hp, 
        amb, 
        charm, 
        spirit, 
        greed, 
        constants.START_RADIUS+spirit, 
        class, 
        constants.PLAYERSTATE_NORMAL,
        make([]*item.Item, INVENTORY_MAX),
		-1,
		-1,
		-1,
		-1,
    }
    
    return &e
}


//INVENTORY 
//*********
func (e *Player) AddInventoryItem(it *item.Item) int {
    for i:=0; i<INVENTORY_MAX; i++ {
        if e.Inventory[i]==nil {
            e.Inventory[i] = it
            return i
        }
    }
    return -1
}

func (e *Player) StringToInvIndex(c string) int {
    code := int(c[0])
    index := code - 97
    
    if index<0 || index>=INVENTORY_MAX {
        return -1
    } 
    
    return index 
}

func (e *Player) UnequipItem(index int) {
	if index==e.EquippedArmorSuit {
		e.EquippedArmorSuit = -1
	}
	if index==e.EquippedArmorTie {
		e.EquippedArmorTie = -1
	}
	if index==e.EquippedWeaponClose {
		e.EquippedWeaponClose = -1
	}
	if index==e.EquippedWeaponRanged {
		e.EquippedWeaponRanged = -1
	}
}

func (e *Player) RemoveInventoryItem(index int) *item.Item {
    if index<INVENTORY_MAX && index>=0 {
        e.UnequipItem(index)
		
		tmp := e.Inventory[index]
        e.Inventory[index] = nil
        return tmp
    }
    return nil
}

func (e *Player) UseItem(index int) string {
	//todo: potions 
	
	if index<0||index>=INVENTORY_MAX {
		return "Invalid item selected"
	}
	
	switch e.Inventory[index].Kind {
		case item.KIND_WEAPON_BOTH, item.KIND_WEAPON_PROMOTE, item.KIND_WEAPON_DEMOTE:
			if e.Inventory[index].Subkind == item.SUBKIND_WEAPON_RANGED {
				e.EquippedWeaponRanged = index
			} else {
				e.EquippedWeaponClose = index
			}
			return "You equipped the weapon "+e.Inventory[index].Name
			
		case item.KIND_ARMOR_SUIT:
			e.EquippedArmorSuit = index 
			return "You equipped the "+e.Inventory[index].Name
		
		case item.KIND_ARMOR_TIE:
			e.EquippedArmorTie = index 
			return "You equipped the "+e.Inventory[index].Name
	}
	return "[NOT IMPLEMENTED]"
}

func (e *Player) ListInventory(mesg string) {
    ansiterm.SetFGColor(7)
    
    ansiterm.MoveToXY(2,3)
    fmt.Print(mesg) //key int(s) == 97 for 'a', 'A'==65
    
    for i:=0; i<INVENTORY_MAX/2; i++ {
        ansiterm.MoveToXY(2,i+5)
        if e.Inventory[i]!=nil {
            if i==e.EquippedArmorSuit || i==e.EquippedArmorTie || i==e.EquippedWeaponClose || i==e.EquippedWeaponRanged {
				fmt.Printf("*%s: %s", string(97+i), e.Inventory[i].Name)
			} else {
				fmt.Printf(" %s: %s", string(97+i), e.Inventory[i].Name)
			}
        } else {
            fmt.Print("                           ")
        }
    }
    
    for i:=INVENTORY_MAX/2; i<INVENTORY_MAX; i++ {
        ansiterm.MoveToXY(31,i+5-INVENTORY_MAX/2)
        if e.Inventory[i]!=nil {
            if i==e.EquippedArmorSuit || i==e.EquippedArmorTie || i==e.EquippedWeaponClose || i==e.EquippedWeaponRanged {
				fmt.Printf("*%s: %s", string(97+i), e.Inventory[i].Name)
			} else {
				fmt.Printf(" %s: %s", string(97+i), e.Inventory[i].Name)
			}
        } else {
            fmt.Print("                           ")
        }
    }
}


//POSITION
//********
func (e *Player) GetXY() (int,int) {
    return e.X, e.Y
}

//currently ignores IsBlocked state
func (e *Player) SetXY(x,y int) {
    e.X=x
    e.Y=y
}

//MOVEMENT
//********
func (e *Player) MoveUp() bool {
    e.Y--
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.Y++
        return false
    }
    return true
}

func (e *Player) MoveUpLeft() bool {
    e.Y--
    e.X--
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.Y++
        e.X++
        return false
    }
    return true
}

func (e *Player) MoveUpRight() bool {
    e.Y--
    e.X++
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.Y++
        e.X--
        return false
    }
    return true
}

func (e *Player) MoveDownLeft() bool {
    e.Y++
    e.X--
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.Y--
        e.X++
        return false
    }
    return true
}

func (e *Player) MoveDownRight() bool {
    e.Y++
    e.X++
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.Y--
        e.X--
        return false
    }
    return true
}

func (e *Player) MoveDown() bool {
    e.Y++
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.Y--
        return false
    }
    return true
}

func (e *Player) MoveLeft() bool {
    e.X--
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.X++
        return false
    }
    return true 
}

func (e *Player) MoveRight() bool {
    e.X++
    if (world.CurrentMap.IsBlocked(e.X, e.Y)) {
        e.X--
        return false
    }
    return true
}

//DISPLAY 
//*******
func (e *Player) Display(wx, wy int) {
    ansiterm.SetFGColor(constants.COLOR_PLAYER)
    ansiterm.MoveToXY(wx/2,wy/2)
    fmt.Print(constants.RUNE_PLAYER)
}

