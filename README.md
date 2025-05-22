# Conway's Game of Life (64-bit Integer Space)

This is an implementation of [Conway’s Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) written in Go, designed to handle coordinates anywhere in the signed 64-bit integer space.

Live cells are read from a file in the `Life 1.06` format.

---

## 📦 Features

- Handles massive coordinate spaces (signed 64-bit integers).
- Parses input in `#Life 1.06` format.
- Terminal-aware rendering.
- Optional terminal clearing for animation.
- Adjustable update rate and generation count via command-line flags.

---

## 🧱 Project Structure

```
conway-game-of-life/
├── main.go               # Entry point
├── go.mod
├── life/
│   ├── board.go          # Next generation logic
│   └── loader.go         # Parser for Life 1.06 input format
├── render/
│   ├── factory.go        # Chooses renderer (terminal vs standard)
│   ├── terminal.go       # Renderer for terminal (centered on screen)
│   └── standard.go       # Renderer for non-terminal output
└── types/
    └── types.go          # Point and Board definitions
```

---

## 🚀 How to Run

### 1. Clone the repository

```bash
git clone https://github.com/jonasssilveira/gameoflife.git
cd conway-game-of-life
```

### 2. Build or run directly with Go

```bash
go run main.go
```

---

## ⚙️ Command-Line Flags

| Flag      | Description                                      | Default            |
|-----------|--------------------------------------------------|--------------------|
| `-file`   | Path to the Life 1.06 input file                 | `life_106.txt`     |
| `-gens`   | Number of generations to simulate                | `10`               |
| `-clear`  | Whether to clear the terminal each generation    | `false`            |
| `-rate`   | Delay between generations (milliseconds)         | `500`              |

### Example:

```bash
go run main.go -file=sample.lif -gens=20 -clear=true -rate=100
```

---

## 📄 Life 1.06 Format

```
#Life 1.06
0 1
1 2
2 0
2 1
2 2
```

Each line represents a live cell using two integers for X and Y coordinates.

---

## 🛠 Requirements

- Go 1.20 or newer
- Terminal supporting ANSI escape sequences (for `-clear=true`)

---

## 📋 License

MIT License