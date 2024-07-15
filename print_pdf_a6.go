package characters

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func PDFA6CharacterSheet(c Character, printKarmaLog, printCharacterLog, printPopularityLog bool) (pdf *gofpdf.Fpdf) {
	orientation := "P"
	sheetSize := "A6"
	pdf = gofpdf.New(orientation, "mm", sheetSize, ".")
	pdf.AddPage()

	// pdf.AddFont()
	// out = append(out, []byte(fmt.Sprintf("Name: %s\n", c.Name))...)

	return pdf
}

func attLinePDF6x4(att Rank, name string, padding int) []byte {
	attFmt := " %-*s %2s(%2d) [ %2d / %2d / %2d ]"
	green := att.Entry().Green
	yellow := att.Entry().Yellow
	red := att.Entry().Red
	attName := name + ":"
	return []byte(fmt.Sprintf(attFmt, padding, attName, att.Abbreviation(), att.Value(), green, yellow, red))
}

func powLinePDF6x4(p PowerEntry) []byte {
	attFmt := " - %s (%s) %s\n"
	return []byte(fmt.Sprintf(attFmt, p.Name, p.Reference, p.Description))
}

func valueLogPDF6x4(vl []ValueLog) (out []byte) {
	for _, l := range vl {
		out = append(out, fmt.Sprintf(" - [%3d] %s", l.Value, l.Notes)...)
		out = append(out, '\n')
	}
	return
}

func creationLogPDF6x4(vl []CreationLog) (out []byte) {
	for _, l := range vl {
		out = append(out, fmt.Sprintf(" - %s ", l.Message)...)
		for k, v := range l.Fields {
			out = append(out, fmt.Sprintf(" %s=%s", k, v)...)
		}
		out = append(out, '\n')
	}
	return
}
