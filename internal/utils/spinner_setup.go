package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

// func NewDefaultSpinner(loading_msg string) *spinner.Spinner {
func NewDefaultSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("yellow")
	// s.Suffix = loading_msg
	return s
}
