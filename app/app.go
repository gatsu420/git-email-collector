package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

// PrintCommitMsgThirdLine will print third line of commit message, since it
// usually comes in this format:
//
//	this is commit title
//
//	This is commit message. It may expand to multiple lines, paragraphs, or
//	bullet points.
//
// Nothing will be printed if third line is nonexistent.
func PrintCommitMsgThirdLine(msg string) (res string) {
	defer func() {
		if recv := recover(); recv != nil {
			res = ""
		}
	}()

	return strings.Split(msg, "\n")[2]
}

func Collect(gitHttpsAddress string, numPrints string) error {
	defer func() {
		if recv := recover(); recv != nil {
			log.Fatal(recv)
		}
	}()
	defer os.RemoveAll("./repo")

	if len(gitHttpsAddress) == 0 {
		return errors.New("argument for git https address must be inputted")
	}

	if len(numPrints) == 0 {
		return errors.New("argument for number of prints must be inputted")
	}

	numPrintsInt, err := strconv.Atoi(numPrints)
	if err != nil {
		return errors.New("unable to parse argument for number of prints")
	}

	repo, err := git.PlainClone("./repo", &git.CloneOptions{
		URL:      gitHttpsAddress,
		Progress: os.Stdout,
	})
	if err != nil {
		return errors.New("unable to clone repo")
	}

	ref, err := repo.Head()
	if err != nil {
		return fmt.Errorf("unable to get reference of head commit: %v", err)
	}

	gitlog, err := repo.Log(&git.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return fmt.Errorf("unable to get log: %v", err)
	}

	commitBuffer := map[string]int{}
	err = gitlog.ForEach(func(c *object.Commit) error {
		commitBuffer[c.Author.Email]++
		commitCount := commitBuffer[c.Author.Email]
		if commitCount <= numPrintsInt {
			fmt.Println(
				c.Author.Name,
				c.Author.Email,
				strings.Split(c.Message, "\n")[0],
				PrintCommitMsgThirdLine(c.Message),
			)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to iterate each log: %v", err)
	}

	return nil
}
