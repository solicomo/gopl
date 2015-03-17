package player

import (
	"fmt"
	"time"
)

type MP3Player struct {
	name    string
	process int
}

func (p *MP3Player) Play() {
	fmt.Println("playing mp3:", p.name)

	for p.process = 0; p.process < 10; p.process++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}

	fmt.Println("\nfinished.")
}
