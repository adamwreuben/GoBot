package main

import (
	"github.com/adamwreuben/GoBot"
)

func main() {
	intents := make(map[string]interface{})
	stories := make(map[string]interface{})

	//Create intents
	intents["kujiunga"] = []string{
		"nataka kujiunga chama",
		"kujiunga chama",
		"jisali kwenye chama",
		"sajili kwenye chama",
		"niongeze kwenye chama",
	}

	intents["lipa"] = []string{
		"lipa ada",
		"nataka kulipa ada",
		"nataka kutuma ada",
		"lipia ada ya uanachama",
		"lipia ada ya chama",
		"tuma ada",
	}

	intents["repoti"] = []string{
		"nataka repoti",
		"nitumie report",
		"naomba repoti ya mwezi huu",
		"nataka repoti",
		"repoti",
		"repoti ipo wapi ya chama",
	}

	intents["sheria"] = []string{
		"sheria za chama",
		"sheria",
		"sheria ni zipi",
		"nataka sheria",
		"naomba sheria",
	}

	intents["cancel"] = []string{
		"sitisha",
		"acha",
		"cancel",
		"sitaki",
		"funga",
	}

	//creating stories

	stories["kujiunga"] = map[string]interface{}{
		"message": []string{
			"Karibu sana, nikuhudumie nini?",
			"Habari, Karibu nikuhudumie",
		},
		"choices": nil,
		"next":    nil, //nil means end
	}

	stories["cancel"] = map[string]interface{}{
		"message": "Usijali, Karibu tena",
		"choices": nil,
		"next":    nil,
	}

	stories["lipa"] = map[string]interface{}{
		"message": []string{
			"Lipa ada ya chama kupitia",
		},
		"lipa_choices": []string{
			"1. Airtel Money",
			"2. Mpesa",
			"3. Tigo Pesa",
		},
		"next":            nil,
		"choice_fallback": "Sorry, lipia kwa mitandao iliyopo!",
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
	goBot := GoBot.NewGoBot(intents, stories)

	goBot.Playground()

	// _, response := goBot.Chat("kujiunga chama")
	// fmt.Println(response)
}
