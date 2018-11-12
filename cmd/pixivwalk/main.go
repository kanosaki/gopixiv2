package main

import (
	"encoding/json"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/sirupsen/logrus"
	"go-pixiv"
	"go-pixiv/api"
	"log"
	"os"
	"strings"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	auth, err := loadConfig(os.ExpandEnv("$HOME/.config/pixiv"))
	if err != nil {
		panic(err)
	}
	client, err := pixiv.NewOAuthClient(auth)
	if err != nil {
		log.Fatal(err)
	}
	a := &app{
		client: client,
		api:    api.New(client),
		commands: map[string]*command{
			"get": {fn: cmdGet},
		},
	}
	args := os.Args
	if len(args) > 1 {
		a.executor(strings.Join(args[1:], " "))
	} else {
		app := prompt.New(a.executor, a.completer)
		app.Run()
	}
}

func loadConfig(path string) (*pixiv.AuthConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	ret := &pixiv.AuthConfig{}
	if err := json.NewDecoder(f).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type app struct {
	client   pixiv.Session
	api      *api.API
	commands map[string]*command
}

func (a *app) executor(cmdLine string) {
	blocks := strings.Split(cmdLine, " ")
	if len(blocks) == 0 {
		return
	}
	cmdName := blocks[0]
	cmd, ok := a.commands[cmdName]
	if !ok {
		fmt.Printf("Command not found: %s", cmdName)
	} else {
		if err := cmd.fn(a, blocks[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func (a *app) completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "get", Description: "do get"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
