package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Constants for the supported commands
const (
	RunCommand     = "run"
	ExplainCommand = "explain"
	AskCommand     = "ask"
)

const (
	CODE_MODEL    = "codestral"
	EXPLAIN_MODEL = "codellama"
	ASK_MODEL     = "llama3"
)

// Cli represents the command-line interface struct.
type Cli struct {
	DispatchedCommand string
	RunCommand        string
	AskCommand        string
	ExplainCommand    string
	ModelCommand      string
	ExplainModel      string // Default: ""
	CodeModel         string // Default: ""
	AskModel          string // Default: ""
}

// New initializes a new Cli instance with parsed command-line arguments.
func New() *Cli {
	cli := &Cli{}

	// Define flags
	flag.StringVar(&cli.RunCommand, "run", "", "Generate cli command based on given description")
	flag.StringVar(&cli.AskCommand, "ask", "", "Ask the llm any provided question string")
	flag.StringVar(&cli.ExplainCommand, "explain", "", "Ask the llm to explain a provided shell command")
	flag.StringVar(&cli.ModelCommand, "model", "", "Overwrite param for all models")
	flag.StringVar(&cli.ExplainModel, "explainmodel", EXPLAIN_MODEL, "Model used for cli command explainatin")
	flag.StringVar(&cli.CodeModel, "codemodel", CODE_MODEL, "Model used for code generating")
	flag.StringVar(&cli.AskModel, "askmodel", ASK_MODEL, "Model used to ask any question")

	// Parse flags
	flag.Parse()

	// Determine the dispatched command based on priority
	if cli.RunCommand != "" {
		cli.DispatchedCommand = RunCommand
	} else if cli.ExplainCommand != "" {
		cli.DispatchedCommand = ExplainCommand
	} else if cli.AskCommand != "" {
		cli.DispatchedCommand = AskCommand
	}

	if cli.DispatchedCommand == "" {
		cli.PrintHelp()
		os.Exit(0)
	}

	// If -model is provided, overwrite -explainmodel, -codemodel, -askmodel
	if cli.ModelCommand != "" {
		cli.ExplainModel = cli.ModelCommand
		cli.CodeModel = cli.ModelCommand
		cli.AskModel = cli.ModelCommand
	}

	return cli
}

func (cli *Cli) PrintHelp() {
	fmt.Println("CLI Tool Summary:")
	fmt.Println("This tool helps you create and understand shell commands based on your needs.")
	fmt.Println("It supports generating shell commands from descriptions, explaining shell commands,")
	fmt.Println("and asking questions using a Language Model (LLM) API.")
	fmt.Println()

	fmt.Println("Usage of this CLI tool:")
	fmt.Println()
	//fmt.Println("Flags:")
	//flag.PrintDefaults()
	//fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  -run string")
	fmt.Println("        Generate a shell command based on the given description")
	fmt.Println("  -ask string")
	fmt.Println("        Ask the LLM any provided question")
	fmt.Println("  -explain string")
	fmt.Println("        Ask the LLM to explain a provided shell command")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -model string")
	fmt.Println("        Overwrite param for all models")
	fmt.Println("  -explainmodel string")
	fmt.Printf("        Model used for shell command explanation (default '%s')\n", EXPLAIN_MODEL)
	fmt.Println("  -codemodel string")
	fmt.Printf("        Model used for code generation (default '%s')\n", CODE_MODEL)
	fmt.Println("  -askmodel string")
	fmt.Printf("        Model used for asking questions (default '%s')\n", ASK_MODEL)
	fmt.Println()
}

// PrintValues prints the current values of the cli instance.
func (cli *Cli) PrintValues() {
	fmt.Println("Dispatched Command:", cli.DispatchedCommand)
	fmt.Println("Run Command:", cli.RunCommand)
	fmt.Println("Ask Command:", cli.AskCommand)
	fmt.Println("Explain Command:", cli.ExplainCommand)
	fmt.Println("Model Command:", cli.ModelCommand)
	fmt.Println("Explain Model:", cli.ExplainModel)
	fmt.Println("Code Model:", cli.CodeModel)
	fmt.Println("Ask Model:", cli.AskModel)
}

// ComplexQuestion prompts the user with a question and possible answers, returning the chosen answer.
// It keeps asking until it gets a valid answer from the provided map.
func (c *Cli) ComplexQuestion(question string, answers map[string]string) string {
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
