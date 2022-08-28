package parser

type Search struct {
	Results  []Result
	ResCount int
	Pages    []int
}

type Result struct {
	Id             int
	Hangul         string
	HomonymNumber  int
	Hanja          string
	TypeKr         string
	TypeEng        string
	Pronounciation string
	Audio          string
	Level          int
	Inflections    string
	Senses         []sense
}
