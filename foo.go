package main

import (
	"encoding/json"
	"fmt"
)

type Capability string

type Capabilities struct {
	// Added capabilities
	Add []Capability `json:"add,omitempty"`
}

func main() {
	var c Capability
	var d Capability
	c = "CAP_CHOWN"
	d = "CAP_FOO"

	caps := Capabilities{}
	adds := make([]Capability, 2)
	adds[0] = c
	adds[1] = d
	caps.Add = adds

	bolB, _ := json.Marshal(caps)
	fmt.Println(string(bolB))

}
