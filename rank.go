package characters

type RankEntry struct {
	Name         RankName
	Abbreviation RankAbbreviation
	Min          int
	Avg          int
	Max          int
	Green        int
	Yellow       int
	Red          int
}

type Rank int
type RankName string
type RankAbbreviation string

func (r Rank) Value() int {
	return int(r)
}

func (r Rank) Name() RankName {
	return r.Entry().Name
}

func (r Rank) Abbreviation() RankAbbreviation {
	return r.Entry().Abbreviation
}

func (r Rank) Entry() RankEntry {
	v := r.Value()
	if re, ok := RanksByAverge[v]; ok {
		return re
	}

	for _, re := range Ranks {
		if re.Min <= v && re.Max >= v {
			return re
		}
	}
	return UnknownRank
}

var (
	UnknownRank    = RankEntry{Name: "Unknown", Abbreviation: "Unknown", Min: -9, Max: -9, Avg: -9, Green: -9, Yellow: -9, Red: -9}
	ShiftZeroRank  = RankEntry{Name: "Shift 0", Abbreviation: "0", Min: 0, Max: 0, Avg: 0, Green: 66, Yellow: 95, Red: 100}
	FeebleRank     = RankEntry{Name: "Feeble", Abbreviation: "Fe", Min: 1, Max: 2, Avg: 2, Green: 61, Yellow: 91, Red: 100}
	PoorRank       = RankEntry{Name: "Poor", Abbreviation: "Pr", Min: 3, Max: 4, Avg: 4, Green: 56, Yellow: 86, Red: 100}
	TypicalRank    = RankEntry{Name: "Typical", Abbreviation: "Ty", Min: 5, Max: 7, Avg: 6, Green: 51, Yellow: 81, Red: 98}
	GoodRank       = RankEntry{Name: "Good", Abbreviation: "Gd", Min: 8, Max: 15, Avg: 10, Green: 46, Yellow: 76, Red: 98}
	ExcellentRank  = RankEntry{Name: "Excellent", Abbreviation: "Ex", Min: 16, Max: 25, Avg: 20, Green: 41, Yellow: 71, Red: 95}
	RemarkableRank = RankEntry{Name: "Remarkable", Abbreviation: "Rm", Min: 26, Max: 36, Avg: 30, Green: 36, Yellow: 66, Red: 95}
	IncredibleRank = RankEntry{Name: "Incredible", Abbreviation: "In", Min: 37, Max: 45, Avg: 40, Green: 31, Yellow: 61, Red: 91}
	AmazingRank    = RankEntry{Name: "Amazing", Abbreviation: "Am", Min: 46, Max: 62, Avg: 50, Green: 26, Yellow: 56, Red: 91}
	MonsterousRank = RankEntry{Name: "Monsterous", Abbreviation: "Mn", Min: 63, Max: 87, Avg: 75, Green: 21, Yellow: 51, Red: 86}
	UnearthlyRank  = RankEntry{Name: "Unearthly", Abbreviation: "Un", Min: 88, Max: 125, Avg: 100, Green: 16, Yellow: 46, Red: 86}
	ShiftXRank     = RankEntry{Name: "Shift X", Abbreviation: "X", Min: 126, Max: 175, Avg: 150, Green: 11, Yellow: 41, Red: 81}
	ShiftYRank     = RankEntry{Name: "Shift Y", Abbreviation: "Y", Min: 176, Max: 350, Avg: 200, Green: 7, Yellow: 41, Red: 81}
	ShiftZRank     = RankEntry{Name: "Shift Z", Abbreviation: "Z", Min: 351, Max: 999, Avg: 500, Green: 4, Yellow: 36, Red: 76}
	Class1000Rank  = RankEntry{Name: "Class 1000", Abbreviation: "C1k", Min: 1000, Max: 1000, Avg: 1000, Green: 2, Yellow: 36, Red: 76}
	Class3000Rank  = RankEntry{Name: "Class 3000", Abbreviation: "C3k", Min: 3000, Max: 3000, Avg: 3000, Green: 2, Yellow: 31, Red: 71}
	Class5000Rank  = RankEntry{Name: "Class 5000", Abbreviation: "C5k", Min: 5000, Max: 5000, Avg: 5000, Green: 2, Yellow: 26, Red: 66}
	BeyondRank     = RankEntry{Name: "Beyond", Abbreviation: "B", Min: -1, Max: -1, Avg: -1, Green: 2, Yellow: 21, Red: 61}
)

var RanksByAverge = map[int]RankEntry{
	ShiftZeroRank.Avg:  ShiftZeroRank,
	Class1000Rank.Avg:  Class1000Rank,
	Class3000Rank.Avg:  Class3000Rank,
	Class5000Rank.Avg:  Class5000Rank,
	BeyondRank.Avg:     Class5000Rank,
	FeebleRank.Avg:     FeebleRank,
	PoorRank.Avg:       PoorRank,
	TypicalRank.Avg:    TypicalRank,
	GoodRank.Avg:       GoodRank,
	ExcellentRank.Avg:  ExcellentRank,
	RemarkableRank.Avg: RemarkableRank,
	IncredibleRank.Avg: IncredibleRank,
	AmazingRank.Avg:    AmazingRank,
	MonsterousRank.Avg: MonsterousRank,
	UnearthlyRank.Avg:  UnearthlyRank,
	ShiftXRank.Avg:     ShiftXRank,
	ShiftYRank.Avg:     ShiftYRank,
	ShiftZRank.Avg:     ShiftZRank,
}

var Ranks []RankEntry = []RankEntry{
	FeebleRank,
	PoorRank,
	TypicalRank,
	GoodRank,
	ExcellentRank,
	RemarkableRank,
	IncredibleRank,
	AmazingRank,
	MonsterousRank,
	UnearthlyRank,
	ShiftXRank,
	ShiftYRank,
	ShiftZRank,
}
