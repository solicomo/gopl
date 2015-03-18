package main

import (
	"bufio"
	"fmt"
	"gopl/chap03/smplayer/library"
	"gopl/chap03/smplayer/player"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1

func handleLibCommands(tokens []string) {
	if len(tokens) < 2 {
		fmt.Println("spported commands: list, add, del.")
		return
	}

	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			m, err := lib.Get(i)
			if m == nil || err != nil {
				fmt.Println("get music info failed:", err)
			}

			fmt.Println(i+1, ":", m.Name, m.Artist, len(m.Type), m.Type, m.Source)
		}
	case "add":
		if len(tokens) != 6 {
			fmt.Println("usage: lib add <name> <aritist> <type> <source>")
			break
		}

		id++
		lib.Add(&library.Music{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
	case "del":
		if len(tokens) != 3 {
			fmt.Println("usage: lib del <id>")
			break
		}

		idx, err := strconv.Atoi(tokens[2])
		if err != nil {
			fmt.Println("invalid argument:", err)
			break
		}

		m, err := lib.Del(idx-1)
		if err != nil {
			fmt.Println("delete failed:", err)
			break
		}

		fmt.Println(m.Name, "is deleted.")
	default:
		fmt.Println("unrecognized lib command:", tokens[1])
	}

}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("usage: play <name>")
		return
	}

	m := lib.Find(tokens[1])
	if m == nil {
		fmt.Println("no such music")
		return
	}

	player.Play(m.Source, m.Type)
}

func main() {
	fmt.Println(`
		* Simple Media Player *
		Enter following commands to control player:
		lib list -- list the musics in the library
		lib add <name> <artist> <type> <source> -- add music to the library
		lib del <id> -- delete music from the library
		play <name> -- play the specified music
		quit -- quit
	`)

	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command -> ")

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
}
