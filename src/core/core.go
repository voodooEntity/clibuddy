package core

import (
	"fmt"
	"github.com/voodooEntity/go-clibuddy/src/cli"
	"github.com/voodooEntity/go-clibuddy/src/envinfo"
	"github.com/voodooEntity/go-clibuddy/src/executor"
	"github.com/voodooEntity/go-clibuddy/src/ollamapi"
	"github.com/voodooEntity/go-clibuddy/src/promptbuilder"
	"os"
)

type Core struct {
	Api           *ollamapi.OllamApi
	Cli           *cli.Cli
	Environment   *envinfo.EnvInfo
	PromptBuilder *promptbuilder.PromptBuilder
	Executor      *executor.Executor
}

func New() *Core {
	ei := envinfo.New()

	c := Core{
		Api:           ollamapi.New("http://localhost:11434/api/generate"),
		Cli:           cli.New(),
		Environment:   ei,
		PromptBuilder: promptbuilder.New(ei),
		Executor:      executor.New(ei),
	}

	return &c
}

func (c *Core) Execute() {

	switch c.Cli.DispatchedCommand {
	case cli.RunCommand:
		c.Generate(c.Cli.RunCommand)
	case cli.ExplainCommand:
		c.ExplainCommand(c.Cli.ExplainCommand)
	case cli.AskCommand:
		c.Ask(c.Cli.AskCommand)
	}

}

func (c *Core) Generate(command string) {
	prompt := c.PromptBuilder.BuildCommandPrompt(c.Cli.RunCommand)

	// Ask the question using the API
	command, err := c.Api.Ask(c.Cli.CodeModel, prompt)
	if err != nil {
		fmt.Println("Error asking question:", err)
		return
	}
	fmt.Println("The generated command is:", command)

	options := map[string]string{"explain": "Explains the generated command", "run": "Executes the generated command", "exit": "End the clibuddy conversation"}
	for {
		action := c.Cli.ComplexQuestion("How you want to proceed?", options)
		switch action {
		case "explain":
			c.ExplainCommand(command)
		case "run":
			result, err := c.Executor.Do(command)
			if err != nil {
				fmt.Println("Error while executing command - exiting:", err)
				os.Exit(0)
			}
			fmt.Println("Command executed successfully:")
			fmt.Println(result)
		case "exit":
			fmt.Println("\n")
			fmt.Println("So long and thanks for all the fish!")
			os.Exit(0)
		}
	}
}

func (c *Core) ExplainCommand(command string) {
	explainPrompt := c.PromptBuilder.BuildExplanationPrompt(command)
	explaination, err := c.Api.Ask(c.Cli.ExplainModel, explainPrompt)
	if err != nil {
		fmt.Println("Error while trying to ask the llm - exiting:", err)
		os.Exit(0)
	}
	fmt.Println(explaination)
}

func (c *Core) Ask(question string) {
	answer, err := c.Api.Ask(c.Cli.AskModel, question)
	if err != nil {
		fmt.Println("Error while trying to ask the llm - exiting:", err)
		os.Exit(0)
	}
	fmt.Println(answer)
}
