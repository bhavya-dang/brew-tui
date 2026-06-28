# pkgui

A TUI to manage multiple package managers and packages installed by them.

> Written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea)

## Features

- List installed Homebrew formulae
- Fuzzy search installed packages (`/` to search)
- View formula details: version, description, homepage, license, dependencies
- Scrollable package list with keyboard navigation

## Prerequisites

- [Homebrew](https://brew.sh)
- Go 1.25+ (if building from source)

## Installation

### Using Go

```bash
go install github.com/bhavyadang/pkgui@latest
```

### Using install.sh

```bash
curl -sSL https://raw.githubusercontent.com/bhavyadang/pkgui/main/install.sh | sh
```

### Using Makefile

```bash
git clone https://github.com/bhavyadang/pkgui.git
cd pkgui
make install
```

### Manual

```bash
git clone https://github.com/bhavyadang/pkgui.git
cd pkgui
go build -o build/pkgui .
cp build/pkgui "$GOPATH/bin/pkgui"
```

## Usage

```bash
pkgui
```

### Keybindings

| Key | Action |
|-----|--------|
| `↑` / `↓` | Navigate package list |
| `/` | Start search (type to filter) |
| `Esc` | Exit search |
| `Enter` | Confirm search |
| `q` / `Ctrl+C` | Quit |

## Support

- Homebrew
  - formulae (with detail view from the Homebrew API)

## Preview

![Screenshot 1](./preview/1.png)

## License

MIT
