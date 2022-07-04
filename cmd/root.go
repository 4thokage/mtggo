package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var deckFile string

var rootCmd = &cobra.Command{
	Use:   "mtggo",
	Short: "Converts cards/deck list from/to another formats",
	Long:  `Simple CLI to help with MTG card management`,
	Run: func(cmd *cobra.Command, args []string) {
		if deckFile != "" {

		} else {
			log.Println("Command ignored. No deck file passed")
		}
	},
}

// Execute will terminate the app if any command returns an error
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&deckFile, "deck", "d", "", "Deck list file")

}
