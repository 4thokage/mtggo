package cmd

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

// proxyCmd represents the price command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Generates a pdf with card proxy images",
	Long:  `Generates a pdf with card proxy images`,
	Run: func(cmd *cobra.Command, args []string) {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		var name = strings.Join(args, "_")
		if deckFile != "" {
			var extension = filepath.Ext(deckFile)
			name = deckFile[0 : len(deckFile)-len(extension)]
		} else {
			cardCounter := 0
			rowNum := float64(0)
			for index, element := range args {

				if index != 0 && index%3 == 0 {
					rowNum++
					cardCounter = 0
				}
				if index != 0 && index%9 == 0 {
					rowNum = 0
					cardCounter = 0
					pdf.AddPage()
				}
				y := rowNum * 88.9
				x := float64(cardCounter) * 63.5
				url := scry(element)[0].ImageURIs.Normal
				httpimg.Register(pdf, url, "")
				pdf.ImageOptions(
					url,
					x, y,
					63.5, 88.9,
					false,
					gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
					0,
					url,
				)

				cardCounter++

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
