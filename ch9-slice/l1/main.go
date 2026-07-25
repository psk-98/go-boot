package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var costs [3]int
	costs[0] = len(primary)
	costs[1] = len(primary) + len(secondary)
	costs[2] = len(primary) + len(secondary) + len(tertiary)

	return messages, costs
}
