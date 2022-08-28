package parser

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func ParseView(result_html io.Reader, id string, l string) (res *View, err error) {
	var v View
	v.Id, _ = strconv.Atoi(id)
	doc, err := html.Parse(result_html)
	dfsvView(doc, &v, l)
	return &v, err
}

func dfsvView(n *html.Node, in *View, l string) *html.Node {
	// Get the Hangul
	if CheckClass(n, "word_head") {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c != nil {
				if c.Type == html.TextNode {
					in.Hangul = c.Data
				} else if c.Type == html.ElementNode {
					in.HomonymNumber, _ = strconv.Atoi(GetTextAll(c))
					break
				}
			}
		}

	} else if CheckClass(n, "chi_info ml5") {
		// Get the Hanja
		hanja := MatchBetween(GetTextAll(n), "(", " )")
		in.Hanja = hanja

	} else if CheckClass(n, "search_sub") {
		// Get the pronounciation and audio file
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && strings.TrimSpace(c.Data) != "" && c.Data[0:1] == "[" {
				in.Pronounciation = strings.TrimSpace(c.Data[1:])
			} else if c.Data == "a" && c.Attr[1].Val == "sound" {
				for _, a := range c.Attr {
					if a.Key == "href" {
						in.Audio = MatchBetween(a.Val, "'", "');")
						break
					}
				}
				break
			}
		}

	} else if CheckClass(n, "word_att_type1") {
		// Get the Korean word type
		match := MatchBetween(GetTextAll(n.FirstChild), "「", "」")
		in.TypeKr = strings.ToValidUTF8(match, "")

	} else if CheckClass(n, fmt.Sprintf("manyLang%s ml5", l)) {
		// Get the English word type
		in.TypeEng = strings.TrimSpace(GetTextSingle(n.FirstChild))

	} else if CheckClass(n, "ri-star-s-fill") {
		// Get the level of the word
		in.Level++

	} else if CheckClass(n, "printArea") {
		// Get the Inflections
		in.Inflections = (GetTextSingle(n))
		inf := InitInflectionLinks()
		ind := 0
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Data == "a" {
				in.InflectionLinks = append(in.InflectionLinks, inf)
				re := regexp.MustCompile("[0-9]+")
				id := c.Attr[1].Val
				id = re.FindAllString(id, -1)[0]
				in.InflectionLinks[ind].Id, _ = strconv.Atoi(id)
				in.InflectionLinks[ind].Hangul = GetTextSingle(c) + "<sup>" + GetTextSingle(c.NextSibling) + "</sup>"
				ind++
			}
		}

	} else if CheckClass(n, "explain_list") && !CheckClass(n.Parent, "explain_list_wrap") {
		in.Senses = append(in.Senses, InitSense())
		n = getSenses(n, in.Senses, l)

	} else if CheckClass(n, "idiom_adage") {
		in.Proverbs = append(in.Proverbs, InitProverb())
		getProverbs(n, in.Proverbs, l)
	}

	// Traverse the tree of nodes vi depth-first search
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if c.Type == html.CommentNode || c.Data == "script" || len(strings.TrimSpace(c.Data)) == 0 {
			continue
		} else {
			dfsvView(c, in, l)
		}
	}
	return n
}
