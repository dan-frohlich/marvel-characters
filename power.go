package characters

type PowerName string
type PowerRef string
type PowerDesc string

type PowerEntry struct {
	Name         PowerName
	Reference    PowerRef
	Description  string
	IsInitiative bool
}

type Power struct {
	Name      string    `yaml:"name,omitempty"`
	EntryName PowerName `yaml:"entryname,omitempty"`
	Rank      Rank      `yaml:"rank,omitempty"`
}

var PowerEntries = map[PowerName]PowerEntry{
	"Infravision":                  {Name: "Infravision", Description: "Use to see in the dark", Reference: "APB-71"},
	"Enhanced Senses":              {Name: "Enhanced Senses", Description: "Use to discover clues, spot items and initiative", Reference: "APB-71", IsInitiative: true},
	"Body Alterations - Offensive": {Name: "Body Alterations - Offensive", Description: "adds combat powers", Reference: "APB-86"},
	"Optic Blasts - cyclops":       {Name: "Optic Blasts", Description: "Ruby beams with a range of 1 area", Reference: "AJB-26"},
	"Visored Blasts - cyclops":     {Name: "Visored Blasts", Description: "A specially-constructed visor grants a range of 3 areas", Reference: "AJB-26"},
	"Invulnerability":              {Name: "Invulnerability", Description: "Totally unaffected by the listed attack form", Reference: "APB-73"},
}

type BookRef struct {
	Name        string
	Description string
}

var BookRefs = []BookRef{
	{Name: "APB", Description: "Advanced Set Players Book"},
	{Name: "ARB", Description: "Advanced Set Referees Book"},
	{Name: "UPB", Description: "Ultimate Powers Book"},
}
