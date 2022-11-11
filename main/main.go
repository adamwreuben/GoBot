package main

import (
	"github.com/adamwreuben/GoBot"
)

func main() {
	intents := make(map[string]interface{})
	stories := make(map[string]interface{})
	// states := make(map[string]interface{})

	//Create intents
	intents["salam"] = []string{
		"mambo",
		"vipi",
		"upo poa",
		"fresh",
		"nambie",
		"nakuona",
		"salama",
		"za asubuhi",
		"za mchana",
		"uko poa",
		"niaje",
		"hello",
		"oya",
		"kama kawa",
	}

	intents["cancel"] = []string{
		"sitisha",
		"acha",
		"cancel",
		"sitaki",
		"funga",
	}

	intents["sajili"] = []string{
		"nataka kujisajili",
		"nataka kujiunga",
		"add me",
	}

	intents["tuma"] = []string{
		"tuma",
		"ndio",
		"sawa",
		"forward",
		"tuma sasa hivi",
		"send",
	}

	//creating stories

	stories["salam"] = map[string]interface{}{
		"message": `Habari karibu sana TBC, sasa unaweza tuma ujumbe wako moja kwa moja kwenye kipindi ukipendacho`,
		"salam_choices": GoBot.GoBotChoice{
			Header:               "Chagua kipindi?",
			SuccessChoiceMessage: "Umechagua ",
			ErrorChoiceMessage:   "Samahani hauna hicho kipindi!",
			Choices: []string{
				"Busati",
				"Sasambu",
				"Sekeseke",
				"Millazo EP",
				"Simela",
				"Kinaganaga",
				"Papaso",
				"Ligi Kuu Tanzania",
			},
			IntentAction: intents["tuma"].([]string),
			IntentCancel: intents["cancel"].([]string),
		},
		"type": "choices",
		"next": nil, //nil means end
	}

	stories["sajili"] = map[string]interface{}{
		"message": "Tafadhali nipe majibu ya haya maswali?",
		"type":    "form",
		"sajili_form": GoBot.GoBotForm{
			Header: "Karibu jaza maswali yafuatayo kwa usahihi",
			Form: []GoBot.Form{
				{
					Variable: "name",
					Hint:     "Jina lako nani?",
				},
				{
					Variable: "age",
					Hint:     "Una umri gani?",
				},
			},
			IntentAction:   intents["tuma"].([]string),
			IntentCancel:   intents["cancel"].([]string),
			ConfirmMessage: "Tafadhali nijibu, nitume au nisitume?",
			ActionMessage:  "Asante, taarifa zako, zimetumwa",
			CancelMessage:  "Sijatuma taarifa zako, ukitaka kutuma tena, karibu sana!",
		},
	}

	stories["cancel"] = map[string]interface{}{
		"message": "Usijali, Karibu tena",
		"type":    "default",
		"choices": nil,
		"next":    nil,
	}

	stories["fallback"] = map[string]interface{}{
		"message": []string{
			"Sijaelewa unataka nini?",
			"Samahani sijakuelewa!, Jaribu kuandika saada?",
		},
		"type":    "default",
		"choices": nil,
		"next":    nil,
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
