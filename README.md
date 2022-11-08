GoBot
=======
![GoPesa](https://storage.googleapis.com/gopherizeme.appspot.com/gophers/11fa3afd6080dae6903dc359e42990d68657e17a.png)

``GoBot`` is a Go library for building chat bots super fast and easily with only interfaces(AKA JSON) inspired by [Sarufi](https://docs.sarufi.io/). 
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/adamwreuben/GoBot)

GoBot Terminology
====================

``Intent``  Is a user's intention or goal when interacting with a chatbot. For example, if a user asks a chatbot a question, the intent of the user is to get an answer to the question. If a user asks a chatbot to order food, the intent of the user is to order food.

``Stories`` Is conversation between a user and a chatbot. Stories is the one that controls the flow of conversation.



## Installation

Simply install with the `go get` command:
```
go get github.com/adamwreuben/GoBot
```
Then import it to your main package as:
```
package main

import (
	gobot "github.com/adamwreuben/GoBot"
)
```


## Usage
First you need to create ``Intent`` and ``stories`` variable of type ``map[string]interface{}`` inside your ``main.go`` as follows: