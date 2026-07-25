package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	returnMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i, name := range names {
		returnMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i]}
	}

	return returnMap, nil

}

type user struct {
	name        string
	phoneNumber int
}
