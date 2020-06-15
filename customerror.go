package orderedmap

import (
	"fmt"
	"unicode/utf8"
)

// ErrorChainIsEmpty
type ErrorChainIsEmpty struct {
	message string
}

func (err ErrorChainIsEmpty) Error() string {

	return err.message
}

func CreateErrorChainIsEmpty(postMessage string) *ErrorChainIsEmpty {

	message := "Chain is empty"

	if utf8.RuneCountInString(postMessage) > 0 {
		message += fmt.Sprintf(": %s", postMessage)
	}

	return &ErrorChainIsEmpty{
		message: message,
	}
}
