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
		"nataka pizza",
		"nataka piza",
		"nilete pizza",
		"nilete piza",
		"i want pizza",
		"i want piza",
		"i want piz",
		"give me piza",
		"give me pizza",
		"pizza",
		"piza",
		"piz",
	}

	intents["event"] = []string{
		"events",
		"event",
		"tukio",
		"matukio",
		"ticket",
		"tickets",
		"tiketi",
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
			"Karibu sana Wamazengo event, ukitaka kununua ticket andika neno ticket au tiketi",
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

	stories["event"] = map[string]interface{}{
		"message": `Changua event uipendayo?`,
		"event_choices": GoBot.GoBotChoice{
			Header:               "Karibu Wamazengo event show\n\n Kuna ticket aina tatu\n\n1. Normal --- Tsh 20,000/=\n2. VIP --- Tsh 40,000/=\n3. Platnum --- Tsh 70,000/=",
			SuccessChoiceMessage: "Asante kwa kuchagua event hii.",
			ErrorChoiceMessage:   "Samahani chagua aina ya ticket uipendayo, Uliochagua mara ya kwanza haipo!",
			Choices: []string{
				"Normal",
				"VIP",
				"Platnum",
			},
			IntentAction: intents["action"].([]string),
			IntentCancel: intents["cancel"].([]string),
		},
		"type": "choices",
		"next": "lipa", //nil means end
	}

	stories["lipa"] = map[string]interface{}{
		"message": `Lipa kwa?`,
		"lipa_choices": GoBot.GoBotChoice{
			Header:               "Lipa kupitia\n\n1. Airtel Money\n2. Mpesa\n3. Tigo Pesa",
			SuccessChoiceMessage: "Asante kwa kuchagua",
			ErrorChoiceMessage:   "Samahani!, ombi ulilochagua halipo",
			Choices: []string{
				"Airtel Money",
				"Mpesa",
				"Tigo Pesa",
			},
			IntentAction: intents["action"].([]string),
			IntentCancel: intents["cancel"].([]string),
		},
		"type": "choices",
		"next": nil, //nil means end
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

	//Create GoBot runnning instance
	goBot := GoBot.NewGoBot(intents, stories)
	// fmt.Println(goBot)
	goBot.Playground("adam")
	// listId := "busatiId"
	// kipind := listId[:len(listId)-2]
	// fmt.Println(kipind)

	// _, response := goBot.Chat("kujiunga chama")
	// fmt.Println(response)
}
