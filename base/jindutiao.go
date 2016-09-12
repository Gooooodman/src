package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Bar(vl int, width int) string {
	return fmt.Sprintf("%s%*c", strings.Repeat("█", vl/10), vl/10-width+1, ([]rune(" ▏▎▍▌▋▋▊▉█"))[vl%10])
}
func main() {
	for i := 0; i <= 100; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("\r%s%d%%", Bar(i, 10), i)
		os.Stdout.Sync()

	}
}
