package main

import (
	"bufio"
	"fmt"
	"gobot/core"
	"os"
	"strings"
)

var userInputs string

func main() {
	intents := make(map[string]interface{})
	stories := make(map[string]interface{})

	//Create Intents
	intents["greets"] = []string{
		"Hello",
		"Hi",
		"Mambo",
		"Za asubuhi",
		"Za mchana",
		"Za usiku",
		"Hola",
	}

	intents["goodbye"] = []string{
		"bye",
		"goodbye",
		"see you later",
		"see you soon",
		"see you",
		"baadae",
		"bye",
		"kwaheri",
	}

	intents["cancel"] = []string{
		"sitisha",
		"acha",
		"cancel",
		"sitaki",
	}

	intents["order_pizza"] = []string{
		"Nataka pizza",
		"I need pizza",
		"I want Pizza",
	}

	stories["greets"] = map[string]interface{}{
		"message": "Helo, Karibu nikusaidiaje?",
		"choices": nil,
		"next":    nil, //nil means end
	}

	stories["goodbye"] = map[string]interface{}{
		"message": "Karibu tena",
		"choices": nil,
		"next":    nil,
	}

	stories["cancel"] = map[string]interface{}{
		"message": "Karibu tena",
		"choices": nil,
		"next":    nil,
	}

	stories["order_pizza"] = map[string]interface{}{
		"message": "Unataka pizza gani?",
		"order_pizza_choices": []string{
			"1. Chicken Pizza",
			"2. Cheese Pizza",
			"3. Mixture Pizza",
			"4. Skyline Pizza",
		},
		"next":            "soda",
		"choice_fallback": "Sorry, hatuna aina hio ya pizza!",
	}

	stories["soda"] = map[string]interface{}{
		"message": "Unakata kinywaji gani?",
		"soda_choices": []string{
			"1. Pepsi",
			"2. Cocacola",
			"3. Sprite",
			"4. Mirinda",
		},
		"choice_fallback": "Sorry, hatuna aina hio ya soda!",
	}

	stories["fallback"] = map[string]interface{}{
		"message": "Sijaelewa unataka nini?",
		"choices": nil,
		"next":    nil,
	}

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
