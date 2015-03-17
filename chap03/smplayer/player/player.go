package player

import "fmt"

type Player interface {
	Play()
}

func Play(source, mtype string) {
	var player Player

	switch mtype {
	case "mp3":
		player = &MP3Player{source, 0}
	case "wav":
		player = &WAVPlayer{source, 0}
	default:
		fmt.Println(mtype, "is not spported.")
		return
	}

	player.Play()
}
