package characters

import (
	"fmt"
	"log"
	"sort"

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
	image        bool
}

func (c cell) String() string {
	return fmt.Sprintf(`[x:%.1f y:%.1f w:%.1f h:%.1f align:%s border:%s fixed:%s image: %v]`, c.x, c.y, c.w, c.h, c.align, c.border, c.fixedText, c.image)
}

func (l layoutPage) String() (out string) {
	keys := make([]string, 0, len(l))
	for k := range l {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := l[layoutKey(k)]
		out += fmt.Sprintf("%s -> %s\n", k, v)
	}
	return out
}

func (l layout) String() (out string) {
	for i, p := range l {
		out += fmt.Sprintf("-%d----------\n", i)
		out += p.String()
	}
	return out
}

func (l layoutPage) drawFixedText(pdf *gofpdf.Fpdf, border bool) {
	for _, v := range l {
		if v.fixedText != "" {
			if v.image {
				continue
			}
			v.drawCell(pdf, v.fixedText, border)
		}
	}
}

func (l layoutPage) draw(pdf *gofpdf.Fpdf, key layoutKey, text string, border bool) bool {
	c, ok := l[key]
	if ok {
		if c.image {
			log.Printf("found image %s: %s", key, text)
		}
		c.drawCell(pdf, text, border)
		// } else {
		// 	log.Printf("skipping %s: no cell found", key)
	}
	return ok
}

func (c *cell) drawCell(pdf *gofpdf.Fpdf, text string, border bool) {
	pdf.SetXY(c.x, c.y)
	if c.fillColor != nil && !c.image {
		r, g, b := pdf.GetFillColor()
		defer pdf.SetFillColor(r, g, b) //reset wehen done
		pdf.SetFillColor(c.fillColor.r, c.fillColor.g, c.fillColor.b)
	}

	brdr := c.border
	if border {
		brdr = "1"
	}

	if c.image {
		log.Printf("found image cell with url len %d : %s", len(text), text)
		info := pdf.GetImageInfo(text)
		if info != nil {
			w, h := info.Extent()
			scale := c.w / w
			if h > w {
				//scale by height
				scale = c.h / h
			}
			log.Printf("image cell: [%.1f, %.1f]: %.2f", w, h, scale)
			w *= scale
			h *= scale
			log.Printf("drawing %s at [%.1f,%.1f,%.1f,%.1f]", "image", c.x, c.y, w, h)

			//centering the image....
			dx := float64(0)
			dy := float64(0)
			if c.h > h {
				dy = (c.h - h)
			}
			if c.w > w {
				dx = (c.w - w) / 2
			}

			// pdf.CellFormat(c.w, c.h, "", "1", 2, "CM", false, -1, "")
			pdf.ImageOptions(text, c.x+dx, c.y+dy, w, h, false, gofpdf.ImageOptions{}, 0, "")
			return
		} else {
			log.Println("image had no info!")
		}
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
