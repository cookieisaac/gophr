package main

import (
	"crypto/rand"
	"fmt"
)

const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const idSourceLen = byte(len(idSource))

func GenerateID(prefix string, length int) string {
	id := make([]byte, length)
	
	rand.Read(id)
	
	for i, b := range id {
		id[i] = idSource[b%idSourceLen]
	}
	
	return fmt.Sprintf("%s_%s", prefix, string(id))
}