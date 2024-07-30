package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andrew-1234/shimmy/mover"
	hook "github.com/robotn/gohook"
)

var movePropsMap map[string]mover.MoveProps

func main() {
	distance := flag.Int("p", 10, "Pixel distance to move the window")
	debug := flag.Bool("d", false, "Debug mode")
	flag.Parse()
	movePropsMap = initMovePropsMap(*distance)
	shimmy(*debug, *distance)
}

func shimmy(debug bool, distance int) {
	if debug {
		fmt.Println("Debug mode enabled, logging keypress events. Press q to exit the program")
		evChan := hook.Start()
		defer hook.End()

		for ev := range evChan {
			if ev.Rawcode == 12 {
				break
			}
			fmt.Println("hook: ", ev)
		}
		return
	}

	var CyanBright = "\033[36;1m"
	var Yellow = "\033[0;35m"
	var Reset = "\033[0m"
	var desc = `
	+-----------------------------+--------------------------+
	|             Key             |          Desc.           |
	+-----------------------------+--------------------------+
	| ctrl shift q                | exit                     |
	| ctrl cmd up/down/left/right | move window in direction |
	+-----------------------------+--------------------------+
	`
	fmt.Println("\n" + CyanBright + "Starting shimmy" + Reset)
	print("---------------\n\n")
	fmt.Println(Yellow+"        Pixel distance =", distance, Reset)
	fmt.Println(desc)
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(sch hook.Event) {
		fmt.Println("you pressed ctrl-shift-q, exiting...")
		hook.End()
	})
	registerMoveHooks()

	start := hook.Start()
	<-hook.Process(start)
}

func initMovePropsMap(distance int) map[string]mover.MoveProps {
	movePropsMap = make(map[string]mover.MoveProps)
	directions := []string{"right", "left", "up", "down"}

	for _, direction := range directions {
		props, err := mover.NewMoveProps(distance, direction)
		if err != nil {
			log.Fatalf("Error initializing MoveProps for direction %s: %v", direction, err)
		}
		movePropsMap[direction] = props
	}
	return movePropsMap
}

func registerMoveHooks() {
	hook.Register(hook.KeyDown, []string{"cmd", "ctrl", "right"}, func(sch hook.Event) {
		movePropsRight := movePropsMap["right"]
		mover.MoveWindow(&movePropsRight)
	})

	hook.Register(hook.KeyDown, []string{"cmd", "ctrl", "left"}, func(sch hook.Event) {
		movePropsLeft := movePropsMap["left"]
		mover.MoveWindow(&movePropsLeft)
	})

	hook.Register(hook.KeyDown, []string{"cmd", "ctrl", "up"}, func(sch hook.Event) {
		movePropsUp := movePropsMap["up"]
		mover.MoveWindow(&movePropsUp)
	})

	// My down arrow doesn't register a KeyDown event but using KeyHold works
	hook.Register(hook.KeyHold, []string{"cmd", "ctrl", "down"}, func(sch hook.Event) {
		movePropsDown := movePropsMap["down"]
		mover.MoveWindow(&movePropsDown)
	})

}

// start move mode with function tab
// stop move mode with function tab
// move window with hjkl
// Switching to an app that doesnt have a visible window and trying to move it
// crashes the program
