package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var costs [3]int
	for index := range messages {
		if index == 0 {
			costs[index] = len(messages[index])
		} else {
			costs[index] = costs[index-1] + len(messages[index])
		}
	}
	return messages, costs
}
