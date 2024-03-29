package cmd

import (
	"context"
	"fmt"
	"github.com/BlueMonday/go-scryfall"
	"log"
	"strings"
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

func scryExact(cardName string) scryfall.Card {
	ctx := context.Background()
	client, err := scryfall.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	sco := scryfall.GetCardByNameOptions{}

	result, err := client.GetCardByName(ctx, cardName, true, sco)
	if err != nil {
		log.Fatal(err)
	}

	return result

}

func scrySpecific(card MTGCard) []scryfall.Card {
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
	lastCharNumber := card.Number[len(card.Number)-1:]
	card.Number = strings.TrimSuffix(card.Number, "p")
	card.Number = strings.TrimSuffix(card.Number, "d")

	query := fmt.Sprintf("!\"%s\" e:%s is:%s number:%s",
		card.Name,
		card.Edition,
		strings.Replace(strings.ToLower(card.Finish), "-", "", -1),
		card.Number)

	if lastCharNumber == "p" {
		query += "st:promo"
	}

	result, err := client.SearchCards(ctx, query, sco)
	if err != nil {
		log.Println(err)
		log.Fatal(query)

		return nil
	}
	return result.Cards

}
