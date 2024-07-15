package characters

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestAsciisheet(t *testing.T) {
	var c = Character{}

	err := json.Unmarshal(sabertoothJSON, &c)
	if err != nil {
		t.Errorf("Error reading sabertooth: %s", err)
	}
	data := AsciiCharacterSheet(c, false, false, false)
	datas := string(data)

	t.Logf("made sheet:\n%s", datas)

	variable := "name"
	expected := c.Name
	if !strings.Contains(datas, c.Name) {
		t.Errorf("sheet does not contain %s [%s]: \n%s", variable, expected, data)
	}
}

func TestLoadSabertooth(t *testing.T) {
	var c = Character{}

	err := json.Unmarshal(sabertoothJSON, &c)
	if err != nil {
		t.Errorf("Error reading sabertooth: %s", err)
	}

	t.Logf("json: %s", sabertoothJSON)
	t.Logf("unmarshaled: %#v", c)

	var b []byte
	b, err = yaml.Marshal(c)
	if err != nil {
		t.Errorf("Error marhsaling sabertooth yaml: %s", err)
	}
	t.Log(string(b), '\n')

	dc := c.ToDisplayCharacter()
	b, err = yaml.Marshal(dc)
	if err != nil {
		t.Errorf("Error marhsaling display yaml: %s", err)
	}
	t.Log(string(b), '\n')

	variable := "Name"
	actual := string(c.Name)
	expected := "Sabertooth"
	if actual != expected {
		t.Errorf("expcted %s %s but got %s", variable, expected, actual)

	}

	variable = "Fighting"
	actual = string(c.Attributes.Fighting.Abbreviation())
	expected = "In"
	if actual != expected {
		t.Errorf("expcted %s %s but got %s", variable, expected, actual)

	}

	variable = "Strength"
	actual = string(c.Attributes.Strength.Abbreviation())
	expected = "Gd"
	if actual != expected {
		t.Errorf("expcted %s %s but got %s", variable, expected, actual)

	}
}

func TestSaveVermin(t *testing.T) {
	c := Character{
		Name: "Vermin",
		Attributes: Attributes{
			Fighting:  30,
			Agility:   21,
			Strength:  22,
			Endurance: 53,
			Reason:    2,
			Intuition: 6,
			Psyche:    4,
		},
		Log: []CreationLog{
			{Message: "testing"},
		},
	}
	checkVerminAttributes(c, t)

	vb, _ := json.Marshal(c)
	t.Logf("vermin: %s", vb)

	_ = json.Unmarshal(vb, &c)
	checkVerminAttributes(c, t)

}

func checkVerminAttributes(c Character, t *testing.T) {
	actual := c.Attributes.Fighting.Abbreviation()
	expected := RankAbbreviation("Rm")
	if expected != actual {
		t.Errorf("expcted Fighting %s but got %s", expected, actual)
	}

	actual = c.Attributes.Agility.Abbreviation()
	expected = RankAbbreviation("Ex")
	if expected != actual {
		t.Errorf("expcted Agility %s but got %s", expected, actual)
	}

	actual = c.Attributes.Strength.Abbreviation()
	expected = RankAbbreviation("Ex")
	if expected != actual {
		t.Errorf("expcted Strength %s but got %s", expected, actual)
	}

	actual = c.Attributes.Endurance.Abbreviation()
	expected = RankAbbreviation("Am")
	if expected != actual {
		t.Errorf("expcted Endurance %s but got %s", expected, actual)
	}

	actual = c.Attributes.Reason.Abbreviation()
	expected = RankAbbreviation("Fe")
	if expected != actual {
		t.Errorf("expcted Reason %s but got %s", expected, actual)
	}

	actual = c.Attributes.Intuition.Abbreviation()
	expected = RankAbbreviation("Ty")
	if expected != actual {
		t.Errorf("expcted Intuition %s but got %s", expected, actual)
	}

	actual = c.Attributes.Psyche.Abbreviation()
	expected = RankAbbreviation("Pr")
	if expected != actual {
		t.Errorf("expcted Psyche %s but got %s", expected, actual)
	}
}

//go:embed test_data/sabertooth.json
var sabertoothJSON []byte

//go:embed test_data/sabertooth.yaml
var sabertoothYAML []byte

//// go:embed test_data/sabertooth.png
// var sabertoothImage []byte
