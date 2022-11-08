package model

import (
	"encoding/json"
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
func (p *GoBot) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *GoBot) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
