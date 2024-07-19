package characters

import (
	"fmt"
	"log"
	"strconv"
)

type DisplayCharacter map[string]string

func (c Character) ToDisplayCharacter() DisplayCharacter {
	dc := make(DisplayCharacter)
	dc["name"] = c.Name
	dc["health"] = strconv.Itoa(c.Health())
	dc["resources"] = strconv.Itoa(c.Resources())
	dc["karma"] = strconv.Itoa(c.Karma())
	dc["popularity"] = strconv.Itoa(c.Popularity())
	dc["move"] = fmt.Sprintf("%4d areas", c.Move())
	dc["initiative"] = strconv.Itoa(c.InitiativeMod())
	dc["image"] = c.Image
	if c.Image == "" {
		log.Fatal("no character image found")
	}

	c.addAttributesToDisplay(dc)
	c.addPowersToDisplay(dc)
	return dc
}

func (c Character) addPowersToDisplay(dc DisplayCharacter) {
	for i, p := range c.Powers {
		name := string(p.EntryName)
		if p.Name != "" {
			name = p.Name
		}
		baseKey := fmt.Sprintf("powers.%d", i)
		dc[baseKey+".name"] = name
		dc[baseKey+".rank.value"] = strconv.Itoa(p.Rank.Value())
		dc[baseKey+".rank.name"] = string(p.Rank.Name())
		dc[baseKey+".rank.abvr"] = string(p.Rank.Abbreviation())
		dc[baseKey+".rank.green"] = strconv.Itoa(p.Rank.Entry().Green)
		dc[baseKey+".rank.yellow"] = strconv.Itoa(p.Rank.Entry().Yellow)
		dc[baseKey+".rank.red"] = strconv.Itoa(p.Rank.Entry().Red)

		if pe, ok := PowerEntries[p.EntryName]; ok {
			dc[baseKey+".entry.name"] = string(pe.Name)
			dc[baseKey+".entry.ref"] = string(pe.Reference)
			dc[baseKey+".entry.desc"] = string(pe.Description)
		}
	}
}

func (c Character) addAttributesToDisplay(dc DisplayCharacter) {
	dc["attribute.0.name"] = "Fighting"
	dc["attribute.0.abvr"] = "F"
	dc["attribute.0.rank.value"] = strconv.Itoa(c.Attributes.Fighting.Value())
	dc["attribute.0.rank.name"] = string(c.Attributes.Fighting.Name())
	dc["attribute.0.rank.abvr"] = string(c.Attributes.Fighting.Abbreviation())
	dc["attribute.0.rank.green"] = strconv.Itoa(c.Attributes.Fighting.Entry().Green)
	dc["attribute.0.rank.yellow"] = strconv.Itoa(c.Attributes.Fighting.Entry().Yellow)
	dc["attribute.0.rank.red"] = strconv.Itoa(c.Attributes.Fighting.Entry().Red)

	dc["attribute.1.name"] = "Agility"
	dc["attribute.1.abvr"] = "A"
	dc["attribute.1.rank.value"] = strconv.Itoa(c.Attributes.Agility.Value())
	dc["attribute.1.rank.name"] = string(c.Attributes.Agility.Name())
	dc["attribute.1.rank.abvr"] = string(c.Attributes.Agility.Abbreviation())
	dc["attribute.1.rank.green"] = strconv.Itoa(c.Attributes.Agility.Entry().Green)
	dc["attribute.1.rank.yellow"] = strconv.Itoa(c.Attributes.Agility.Entry().Yellow)
	dc["attribute.1.rank.red"] = strconv.Itoa(c.Attributes.Agility.Entry().Red)

	dc["attribute.2.name"] = "Strength"
	dc["attribute.2.abvr"] = "S"
	dc["attribute.2.rank.value"] = strconv.Itoa(c.Attributes.Strength.Value())
	dc["attribute.2.rank.name"] = string(c.Attributes.Strength.Name())
	dc["attribute.2.rank.abvr"] = string(c.Attributes.Strength.Abbreviation())
	dc["attribute.2.rank.green"] = strconv.Itoa(c.Attributes.Strength.Entry().Green)
	dc["attribute.2.rank.yellow"] = strconv.Itoa(c.Attributes.Strength.Entry().Yellow)
	dc["attribute.2.rank.red"] = strconv.Itoa(c.Attributes.Strength.Entry().Red)

	dc["attribute.3.name"] = "Endurance"
	dc["attribute.3.abvr"] = "E"
	dc["attribute.3.rank.value"] = strconv.Itoa(c.Attributes.Endurance.Value())
	dc["attribute.3.rank.name"] = string(c.Attributes.Endurance.Name())
	dc["attribute.3.rank.abvr"] = string(c.Attributes.Endurance.Abbreviation())
	dc["attribute.3.rank.green"] = strconv.Itoa(c.Attributes.Endurance.Entry().Green)
	dc["attribute.3.rank.yellow"] = strconv.Itoa(c.Attributes.Endurance.Entry().Yellow)
	dc["attribute.3.rank.red"] = strconv.Itoa(c.Attributes.Endurance.Entry().Red)

	dc["attribute.4.name"] = "Reason"
	dc["attribute.4.abvr"] = "R"
	dc["attribute.4.rank.value"] = strconv.Itoa(c.Attributes.Reason.Value())
	dc["attribute.4.rank.name"] = string(c.Attributes.Reason.Name())
	dc["attribute.4.rank.abvr"] = string(c.Attributes.Reason.Abbreviation())
	dc["attribute.4.rank.green"] = strconv.Itoa(c.Attributes.Reason.Entry().Green)
	dc["attribute.4.rank.yellow"] = strconv.Itoa(c.Attributes.Reason.Entry().Yellow)
	dc["attribute.4.rank.red"] = strconv.Itoa(c.Attributes.Reason.Entry().Red)

	dc["attribute.5.name"] = "Intuition"
	dc["attribute.5.abvr"] = "I"
	dc["attribute.5.rank.value"] = strconv.Itoa(c.Attributes.Intuition.Value())
	dc["attribute.5.rank.name"] = string(c.Attributes.Intuition.Name())
	dc["attribute.5.rank.abvr"] = string(c.Attributes.Intuition.Abbreviation())
	dc["attribute.5.rank.green"] = strconv.Itoa(c.Attributes.Intuition.Entry().Green)
	dc["attribute.5.rank.yellow"] = strconv.Itoa(c.Attributes.Intuition.Entry().Yellow)
	dc["attribute.5.rank.red"] = strconv.Itoa(c.Attributes.Intuition.Entry().Red)

	dc["attribute.6.name"] = "Psyche"
	dc["attribute.6.abvr"] = "P"
	dc["attribute.6.rank.value"] = strconv.Itoa(c.Attributes.Psyche.Value())
	dc["attribute.6.rank.name"] = string(c.Attributes.Psyche.Name())
	dc["attribute.6.rank.abvr"] = string(c.Attributes.Psyche.Abbreviation())
	dc["attribute.6.rank.green"] = strconv.Itoa(c.Attributes.Psyche.Entry().Green)
	dc["attribute.6.rank.yellow"] = strconv.Itoa(c.Attributes.Psyche.Entry().Yellow)
	dc["attribute.6.rank.red"] = strconv.Itoa(c.Attributes.Psyche.Entry().Red)
}
