/*

    Cursor.go - a movable cursor

*/
package cursor

import (
    "github.com/blamarche/ansiterm"
    "fmt"
)

type Cursor struct {
	X int 
	Y int
	Bgcolor int
	Fgcolor int
	Rune string
	leftlim int
	rightlim int
	uplim int
	downlim int
	offsetX int
	offsetY int
}

//NEW
func NewCursor(x, y, leftlim, rightlim, uplim, downlim, offsetX, offsetY int) *Cursor {
	c := Cursor{x, y, 2, 0, "+", leftlim, rightlim, uplim, downlim, offsetX, offsetY}

	return &c
}

//POSITION
func (c *Cursor) SetXYWithOffset(x, y int) {
	c.X = x + c.offsetX
	c.Y = y + c.offsetY
}

//these are probably somewhat wrong overall, but work for our current setup
func (c *Cursor) GetXYWithOffset() []int {
	return []int{c.X - c.offsetX, c.Y - c.offsetY}
}

func (c *Cursor) GetMapXY(px, py int) []int {
	return []int{c.X - c.offsetX - ((c.rightlim - c.leftlim)/2) + px, c.Y - c.offsetY - ((c.downlim - c.uplim)/2) + py}
}

//MOVEMENT
func (c *Cursor) MoveUp() {
	c.Y--
	if c.Y < c.uplim {
		c.Y++
	}
}

func (c *Cursor) MoveDown() {
	c.Y++
	if c.Y > c.downlim {
		c.Y--
	}
}

func (c *Cursor) MoveLeft() {
	c.X--
	if c.X < c.leftlim {
		c.X++
	}
}

func (c *Cursor) MoveRight() {
	c.X++
	if c.X > c.rightlim {
		c.X--
	}
}

func (c *Cursor) MoveUpRight() {
	c.X++
	c.Y--
	if c.X > c.rightlim || c.Y < c.uplim {
		c.X--
		c.Y++
	}
}

func (c *Cursor) MoveUpLeft() {
	c.X--
	c.Y--
	if c.X < c.leftlim || c.Y < c.uplim {
		c.X++
		c.Y++
	}
}

func (c *Cursor) MoveDownRight() {
	c.X++
	c.Y++
	if c.X > c.rightlim || c.Y > c.downlim {
		c.X--
		c.Y--
	}
}

func (c *Cursor) MoveDownLeft() {
	c.X--
	c.Y++
	if c.X < c.leftlim || c.Y > c.downlim {
		c.X++
		c.Y--
	}
}

//OTHER
func (c *Cursor) Display() {
	ansiterm.SetBGColor(c.Bgcolor)
	ansiterm.SetFGColor(c.Fgcolor)
	
    ansiterm.MoveToXY(c.X, c.Y)
    fmt.Print(c.Rune)
    ansiterm.SetBGColor(0)
}