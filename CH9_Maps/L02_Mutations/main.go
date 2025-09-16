package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	if user, exists := users[name]; exists {
		if user.scheduledForDeletion {
			delete(users, name)
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("not found")
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
