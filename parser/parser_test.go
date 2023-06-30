package parser

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestFindAnchorElementsInTree(t *testing.T) {
	var html_str string = `
	<html>
	<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
	</body>
	</html>
	`
	doc, _ := html.Parse(strings.NewReader(html_str))
	anchors_arr := []*html.Node{}
	findElementInTree(doc, &anchors_arr, "a")

	got := len(anchors_arr)
	want := 1
	assert.Equal(t, got, want)
}

func TestParse(t *testing.T) {
	file := "../ex2.html"

	fopen, err := os.Open(file)
	assert.Nil(t, err)

	links, err := Parse(fopen)
	assert.Nil(t, err)

	assert.NotEmpty(t, len(links))
	assert.Equal(t, links[0].Href, "https://www.twitter.com/joncalhoun")
	assert.Equal(t, links[0].Text, "Check me out on twitter")
	assert.Equal(t, links[1].Href, "https://github.com/gophercises")
	assert.Equal(t, links[1].Text, "Gophercises is on Github!")
}
