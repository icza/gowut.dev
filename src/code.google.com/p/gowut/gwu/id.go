// Copyright 2013 Andras Belicza. All rights reserved.

// ID type definition, and unique ID generation.

package gwu

import (
	"strconv"
)

// The type of the ids of the components.
type ID int

// Converts an ID to a string.
func (id ID) String() string {
	return strconv.Itoa(int(id))
}

// Converts a string to ID
func AtoID(s string) (ID, error) {
	id, err := strconv.Atoi(s)

	if err != nil {
		return ID(-1), err
	}
	return ID(id), nil
}

// Component id generation and provider

// A channel used to generate unique ids
var idChan chan ID = make(chan ID)

// init stats a new go routine to generate unique ids
func init() {
	go func() {
		for i := 0; ; i++ {
			idChan <- ID(i)
		}
	}()
}

// nextCompId returns a unique component id
func nextCompId() ID {
	return <-idChan
}
