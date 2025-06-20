package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

// printCommitMsgThirdLine will print third line of commit message, since it
// usually comes in this format:
//
//	this is commit title
//
//	This is commit message. It may expand to multiple lines, paragraphs, or
//	bullet points.
//
// Nothing will be printed if third line is nonexistent.
func printCommitMsgThirdLine(msg string) string {
	defer func() {
		if recv := recover(); recv != nil {
		}
	}()

	return strings.Split(msg, "\n")[2]
}

func main() {
	defer func() {
		if recv := recover(); recv != nil {
			log.Fatal(recv)
		}
	}()
	defer os.RemoveAll("./repo")

	gitHttpsAddress := os.Args[1]
	numPrints, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("unable to parse argument for number of prints: %v", err)
	}

	repo, err := git.PlainClone("./repo", &git.CloneOptions{
		URL:      gitHttpsAddress,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatalf("unable to clone repo: %v", err)
	}

	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("unable to get reference of head commit: %v", err)
	}

	gitlog, err := repo.Log(&git.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		log.Fatalf("unable to get log: %v", err)
	}

	commitBuffer := map[string]int{}
	err = gitlog.ForEach(func(c *object.Commit) error {
		commitBuffer[c.Author.Email]++
		commitCount := commitBuffer[c.Author.Email]
		if commitCount <= numPrints {
			fmt.Println(
				c.Author.Name,
				c.Author.Email,
				strings.Split(c.Message, "\n")[0],
				printCommitMsgThirdLine(c.Message),
			)
		}

		return nil
	})
	if err != nil {
		log.Fatalf("unable to iterate each log: %v\n", err)
	}
}
