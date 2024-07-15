package characters

type CreationLog struct {
	Message string            `json:"msg,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`
}

type ValueLog struct {
	Notes string `json:"notes,omitempty"`
	Value int    `json:"value,omitempty"`
}

type Character struct {
	Name             string        `json:"name,omitempty"`
	Attributes       Attributes    `json:"attributes,omitempty"`
	InitialResources Rank          `json:"initial_resources,omitempty"`
	Powers           []Power       `json:"powers,omitempty"`
	Log              []CreationLog `json:"log,omitempty"`
	KarmaLog         []ValueLog    `json:"karma_log,omitempty"`
	PopularityLog    []ValueLog    `json:"popularity_log,omitempty"`
	ResourcesLog     []ValueLog    `json:"resources_log,omitempty"`
}

type Attributes struct {
	Fighting  Rank `json:"fighting,omitempty"`
	Agility   Rank `json:"agility,omitempty"`
	Strength  Rank `json:"strength,omitempty"`
	Endurance Rank `json:"endurance,omitempty"`
	Reason    Rank `json:"reason,omitempty"`
	Intuition Rank `json:"intuition,omitempty"`
	Psyche    Rank `json:"psyche,omitempty"`
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
