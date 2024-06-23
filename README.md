# clibuddy

clibuddy is a command-line tool written in Golang that helps you to create complex shell commands by textual description or analyze complex shell commands using a Language Model (LLM) provided by ollama. It is currently in its Alpha state.

## Features

- Generate shell commands based on descriptions.
- Explain shell commands using LLM.
- Ask questions and receive answers using LLM.

## Privacy Assurance

clibuddy operates entirely on your local machine, utilizing the ollama API to ensure that your privacy is maintained. No information is exposed to online or company APIs, keeping your data secure and private. This tool is completely free and relies solely on your local hardware, ensuring that you have full control over your environment and data.

## Installation

### Prerequisites

- Golang installed on your machine.
- Local [ollama](https://ollama.com/download) API running.
- Models `codestral`, `codellama`, and `llama3` available locally.

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/cli.git
   cd cli/cmd/client
   ```

2. Build the executable:
   ```bash
   go build -o clibuddy
   ```

3. Copy it into a PATH registered location: (might differ on your system)
   ```bash
   cp clibuddy ~/go/bin
   ```

## Usage

Run the executable with the following options:

```bash
clibuddy -[command] [commandparam] [-[option] [optionvalue]]
```

### Commands:
At least one command must be given. On no command given help is printed.
- `-run string`: Generate a shell command based on the given description.
- `-ask string`: Ask the LLM any provided question.
- `-explain string`: Ask the LLM to explain a provided shell command.

### Options
You can use any number of options. (optional)
- `-model string`: Overwrite parameter for all models.
- `-explainmodel string`: Model used for shell command explanation (default 'codellama').
- `-codemodel string`: Model used for code generation (default 'codestral').
- `-askmodel string`: Model used for asking questions (default 'llama3').

### Examples

#### Generate a Shell Command:
To generate a shell command you can use clibuddy like :
```bash
clibuddy -run "Search for files with '.php' extension and count lines"
```
The -run command is an interactive command. Since there is always a certain danger about an llm creating and running command by itself, the clibuddy will first print the command and give you the option to decide how to go on. You can either choose "run" which will execute the command, "explain" which will ask the llm for an detailed explaination of the command and print it, or "exit" which - well will exit the application. Such a response to the example command would be: 
Will result in an answer like ''
```
The generated command is:  find . -name '*.php' -type f -exec wc -l {} +
How you want to proceed?
- explain ( Explains the generated command )
- run ( Executes the generated command )
- exit ( End the clibuddy conversation )
Your answer: 
```
If the resulting command is to complex, you simply respond with "explain" and get an output like this: 

```
You chose 'explain' (Explains the generated command)
The `find` command is used to search for files in a directory tree. In this case, it searches for files with the `.php` extension in the current directory and its subdirectories. The `-type f` option tells `find` to only consider regular files (i.e., not directories or symbolic links).

The `-exec` option allows you to execute a command on each file that is found. In this case, the command is `wc -l {} +`, which counts the number of lines in each file and outputs the result. The `{}` placeholder is replaced with the name of the current file for each iteration of the `-exec` option.

The `+` at the end of the command tells `find` to execute the command once per group of files that match the search criteria, rather than once per file. This can be useful when you want to save resources by running a command only once on multiple files, rather than having it run separately for each file.
How you want to proceed?
- explain ( Explains the generated command )
- run ( Executes the generated command )
- exit ( End the clibuddy conversation )
Your answer: 
```
As we can see the llm explained in detail how this command works. Since we think this is fine we can execute the command using "run"
```
You chose 'run' (Executes the generated command)
Command executed successfully:
  55 ./index.php
   2 ./some/directory/php.php
  57 total

How you want to proceed?
- explain ( Explains the generated command )
- run ( Executes the generated command )
- exit ( End the clibuddy conversation )
Your answer:
```
At this point we successful used clibuddy to generate a cli command based on a short description, let the llm explain in detail what it actually does, and execute it.

#### Explain a Shell Command:
```bash
clibuddy -explain "find . -type f -name '*.php' -exec wc -l {} \;"
```
Using the -explain command, an LLM will be prompted with an enhanced prompt to get a proper explanation of how the provided cli command works. The example provided command might return an answer like
```
This command uses the `find` utility to search for files with a `.php` extension in the current directory and its subdirectories. The `-type f` option tells `find` to only consider files (and not directories), and the `-name '*.php'` option restricts the search to files whose name ends with `.php`.

The `-exec` option tells `find` to execute a command on each file that matches the search pattern. In this case, the command is `wc -l {} \;`, which invokes the `wc` utility to count the number of lines in each matching file. The `{}` placeholder represents the current file being processed by `find`.

The resulting output will be a list of files with their corresponding line counts. For example:

$ find . -type f -name '*.php' -exec wc -l {} \;
./file1.php: 50
./file2.php: 100
./subdirectory/file3.php: 200

This command is useful for quickly checking the number of lines in a large number of PHP files, or for finding files with a high line count that may indicate issues such as code complexity or large functions.
```
This function should help people that try to dive into cli usage. Sometimes example commands can be quite complex, and using the -explain method you can step by step understand what a command would do, and based on this decide if you really want to execute it on your system. The -run command also implements and interactive option the result of a generated command before running it.

#### Ask a Question:
```bash
clibuddy -ask "What are the 10 biggest capitals by inhabitants?"
```
Well, since we already have implemented a simple cli interface and the necessary stuff to requrest a local ollama instance, why not add a method to simply send a "question" to an llm and print the answer. This is just as "goody" since were already there. The answer for this one might be like: 
```
Here are the top 10 most populous capital cities in the world, based on the latest available estimates:

1. **Tokyo, Japan**: approximately 38,140,000 people
2. **Delhi, India**: approximately 29,280,000 people
3. **Sao Paulo, Brazil**: approximately 21,310,000 people
4. **Mexico City, Mexico**: approximately 21,180,000 people
5. **Osaka, Japan**: approximately 20,310,000 people (note: Osaka is often considered a part of the larger Keihanshin metropolitan area, which includes Kobe and Kyoto)
6. **Cairo, Egypt**: approximately 20,450,000 people
7. **Beijing, China**: approximately 21,540,000 people
8. **Dhaka, Bangladesh**: approximately 19,440,000 people
9. **Shanghai, China**: approximately 18,580,000 people (note: Shanghai is often considered a separate entity from the provincial capital Nanjing)
10. **Lagos, Nigeria**: approximately 17,780,000 people

Please note that these numbers are estimates and may vary depending on the source and methodology used. Additionally, it's worth noting that some of these cities may be part of larger metropolitan areas or urban agglomerations, which can affect the population figures.
```

### Known issues

Due to the way Go executes the generated commands on the system, some commands might not work properly. I plan to update the execution method to eliminate these issues. For now, some commands may produce errors when executed through the tool, but they should still work if you copy and paste them directly into your shell and run them natively.


### Environment Information

During command generation (`-run` mode), clibuddy collects environment information. See [EnvInfo](src/envinfo/envinfo.go) for details. This information is used to achieve a better quality of generated cli commands. The data is only used locally.

### Disclaimer

While clibuddy can be a useful tool for generating and explaining shell commands using LLMs, please exercise caution when executing any commands generated by the tool on your machine. There is always some inherent risk associated with running commands that are generated by automated systems. Always review and understand the commands before executing them to ensure they do not pose any security or stability risks to your system.

## Contributing

Feel free to contribute by opening issues or pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the  Apache License Version 2.0 - see the [LICENSE](LICENSE) file for details.
```