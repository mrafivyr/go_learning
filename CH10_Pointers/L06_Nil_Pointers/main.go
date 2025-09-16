package main

import "strings"

/*
func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}

*/

var replacer = strings.NewReplacer(
	"fubb", "****",
	"shiz", "****",
	"witch", "*****",
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	*message = replacer.Replace(*message)
}
