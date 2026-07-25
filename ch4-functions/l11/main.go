package main

func reformat(message string, formatter func(string) string) string {
	firstForm := formatter(message)
	secondForm := formatter(firstForm)
	thirdForm := formatter(secondForm)

	return "TEXTIO: " + thirdForm

}
