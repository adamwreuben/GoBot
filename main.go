package main

import (
	"fmt"
	"gobot/core"
)

var answers map[string]interface{}

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

	//list
	stories["order_pizza"] = map[string]interface{}{
		"message": "Unataka pizza gani?",
		"choices": []string{
			"1. Chicken Pizza",
			"2. Cheese Pizza",
			"3. Mixture Pizza",
			"4. Skyline Pizza",
		},
		"next":            "soda_choices",
		"choice_fallback": "Sorry, hatuna aina hio ya pizza!",
	}

	stories["soda_choices"] = map[string]interface{}{
		"message": "Unakata kinywaji gani?",
		"choices": []string{
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

	key := gobot.FindMessageKey("pepsi")

	if key != "" {
		storyObj := stories[key]
		if key != "cancel" {
			if storyObj != nil {
				story := storyObj.(map[string]interface{})

				//Becase choice depend from next
				fmt.Println("Key: ", key)
				fmt.Println(story["message"])

			} else {
				fmt.Println("Key: ", key)
				fmt.Println("interface is nil")
			}
		} else {
			fmt.Println("Key: ", key)
			fmt.Println("Karibu tena")
		}

	} else {
		fmt.Println("Key: ", key)
		fallbackObject := stories["fallback"].(map[string]interface{})
		message := fallbackObject["message"]
		fmt.Println(message)
	}

}
