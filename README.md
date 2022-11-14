GoBot
=======
![GoPesa](https://storage.googleapis.com/gopherizeme.appspot.com/gophers/11fa3afd6080dae6903dc359e42990d68657e17a.png)

``GoBot`` is a Go library for building chat bots super fast and easily. 
[![Made in Tanzania](https://img.shields.io/badge/made%20in-tanzania-008751.svg?style=flat-square)](https://github.com/Tanzania-Developers-Community/made-in-tanzania)


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
go get github.com/adamwreuben/GoBot@v0.0.6
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

# GoBot story types
``echo`` - This is the default story that echo back to user

Example of ``echo`` story below:
```
stories["greeting"] = map[string]interface{}{
		"message": []string{
			"Welcom to Pizza plaza, what can i help you",
			"Hello, welcome, what can i do for you?",
		},
		"type": "echo", //GoBot stories must have types
		"next": nil,
	}

```



``form`` - This is the story that ask user a couple of questions and save it in memory
Example of ``form`` story below:
```
stories["registration"] = map[string]interface{}{
		"message": "Welcome to Pizza social club",
		"type":    "form",
		"sajili_form": GoBot.GoBotForm{
			Header: "Please answer the following question to register you, to our club!",
			Form: []GoBot.Form{
				{
					Variable: "name",
					Hint:     "Your full name",
				},
				{
					Variable: "age",
					Hint:     "Your age?",
				},
			},
			IntentAction:   intents["send"].([]string),
			IntentCancel:   intents["cancel"].([]string),
			ConfirmMessage: "Do you confirm that you have provided, real information",
			ActionMessage:  "Thanks, your information has been sent!",
			CancelMessage:  "Sorry, answer all questions needed!",
		},
		"next": nil,
	}

```


``choices`` - This is the story where user has to select one choice out of many options
Example of ``choices`` story below:
```
stories["order"] = map[string]interface{}{
		"message": `Please choose pizza you want?`,
		"order_choices": GoBot.GoBotChoice{ // Note here key of story is written in this format key_name_choices
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
		"type": "choices", // type of choices
		"next": nil, //nil means end
	}

```


``input`` - This is the story that save user answer in memory GoBot Instance
Example of ``input`` story below:
```
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

```


## Creating Intents
Then start creating intent, by providing key to intents as intent name as follows, In this demo we will create an ordering pizza chatbot, The followings are its intents.

```
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

```

From above we have create 3 intents namely ``greets``,   ``cancel``,  ``order_pizza``

## Creating stories (Dialog flow)
Creating stories is simple, make sure that the ``Key name of the intent`` matches the ``key name of the story``, by doing so GoBot we know exactly which intent corresponds to what story.

Example of stories according to our ordering pizza chatbot

```
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
