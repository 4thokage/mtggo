package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

// manaCmd represents the mana calculation command
var manaCmd = &cobra.Command{
	Use:   "mana",
	Short: "Calculates deck mana needs",
	Long:  `Simple calculation of mana amounts to make the deck work`,
	Run: func(cmd *cobra.Command, args []string) {

		var wPips, bPips, rPips, uPips, gPips = 0, 0, 0, 0, 0
		if deckFile != "" {
			for _, element := range fromFile(deckFile, fileType) {
				cardColorMana := scryExact(element.Name).ManaCost
				cardCount, _ := strconv.Atoi(element.Count)

				wPips += strings.Count(cardColorMana, "W") * cardCount
				bPips += strings.Count(cardColorMana, "B") * cardCount
				rPips += strings.Count(cardColorMana, "R") * cardCount
				uPips += strings.Count(cardColorMana, "U") * cardCount
				gPips += strings.Count(cardColorMana, "G") * cardCount
			}

			if wPips != 0 {
				wPercent := float64(wPips) / deckSize
				log.Println("Plains: ", wPercent*landCount)
			}
			if bPips != 0 {
				bPercent := float64(bPips) / deckSize
				log.Println("Swamps: ", bPercent*landCount)
			}
			if rPips != 0 {
				rPercent := float64(rPips) / deckSize
				log.Println("Mountains: ", rPercent*landCount)
			}
			if uPips != 0 {
				uPercent := float64(uPips) / deckSize
				log.Println("Islands: ", uPercent*landCount)
			}
			if gPips != 0 {
				gPercent := float64(gPips) / deckSize
				log.Println("Forests :", gPercent*landCount)
			}

		}
	},
}

var landCount float64
var deckSize float64

func init() {
	manaCmd.Flags().Float64Var(&landCount, "lands", 24, "Land count")
	manaCmd.Flags().Float64Var(&deckSize, "size", 60, "Deck")
	rootCmd.AddCommand(manaCmd)
}
