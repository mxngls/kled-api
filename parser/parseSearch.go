package parser

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func ParseSearch(result_html io.Reader, l string) (res Search, err error) {
	var s Search
	doc, err := html.Parse(result_html)
	dfsvSearch(doc, &s, 0, l)
	return s, err
}

func dfsvSearch(n *html.Node, in *Search, i int, l string) *html.Node {
	rl := len(in.Results)
	if CheckClass(n, "blue ml5") {
		// Get the number of results
		str := GetTextAll(n)
		arr := strings.Split(str, " ")
		for x := 0; x < len(arr); x++ {
			conv, err := strconv.Atoi(arr[0])
			if err != nil {
				continue
			}
			in.ResCount = conv
		}

	} else if n.Data == "dl" {
		// Append the results array
		var r Result
		in.Results = append(in.Results, r)

	} else if CheckClass(n, "") && n.Data == "dd" {
		in.Results[rl-1].Inflections = GetContent(n, "sup")

	} else if CheckClass(n, "word_type1_17") {
		// Get the Hangul
		// Get the Id

		in.Results[rl-1].Hangul = GetTextSingle(n)
		re := regexp.MustCompile("[0-9]+")
		id, _ := strconv.Atoi(re.FindString(n.Parent.Attr[0].Val))
		in.Results[rl-1].Id = id

	} else if n.Data == "span" && n.FirstChild != nil &&
		n.FirstChild.Type == html.TextNode &&
		n.FirstChild.Data[0:1] == "(" {
		// Get the Hanja
		// Get the Hanja (if there is one)

		hanja := MatchBetween(n.FirstChild.Data, "(", ")")
		in.Results[rl-1].Hanja = hanja

	} else if CheckClass(n, "search_sub") {
		// Get the pronounciation and audio file

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && strings.TrimSpace(c.Data) != "" && c.Data[0:1] == "[" {
				in.Results[rl-1].Pronounciation = strings.TrimSpace(c.Data[1:])
			} else if c.Data == "a" && c.Attr[1].Val == "sound" {
				for _, a := range c.Attr {
					if a.Key == "href" {
						in.Results[rl-1].Audio = MatchBetween(a.Val, "'", "');")
						break
					}
				}
				break
			}
		}

	} else if CheckClass(n, "word_att_type1") {
		// Get the Korean word type

		match := MatchBetween(GetTextAll(n.FirstChild), "「", "」")
		in.Results[rl-1].TypeKr = strings.ToValidUTF8(match, "")

	} else if CheckClass(n, fmt.Sprintf("manyLang%s", l)) &&
		CheckClass(n.Parent, "word_att_type1") {
		// Get the English word type

		match := GetTextSingle(n.FirstChild)
		in.Results[rl-1].TypeEng = strings.TrimSpace(cleanStringSpecial([]byte(match)))

	} else if CheckClass(n, "ri-star-s-fill") {
		// Get the level of the word

		in.Results[rl-1].Level++

	} else if n.Data == "dd" && (CheckClass(n, fmt.Sprintf("manyLang%s mt15", l)) || CheckClass(n, fmt.Sprintf("manyLang%s ", l))) {
		// Get the english translation

		s := InitSense()
		in.Results[rl-1].Senses = append(in.Results[rl-1].Senses, s)

		len := len(in.Results[rl-1].Senses)
		in.Results[rl-1].Senses[len-1].Translation = cleanStringSpecial([]byte(GetTextAll(n)))

	} else if CheckClass(n, "ml20") {
		// Get the korean definition

		len := len(in.Results[rl-1].Senses)
		in.Results[rl-1].Senses[len-1].KrDefinition = GetTextAll(n)

	} else if CheckClass(n, fmt.Sprintf("manyLang%s ml20", l)) && GetTextAll(n) != "" {
		// Get the english definition

		len := len(in.Results[rl-1].Senses)
		in.Results[rl-1].Senses[len-1].Definition = GetTextAll(n)

	} else if CheckClass(n, fmt.Sprintf("manyLang%s ml20", l)) && n.NextSibling.NextSibling == nil {
		// Increment the index by one
		i++

	} else if CheckClass(n, "paging_area") {
		// Get the number of pages
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if CheckClass(c, "btn_first") {
				in.Pages = append(in.Pages, -4)
			} else if CheckClass(c, "btn_prev") {
				in.Pages = append(in.Pages, -3)
			} else if CheckClass(c, "paging_num") || CheckClass(c, "paging_num on") {
				page, err := strconv.Atoi(GetTextAll(c))
				if err != nil {
					panic(err)
				}
				in.Pages = append(in.Pages, page)
			} else if CheckClass(c, "btn_next") {
				in.Pages = append(in.Pages, -2)
			} else if CheckClass(c, "btn_last") {
				in.Pages = append(in.Pages, -1)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Traverse the tree of nodes vi depth-first search
		// Skip all commment nodes or nodes whose type is "script"
		if c.Type == html.CommentNode || c.Data == "script" || len(strings.TrimSpace(c.Data)) == 0 {
			continue
		} else {
			dfsvSearch(c, in, i, l)
		}
	}
	return n
}
