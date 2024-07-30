# shimmy

![Build](https://github.com/andrew-1234/shimmy/actions/workflows/go.yml/badge.svg)

Shimmy provides directional window position adjustments using
keyboard shortcuts on macOS.

https://github.com/user-attachments/assets/b65d4798-456e-42fc-a999-f6292e7d722b

This isn't a complete window management tool, but rather a way to simply nudge
windows in a direction without having to use the mouse. Best used in conjunction
with a window management tool or native macOS window arrangement keyboard
shortcuts.

## Requirements

- macOS
- Go 1.22.1
- A C compiler (to build)

## Installation

1. git clone the repository
2. Build the binary with `go build`

## Usage

Running the binary in the terminal activates the following global shortcuts:

- `Cmd + Shift + Q` - Quit the application
- `Cmd + Shift + Up` - Move active window up
- `Cmd + Shift + Down` - Move active window down
- `Cmd + Shift + Left` - Move active window left
- `Cmd + Shift + Right` - Move active window right

Run in the foreground:

```bash
# Set movement distance to 10 pixels (default)
./shimmy -p 10
```

Run in the background:

```bash
nohup ./shimmy -p 10 &
```

## Future development

- [ ] Add support for custom keybinds
- [ ] Run in the background as a menu bar app

## Related Projects

- [Rectangle](https://github.com/rxhanson/Rectangle/) - A window management app for macOS.
- [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform system automation.
