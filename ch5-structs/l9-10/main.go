package main

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	} else {
		return "", false
	}
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {

	if membershipType == "premium" {
		return User{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: 1000}}
	}
	return User{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: 100}}
}
