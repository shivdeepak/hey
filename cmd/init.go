/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		possibleLlms := [4]string{"1", "2", "3", "4"}

		llmSelection := ""

		for llmSelection == "" {
			fmt.Println("Which LLM would you like to configure?")
			fmt.Println("[1] ChatGPT from OpenAI")
			fmt.Println("[2] Perplexity")
			fmt.Println("[3] Pi from Inflection.ai")
			fmt.Println("[4] Claude from Anthropic")
			fmt.Printf("Make a selection [Default 1]: ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				input = "1"
			}

			for _, possibleLlm := range possibleLlms {
				if possibleLlm == input {
					llmSelection = input
				}
			}

			if llmSelection == "" {
				fmt.Println("\nInvalid input, please enter 1-4")
			}
		}

		if llmSelection == "1" {
			fmt.Println("You have selected ChatGPT from OpenAI")
			openAiApiKey := ""

			for openAiApiKey == "" {
				fmt.Println("We need your OpenAI API Key to make requests to ChatGPT")
				fmt.Println("This API Key will be stored locally on your machine")
				fmt.Println("Your OpenAI API Key will have following format: sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				fmt.Println("You can find your API key at https://platform.openai.com/account/api-keys")
				fmt.Printf("Enter Your API Key: ")
				reader := bufio.NewReader(os.Stdin)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.HasPrefix(input, "sk-") && len(input) >= 50 {
					openAiApiKey = input
				} else {
					fmt.Println("\nInvalid input, please enter a valid API Key")
				}
			}

			initViperConfigFile()

			viper.Set("llm.default", "chatgpt")
			viper.Set("llm.chatgpt.api_key", openAiApiKey)
			saveCurrentViperConfig()
		}

	},
}

func init() {
	configCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
