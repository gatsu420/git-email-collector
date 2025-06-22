package app_test

import (
	"errors"
	"testing"

	"github.com/gatsu420/git-email-collector/app"
	"github.com/stretchr/testify/assert"
)

func Test_PrintCommitMsgThirdLine(t *testing.T) {
	testCases := []struct {
		testName       string
		msg            string
		expectedString string
	}{
		{
			testName:       "third line is nonexistent",
			msg:            "this msg consists of only single line",
			expectedString: "",
		},
		{
			testName:       "msg is empty string",
			msg:            "",
			expectedString: "",
		},
		{
			testName: "third line exists",
			msg: `this is commit title

This is commit message. It may expand to multiple lines, paragraphs, or
bullet points.`,
			expectedString: "This is commit message. It may expand to multiple lines, paragraphs, or",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			thirdLine := app.PrintCommitMsgThirdLine(tc.msg)
			assert.Equal(t, tc.expectedString, thirdLine)
		})
	}
}

func Test_Collect(t *testing.T) {
	testCases := []struct {
		testName        string
		gitHttpsAddress string
		numPrints       string
		expectedErr     error
	}{
		{
			testName:        "gitHttpsAddress is empty",
			gitHttpsAddress: "",
			numPrints:       "2",
			expectedErr:     errors.New("argument for git https address must be inputted"),
		},
		{
			testName:        "numPrints is empty",
			gitHttpsAddress: "https://github.com/gatsu420/mary.git",
			numPrints:       "",
			expectedErr:     errors.New("argument for number of prints must be inputted"),
		},
		{
			testName:        "numPrints can't be converted into int",
			gitHttpsAddress: "https://github.com/gatsu420/mary.git",
			numPrints:       "2q",
			expectedErr:     errors.New("unable to parse argument for number of prints"),
		},
		{
			testName:        "repo can't be cloned because it is private or nonexistent",
			gitHttpsAddress: "https://github.com/gatsu420/this-repo-is-nonexistent.git",
			numPrints:       "2",
			expectedErr:     errors.New("unable to clone repo"),
		},
		{
			testName:        "success",
			gitHttpsAddress: "https://github.com/gatsu420/mary.git",
			numPrints:       "2",
			expectedErr:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := app.Collect(tc.gitHttpsAddress, tc.numPrints)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
