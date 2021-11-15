package cmd

import (
	"context"
	"github.com/BlueMonday/go-scryfall"
	"log"
)

func scry(cardName string) []scryfall.Card {
	ctx := context.Background()
	client, err := scryfall.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	sco := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrints,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: false,
	}

	result, err := client.SearchCards(ctx, cardName, sco)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return result.Cards

}
