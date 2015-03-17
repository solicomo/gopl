package player

import "testing"

func TestMP3Player(t *testing.T) {
	Play("~/Music/Take Five.mp3", "mp3")
	Play("~/Music/So What.wav", "wav")
}
