package GoBot

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

var userInputs string

type GoBot struct {
	Intent interface{} //This defines user intention
	Story  interface{} //This drives conversational flow (dialog)
	State  GoBotLifecycle
}

type GoBotForm struct {
	Header         string
	Form           []Form
	Answers        []FormAnswer
	IntentAction   []string
	IntentCancel   []string
	ConfirmMessage string
	ActionMessage  string
	CancelMessage  string
}

type GoBotInput struct {
	Header              string
	Form                Form
	IntentAction        []string
	IntentCancel        []string
	SuccessInputMessage string
	ErrorInputMessage   string
}

type GoBotChoice struct {
	Header               string
	SuccessChoiceMessage string
	ErrorChoiceMessage   string
	IntentAction         []string
	IntentCancel         []string
	Choices              []string
}

type Form struct {
	Variable string
	Hint     string
}

type FormAnswer struct {
	Variable string
	Value    string
}

func NewGoBot(intents interface{}, stories interface{}) *GoBot {
	return &GoBot{
		Intent: intents,
		Story:  stories,
		State:  *NewLifecycle(),
	}
}

func (gobot *GoBot) FindMessageKey(message string) string {
	matchedKey := ""
	storyType := ""
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

				//Find story type
				goBotStories := gobot.Story.(map[string]interface{})
				storyObj := goBotStories[key]
				if storyObj != nil {
					story := storyObj.(map[string]interface{})
					storyType = story["type"].(string)
					if storyType == "form" {
						storyForm := story[matchedKey+"_form"].(GoBotForm)
						//Set GoBot Lifecycle Key, and type
						gobot.State.SetState(key, storyType)

						//Get Next story if any
						nextStory := story["next"]
						if nextStory != nil {
							//Find type of next story
							goBotStories := gobot.Story.(map[string]interface{})
							storyObj := goBotStories[nextStory.(string)]
							if storyObj != nil {
								story := storyObj.(map[string]interface{})
								gobot.State.SetNextStory(nextStory.(string), story["type"].(string))
							} else {
								fmt.Println("Please set type on ", nextStory.(string), " story")
							} // TODO HANDLING ERRORs
						} else {
							gobot.State.SetNextStory("", "")
						}

						//Get all form ids and set them to gobot lifecycle
						for _, form := range storyForm.Form {
							gobot.State.ActiveFormIds = append(gobot.State.ActiveFormIds, form.Variable)
						}

					} else if storyType == "choices" {
						//Set GoBot Lifecycle Key, and type
						gobot.State.SetState(key, storyType)

						//Get Next story if any
						nextStory := story["next"]
						if nextStory != nil {
							//Find type of next story
							goBotStories := gobot.Story.(map[string]interface{})
							storyObj := goBotStories[nextStory.(string)]
							if storyObj != nil {
								story := storyObj.(map[string]interface{})
								gobot.State.SetNextStory(nextStory.(string), story["type"].(string))

							} else {
								fmt.Println("Please set type on ", nextStory.(string), " story")
							}
						} else {
							gobot.State.SetNextStory("", "")
						}

					} else if storyType == "input" {
						gobot.State.SetState(key, storyType)
						//Get Next story if any
						nextStory := story["next"]
						if nextStory != nil {
							//Find type of next story
							goBotStories := gobot.Story.(map[string]interface{})
							storyObj := goBotStories[nextStory.(string)]
							if storyObj != nil {
								story := storyObj.(map[string]interface{})
								gobot.State.SetNextStory(nextStory.(string), story["type"].(string))

							} else {
								fmt.Println("Please set type on ", nextStory.(string), " story")
							}
						}
					} else {
						//Echo response to user only
						//Set GoBot Lifecycle Key, and type
						gobot.State.SetState(key, "response")

						//Get Next story if any
						nextStory := story["next"]
						if nextStory != nil {
							gobot.State.NextStory = nextStory.(string)
						}
					}
				} else {
					//Return Error() --> no such story
					fmt.Println()
				}
				break
			}

		}

		if outerMatched {
			break
		}

	}

	return matchedKey
}

