package characters

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func PDFA6CharacterSheet(ch Character, printKarmaLog, printCharacterLog, printPopularityLog bool, drawBorders bool) (pdf *gofpdf.Fpdf) {
	orientation := "P"
	sheetSize := "A6"
	pdf = gofpdf.New(orientation, "mm", sheetSize, ".")

	a6LayoutAddAttributes()

	myLayout := a6LayoutAddPowersAndTalents(ch)

	drawCharacter(pdf, ch, myLayout, drawBorders)

	return pdf
}

func drawCharacter(pdf *gofpdf.Fpdf, ch Character, l layout, drawBorders bool) {
	if ch.Image == "" {
		log.Fatal("character image string empty")
	}

	info := pdf.RegisterImage(ch.Image, "")
	if info == nil {
		log.Fatal("failed to register image")
	}

	log.Printf("image is [%.1f, %.1f]", info.Width(), info.Height())

	dc := ch.ToDisplayCharacter()
	for _, pageLayout := range l {
		pdf.AddPage()
		w, h := pdf.GetPageSize()
		fmt.Printf("page size: [%.1f , %.1f]", w, h)

		pageLayout.drawFixedText(pdf, drawBorders)
		for key, _ := range dc {
			value := dc[key]
			pageLayout.draw(pdf, layoutKey(key), value, drawBorders)
		}
	}
}

func a6LayoutAddPowersAndTalents(ch Character) (out layout) {

	out = deepCopy(a6Layout)
	page := layoutPage{}
	out = append(out, page)

	// dc := ch.ToDisplayCharacter()
	var (
		baseX = float64(1)
		baseY = float64(1)
		incY  = float64(7)
		col   *color
		align = "RM"
		// weight     = ""
		// fontSize   = float64(12)
		w = float64(16)
		h = float64(6)
		// border     = "B"
		// fontFamily = a6LayoutFont
		fontSize   = float64(12)
		fontWeight = "B"
	)
	c := &cell{
		x: baseX, y: baseY, w: w, h: h,
		align: align, border: "", fontFamily: a6LayoutFont,
		fontSize: fontSize, fontWeight: fontWeight, fillColor: col, fixedText: "Powers",
	}
	page["powers.label"] = c

	c = &cell{
		x: 0, y: baseY + incY, w: 105, h: 1,
		align: align, border: "B", fontFamily: a6LayoutFont,
		fontSize: fontSize, fontWeight: fontWeight, fillColor: col, fixedText: ".",
	}
	page["powers.hr"] = c

	baseY += incY + 3
	fontSize -= 2
	w += 40
	// powers.0.entry.desc: adds combat powers
	// powers.0.entry.name: Body Alterations - Offensive
	// powers.0.entry.ref: APB-86
	// powers.0.name: Claw and Bite
	// powers.0.rank.abvr: Rm
	// powers.0.rank.green: "36"
	// powers.0.rank.name: Remarkable
	// powers.0.rank.red: "95"
	// powers.0.rank.value: "30"
	// powers.0.rank.yellow: "66"

	incY = float64(9)
	y := baseY
	for i := range ch.Powers {
		if i > 0 {
			y += 9
		}
		c = &cell{
			x: baseX, y: y, w: w, h: h,
			align: align, border: "L", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: fontWeight, fillColor: col,
		}
		baseKey := fmt.Sprintf("powers.%d.", i)
		page[layoutKey(baseKey+"name")] = c
		cellWidth := float64(8)
		c = &cell{
			x: baseX + w + 1, y: y, w: cellWidth, h: h,
			align: align, border: "LB", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "", fillColor: col,
		}
		page[layoutKey(baseKey+"rank.abvr")] = c
		c = &cell{
			x: baseX + w + cellWidth + 1, y: y, w: cellWidth, h: h,
			align: align, border: "B", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "", fillColor: col,
		}
		page[layoutKey(baseKey+"rank.value")] = c

		c = &cell{
			x: baseX + w + 2*(cellWidth+1), y: y, w: cellWidth, h: h,
			align: align, border: "LB", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "I", fillColor: green,
		}
		page[layoutKey(baseKey+"rank.green")] = c

		c = &cell{
			x: baseX + w + 3*(cellWidth+1), y: y, w: cellWidth, h: h,
			align: align, border: "LB", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "I", fillColor: yellow,
		}
		page[layoutKey(baseKey+"rank.yellow")] = c

		c = &cell{
			x: baseX + w + 4*(cellWidth+1), y: y, w: cellWidth, h: h,
			align: align, border: "LB", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "I", fillColor: red,
		}
		page[layoutKey(baseKey+"rank.red")] = c

		// c = &cell{
		// 	x: baseX + w + 5*(cellWidth+1), y: y, w: w, h: h,
		// 	align: "LM", border: "", fontFamily: a6LayoutFont,
		// 	fontSize: fontSize, fontWeight: "I", fillColor: col,
		// }
		// page[layoutKey(baseKey+"entry.ref")] = c

		y += incY - 3
		c = &cell{
			x: 1, y: y, w: 87, h: 5,
			align: "RM", border: "BL", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "I", fillColor: col,
		}
		page[layoutKey(baseKey+"entry.desc")] = c

		c = &cell{
			x: 88, y: y, w: w, h: 5,
			align: "LM", border: "LB", fontFamily: a6LayoutFont,
			fontSize: fontSize, fontWeight: "I", fillColor: col,
		}
		page[layoutKey(baseKey+"entry.ref")] = c
		// if i == 1 {
		// 	page = layoutPage{}
		// 	out = append(out, page)
		// 	y = 1
		// }
	}
	y += incY
	c = &cell{
		x: 1, y: y, w: 103, h: h,
		align: "L", border: "", fontFamily: a6LayoutFont,
		fontSize: fontSize + 2, fontWeight: fontWeight, fillColor: col, fixedText: "Talents",
	}
	page["talents.label"] = c
	y += 6
	c = &cell{
		x: 0, y: y, w: 105, h: 1,
		align: align, border: "B", fontFamily: a6LayoutFont,
		fontSize: fontSize, fontWeight: fontWeight, fillColor: col, fixedText: ".",
	}
	page["talents.hr"] = c

	y += 3
	for i := range ch.Talents {
		baseKey := fmt.Sprintf("talents.%d.", i)
		c = &cell{
			x: 1, y: y, w: 103, h: h,
			align: "L", border: "L", fontFamily: a6LayoutFont,
			fontSize: fontSize - 2, fontWeight: fontWeight, fillColor: col,
		}
		page[layoutKey(baseKey+"name")] = c
		y += 4
		c = &cell{
			x: 1, y: y, w: 87, h: 5,
			align: "LM", border: "BL", fontFamily: a6LayoutFont,
			fontSize: fontSize - 4, fontWeight: "I", fillColor: col,
		}
		page[layoutKey(baseKey+"desc")] = c
		c = &cell{
			x: 88, y: y, w: w, h: 5,
			align: "LM", border: "LB", fontFamily: a6LayoutFont,
			fontSize: fontSize - 4, fontWeight: "I", fillColor: col,
		}
		page[layoutKey(baseKey+"ref")] = c
		y += 6
		// if y > 120 {
		// 	page = layoutPage{}
		// 	out = append(out, page)
		// }
		log.Println("Y", y, "talent", i, ch.Talents[i].Name)
	}

	fmt.Printf("layout:\n%s", out)
	return out
}
