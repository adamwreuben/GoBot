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

	//creating stories

	stories["salam"] = map[string]interface{}{
		"message": `Habari karibu sana TBC, sasa unaweza tuma ujumbe wako moja kwa moja kwenye kipindi ukipendacho`,
		"salam_choices": []string{
			"1. Busati",
			"2. Sasambu",
			"3. Sekeseke",
			"4. Millazo EP",
			"5. Simela",
			"6. Kinaganaga",
			"7. Papaso",
			"8. Ligi Kuu Tanzania",
		},
		"next": nil, //nil means end
	}

	stories["cancel"] = map[string]interface{}{
		"message": "Usijali, Karibu tena",
		"choices": nil,
		"next":    nil,
	}


	stories["fallback"] = map[string]interface{}{
		"message": []string{
			"Sijaelewa unataka nini?",
			"Samahani sijakuelewa!, Jaribu kuandika saada?",
		},
		"choices": nil,
		"next":    nil,
	}

	//Create GoBot instance
	goBot := GoBot.NewGoBot(intents, stories, nil)

	goBot.Playground()

	// _, response := goBot.Chat("kujiunga chama")
	// fmt.Println(response)
}
