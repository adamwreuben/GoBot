package core

import (
	"encoding/json"
	"io"
	"strings"
)

type GoBot struct {
	Intent interface{} //This defines user intention
	Story  interface{} //This drives conversational flow
}

var matchedKey string

func NewGoBot(intents interface{}, stories interface{}) *GoBot {
	return &GoBot{
		Intent: intents,
		Story:  stories,
	}
}

func (gobot *GoBot) FindMessageKey(message string) string {

	matchMessage := NewMatch("*" + strings.ToLower(message) + "*")
	//Find a story where the message belongs

	//Loop through all keys of intents
	outerMatched := false
	for key, intent := range gobot.Intent.(map[string]interface{}) {
		//From Key loop in key slice to find message match
		for _, intentMessage := range intent.([]string) {
			// fmt.Println(intentMessage)
			outerMatched = matchMessage.Matches(strings.ToLower(intentMessage))
			if outerMatched {
				matchedKey = key
				break
			}

		}

		if outerMatched {
			break
		}

	}

	return matchedKey
}

func (gobot *GoBot) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(gobot)
}

func (gobot *GoBot) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(gobot)
}
