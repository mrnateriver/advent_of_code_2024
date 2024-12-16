package shared

import (
	"fmt"
	"os"
	"sync"
)

const (
	ColorRed = 31 + iota
	ColorGreen
	ColorYellow
	ColorBlue
	ColorPurple
	ColorCyan
	ColorGray
	ColorWhite = 97
)

type ConsoleColor int

var printLock sync.Mutex
var printDone bool

func PrintGrid[T any](grid [][]T) {
	printLock.Lock()
	defer printLock.Unlock()

	if printDone {
		setCursorColumn(1)
		moveCursorUp(len(grid))
	}

	for _, row := range grid {
		for _, cell := range row {
			print(cell)
		}
		println()
	}

	printDone = true
}

func Colored(c ConsoleColor, s string) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", c, s)
}

func AwaitInput() {
	var b = make([]byte, 1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		panic(fmt.Errorf("failed to read input: %w", err))
	}

	moveCursorUp(1)
}

func moveCursorUp(c int) {
	fmt.Printf("\033[%dA", c)
}

func setCursorColumn(c int) {
	fmt.Printf("\033[%dG", c)
}