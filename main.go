package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Link is a link from the awesome file
type Link struct {
	Heading string
	Name    string
	URL     string
}

var githubRE regexp

func init() {
	githubRE = regexp.MustCompile("httpgithub")
}

//func (l Link)IsGithub() {
//	l.URL
//}

func main() {
	file, err := os.Open("awesome.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	headingRE := regexp.MustCompile(`##.*`)
	linkRE := regexp.MustCompile(`\[(.*)\]\((.*)\)`)
	heading := "top"
	var links []*Link

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		if headingRE.MatchString(t) {
			fmt.Println(t)
		}

		if linkRE.MatchString(t) {
			l := linkRE.FindStringSubmatch(t)
			links = append(links, &Link{
				Heading: heading,
				Name:    l[1],
				URL:     l[2],
			})
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for _, l := range links {
		fmt.Println(l.URL)
	}
}
