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

## Features

- Parses `#Life 1.06` formatted files
- Handles cells anywhere in 64-bit coordinate space
- Supports two rendering modes:
    - **Terminal View** (centered at `0,0` based on terminal size)
    - **Standard Output View**
- CLI flags to control:
    - Input file
    - Number of generations
    - Refresh rate
    - Clear terminal between generations

---
## 🧠 How I Arrived at the Solution

To ensure performance and clarity, I used:

- A `map[Point]bool` to store only alive cells (sparse representation).
- Logic split across multiple packages (`board`, `render`, `main`) for separation of concerns.
- A factory method to choose the renderer dynamically based on output type (terminal vs non-terminal).
- CLI flags to allow runtime flexibility.

During development:
- I started by reading and parsing the `Life 1.06` format.
- Then, implemented the Game of Life rules.
- Added rendering logic to print in a centered terminal grid.
- Added flexibility with CLI parameters.

---

## ✅ Testing

I wrote unit tests focusing on input parsing edge cases and file handling. The tests use a **table-driven approach**.

### What’s tested:

- Valid file content is parsed correctly into a `Board`.
- Errors are properly returned when files are missing.
- Parsing logic ignores comments and malformed lines.

To run the tests:

```bash
go test ./...
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

## Thought Process & Testing

This solution was developed with scalability and clarity in mind. The input is parsed from a Life 1.06 format file into a sparse data structure (map of live cells). Since coordinates can be anywhere in the 64-bit space, using a matrix would be infeasible — so only live cells are tracked.

Rendering is handled in two ways depending on the output environment: if running in a terminal, the output is centered on the screen using terminal size detection; otherwise, a standard bounding box is printed.

To ensure correctness, I wrote unit tests for the following:
- The core game logic (`NextGen`) using a known oscillating pattern (blinker).
- The Life 1.06 parser, ensuring coordinates are read correctly and comments are ignored.

Tests are located in `board/board_test.go`. You can run them with:

```bash
go test ./...

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