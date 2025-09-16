package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(analytics *Analytics, message Message) {
	analytics.MessagesTotal++
	if message.Success {
		analytics.MessagesSucceeded++
	} else {
		analytics.MessagesFailed++
	}
}
