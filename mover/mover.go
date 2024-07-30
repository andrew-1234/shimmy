package mover

/*
#cgo LDFLAGS: -framework ApplicationServices
#include "mover.h"
*/
import "C"
import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// -L${SRCDIR} ${SRCDIR}/mover.o
type MoveProps struct {
	X         int
	Y         int
	Direction string
}

func NewMoveProps(distance int, direction string) (MoveProps, error) {
	if err := validateDirection(direction); err != nil {
		return MoveProps{}, fmt.Errorf("invalid direction: %v", err)
	}

	x, y, err := calculateOffset(distance, direction)
	if err != nil {
		return MoveProps{}, fmt.Errorf("invalid distance: %v", err)
	}

	return MoveProps{x, y, direction}, nil
}

func MoveWindow(props *MoveProps) {
	// Get the active window
	// Returns MData struct with field AxID: contains AXUIElementRef
	activeWindow := robotgo.GetActive()
	windowRef := activeWindow.AxID
	// check if the window is valid
	if windowRef == 0 {
		fmt.Println("Error: could not get active window")
		return
	}
	callMoveWindow(C.AXUIElementRef(windowRef), props.X, props.Y)
}

func callMoveWindow(window C.AXUIElementRef, x, y int) {
	C.move_window(window, C.int(x), C.int(y))
}

func validateDirection(direction string) error {
	switch direction {
	case "right", "left", "up", "down":
		return nil
	default:
		return fmt.Errorf("invalid direction <%s>: must be one of: <left right up down>", direction)
	}
}

func calculateOffset(distance int, direction string) (int, int, error) {
	if distance > 4000 {
		return 0, 0, fmt.Errorf("distance <%d> is too large, must be less than <4000>", distance)
	}
	switch direction {
	case "right":
		return distance, 0, nil
	case "left":
		return -distance, 0, nil
	case "up":
		return 0, -distance, nil
	case "down":
		return 0, distance, nil
	default:
		return 0, 0, nil
	}
}
