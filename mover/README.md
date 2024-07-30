# Mover

The `mover` package provides an API for moving windows on the screen. It exposes
the `MoveWindow` function, which is a Cgo wrapper around the
C function `move_window` defined in `mover.c`.

## API

### `MoveWindow(props MoveProps) error`

Moves a window on the screen. The movement is defined by the `MoveProps` parameter, which specifies the distance and direction of the move.

### `MoveProps`

A struct that defines the properties of a window move. It includes the following fields:

- `x`: The distance to move the window along the x-axis.
- `y`: The distance to move the window along the y-axis.
- `direction`: The direction to move the window. Valid values are "up", "down", "left", and "right".

### `NewMoveProps(distance int, direction string) (MoveProps, error)`

A helper function that creates a new `MoveProps` struct with the specified distance and
direction.

## Usage

Here's a basic example of how to use the `MoveWindow` function:

```go
props := mover.NewMoveProps(100, "up")
err := mover.MoveWindow(props)
if err != nil {
    fmt.Println("Error moving window:", err)
}
```
