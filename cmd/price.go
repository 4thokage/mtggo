package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Fetches card prices",
	Long:  `Gets card/ deck list price information from the Scryfall API`,
	Run: func(cmd *cobra.Command, args []string) {

		if deckFile != "" {
			cards := fromFile(deckFile, fileType)
			var totalPriceEur float64

			// if card list is bigger than 100 cards (cube or collection) use scryfall bulk data instead
			if len(cards) > 100 {
				log.Println(getPriceFromBulkData(cards))
				return
			}
			for _, element := range cards {
				for _, card := range scrySpecific(element) {
					priceEur, _ := strconv.ParseFloat(card.Prices.EUR, 64)
					if element.Status != "Proxied" {
						totalPriceEur += priceEur
					}
				}
			}
			log.Println(totalPriceEur)
		} else {
			for _, element := range args {
				for _, card := range scry(element) {
					info := fmt.Sprintf("%s%s - (%s) %s: [EUR: %s, USD: %s, TIX: %s]", card.Name, card.ManaCost, card.SetName, strings.Replace(strings.Trim(fmt.Sprint(card.Finishes), "[]"), " ", ",", -1), card.Prices.EUR, card.Prices.USD, card.Prices.Tix)
					log.Println(info)
				}

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}
