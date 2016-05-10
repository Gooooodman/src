package main

import (
	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"

	"muisc/mplayer/library"

	"muisc/mplayer/mp"
)

func handleLibCommands(lib *library.MusicManager, tokens []string) {

	if len(tokens) < 2 {

		fmt.Println(` 

      Enter following commands to control the player: 

      lib list -- View the existing music lib 

      lib add <name><artist><source><type> -- Add a music to the music lib 
      # 字符按一个空格分离
      # type 只支持 2种格式 MP3 WAV 都要大写
      #例如：lib add loveyou 伟哥 loveyou MP3
      #      lib add loveyou 伟哥 loveyou  WAV

      lib remove 序号 -- Remove the specified music from the lib 

      `)

		return

	}

	switch tokens[1] {

	case "list":
		fmt.Println("序号  MP3_id    名字        作者          路径                           类型")
		for i := 0; i < lib.Len(); i++ {

			e, _ := lib.Get(i)
			fmt.Printf("%-4d  %-8s  %-10s  %-12s  %-20s           %-5s\n", i+1, e.Id, e.Name, e.Artist, e.Source, e.Type)
			//fmt.Println(" ", i+1, ":", " ", e.Id, "   ", e.Name, "     ", e.Artist, "   ", e.Source, "   ", e.Type)

		}

	case "add":

		{

			if len(tokens) == 6 {

				id++

				lib.Add(&library.MusicEntry{strconv.Itoa(id),

					tokens[2], tokens[3], tokens[4], tokens[5]})

			} else {

				fmt.Println("USAGE: lib add <name><artist><source><type>")

			}

		}

	case "remove":

		if len(tokens) == 3 {

			index, _ := strconv.Atoi(tokens[2])
			//fmt.Println(index)
			lib.Remove(index)
			fmt.Println("序号  MP3_id    名字        作者          路径                           类型")
			for i := 0; i < lib.Len(); i++ {

				e, _ := lib.Get(i)

				fmt.Printf("%-4d  %-8s  %-10s  %-12s  %-20s           %-5s\n", i+1, e.Id, e.Name, e.Artist, e.Source, e.Type)

			}

		} else {

			fmt.Println("USAGE: lib remove <id>")

		}

	default:

		fmt.Println("Unrecognized lib command:", tokens[1])

	}

}

func handlePlayCommand(lib *library.MusicManager, tokens []string) {

	if len(tokens) != 2 {

		fmt.Println("USAGE: play <name>")

		return

	}

	e := lib.Find(tokens[1])

	if e == nil {

		fmt.Println("The music", tokens[1], "does not exist.")

		return

	}

	mp.Play(e.Source, e.Type)

}

//var lib *library.MusicManager

var id int = 0

func main() {
	//放在main中要传参
	var lib *library.MusicManager
	lib = library.NewMusicManager()
	//lib := library.NewMusicManager()
	fmt.Println(` 

      Enter following commands to control the player: 

      lib list -- View the existing music lib 
      #lib list  列出所有音乐
      lib add <name><artist><source><type> -- Add a music to the music lib 
      # 字符按一个空格分离
      # type 只支持 2种格式 MP3 WAV 都要大写
      #例如：lib add loveyou 伟哥 loveyou MP3
      #      lib add loveyou 伟哥 loveyou  WAV
      lib remove <序号> -- Remove the specified music from the lib 

      play <name> -- Play the specified music 
      #play loveyou 播放loveyou
      q | e  -- quit | exit 

 `)

	r := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter command-> ")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {

			break

		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {

			handleLibCommands(lib, tokens)

		} else if tokens[0] == "play" {

			handlePlayCommand(lib, tokens)

		} else {

			fmt.Println("Unrecognized command:", tokens[0])

		}

	}

}
