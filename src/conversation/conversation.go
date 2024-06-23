package conversation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Conversation struct {
	Argv1 string
}

// New creates a new instance of Conversation and initializes Argv1
func New() *Conversation {
	argv1 := ""
	if len(os.Args) > 1 {
		argv1 = os.Args[1]
	}
	return &Conversation{
		Argv1: argv1,
	}
}

// Question prompts the user with a yes or no question and returns true for yes and false for no.
// It keeps asking until it gets a valid yes or no answer.
func (c *Conversation) Question(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(question + " (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "yes" || response == "y" {
			return true
		} else if response == "no" || response == "n" {
			return false
		} else {
			fmt.Println("Please answer 'yes' or 'no'.")
		}
	}
}

// SimpleQuestion prompts the user with a yes or no question and returns true for yes and false for no.
// It keeps asking until it gets a valid yes or no answer.
func (c *Conversation) SimpleQuestion(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(question + " (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "yes" || response == "y" {
			return true
		} else if response == "no" || response == "n" {
			return false
		} else {
			fmt.Println("Please answer 'yes' or 'no'.")
		}
	}
}

// ComplexQuestion prompts the user with a question and possible answers, returning the chosen answer.
// It keeps asking until it gets a valid answer from the provided map.
func (c *Conversation) ComplexQuestion(question string, answers map[string]string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(question)
		for key, explanation := range answers {
			fmt.Printf("- %s ( %s )\n", key, explanation)
		}
		fmt.Print("Your answer: ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if explanation, ok := answers[response]; ok {
			fmt.Printf("You chose '%s' (%s)\n", response, explanation)
			return response
		} else {
			fmt.Println("Invalid answer. Please choose a valid option.")
		}
	}
}
