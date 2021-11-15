package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// manaCmd represents the price command
var manaCmd = &cobra.Command{
	Use:   "mana",
	Short: "Calculates deck mana needs",
	Long:  `Simple calculation of mana amounts to make the deck work`,
	Run: func(cmd *cobra.Command, args []string) {

		if deckFile != "" {
			cardNamesFromFile(deckFile)
		} else {
			for _, element := range args {
				cardColorMana := scry(element)[0].ManaCost
				log.Println(cardColorMana)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(manaCmd)
}
