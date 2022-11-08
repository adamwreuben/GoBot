package main

import (
	"fmt"
	"gobot/core"
)

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

	//Create Stories(Dialogs) --> stories key must match intents key

	stories["greets"] = map[string]interface{}{
		"message": "Helo, Karibu nikusaidiaje?",
		"next":    nil, //nil means end
		"action":  nil,
	}

	stories["goodbye"] = map[string]interface{}{
		"message": "Karibu tena",
		"next":    nil,
		"action":  nil,
	}

	stories["cancel"] = map[string]interface{}{
		"message": "Karibu tena",
		"next":    nil,
		"action":  nil,
	}

	stories["order_pizza"] = map[string]interface{}{
		"message": "Unataka pizza gani?",
		"next":    nil,
		"action":  nil,
	}

	stories["fallback"] = map[string]interface{}{
		"message": "Sijaelewa unataka nini?",
		"next":    nil,
		"action":  nil,
	}

	gobot := core.NewGoBot(intents, stories)

	key := gobot.FindMessageKey("sitaki")

	if key != "" {
		storyObj := stories[key]
		if key != "cancel" {
			if storyObj != nil {
				story := storyObj.(map[string]interface{})
				fmt.Println(story["message"])
			} else {
				fmt.Println("interface is nil")
			}
		} else {
			fmt.Println("Karibu tena")
		}

	} else {
		fallbackObject := stories["fallback"].(map[string]interface{})
		message := fallbackObject["message"]
		fmt.Println(message)
	}

}
