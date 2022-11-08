package main

import "github.com/adamwreuben/GoBot"

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
		"message": "Karibu sana ACT wazalendo, kujiunga na chama jaza taarifa kwenye fomu hii hapa chini\n\n https://www.google.com",
		"choices": nil,
		"next":    nil, //nil means end
	}

	stories["cancel"] = map[string]interface{}{
		"message": "Usijali, Karibu tena",
		"choices": nil,
		"next":    nil,
	}

	stories["repoti"] = map[string]interface{}{
		"message": "https://firebasestorage.googleapis.com/v0/b/song-d1bd1.appspot.com/o/Reports%2FDebit%20Mandate_API%20Specification%20Document_Biller_v1.2.pdf?alt=media&token=9de278d6-c8ae-45a4-9a44-0db7f72bfd2b",
		"choices": nil,
		"next":    nil,
	}

	stories["lipa"] = map[string]interface{}{
		"message": "Lipa ada ya chama kupitia",
		"lipa_choices": []string{
			"1. Airtel Money",
			"2. Mpesa",
			"3. Tigo Pesa",
		},
		"next":            nil,
		"choice_fallback": "Sorry, lipia kwa mitandao iliyopo!",
	}

	stories["fallback"] = map[string]interface{}{
		"message": "Sijaelewa unataka nini?",
		"choices": nil,
		"next":    nil,
	}

	//Create GoBot instance
	goBot := GoBot.NewGoBot(intents, stories)

	goBot.Playground()
}
