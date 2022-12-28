package cmd

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"time"
)

type coordinate struct {
	x float64
	y float64
}

var cardGrid = []coordinate{
	{
		x: 0,
		y: 0,
	},
	{
		x: 0,
		y: 88.9,
	},
	{
		x: 0,
		y: 177.8,
	},
	{
		x: 63.5,
		y: 0,
	},
	{
		x: 63.5,
		y: 88.9,
	},
	{
		x: 63.5,
		y: 177.8,
	},
	{
		x: 127,
		y: 0,
	},
	{
		x: 127,
		y: 88.9,
	},
	{
		x: 127,
		y: 177.8,
	},
}

var counter = 0

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Generates a pdf with card proxy images",
	Long:  `Generates a pdf with card proxy images`,
	Run: func(cmd *cobra.Command, args []string) {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		var name = time.Now().Format("01-01-1998")
		log.Println(name)
		if deckFile != "" {

			name = deckFile[0 : len(deckFile)-len(filepath.Ext(deckFile))]

			cards := fromFile(deckFile, fileType)

			for _, element := range cards {

				card := scryExact(element.Name)
				url := card.ImageURIs

				if len(card.CardFaces) > 0 {
					for _, face := range card.CardFaces {
						addCardImage(pdf, face.ImageURIs.Normal)
					}

				} else {
					if url == nil {
						continue
					}
					addCardImage(pdf, url.Normal)
				}

			}
			err := pdf.OutputFileAndClose(name + "_PROXIES.pdf")
			if err != nil {
				panic(err)
			}

		} else {

			for _, element := range args {

				url := scry(element)[0].ImageURIs.Normal
				addCardImage(pdf, url)

			}
			err := pdf.OutputFileAndClose(name + "_PROXIES.pdf")
			if err != nil {
				panic(err)
			}
		}

	},
}

func addCardImage(pdf *gofpdf.Fpdf, urlNormal string) {
	httpimg.Register(pdf, urlNormal, "")
	pdf.ImageOptions(
		urlNormal,
		cardGrid[counter].x, cardGrid[counter].y,
		63.5, 88.9,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		urlNormal,
	)

	counter++
	if counter == 9 {
		counter = 0
		pdf.AddPage()
	}
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
