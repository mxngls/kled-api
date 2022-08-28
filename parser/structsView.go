package parser

type View struct {
	Id              int
	Hangul          string
	HomonymNumber   int
	Hanja           string
	TypeKr          string
	TypeEng         string
	Pronounciation  string
	Audio           string
	Level           int
	Inflections     string
	InflectionLinks []inflectionLink
	Proverbs        []proverb
	Senses          []sense
}

type proverb struct {
	Hangul string
	Type   string
	Senses []sense
}

type example struct {
	Plain string
	Html  string
}

type ref struct {
	Type  string
	Value []string
	Id    []int
}

func InitProverb() proverb {
	return proverb{}
}

func InitRef() ref {
	return ref{}
}
