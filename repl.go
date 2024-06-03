package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
  
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("pokedex > ")

    // Wait for input
    scanner.Scan()

    // Read input
    cleanedInput := cleanInput(scanner.Text())
	if len(cleanedInput) == 0 {
		continue
	}

	cmdFromUser := cleanedInput[0]

	args := []string{}
	if len(cleanedInput) > 1 {
		args = cleanedInput[1:]
	}
	availableCommands := getCommands()

	cmdToExecute, ok := availableCommands[cmdFromUser]
	if ok {
		err := cmdToExecute.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
		continue
	} else {
		fmt.Println("Unknown command")
		continue
	}
  }
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}	

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "A command for when you don't know wtf to do.",
			callback: callbackHelp,
		},
		"map": {
			name: "map",
			description: "A command for when you want to view the locations.",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "A command for when you want to go back one page.",
			callback: callbackMapb,
		},
		"explore": {
			name: "explore",
			description: "A command for exploring shit.",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch",
			description: "A command for catching shit.",
			callback: callbackCatch,
		},
		"inspect": {
			name: "inspect",
			description: "A command for inspecting shit.",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "A based command for being a pokebitch.",
			callback: callbackPokedex,
		},
		"exit": {
			name: "exit",
			description: "A command for when you want to gtfo of this repl.",
			callback: callbackExit,
		},
	}
}