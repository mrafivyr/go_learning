package main

import (
	"fmt"
	"reflect"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func validateUser(mToSend messageToSend, userMode string) bool {
	// Get the reflect.Type of the Message struct itself
	messageType := reflect.TypeOf(mToSend)
	// Get the reflect.Value of the Message struct instance
	messageValue := reflect.ValueOf(mToSend)

	// Get the reflect.Type of the User struct itself
	userType := reflect.TypeOf(user{})

	// --- Check Sender ---
	userField, foundUser := messageType.FieldByName(userMode) // Note: "Sender" (capitalized)
	if !foundUser {
		fmt.Printf("Error: '%s' field not found in messageToSend struct.", userMode)
		return false
	}
	if userField.Type != userType { // Check if the field's type is indeed 'user'
		fmt.Printf("Error: '%s' field type (%s) does not match 'user' type (%s).\n", userMode, userField.Type, userType)
		return false
	}

	// Get the actual Value of the Sender field
	userValue := messageValue.FieldByName(userMode) // Note: "Sender" (capitalized)

	// Access the fields of the senderValue (which is a reflect.Value representing the 'user' struct)
	userNameValue := userValue.FieldByName("name")     // Note: "Name" (capitalized)
	userNumberValue := userValue.FieldByName("number") // Note: "Number" (capitalized)

	// Check if the Name and Number fields exist and are not their zero values
	if !userNameValue.IsValid() || userNameValue.String() == "" {
		fmt.Printf("Error: %s 'Name' field is invalid or empty.", userMode)
		return false
	}
	if !userNumberValue.IsValid() || userNumberValue.Int() == 0 {
		fmt.Printf("Error: %s 'Number' field is invalid or zero.", userMode)
		return false
	}
	return true
}

func canSendMessage(mToSend messageToSend) bool {

	if !validateUser(mToSend, "sender") {
		return false
	}

	if !validateUser(mToSend, "recipient") {
		return false
	}

	return true // If all checks pass
	/*
		// Get the reflect.Type of the Message struct itself
		messageType := reflect.TypeOf(mToSend)
		// Get the reflect.Value of the Message struct instance
		messageValue := reflect.ValueOf(mToSend)

		// Get the reflect.Type of the User struct itself
		userType := reflect.TypeOf(user{})

		// --- Check Sender ---
		senderField, foundSender := messageType.FieldByName("sender") // Note: "Sender" (capitalized)
		if !foundSender {
			fmt.Println("Error: 'Sender' field not found in messageToSend struct.")
			return false
		}
		if senderField.Type != userType { // Check if the field's type is indeed 'user'
			fmt.Printf("Error: 'Sender' field type (%s) does not match 'user' type (%s).\n", senderField.Type, userType)
			return false
		}

		// Get the actual Value of the Sender field
		senderValue := messageValue.FieldByName("sender") // Note: "Sender" (capitalized)

		// Access the fields of the senderValue (which is a reflect.Value representing the 'user' struct)
		senderNameValue := senderValue.FieldByName("name")     // Note: "Name" (capitalized)
		senderNumberValue := senderValue.FieldByName("number") // Note: "Number" (capitalized)

		// Check if the Name and Number fields exist and are not their zero values
		if !senderNameValue.IsValid() || senderNameValue.String() == "" {
			fmt.Println("Error: Sender 'Name' field is invalid or empty.")
			return false
		}
		if !senderNumberValue.IsValid() || senderNumberValue.Int() == 0 {
			fmt.Println("Error: Sender 'Number' field is invalid or zero.")
			return false
		}

		// --- Check Recipient ---
		recipientField, foundRecipient := messageType.FieldByName("recipient") // Note: "Recipient" (capitalized)
		if !foundRecipient {
			fmt.Println("Error: 'Recipient' field not found in messageToSend struct.")
			return false
		}
		if recipientField.Type != userType { // Check if the field's type is indeed 'user'
			fmt.Printf("Error: 'Recipient' field type (%s) does not match 'user' type (%s).\n", recipientField.Type, userType)
			return false
		}

		// Get the actual Value of the Recipient field
		recipientValue := messageValue.FieldByName("recipient") // Note: "Recipient" (capitalized)

		// Access the fields of the recipientValue (which is a reflect.Value representing the 'user' struct)
		recipientNameValue := recipientValue.FieldByName("name")     // Note: "Name" (capitalized)
		recipientNumberValue := recipientValue.FieldByName("number") // Note: "Number" (capitalized)

		// Check if the Name and Number fields exist and are not their zero values
		if !recipientNameValue.IsValid() || recipientNameValue.String() == "" {
			fmt.Println("Error: Recipient 'Name' field is invalid or empty.")
			return false
		}
		if !recipientNumberValue.IsValid() || recipientNumberValue.Int() == 0 {
			fmt.Println("Error: Recipient 'Number' field is invalid or zero.")
			return false
		}
		return true // If all checks pass
	*/

}

func main() {
	tempMsg := messageToSend{
		message:   "you have an appointment tomorrow",
		sender:    user{name: "Brenda Halafax", number: 16545550987},
		recipient: user{name: "Sally Sue", number: 19035558973},
	}

	if canSendMessage(tempMsg) {
		fmt.Println("Message can be sent!")
	} else {
		fmt.Println("Message cannot be sent.")
	}

	// Test with a problematic message
	badMsg := messageToSend{
		message:   "test",
		sender:    user{name: "", number: 0}, // Zero values
		recipient: user{name: "Valid", number: 123},
	}
	if canSendMessage(badMsg) {
		fmt.Println("Bad message can be sent!")
	} else {
		fmt.Println("Bad message cannot be sent.")
	}
}

/*
package main

import (
	"fmt"
	"reflect"
)

type messageToSend struct {
	Message   string // Exported for broader reflection access if needed
	Sender    user   // Exported
	Recipient user   // Exported
}

type user struct {
	Name   string // Exported
	Number int    // Exported
}

func canSendMessage(mToSend messageToSend) bool {
	// Get the reflect.Type of the Message struct itself
	messageType := reflect.TypeOf(mToSend)
	// Get the reflect.Value of the Message struct instance
	messageValue := reflect.ValueOf(mToSend)

	// Get the reflect.Type of the User struct itself
	userType := reflect.TypeOf(user{})

	// --- Check Sender ---
	senderField, foundSender := messageType.FieldByName("Sender") // Note: "Sender" (capitalized)
	if !foundSender {
		fmt.Println("Error: 'Sender' field not found in messageToSend struct.")
		return false
	}
	if senderField.Type != userType { // Check if the field's type is indeed 'user'
		fmt.Printf("Error: 'Sender' field type (%s) does not match 'user' type (%s).\n", senderField.Type, userType)
		return false
	}

	// Get the actual Value of the Sender field
	senderValue := messageValue.FieldByName("Sender") // Note: "Sender" (capitalized)

	// Access the fields of the senderValue (which is a reflect.Value representing the 'user' struct)
	senderNameValue := senderValue.FieldByName("Name")     // Note: "Name" (capitalized)
	senderNumberValue := senderValue.FieldByName("Number") // Note: "Number" (capitalized)

	// Check if the Name and Number fields exist and are not their zero values
	if !senderNameValue.IsValid() || senderNameValue.String() == "" {
		fmt.Println("Error: Sender 'Name' field is invalid or empty.")
		return false
	}
	if !senderNumberValue.IsValid() || senderNumberValue.Int() == 0 {
		fmt.Println("Error: Sender 'Number' field is invalid or zero.")
		return false
	}

	// --- Check Recipient ---
	recipientField, foundRecipient := messageType.FieldByName("Recipient") // Note: "Recipient" (capitalized)
	if !foundRecipient {
		fmt.Println("Error: 'Recipient' field not found in messageToSend struct.")
		return false
	}
	if recipientField.Type != userType { // Check if the field's type is indeed 'user'
		fmt.Printf("Error: 'Recipient' field type (%s) does not match 'user' type (%s).\n", recipientField.Type, userType)
		return false
	}

	// Get the actual Value of the Recipient field
	recipientValue := messageValue.FieldByName("Recipient") // Note: "Recipient" (capitalized)

	// Access the fields of the recipientValue (which is a reflect.Value representing the 'user' struct)
	recipientNameValue := recipientValue.FieldByName("Name")     // Note: "Name" (capitalized)
	recipientNumberValue := recipientValue.FieldByName("Number") // Note: "Number" (capitalized)

	// Check if the Name and Number fields exist and are not their zero values
	if !recipientNameValue.IsValid() || recipientNameValue.String() == "" {
		fmt.Println("Error: Recipient 'Name' field is invalid or empty.")
		return false
	}
	if !recipientNumberValue.IsValid() || recipientNumberValue.Int() == 0 {
		fmt.Println("Error: Recipient 'Number' field is invalid or zero.")
		return false
	}

	return true // If all checks pass
}

func main() {
	tempMsg := messageToSend{
		Message:   "you have an appointment tomorrow",
		Sender:    user{Name: "Brenda Halafax", Number: 16545550987},
		Recipient: user{Name: "Sally Sue", Number: 19035558973},
	}

	if canSendMessage(tempMsg) {
		fmt.Println("Message can be sent!")
	} else {
		fmt.Println("Message cannot be sent.")
	}

	// Test with a problematic message
	badMsg := messageToSend{
		Message:   "test",
		Sender:    user{Name: "", Number: 0}, // Zero values
		Recipient: user{Name: "Valid", Number: 123},
	}
	if canSendMessage(badMsg) {
		fmt.Println("Bad message can be sent!")
	} else {
		fmt.Println("Bad message cannot be sent.")
	}
}
*/
