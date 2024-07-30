package main

import (
	"reflect"
	"testing"

	"github.com/andrew-1234/shimmy/mover"
)

func Test_initMovePropsMap(t *testing.T) {
	type args struct {
		distance int
	}
	tests := []struct {
		name string
		args args
		want map[string]mover.MoveProps
	}{
		// TODO: Add test cases.
		{
			name: "Test initMovePropsMap",
			args: args{
				distance: 10,
			},
			want: map[string]mover.MoveProps{
				"right": {
					X:         10,
					Y:         0,
					Direction: "right",
				},
				"left": {
					X:         -10,
					Y:         0,
					Direction: "left",
				},
				"up": {
					X:         0,
					Y:         -10,
					Direction: "up",
				},
				"down": {
					X:         0,
					Y:         10,
					Direction: "down",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initMovePropsMap(tt.args.distance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initMovePropsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registerMoveHooks(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerMoveHooks()
		})
	}
}
