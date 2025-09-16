package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user, len(names))
	for index, name := range names {
		userMap[name] = user{name: name, phoneNumber: phoneNumbers[index]}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
