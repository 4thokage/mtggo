package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// manaCmd represents the price command
var manaCmd = &cobra.Command{
	Use:   "mana",
	Short: "Calculates deck mana needs",
	Long:  `Simple calculation of mana amounts to make the deck work`,
	Run: func(cmd *cobra.Command, args []string) {

		var landCount = cmd.Flag("lands").Value
		var wPips, bPips, rPips, uPips, gPips = 0, 0, 0, 0, 0
		if deckFile != "" {
			cardNamesFromFile(deckFile)
		} else {
			for _, element := range args {
				cardColorMana := scry(element)[0].ManaCost
				wPips += strings.Count(cardColorMana, "W")
				bPips += strings.Count(cardColorMana, "B")
				rPips += strings.Count(cardColorMana, "R")
				uPips += strings.Count(cardColorMana, "U")
				gPips += strings.Count(cardColorMana, "G")
			}

			info := fmt.Sprintf("Mana source count consistency: [W:%d, B:%d, R:%d, U:%d, G:%d] - [Land count: %d]", wPips/2, bPips/2, rPips/2, uPips/2, gPips/2, landCount)
			log.Println(info)
		}
	},
}

func init() {
	manaCmd.Flags().IntP("lands", "l", 24, "Land count")
	rootCmd.AddCommand(manaCmd)
}
