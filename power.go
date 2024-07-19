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
	Name      string    `json:"name,omitempty"`
	EntryName PowerName `json:"power_name,omitempty"`
	Rank      Rank      `json:"rank,omitempty"`
}

var PowerEntries = map[PowerName]PowerEntry{
	"Infravision":                  {Name: "Infravision", Description: "Use to see in the dark", Reference: "APB-71"},
	"Enhanced Senses":              {Name: "Enhanced Senses", Description: "Use to discover clues, spot items and initiative", Reference: "APB-71", IsInitiative: true},
	"Body Alterations - Offensive": {Name: "Body Alterations - Offensive", Description: "adds combat powers", Reference: "APB-86"},
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
