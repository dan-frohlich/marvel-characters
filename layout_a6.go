package characters

import "fmt"

const (
	a6LayoutFont = "Times"
)

var (
	cornflowerBlue = &color{r: 100, g: 149, b: 237}
	green          = &color{r: 127, g: 254, b: 126}
	yellow         = &color{r: 251, g: 254, b: 126}
	red            = &color{r: 255, g: 122, b: 94}
)

var a6Layout = layout{
	layoutPage{
		"name": &cell{
			x: 1, y: 1, w: 103, h: 6,
			align:      "LM",
			fontFamily: a6LayoutFont,
			fontSize:   16,
			fontWeight: "B",
			border:     "B",
		},
		// 	   Health:  116
		//  	Karma:   52
		//  Resources:    6
		// Popularity:    0
		//   	 Move:    3 areas
		// Initiative:    0

		"health.label": &cell{
			x: 72, y: 10, w: 22, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "B",
			border:     "LB",
			fixedText:  "Health",
		},
		"health": &cell{
			x: 95, y: 10, w: 8, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "B",
		},
		"move.label": &cell{
			x: 72, y: 17, w: 22, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "B",
			border:     "LB",
			fixedText:  "Move (area)",
		},
		"move": &cell{
			x: 95, y: 17, w: 8, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "B",
		},
		"initiative.label": &cell{
			x: 72, y: 24, w: 22, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "B",
			border:     "LB",
			fixedText:  "Initiative",
		},
		"initiative": &cell{
			x: 95, y: 24, w: 8, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "B",
		},
		"karma.label": &cell{
			x: 72, y: 31, w: 22, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "B",
			border:     "LB",
			fixedText:  "Karma",
		},
		"karma": &cell{
			x: 95, y: 31, w: 8, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "B",
		},
		"popularity.label": &cell{
			x: 72, y: 38, w: 22, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "B",
			border:     "LB",
			fixedText:  "Popularity",
		},
		"popularity": &cell{
			x: 95, y: 38, w: 8, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "B",
		},
		"resources.label": &cell{
			x: 72, y: 45, w: 22, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "B",
			border:     "LB",
			fixedText:  "Resources",
		},
		"resources": &cell{
			x: 95, y: 45, w: 8, h: 6,
			align:      "RM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "B",
		},
		"image": &cell{
			x: 1, y: 60, w: 103, h: 68,
			align:      "CM",
			fontFamily: a6LayoutFont,
			fontSize:   10,
			fontWeight: "",
			border:     "LRTB",
			fillColor:  cornflowerBlue,
			fixedText:  "character image",
			image:      true,
		},
	},
}

func a6LayoutAddAttributes() {
	xOrigin := float64(1)
	x := xOrigin
	y := float64(10)
	yInc := float64(7)
	for i := 0; i < 8; i++ {
		for _, v := range []string{"name", "rank.abvr", "rank.value", "rank.green", "rank.yellow", "rank.red"} {
			key := layoutKey(fmt.Sprintf("attribute.%d.%s", i, v))
			var (
				col      *color
				align    = "RM"
				weight   = ""
				fontSize = float64(12)
				w        = float64(8)
				h        = float64(6)
				border   = "B"
			)

			switch v {
			case "name":
				weight = "B"
				align = "LM"
				col = nil
				fontSize = 10
				w = 20
				border = "LB"
			case "rank.abvr":
				fontSize = 10
				align = "CM"
				border = "LB"
			case "rank.green":
				border = "LB"
				col = green
				weight = "I"
			case "rank.yellow":
				border = "LB"
				col = yellow
				weight = "I"
			case "rank.red":
				border = "LB"
				col = red
				weight = "I"
			}
			c := &cell{
				x: x, y: y, w: w, h: h,
				align: align, fontFamily: a6LayoutFont, fontSize: fontSize, fontWeight: weight,
				fillColor: col,
				border:    border,
			}
			a6Layout[0][key] = c
			x += w + 1
			if "rank.abvr" == v {
				x -= 1
			}
		}
		x = xOrigin
		y += yInc
	}
}

func deepCopy(in layout) (out layout) {
	out = make(layout, 0, len(in))
	for _, pIn := range in {
		pOut := layoutPage{}
		for k, v := range pIn {
			pOut[k] = v
		}
		out = append(out, pOut)
	}
	return out
}
