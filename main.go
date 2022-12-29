package main

import "fmt"

var questions map[string]string = map[string]string{
	"question1": `In the famous netflix series, "Wednesday" who is the main character's real name?
	a. Jenna Ortega
	b. Joy sunday
	c. Moosa Mostafa
	`,
	"question2": `In the famous netflix series, "Wednesday" what is the name of Wednesday's sidekick?
	a. Enid
	b. Thing
	c. Tyler
	`,
	"question3": `In the famous netflix series, "Wednesday" what is the name of Wednesday's mother?
	a. Gwendolyn
	b. Morticia
	c. Martisha
	`,
}

var answers map[string]string = map[string]string{
	"question1": "a",
	"question2": "b",
	"question3": "b",
}

func main() {

	fmt.Println("Welcome To Top House Trivia!")
	for questionId, q := range questions {
		triviaQuestion(q, answers[questionId])
	}
}

func triviaQuestion(question string, answer string) {
	fmt.Println(question)
	// get user input and compare
	var userAnswer string
	fmt.Scanln(&userAnswer)
	if userAnswer == answer {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}
}


