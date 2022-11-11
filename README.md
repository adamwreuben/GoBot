GoBot
=======
![GoPesa](https://storage.googleapis.com/gopherizeme.appspot.com/gophers/11fa3afd6080dae6903dc359e42990d68657e17a.png)

``GoBot`` is a Go library for building chat bots super fast and easily. 
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/adamwreuben/GoBot)


# Why I made GoBot

It’s no longer about a race car, it’s a boat. And steering a boat requires too much damn planning and patience, So if ``GoBot`` succeeds, it really never was me, it’s the people there who can drive boats (open community). As I am with the world, I am a cheerleader.

Doing all this is just a hobby, is just a way to write some open source infrastructure so I’m not doing nothing. It’s not the future.

``Lets make our open community safe``



GoBot Terminology
====================

``Intent``  Is a user's intention or goal when interacting with a chatbot. For example, if a user asks a chatbot a question, the intent of the user is to get an answer to the question. If a user asks a chatbot to order food, the intent of the user is to order food.

``Stories`` Is conversation between a user and a chatbot. Stories is the one that controls the flow of conversation.



## Installation

Simply install with the `go get` command:
```
go get github.com/adamwreuben/GoBot@v0.0.4
```
Then import it to your main package as:
```
package main

import (
	gobot "github.com/adamwreuben/GoBot"
)
```

## Uninstallation
Simply uninstall with the `go get` command:
```
go get github.com/adamwreuben/GoBot@none
```

# Usage
First you need to create ``Intent`` and ``stories`` variable of type ``map[string]interface{}`` inside your ``main.go`` as follows:
```
intents := make(map[string]interface{})
stories := make(map[string]interface{})

```

## Creating Intents
Then start creating intent, by providing key to intents as intent name as follows, In this demo we will create an ordering pizza chatbot, The followings are its intents.

```
intents["greets"] = []string{
		"Hello",
		"Hi",
		"Mambo",
		"Za asubuhi",
		"Za mchana",
		"Za usiku",
		"Hola",
}

intents["sajili"] = []string{
		"nataka kujisajili",
		"nataka kujiunga",
		"add me",
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

```

From above we have create 3 intents namely ``greets``,   ``cancel``,  ``order_pizza``

## Creating stories (Dialog flow)
Creating stories is simple, make sure that the ``Key name of the intent`` matches the ``key name of the story``, by doing so GoBot we know exactly which intent corresponds to what story.

Example of stories according to our ordering pizza chatbot

```
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
		"order_pizza_choices": GoBot.GoBotChoice{
			Header:               "Unataka pizza gani?",
			SuccessChoiceMessage: "Karibu tena!",
			ErrorChoiceMessage:   "Samahani pizza uliochagua haipo!",
			Choices: []string{
				"Chicken Pizza",
				"Cheese Pizza",
				"Mixture Pizza",
				"Skyline Pizza",
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
			CancelMessage:  "Sijatuma taarifa zako, ukitaka kutuma tena, karibu sana",
		},
}

stories["fallback"] = map[string]interface{}{
		"message": "Sijaelewa unataka nini?",
		"type":"default",
		"choices": nil,
		"next":    nil,
}

```

But also if you want you stories to answer user in random way, you do this. For example here, when the use type bye, GoBot should respond with words like ``Karibu tena`` or ``Bye, welcome again`` etc... To do so u need to use ``[]string struct`` other than ``string``

Consider an example below

```
stories["goodbye"] = map[string]interface{}{
		"message": []string{
			"Karibu tena",
			"Bye, welcome again",
		},
		"choices": nil,
		"next":    nil,
	}

```

``NB`` The following must be provided as ``AddOns`` to GoBot to ensure it can cancel execution when the user type ``Intents from cancel intent``
From above stories, the following are Special stories key which are recognized by ``GoBot`` which are ``cancel`` - which is responsible to cancel any excution of the bot. and ``fallback`` - This is a default answer ``GoBot`` Gives when it doesn't understand the user intent.

## Understanding story keys
```
If key are not provided, GoBot will assume all keys are nil
```
As you can see story has following keys
1. message - This is the response, which will be sent by GoBot.
2. next - This link another story.
3. choices - This acts as creating list of options to choose for (For example in our example Pizza type has choice)
4. choice_fallback - Is what ``GoBot`` will say when user entered choices that are not present.


## Linking intents with stories
```
goBot := GoBot.NewGoBot(intents, stories)

```

## PlayGround
This is a commandline playground of your created GoBot
```
goBot.Playground()

```

## Chatting with GoBot
```
key, response := goBot.Chat("hello") --> Chat(your message)
fmt.Println(key, response)

```

# Integration with WhatsApp
```
In progress
```

# Integration with Telegram
Integrating with Telegram, to create its chat bot we will use this awosome Open Source Library from [Golang-Tanzania](https://github.com/Golang-Tanzania/Group-Bot).

```
In Progress

```


## Authors

This package is authored and maintained by [who can drive boats](https://github.com/community)

## License

MIT License

Copyright (c) 2022 Adam Reuben
