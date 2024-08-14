# shimmy

![Build](https://github.com/andrew-1234/shimmy/actions/workflows/go.yml/badge.svg)

Shimmy is a lightweight utility designed for macOS users who need quick and
precise control over their window positions. It allows users to nudge windows in
any direction using customisable global keyboard shortcuts. Unlike full-fledged
window management tools, Shimmy focuses on providing simple directional
adjustments, making it a perfect companion to other window management solutions.


https://github.com/user-attachments/assets/b65d4798-456e-42fc-a999-f6292e7d722b


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
