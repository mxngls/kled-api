package parser

type sense struct {
	Translation  string
	Definition   string
	KrDefinition string
	Examples     []example
	Reference    []ref
}

type inflectionLink struct {
	Id     int
	Hangul string
}

func InitSense() sense {
	return sense{}
}

func InitInflectionLinks() inflectionLink {
	return inflectionLink{}
}
