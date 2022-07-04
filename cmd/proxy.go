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

// proxyCmd represents the price command
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
			var extension = filepath.Ext(deckFile)
			name = deckFile[0 : len(deckFile)-len(extension)]
		} else {

			var counter = 0
			for _, element := range args {

				url := scry(element)[0].ImageURIs.Normal
				httpimg.Register(pdf, url, "")
				pdf.ImageOptions(
					url,
					cardGrid[counter].x, cardGrid[counter].y,
					63.5, 88.9,
					false,
					gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
					0,
					url,
				)

				counter++
				if counter == 9 {
					counter = 0
					pdf.AddPage()
				}

			}
		}

		err := pdf.OutputFileAndClose(name + "_PROXIES.pdf")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
