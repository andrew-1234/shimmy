package mover

import (
	"reflect"
	"testing"
)

func TestNewMoveProps(t *testing.T) {
	type args struct {
		distance  int
		direction string
	}
	tests := []struct {
		name    string
		args    args
		want    MoveProps
		wantErr bool
	}{
		{
			name: "Test NewMoveProps with valid direction",
			args: args{
				distance:  10,
				direction: "right",
			},
			want: MoveProps{
				X:         10,
				Y:         0,
				Direction: "right",
			},
			wantErr: false,
		},
		{
			name: "Test NewMoveProps with invalid direction",
			args: args{
				distance:  10,
				direction: "invalid",
			},
			want:    MoveProps{},
			wantErr: true,
		},
		{
			name: "Test NewMoveProps with invalid distance",
			args: args{
				distance:  4001,
				direction: "right",
			},
			want:    MoveProps{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMoveProps(tt.args.distance, tt.args.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMoveProps() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMoveProps() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_calculateOffset(t *testing.T) {
	// Define the test cases in a table
	tests := []struct {
		name      string
		distance  int
		direction string
		wantX     int
		wantY     int
		wantErr   bool
	}{
		{
			name:      "right direction",
			distance:  10,
			direction: "right",
			wantX:     10,
			wantY:     0,
			wantErr:   false,
		},
		{
			name:      "left direction",
			distance:  10,
			direction: "left",
			wantX:     -10,
			wantY:     0,
			wantErr:   false,
		},
		{
			name:      "up direction",
			distance:  10,
			direction: "up",
			wantX:     0,
			wantY:     -10,
			wantErr:   false,
		},
		{
			name:      "down direction",
			distance:  10,
			direction: "down",
			wantX:     0,
			wantY:     10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, err := calculateOffset(tt.distance, tt.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("calculateOffset() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("calculateOffset() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
