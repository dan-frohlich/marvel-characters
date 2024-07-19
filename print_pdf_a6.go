package characters

import (
	"github.com/jung-kurt/gofpdf"
)

func PDFA6CharacterSheet(ch Character, printKarmaLog, printCharacterLog, printPopularityLog bool, drawBorders bool) (pdf *gofpdf.Fpdf) {
	orientation := "P"
	sheetSize := "A6"
	pdf = gofpdf.New(orientation, "mm", sheetSize, ".")

	a6LayoutAddAttributes()

	drawCharacter(pdf, ch, []layoutPage{a6Layout}, drawBorders)

	return pdf
}

func drawCharacter(pdf *gofpdf.Fpdf, ch Character, l layout, drawBorders bool) {
	dc := ch.ToDisplayCharacter()
	for _, pageLayout := range l {
		pdf.AddPage()
		pageLayout.drawFixedText(pdf, drawBorders)
		for key, _ := range dc {
			value := dc[key]
			pageLayout.draw(pdf, layoutKey(key), value, drawBorders)
		}
	}
}
