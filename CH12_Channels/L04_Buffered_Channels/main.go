package main

/*
func addEmailsToQueue(emails []string) chan string {
	mailChan := make(chan string, len(emails))
	for _, email := range emails {
		mailChan <- email
	}
	return mailChan
}
*/

func addEmailsToQueue(emails []string) chan string {
	if len(emails) == 0 {
		return nil // Or return a closed channel
	}

	mailChan := make(chan string, len(emails))
	go func() {
		defer close(mailChan)
		for _, email := range emails {
			mailChan <- email
		}
	}()
	return mailChan
}