func (gobot *GoBot) Playground(id string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if userInputs == "" {
			fmt.Print("Start chatting with your Bot:\n")
		}
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			// fmt.Println(text)
			userInputs = text
			key, response := gobot.Chat(text, id)
			fmt.Println(gobot.State.SavedResults)
			fmt.Println("Intent: " + key + " ----- " + "Response: " + response)

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

func (gobot *GoBot) Chat(message string, id string) (string, string) {

	key := gobot.FindMessageKey(message)
	goBotStories := gobot.Story.(map[string]interface{})

	if key != "" {
		storyObj := goBotStories[key]
		if key != "cancel" {
			if storyObj != nil {
				story := storyObj.(map[string]interface{})

				activeKey, activeStoryType := gobot.State.GetState()

				if activeKey == key && activeStoryType == "form" {
					return gobot.CheckingGoBotState(key, activeKey, activeStoryType)
				} else if activeKey == key && activeStoryType == "choices" {
					return gobot.CheckingGoBotState(key, activeKey, activeStoryType)
				} else if activeKey == key && activeStoryType == "input" {
					return gobot.CheckingGoBotState(key, activeKey, activeStoryType)
				} else {

				}

				//Check message type (string or []string{})
				switch messageType := story["message"].(type) {
				case string:
					return key, story["message"].(string)
				case []string:
					messageSlice := story["message"].([]string)

					//Randomize response (not suitable for security but to randomize our response is OK)
					rand.Seed(time.Now().UnixNano())
					randomValue := rand.Intn(len(messageSlice))
					return key, messageSlice[randomValue]
				default:
					fmt.Printf("[]string: %v", messageType)
					return key, "GoBot only support string and []string!"
				}

			} else {
				// fmt.Println("Key when no interface: ", key)

				// fmt.Println(key)
				if !reflect.DeepEqual(gobot.State.ActiveForm, GoBotForm{}) {
					//check if user has entered value for IntentAction OR IntentCancel
					hasAction := containsInSlices(gobot.State.ActiveForm.IntentAction, message)
					hasCancel := containsInSlices(gobot.State.ActiveForm.IntentCancel, message)

					//Do this before clearly gobot lifecycle in memory
					actionMessage := gobot.State.ActiveForm.ActionMessage
					actionCancel := gobot.State.ActiveForm.CancelMessage
					actionConfirm := gobot.State.ActiveForm.CancelMessage

					gobot.State.ActiveForm = GoBotForm{}
					gobot.State.ActiveCounter = 0
					gobot.State.ActiceStoryType = ""
					gobot.State.ActiveStory = ""
					gobot.State.ActiveFormIds = []string{}

					if hasAction {
						return key, actionMessage
					} else if hasCancel {
						return key, actionCancel
					} else {
						return key, actionConfirm
					}

				} else if !reflect.DeepEqual(gobot.State.ActiveChoice, GoBotChoice{}) {
					hasAction := containsInSlices(gobot.State.ActiveChoice.IntentAction, message)
					hasCancel := containsInSlices(gobot.State.ActiveChoice.IntentCancel, message)

					//Do this before clearly gobot lifecycle in memory
					// successChoiceMessage := gobot.State.ActiveChoice.SuccessChoiceMessage
					errorChoiceMessage := gobot.State.ActiveChoice.ErrorChoiceMessage
					choice := gobot.State.ActiveChoiceValue

					gobot.State.ActiveChoice = GoBotChoice{}
					gobot.State.ActiveChoiceValue = ""

					if hasAction {
						fmt.Print("Has Action")
						return key, choice
					} else if hasCancel {
						fmt.Print("Has Cancel")
						return key, errorChoiceMessage
					} else {
						fmt.Print("Has Else nothing!")
						return key, errorChoiceMessage
					}
				} else if !reflect.DeepEqual(gobot.State.ActiveInput, GoBotInput{}) {
					hasAction := containsInSlices(gobot.State.ActiveInput.IntentAction, message)
					hasCancel := containsInSlices(gobot.State.ActiveInput.IntentCancel, message)

					//Do this before clearly gobot lifecycle in memory
					// successChoiceMessage := gobot.State.ActiveChoice.SuccessChoiceMessage
					errorInputMessage := gobot.State.ActiveInput.ErrorInputMessage
					input := gobot.State.ActiveInputValue

					gobot.State.ActiveInput = GoBotInput{}
					gobot.State.ActiveChoiceValue = ""

					if hasAction {
						return key, input
					} else if hasCancel {
						return key, errorInputMessage
					} else {
						return key, errorInputMessage
					}
				} else {
					fmt.Println("interface is nil")
					return key, ""

				}

			}
		} else {
			// fmt.Println("Key: ", key)
			// fmt.Println("Karibu tena")
			//Reset GoBotLifecycle
			gobot.State = GoBotLifecycle{}
			return key, "Karibu tena"
		}

	} else {

		if !reflect.DeepEqual(gobot.State.ActiveForm, GoBotForm{}) {

			gobot.State.ActiveForm.Answers = append(gobot.State.ActiveForm.Answers, FormAnswer{
				Variable: gobot.State.ActiveFormIds[gobot.State.ActiveCounter],
				Value:    message,
			})

			//Increment so we can go to next form data
			if gobot.State.ActiveCounter != len(gobot.State.ActiveFormIds)-1 {
				gobot.State.ActiveCounter += 1
				return key, gobot.State.ActiveForm.Form[gobot.State.ActiveCounter].Hint
			} else {
				//End of the form data is reached, so send the summary to user so as to comfirm
				/*
					Your name?
					Adam

					Your age?
					20
				*/

				answerObject := make(map[string]interface{})

				prevKey, _ := gobot.State.GetState()
				nextKey, nextType := gobot.State.GetNextStory()

				summary := gobot.State.ActiveForm.ConfirmMessage + "\n\n"
				for i := 0; i < len(gobot.State.ActiveForm.Answers); i++ {

					if nextKey != "" {
						answerObject[key] = map[string]interface{}{
							gobot.State.ActiveForm.Answers[i].Variable: gobot.State.ActiveForm.Answers[i].Value,
						}
					}

					summary += gobot.State.ActiveForm.Answers[i].Variable + "\n" + gobot.State.ActiveForm.Answers[i].Value + "\n"
				}

				//If it meant all requirement.... it goes to next story
				if nextKey != "" {
					//save the result according to story key in Gobot lifecycle object
					gobot.State.SavedResults[prevKey] = answerObject

					//After saving the results ,,,,, Free GoBotForm to be used by another story
					if nextKey != prevKey {
						gobot.State.ActiveForm = GoBotForm{}
						gobot.State.ActiveCounter = 0
						//Set state for upcoming story
						gobot.State.SetState(nextKey, nextType)

						activeKey, activeStoryType := gobot.State.GetState()
						return gobot.CheckingGoBotState(nextKey, activeKey, activeStoryType)
					}

				} else {
					fmt.Println("No next story")
					gobot.State.SetNextStory("", "")
				}

				return key, summary

			}
		} else if !reflect.DeepEqual(gobot.State.ActiveChoice, GoBotChoice{}) {
			answerObject := ""

			prevKey, _ := gobot.State.GetState()

			nextKey, nextType := gobot.State.GetNextStory()

			for _, choice := range gobot.State.ActiveChoice.Choices {
				if strings.ToLower(choice) == strings.ToLower(message) {
					if nextKey != "" {
						answerObject = message
					}
				}

			}

			//If it meant all requirement.... it goes to next story
			if nextKey != "" {
				//save the result according to story key in Gobot lifecycle object
				gobot.State.SavedResults[prevKey] = answerObject

				//After saving the results ,,,,, Free GoBotForm to be used by another story
				if nextKey != prevKey {
					gobot.State.ActiveChoice = GoBotChoice{}
					//Set state for upcoming story
					gobot.State.SetState(nextKey, nextType)
					activeKey, activeStoryType := gobot.State.GetState()
					return gobot.CheckingGoBotState(nextKey, activeKey, activeStoryType)
				}

			} else {
				//save the result according to story key in Gobot lifecycle object
				gobot.State.SavedResults[prevKey] = answerObject
				gobot.State.SetNextStory("", "")
				return prevKey, gobot.State.ActiveChoice.SuccessChoiceMessage + " " + gobot.State.ActiveChoiceValue
			}

			if prevKey == nextKey {
				return prevKey, gobot.State.ActiveChoice.SuccessChoiceMessage + " " + gobot.State.ActiveChoiceValue
			} else {
				return "key", gobot.State.ActiveChoice.ErrorChoiceMessage
			}
		} else if !reflect.DeepEqual(gobot.State.ActiveInput, GoBotInput{}) {

			prevKey, _ := gobot.State.GetState()
			nextKey, nextType := gobot.State.GetNextStory()

			gobot.State.ActiveInputValue = message

			if nextKey != "" {

				//save the result according to story key in Gobot lifecycle object
				gobot.State.SavedResults[prevKey] = message

				//After saving the results ,,,,, Free GoBotForm to be used by another story
				if nextKey != prevKey {
					gobot.State.ActiveInput = GoBotInput{}
					//Set state for upcoming story
					gobot.State.SetState(nextKey, nextType)

					activeKey, activeStoryType := gobot.State.GetState()
					return gobot.CheckingGoBotState(nextKey, activeKey, activeStoryType)
				}

			} else {

				gobot.State.SavedResults[prevKey] = message

				gobot.State.SetNextStory("", "")

			}

			return prevKey, gobot.State.ActiveInput.SuccessInputMessage

			//If it meant all requirement.... it goes to next story

		} else {
			fallbackObject := goBotStories["fallback"].(map[string]interface{})

			//Check message type (string or []string{})
			switch messageType := fallbackObject["message"].(type) {
			case string:
				return key, fallbackObject["message"].(string)
			case []string:
				messageSlice := fallbackObject["message"].([]string)

				//Randomize response (not suitable for security but to randomize our response is OK)
				rand.Seed(time.Now().UnixNano())
				randomValue := rand.Intn(len(messageSlice))
				return key, messageSlice[randomValue]
			default:
				fmt.Printf("[]string: %v", messageType)
				return key, "GoBot only support string and []string!"
			}
		}

	}
}

func (gobot *GoBot) CheckingGoBotState(key string, activeKey string, activeStoryType string) (string, string) {
	goBotStories := gobot.Story.(map[string]interface{})
	storyObj := goBotStories[key]

	story := storyObj.(map[string]interface{})

	if key != "cancel" {

		if activeKey == key && activeStoryType == "form" {
			storyForm := story[key+"_form"].(GoBotForm)

			if reflect.DeepEqual(gobot.State.ActiveForm, GoBotForm{}) {

				gobot.State.ActiveCounter = 0 //This will help us iterate through forms data slice
				gobot.State.ActiveForm = storyForm
				//Ask a first form data?
				return key, storyForm.Form[gobot.State.ActiveCounter].Hint

			}

		} else if activeKey == key && activeStoryType == "choices" {
			storyForm := story[key+"_choices"].(GoBotChoice)

			if reflect.DeepEqual(gobot.State.ActiveForm, GoBotForm{}) {
				gobot.State.ActiveChoice = storyForm
				//Ask a first form data?
				return key, storyForm.Header
			}

			//others
		} else if activeKey == key && activeStoryType == "input" {
			storyForm := story[key+"_input"].(GoBotInput)
			if reflect.DeepEqual(gobot.State.ActiveForm, GoBotForm{}) {
				gobot.State.ActiveInput = storyForm
				//Ask a first form data?
				return key, storyForm.Header
			}
		} else {
			//Other
			return key, ""
		}

		return key, ""
	} else {
		return key, "Karibu tena"
	}

}

func (gobot *GoBot) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(gobot)
}

func (gobot *GoBot) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(gobot)
}
