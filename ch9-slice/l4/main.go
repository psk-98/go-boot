package main

func getMessageCosts(messages []string) []float64 {
	messageCost := make([]float64, len(messages))

	for i, message := range messages {
		messageCost[i] = float64(len(message)) * 0.01
	}

	return messageCost
}
