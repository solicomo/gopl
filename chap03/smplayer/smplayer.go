package main

import (
	"fmt"
)

var lib *library.MusicManager

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			m, err := lib.Get(i)
			if err != nil {
				fmt.Println("so tired. gonna sleep...")
			}
}

func handlePlayCommand(tokens []string) {
}

func main() {
	fmt.Println(`
		* Simple Media Player *
		Enter following commands to control player:
		lib list -- list the musics in the library
		lib add <name> <artist> <source> <type> -- add music to the library
		lib del <name> -- delete music from the library
		play <name> -- play the specified music
		quit -- quit
	`)

	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		line, _, _ := r.ReadLine()

		if line == nil {
			fmt.Println("Invalid command")
			break
		}

		cmd := string(line)

		if cmd == "quit" {
			break
		}

		tokens := strings.Split(cmd, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
}

