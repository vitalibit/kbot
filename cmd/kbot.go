/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var TeleToken string

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)

		tokenBytes, err := ioutil.ReadFile("/etc/app/secret/token")
		if err != nil {
			log.Println("Failed to read token file:", err)
			return
		}

		TeleToken = strings.TrimSpace(string(tokenBytes))
		fmt.Printf("Token: %s, Length: %d, Type: %T\n", TeleToken, len(TeleToken), TeleToken)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		fmt.Println("Bot created.")

		if err != nil {
			log.Println("Failed to create bot:", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {

			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm KBot %s!", appVersion))
			}

			return err

		})

		kbot.Start()

		// Check if token has changed
		currentSecretValue := TeleToken

		log.Println("Starting server on port 8080")

		http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Handling liveness probe request")
			updatedSecretValue := strings.TrimSpace(string(tokenBytes))
			if currentSecretValue != updatedSecretValue {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Token has not changed."))
				log.Println("Token has not changed.")
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Token has changed."))
				log.Println("Token has changed.")
			}
		})

		go func() {
			http.ListenAndServe(":8080", nil)
		}()
		log.Println("End of code")
		fmt.Println("End of code")
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
