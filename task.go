package task

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"io/ioutil"
)

var ErrNotImplemented = errors.New("function not implemented")

// ReadHTMLFromFile should read the file from the current directory, if it exists.
// The file data should be returned as a string.
func ReadHTMLFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
	return "", ErrNotImplemented
}

// CreateBuffer should transfer the contents of a string to a buffer.
func CreateBuffer(data string) bytes.Buffer {
	return bytes.Buffer{}
}

// CreateTree should create the tree representation of HTML represented by the buffer.
func CreateTree(buf bytes.Buffer) (*html.Node, error) {
	return nil, ErrNotImplemented
}

// CountDivTags should return the count of <div> tags in the document tree.
func CountDivTags(node *html.Node) int {
	return -1
}

// dfs is a utility function which will help you count the number of unique tags.
func dfs(node *html.Node, tagsCount map[string]int) {
}

// ExtractAllUniqueTagsInSortedOrder should return the unique tags in the document.
// These tags should also be sorted alphabetically.
func ExtractAllUniqueTagsInSortedOrder(node *html.Node) []string {
	return nil
}

// ExtractAllComments returns the list of all comments as they appear in the document.
// You also need to remove all the leading and trailing spaces in the comments.
// HINT: You might need to read about variadic functions.
func ExtractAllComments(node *html.Node) []string {
	return nil
}

// ExtractAllLinks returns all the links in the document, in order of appearance.
func ExtractAllLinks(node *html.Node) []string {
	return nil
}
