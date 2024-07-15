package characters

import (
	"fmt"
	"strings"
)

func AsciiCharacterSheet(c Character) (out []byte) {

	out = append(out, []byte(fmt.Sprintf("Name: %s\n", c.Name))...)
	out = append(out, []byte("Attributes...\n")...)
	out = append(out, []byte("--------------------------------------------------------------------------------\n")...)

	att := c.Attributes.Fighting
	attName := "Fighting"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, []byte(fmt.Sprintf("%12s %4d", "Health:", c.Health()))...)
	out = append(out, '\n')

	att = c.Attributes.Agility
	attName = "Agility"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, []byte(fmt.Sprintf("%12s %4d", "Karma:", c.Karma()))...)
	out = append(out, '\n')

	att = c.Attributes.Strength
	attName = "Strength"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, []byte(fmt.Sprintf("%12s %4d", "Resources:", c.Resources()))...)
	out = append(out, '\n')

	att = c.Attributes.Endurance
	attName = "Endurance"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, []byte(fmt.Sprintf("%12s %4d", "Popularity:", c.Popularity()))...)
	out = append(out, '\n')

	att = c.Attributes.Reason
	attName = "Reason"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, []byte(fmt.Sprintf("%12s %4d areas", "Move:", c.Move()))...)
	out = append(out, '\n')

	att = c.Attributes.Intuition
	attName = "Intuition"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, []byte(fmt.Sprintf("%12s %4d", "Initiative:", c.InitiativeMod()))...)
	out = append(out, '\n')

	att = c.Attributes.Psyche
	attName = "Psyche"
	out = append(out, attLine(att, attName, 10)...)
	out = append(out, ' ', '|', ' ')
	out = append(out, '\n')

	out = append(out, '\n')
	out = append(out, []byte("Powers...\n")...)
	out = append(out, []byte("--------------------------------------------------------------------------------\n")...)
	padding := 0
	for _, p := range c.Powers {
		if len(p.Name) > padding {
			padding = len(p.Name)
		}
	}
	padding += 3
	refs := make(map[string]struct{})
	for _, p := range c.Powers {
		att = p.Rank
		attName = string(p.EntryName)
		if p.Name != "" {
			attName = p.Name
		}
		out = append(out, attLine(att, attName, padding)...)
		out = append(out, '\n')
		if pe, ok := PowerEntries[p.EntryName]; ok {
			out = append(out, powLine(pe)...)
			for _, r := range BookRefs {
				if strings.HasPrefix(string(pe.Reference), r.Name) {
					refs[r.Name] = struct{}{}
				}
			}
		}
		out = append(out, '\n')
	}

	for _, r := range BookRefs {
		if _, ok := refs[r.Name]; ok {
			out = append(out, fmt.Sprintf(" ** %s - %s", r.Name, r.Description)...)
		}
	}

	return out
}

func attLine(att Rank, name string, padding int) []byte {
	attFmt := " %-*s %3s(%3d) [ %2d / %2d / %2d ]"
	green := att.Entry().Green
	yellow := att.Entry().Yellow
	red := att.Entry().Red
	attName := name + ":"
	return []byte(fmt.Sprintf(attFmt, padding, attName, att.Abbreviation(), att.Value(), green, yellow, red))
}

func powLine(p PowerEntry) []byte {
	attFmt := " - %s (%s) %s\n"
	return []byte(fmt.Sprintf(attFmt, p.Name, p.Reference, p.Description))
}
