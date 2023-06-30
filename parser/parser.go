package parser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func findElementInTree(node *html.Node, elments_arr *[]*html.Node, element string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == element {
			*elments_arr = append(*elments_arr, c)
		}
		findElementInTree(c, elments_arr, element)
	}
}

func findInTextNode(node *html.Node, text_arr *string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			data := c.Data
			*text_arr += data
		}
		findInTextNode(c, text_arr)
	}
}

func Parse(document io.Reader) ([]Link, error) {
	doc_tree, err := html.Parse(document)
	if err != nil {
		return []Link{}, err
	}

	anchors_arr := []*html.Node{}
	findElementInTree(doc_tree, &anchors_arr, "a")

	links := make([]Link, len(anchors_arr))

	for i := 0; i < len(anchors_arr); i++ {
		if anchors_arr[i].Attr[0].Key == "href" {
			links[i].Href = anchors_arr[i].Attr[0].Val
		}
	}

	for i := 0; i < len(anchors_arr); i++ {
		str := ""
		findInTextNode(anchors_arr[i], &str)
		str = strings.Join(strings.Fields(str), " ")
		fmt.Println(str)
		links[i].Text = str
	}

	return links, nil
}
