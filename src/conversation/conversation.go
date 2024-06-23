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
