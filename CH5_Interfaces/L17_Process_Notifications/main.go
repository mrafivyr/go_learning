package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (direct directMessage) importance() int {
	if direct.isUrgent {
		return 50
	}
	return direct.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (group groupMessage) importance() int {
	return group.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (system systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch messageType := n.(type) {
	case directMessage:
		return messageType.senderUsername, messageType.importance()
	case groupMessage:
		return messageType.groupName, messageType.importance()
	case systemAlert:
		return messageType.alertCode, messageType.importance()
	default:
		return "", 0
	}
}
