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

var a6Layout = layoutPage{
	"name": &cell{
		x: 1, y: 1, w: 65, h: 8,
		align:      "LM",
		fontFamily: a6LayoutFont,
		fontSize:   16,
		fontWeight: "B",
	},
	// 	   Health:  116
	//  	Karma:   52
	//  Resources:    6
	// Popularity:    0
	//   	 Move:    3 areas
	// Initiative:    0

	"health.label": &cell{
		x: 72, y: 10, w: 22, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "B",
		border:     "LB",
		fixedText:  "Health",
	},
	"health": &cell{
		x: 95, y: 10, w: 8, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "",
		border:     "B",
	},
	"move.label": &cell{
		x: 72, y: 19, w: 22, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "B",
		border:     "LB",
		fixedText:  "Move (area)",
	},
	"move": &cell{
		x: 95, y: 19, w: 8, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "",
		border:     "B",
	},
	"initiative.label": &cell{
		x: 72, y: 28, w: 22, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "B",
		border:     "LB",
		fixedText:  "Initiative",
	},
	"initiative": &cell{
		x: 95, y: 28, w: 8, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "",
		border:     "B",
	},
	"karma.label": &cell{
		x: 72, y: 37, w: 22, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "B",
		border:     "LB",
		fixedText:  "Karma",
	},
	"karma": &cell{
		x: 95, y: 37, w: 8, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "",
		border:     "B",
	},
	"popularity.label": &cell{
		x: 72, y: 46, w: 22, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "B",
		border:     "LB",
		fixedText:  "Popularity",
	},
	"popularity": &cell{
		x: 95, y: 46, w: 8, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "",
		border:     "B",
	},
	"resources.label": &cell{
		x: 72, y: 55, w: 22, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "B",
		border:     "LB",
		fixedText:  "Resources",
	},
	"resources": &cell{
		x: 95, y: 55, w: 8, h: 8,
		align:      "RM",
		fontFamily: a6LayoutFont,
		fontSize:   10,
		fontWeight: "",
		border:     "B",
	},
}

func a6LayoutInit() {
	//add special abilities
	dy := 6.2
	offsetX := 7.0
	height := 6.0
	tableType := ttSA
	for row := 0; row < 10; row++ {
		key := layoutTableKeyName(tableType, row, "name")
		offset := float64(row) * dy
		a6Layout[key] = &cell{
			x: offsetX, y: 76 + offset, w: 30, h: height,
			align: "LM", fontFamily: a6LayoutFont, fontSize: 10, fontWeight: "B",
		}
		key = layoutTableKeyName(tableType, row, "notes")
		a6Layout[key] = &cell{
			x: 31 + offsetX, y: 76 + offset, w: 38, h: height,
			align: "LM", fontFamily: a6LayoutFont, fontSize: 8, fontWeight: "I",
		}
		key = layoutTableKeyName(tableType, row, "pool")
		a6Layout[key] = &cell{
			x: 69 + offsetX, y: 76 + offset, w: 12, h: height,
			align: "CM", fontFamily: a6LayoutFont, fontSize: 12, fontWeight: "B",
		}
	}

	tableType = ttEq
	offsetX = 101.5
	for row := 0; row < 10; row++ {
		key := layoutTableKeyName(tableType, row, "name")
		offset := float64(row) * dy
		a6Layout[key] = &cell{
			x: offsetX, y: 76 + offset, w: 30, h: height,
			align: "LM", fontFamily: a6LayoutFont, fontSize: 10, fontWeight: "B",
		}
		key = layoutTableKeyName(tableType, row, "notes")
		a6Layout[key] = &cell{
			x: 30.5 + offsetX, y: 76 + offset, w: 38.5, h: height,
			align: "LM", fontFamily: a6LayoutFont, fontSize: 8, fontWeight: "I",
		}
		key = layoutTableKeyName(tableType, row, "mass")
		a6Layout[key] = &cell{
			x: 69 + offsetX, y: 76 + offset, w: 14.5, h: height,
			align: "CM", fontFamily: a6LayoutFont, fontSize: 12, fontWeight: "B",
		}
		key = layoutTableKeyName(tableType, row, "status")
		a6Layout[key] = &cell{
			x: 84 + offsetX, y: 76 + offset, w: 19, h: height,
			align: "LM", fontFamily: a6LayoutFont, fontSize: 8, fontWeight: "I",
		}
	}
}

func a6LayoutAddAttributes() {
	xOrigin := float64(1)
	x := xOrigin
	y := float64(10)
	yInc := float64(9)
	for i := 0; i < 8; i++ {
		for _, v := range []string{"name", "rank.abvr", "rank.value", "rank.green", "rank.yellow", "rank.red"} {
			key := layoutKey(fmt.Sprintf("attribute.%d.%s", i, v))
			var (
				col      *color
				align    = "RM"
				weight   = ""
				fontSize = float64(12)
				w        = float64(8)
				h        = float64(8)
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
				align = "LM"
				border = "LB"
			case "rank.green":
				border = "LB"
				col = green
				weight = "I"
			case "rank.yellow":
				col = yellow
				weight = "I"
			case "rank.red":
				col = red
				weight = "I"
			}
			c := &cell{
				x: x, y: y, w: w, h: h,
				align: align, fontFamily: a6LayoutFont, fontSize: fontSize, fontWeight: weight,
				fillColor: col,
				border:    border,
			}
			a6Layout[key] = c
			x += w + 1
		}
		x = xOrigin
		y += yInc
	}
}

type layoutTableType string

const (
	ttSA = layoutTableType("sa")
	ttEq = layoutTableType("eq")
)

func layoutTableKeyName(tableType layoutTableType, row int, field string) layoutKey {
	return layoutKey(fmt.Sprintf("%s.%d.%s", tableType, row, field))
}
