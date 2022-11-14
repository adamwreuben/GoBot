package main

import (
	"github.com/adamwreuben/GoBot"
)

func main() {

	/**
	Types

	- choices
	- form -- multiple
	- echo --- just echos
	- input  --- single capture user input

	*/

	intents := make(map[string]interface{})
	stories := make(map[string]interface{})
	// states := make(map[string]interface{})

	//Create intents
	intents["greeting"] = []string{
		"hello",
		"hi",
		"how are you",
		"morning",
		"afternoon",
		"evening",
	}

	intents["order"] = []string{
		"order pizza",
		"i want pizza",
		"i want piza",
		"i want piz",
		"give me piza",
		"give me pizza",
		"pizza",
		"piza",
		"piz",
	}

	intents["menu"] = []string{
		"what pizza do you have",
		"menu",
		"menu please",
	}

	intents["action"] = []string{
		"yes",
		"i want this",
		"order this",
		"bring now",
		"give me this",
		"give me",
		"yes please",
		"alright",
		"okay",
	}

	intents["cancel"] = []string{
		"stop",
		"cancel",
	}

	//creating stories

	stories["greeting"] = map[string]interface{}{
		"message": []string{
			"Welcom to Pizza plaza, what can i help you",
			"Hello, welcome, what can i do for you?",
		},
		"type": "echo",
		"next": nil,
	}

	//Order pizza choice
	stories["order"] = map[string]interface{}{
		"message": `Please choose pizza you want?`,
		"order_choices": GoBot.GoBotChoice{
			Header:               "The following is our menu\n1. Cheese Pizza\n2. Chicken Pizza\n3.Sausage Pizza\n\n#Please select our choice",
			SuccessChoiceMessage: "Thanks for choosing, your order is being processed",
			ErrorChoiceMessage:   "Sorry the choice you selected, is not present by now!",
			Choices: []string{
				"Cheese Pizza",
				"Chicken Pizza",
				"Sausage Pizza",
			},
			IntentAction: intents["action"].([]string),
			IntentCancel: intents["cancel"].([]string),
		},
		"type": "choices",
		"next": "order_number", //nil means end
	}

	stories["order_number"] = map[string]interface{}{
		"message": "Unapenda team gani?",
		"type":    "input",
		"order_number_input": GoBot.GoBotInput{
			Header: "How many do you want?",
			Form: GoBot.Form{
				Variable: "size",
				Hint:     "How many pizza do you want?",
			},
			IntentAction:        intents["action"].([]string),
			IntentCancel:        intents["cancel"].([]string),
			SuccessInputMessage: "#Thanks your order is being processed\n#OrderId: 23434",
			ErrorInputMessage:   "Please tell me, how many pizza do you want?",
		},
		"next": nil,
	}

	//story for menu
	stories["menu"] = map[string]interface{}{
		"message": "The following is our menu,\n\n 1. Cheese Pizza\n2. Chicken Pizza\n3.Sausage Pizza\n\n #Welcome",
		"type":    "echo",
		"next":    nil,
	}

	//Defaults one... These are like Addons to GoBot so as it can cancel any action
	stories["cancel"] = map[string]interface{}{
		"message": "I gotch you, welcome again",
		"type":    "echo",
		"next":    nil,
	}

	stories["fallback"] = map[string]interface{}{
		"message": []string{
			"Hello, I didn't understand you?,",
			"Sorry, I didn't get what you want!",
		},
		"type": "echo",
		"next": nil,
	}

	//Create GoBot instance
	goBot := GoBot.NewGoBot(intents, stories)
	// fmt.Println(goBot)
	goBot.Playground()
	// listId := "busatiId"
	// kipind := listId[:len(listId)-2]
	// fmt.Println(kipind)

	// _, response := goBot.Chat("kujiunga chama")
	// fmt.Println(response)
}
