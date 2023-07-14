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

func init() {
	rootCmd.AddCommand(kbotCmd)
}

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
		fmt.Println("Starting server on port 8080")

		go func() {
			err := http.ListenAndServe(":8080", handleRequests())
			if err != nil {
				log.Fatal(err)
			}
		}()
		fmt.Println("Started server on port 8080")

		fmt.Printf("kbot %s started\n", appVersion)

		TeleToken = strings.TrimSpace(string(getTokenBytes()))
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
	},
}

func handleRequests() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/liveness":
			currentTeleToken := strings.TrimSpace(string(getTokenBytes()))
			if currentTeleToken != TeleToken {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Printf("TeleToken: %s", TeleToken)
				fmt.Printf("currentTeleToken: %s", currentTeleToken)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		default:
			http.NotFound(w, r)
		}
	})
}

func getTokenBytes() []byte {
	tokenBytes, err := ioutil.ReadFile("/etc/app/secret/token")
	if err != nil {
		log.Println("Failed to read token file:", err)
	}
	return tokenBytes
}
