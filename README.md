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
go get github.com/adamwreuben/GoBot@v0.0.3
```
Then import it to your main package as:
```
package main

import (
	gobot "github.com/adamwreuben/GoBot"
)
```


# Usage
First you need to create ``Intent`` and ``stories`` variable of type ``map[string]interface{}`` inside your ``main.go`` as follows:
```
intents := make(map[string]interface{})
stories := make(map[string]interface{})

```

## Creating Intents
Then start creating intent, by providing key to intents as intent name as follows, In this demo we will create an ordering pizza chatbot, The followings are its intents, The following example was taken from [Sarfufi doc](https://docs.sarufi.io/docs/Getting%20started%20/create-a-simple-chatbot)

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

This package is authored and maintained by [Adam Reuben](https://github.com/adamwreuben/GoBot)

## License

MIT License

Copyright (c) 2022 Adam Reuben
