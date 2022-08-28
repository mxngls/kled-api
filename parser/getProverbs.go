package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func getProverbs(n *html.Node, p []proverb, l string) {

	if CheckClass(n, "idiom_title printArea") {
		for a := n.FirstChild; a != nil; a = a.NextSibling {
			if a.Type == html.CommentNode || a.Data == "script" || len(strings.TrimSpace(a.Data)) == 0 {
				continue
			} else if a.Data == "dd" {
				p[len(p)-1].Hangul = GetTextAll(a)
				break
			}
		}
	} else if p[len(p)-1].Type == "" && CheckClass(n, "manyLang6") {
		p[len(p)-1].Type = GetTextAll(n)

	} else if CheckClass(n, "explain_list") {
		p[len(p)-1].Senses = append(p[len(p)-1].Senses, InitSense())
		getSenses(n, p[len(p)-1].Senses, l)
	}

	for a := n.FirstChild; a != nil; a = a.NextSibling {
		// Skip all commment nodes or nodes whose type is "script"
		if a.Type == html.CommentNode || a.Data == "script" || len(strings.TrimSpace(a.Data)) == 0 {
			continue
		} else {
			getProverbs(a, p, l)
		}
	}
}
