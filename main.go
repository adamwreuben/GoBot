package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adamwreuben/GoBot/core"
)

var userInputs string

func main() {
	intents := make(map[string]interface{})
	stories := make(map[string]interface{})

	gobot := core.NewGoBot(intents, stories)
	playground(*gobot)

	// key, response := Chat(*gobot, "hello")
	// fmt.Println(key, response)

}

func playground(goBot core.GoBot) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if userInputs == "" {
			fmt.Print("Start chatting:\n")
		}
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			// fmt.Println(text)
			userInputs = text
			key, response := Chat(goBot, text)
			fmt.Println(key, response)
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func Chat(gobot core.GoBot, message string) (string, string) {

	key := gobot.FindMessageKey(message)
	goBotStories := gobot.Story.(map[string]interface{})

	if key != "" {
		storyObj := goBotStories[key]
		if key != "cancel" {
			if storyObj != nil {
				story := storyObj.(map[string]interface{})

				//Becase choice depend from next
				// fmt.Println("Key: real ", key)
				// fmt.Println(story["message"])

				return key, story["message"].(string)

			} else {
				// fmt.Println("Key when no interface: ", key)
				if strings.Contains(key, "choices") {
					key := key[:len(key)-8]
					story := goBotStories[key].(map[string]interface{})
					next := story["next"]

					if next != nil {
						story := goBotStories[next.(string)].(map[string]interface{})
						// fmt.Println(story["message"])
						return key, story["message"].(string)

					} else { // when next is Nil
						// fmt.Println("Asante sana")
						return key, "Asante sana"

					}

				} else {
					fmt.Println("interface is nil")
					return "", ""

				}
			}
		} else {
			// fmt.Println("Key: ", key)
			// fmt.Println("Karibu tena")
			return key, "Karibu tena"
		}

	} else {
		// fmt.Println("Key: ", key)
		fallbackObject := goBotStories["fallback"].(map[string]interface{})
		message := fallbackObject["message"]
		// fmt.Println(message)
		return key, message.(string)
	}

}
