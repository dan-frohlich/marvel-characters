package characters

import "log"

type CreationLog struct {
	Message string            `json:"msg,omitempty" yaml:"msg,omitempty"`
	Fields  map[string]string `json:"fields,omitempty" yaml:"fields,omitempty"`
}

type ValueLog struct {
	Notes string `json:"notes,omitempty" yaml:"notes,omitempty"`
	Value int    `json:"value,omitempty" yaml:"value,omitempty"`
}

type Character struct {
	Name             string        `json:"name,omitempty" yaml:"name,omitempty"`
	Attributes       Attributes    `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	InitialResources Rank          `json:"initial_resources,omitempty" yaml:"initial_resources,omitempty"`
	Powers           []Power       `json:"powers,omitempty" yaml:"powers,omitempty"`
	Talents          []Talent      `json:"talents,omitempty" yaml:"talents,omitempty"`
	Log              []CreationLog `json:"log,omitempty" yaml:"log,omitempty"`
	KarmaLog         []ValueLog    `json:"karma_log,omitempty" yaml:"karma_log,omitempty"`
	PopularityLog    []ValueLog    `json:"popularity_log,omitempty" yaml:"popularity_log,omitempty"`
	ResourcesLog     []ValueLog    `json:"resources_log,omitempty" yaml:"resources_log,omitempty"`

	Image string `json:"portrait,omitempty" yaml:"portrait,omitempty"`
}

type Attributes struct {
	Fighting  Rank `json:"fighting,omitempty" yaml:"fighting,omitempty"`
	Agility   Rank `json:"agility,omitempty" yaml:"agility,omitempty"`
	Strength  Rank `json:"strength,omitempty" yaml:"strength,omitempty"`
	Endurance Rank `json:"endurance,omitempty" yaml:"endurance,omitempty"`
	Reason    Rank `json:"reason,omitempty" yaml:"reason,omitempty"`
	Intuition Rank `json:"intuition,omitempty" yaml:"intuition,omitempty"`
	Psyche    Rank `json:"psyche,omitempty" yaml:"psyche,omitempty"`
}

type Talent struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Reference   string `json:"reference,omitempty" yaml:"reference,omitempty"`
}

func (c Character) Health() int {
	return c.Attributes.Fighting.Value() +
		c.Attributes.Agility.Value() +
		c.Attributes.Strength.Value() +
		c.Attributes.Endurance.Value()
}

func (c Character) Karma() int {
	v := c.Attributes.Reason.Value() +
		c.Attributes.Intuition.Value() +
		c.Attributes.Psyche.Value()

	for _, k := range c.KarmaLog {
		v += k.Value
	}
	return v
}

func (c Character) Popularity() int {
	v := 0
	for _, k := range c.PopularityLog {
		v += k.Value
	}
	return v
}

func (c Character) Resources() int {
	v := 0
	for _, k := range c.ResourcesLog {
		v += k.Value
	}
	return v
}

func (c Character) Move() int {
	switch c.Attributes.Endurance.Entry().Abbreviation {
	case "0":
		return 0
	case "Fe":
		return 1
	case "Pr", "Ty", "Gd", "Ex":
		return 2
	default:
		return 3
	}
}

func (c Character) InitiativeMod() int {
	val := c.Attributes.Intuition.Value()

	for _, p := range c.Powers {
		if pe, ok := PowerEntries[p.EntryName]; ok {
			if pe.IsInitiative && p.Rank.Value() > val {
				log.Printf("%s uses %s rank of %d for Initiative", c.Name, pe.Name, p.Rank.Value())
				val = p.Rank.Value()
			}
		}
	}

	switch {
	case val <= 10:
		return 0
	case 11 <= val && val <= 20:
		return 1
	case 21 <= val && val <= 30:
		return 2
	case 31 <= val && val <= 40:
		return 3
	case 41 <= val && val <= 50:
		return 4
	case 51 <= val && val < 75:
		return 5
	case 75 <= val:
		return 6
	default:
		return 0
	}
}
