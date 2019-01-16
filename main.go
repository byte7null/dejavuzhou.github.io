package main

import (
	"github.com/dejavuzhou/dejavuzhou.github.io/robot"
	"log"
	"time"
)

var gitCount = 1

func createCmds() []robot.Cmd {
	gitCount++
	gifConfig1 := []robot.Cmd{
		{"git", []string{"config", "--global", "user.email", "'dejavuzhou@qq.com'"}},
	}
	gifConfig2 := []robot.Cmd{
		{"git", []string{"config", "--global", "user.email", "'1413507308@qq.com'"}},
	}
	cmds := []robot.Cmd{
		{"git", []string{"config", "--global", "user.name", "'EricZhou'"}},
		{"git", []string{"stash"}},
		{"git", []string{"pull", "origin", "master"}},
		{"git", []string{"stash", "apply"}},
		{"git", []string{"add", "."}},
		{"git", []string{"status"}},
		{"git", []string{"commit", "-am", "hacknews-update" + time.Now().Format(time.RFC3339)}},
		{"git", []string{"status"}},
		{"git", []string{"push", "origin", "master"}},
		//{"netstat", []string{"-lntp"}},
		//{"free", []string{"-m"}},
		//{"ps", []string{"aux"}},
	}
	if gitCount%2 == 0 {
		cmds = append(gifConfig2, cmds...)
	} else {
		cmds = append(gifConfig1, cmds...)
	}
	return cmds
}

func main() {
	for {
		if err := robot.SpiderHackNews(); err != nil {
			log.Println(err)
		}
		if err := robot.SpiderHackShows(); err != nil {
			log.Println(err)
		}
		if err := robot.ParsemarkdownHacknews(); err != nil {
			log.Println(err)
		}
		_, err := robot.RunCmds(createCmds())
		if err != nil {
			log.Println(err)
		}
		time.Sleep(3 * time.Hour)
	}
}



