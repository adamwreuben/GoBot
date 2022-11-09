package GoBot

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var userInputs string

type GoBot struct {
	Intent interface{} //This defines user intention
	Story  interface{} //This drives conversational flow (dialog)
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

func (gobot *GoBot) Playground() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if userInputs == "" {
			fmt.Print("Start chatting:\n")
		}
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			// fmt.Println(text)
			userInputs = text
			_, response := gobot.Chat(text)
			fmt.Println(response)
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func (gobot *GoBot) Chat(message string) (string, string) {

	key := gobot.FindMessageKey(message)
	goBotStories := gobot.Story.(map[string]interface{})

	if key != "" {
		storyObj := goBotStories[key]
		if key != "cancel" {
			if storyObj != nil {
				story := storyObj.(map[string]interface{})

				//Becase choice depend from next
				// fmt.Println("Key: real ", key)
				// fmt.Println(story["message"])

				return key, story["message"].(string)

			} else {
				// fmt.Println("Key when no interface: ", key)
				if strings.Contains(key, "choices") {
					key := key[:len(key)-8]
					story := goBotStories[key].(map[string]interface{})
					next := story["next"]

					if next != nil {
						story := goBotStories[next.(string)].(map[string]interface{})
						// fmt.Println(story["message"])
						return key, story["message"].(string)

					} else { // when next is Nil
						story := goBotStories[key].(map[string]interface{})
						choices := story[key+"_"+"choices"].([]string)

						for _, choice := range choices {
							if strings.Contains(strings.ToLower(choice), message) {
								// fmt.Println(choice)
								return key, choice
							}
						}

						return key, "Asante sana"

					}

				} else {
					fmt.Println("interface is nil")
					return "", ""

				}
			}
		} else {
			// fmt.Println("Key: ", key)
			// fmt.Println("Karibu tena")
			return key, "Karibu tena"
		}

	} else {
		// fmt.Println("Key: ", key)
		fallbackObject := goBotStories["fallback"].(map[string]interface{})
		message := fallbackObject["message"]
		// fmt.Println(message)
		return key, message.(string)
	}
}
