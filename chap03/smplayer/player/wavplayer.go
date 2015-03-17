package player

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	name    string
	process int
}

func (p *WAVPlayer) Play() {
	fmt.Println("playing wav:", p.name)

	for p.process = 0; p.process < 10; p.process++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("+")
	}

	fmt.Println("\nfinished.")
}
