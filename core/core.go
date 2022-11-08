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

func NewGoBot(intents interface{}, stories interface{}) *GoBot {
	return &GoBot{
		Intent: intents,
		Story:  stories,
	}
}

func (gobot *GoBot) FindMessageKey(message string) string {
	matchedKey := ""
	matchMessage := NewMatch("*" + strings.ToLower(message) + "*")
	//Find a story where the message belongs

	//Loop through all keys of intents
	outerMatched := false
	for key, intent := range gobot.Intent.(map[string]interface{}) {
		//From Key loop in key slice to find message match
		for _, intentMessage := range intent.([]string) {
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

	//Loop choices slice  *This is not good practice...Will be fixed as days goes on
	if matchedKey == "" {
		for key, choices := range gobot.Story.(map[string]interface{}) {
			//Check if story has choices
			if strings.Contains(key, "choices") {
				matchedKey = ""
			} else {
				choiceObject := choices.(map[string]interface{})[key+"_choices"]
				if choiceObject != nil {
					for _, choice := range choiceObject.([]string) {
						outerMatched = matchMessage.Matches(strings.ToLower(choice))
						if outerMatched {
							if strings.Contains(key, "choices") {
								matchedKey = key
							} else {
								matchedKey = key + "_choices"
							}
							break
						}
					}
				}

				if outerMatched {
					break
				}
			}

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
