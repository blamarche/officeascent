/*

    interfaces.go - interfaces to pass around

*/
package interfaces


import (
    
)

type IPlayer interface {
	Display(int, int)
	GetXY() (int, int)
}
