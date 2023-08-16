package gask

import (
	"fmt"
)

// Yes : Trueを返す
func Yes(string) *bool {
	v := true
	return &v
}

// No : Falseを返す
func No(string) *bool {
	v := false
	return &v
}

var defaultBehaviors = Behaviors{
	"y":   Yes,
	"yes": Yes,
	"Y":   Yes,
	"n":   No,
	"no":  No,
	"N":   No,
}
var defaultPrompt = New(defaultBehaviors)

// Ask : プロンプトを出してy/nの入力を迫る
func Ask(question string) bool {
	return defaultPrompt.Ask(question)
}

// Askf : プロンプトを出してy/nの入力を迫る
func Askf(format string, args ...interface{}) bool {
	return Ask(fmt.Sprintf(format, args...))
}
