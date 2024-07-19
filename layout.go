package characters

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

type layoutKey string

type layout []layoutPage

type layoutPage map[layoutKey]*cell

type color struct {
	r, g, b int
}
type cell struct {
	x, y, w, h   float64
	align        string
	border       string
	fontFamily   string
	fontSize     float64
	fontWeight   string
	overflowRows int
	fillColor    *color
	fixedText    string
}

func (l layoutPage) drawFixedText(pdf *gofpdf.Fpdf, border bool) {
	for _, v := range l {
		if v.fixedText != "" {
			v.drawCell(pdf, v.fixedText, border)
		}
	}
}

func (l layoutPage) draw(pdf *gofpdf.Fpdf, key layoutKey, text string, border bool) bool {
	c, ok := l[key]
	if ok {
		c.drawCell(pdf, text, border)
	} else {
		log.Printf("skipping %s: no cell found", key)
	}
	return ok
}

func (c *cell) drawCell(pdf *gofpdf.Fpdf, text string, border bool) {
	pdf.SetXY(c.x, c.y)
	if c.fillColor != nil {
		r, g, b := pdf.GetFillColor()
		defer pdf.SetFillColor(r, g, b) //reset wehen done
		pdf.SetFillColor(c.fillColor.r, c.fillColor.g, c.fillColor.b)
	}

	brdr := c.border
	if border {
		brdr = "1"
	}
	pdf.SetFont(c.fontFamily, c.fontWeight, c.fontSize)
	lines := pdf.SplitText(text, c.w)
	rows := c.overflowRows + 1
	if len(lines) > rows {
		log.Printf("WARN cell needs [%d] lines to handle text [%s]", len(lines), text)
	}
	for i, line := range lines {
		if i == rows {
			break
		}
		if i != 0 {
			line = " " + line
		}
		pdf.CellFormat(c.w, c.h, line, brdr, 2, c.align, c.fillColor != nil, -1, "")
		//pdf.Cell(c.w, c.h, line)
	}
}

// func loadBkgrndImg(pdf *gofpdf.Fpdf, useLargeBackground bool) error {

// 	url := smallImageDataURL
// 	name := smImgNm
// 	imgType := smImgType
// 	if useLargeBackground {
// 		url = largeImageDataURL
// 		name = lgImgNm
// 		imgType = lgImgType
// 	}
// 	r, err := dataURLReader(url)
// 	if err != nil {
// 		return err
// 	}
// 	pdf.RegisterImageReader(name, imgType, r)
// 	return nil
// }
