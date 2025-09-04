package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}

type bold struct {
	message string
}

type code struct {
	message string
}

func (plain plainText) format() string {
	return plain.message
}

func (b bold) format() string {
	return fmt.Sprintf("**%s**", b.message)
}

func (c code) format() string {
	return fmt.Sprintf("`%s`", c.message)
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
