package core

import (
	"encoding/json"
	"fmt"
	"io"
)

type GoBot struct {
	Intent interface{} //This defines user intention
	Story  interface{} //This drives conversational flow
}

func NewGoBot(intents interface{}, stories interface{}) *GoBot {
	return &GoBot{
		Intent: intents,
		Story:  stories,
	}
}

func (gobot *GoBot) Chat(message string) string {

	matchMessage := NewMatch("*" + message + "*")
	//Find a story where the message belongs

	//Loop through all keys of intents
	for key, intent := range gobot.Intent.(map[string]interface{}) {
		//From Key loop in key slice to find message match
		fmt.Println(key + "\n ======== \n")
		for _, intentMessage := range intent.([]string) {
			fmt.Println(intentMessage)
			matched := matchMessage.Matches(intentMessage)
			fmt.Println(matched)

		}

	}

	return ""
}

func (gobot *GoBot) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(gobot)
}

func (gobot *GoBot) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(gobot)
}
