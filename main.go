package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

// generates rgb function
func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)

}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("This command is intended to work with pipes")
		fmt.Println("fortune | gorainbow")
	}

	reader := bufio.NewReader(os.Stdin)
	j := 0
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)
		j++
	}

	print(output)
}
