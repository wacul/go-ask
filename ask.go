package gask

import (
	"fmt"
	"sort"
	"strings"
)

// Responder : ユーザーの入力に応じた処理
type Responder func(string) *bool

// Behaviors : ユーザーの入力と対応する処理を取得する
type Behaviors map[string]Responder

// New : ユーザーの入力と対応する処理を指定して、プロンプトを設定します
func New(behaviors Behaviors) *Prompt {
	keys := make(sort.StringSlice, 0, len(behaviors))
	for key := range behaviors {
		keys = append(keys, key)
	}
	sort.Sort(keys)
	return &Prompt{behaviors: behaviors, keys: strings.Join(keys, "/")}
}

// YesNoOrElse : YesとNo以外の選択肢を指定して、プロンプトを設定します
func YesNoOrElse(behaviors Behaviors) *Prompt {
	beh := Behaviors{}
	for key, behavior := range defaultBehaviors {
		beh[key] = behavior
	}
	for key, behavior := range behaviors {
		beh[key] = behavior
	}
	return New(beh)
}

// Prompt : ユーザー入力を促す処理のハンドラ
type Prompt struct {
	behaviors Behaviors
	keys      string
}

// Ask : プロンプトを出してy/nの入力を迫る
func (p *Prompt) Ask(question string) bool {
	// loop until we have an answer
	for {
		var line string
		fmt.Printf("%s [%s] ", question, p.keys)
		fmt.Scanln(&line)
		proc, ok := p.behaviors[line]
		if !ok {
			fmt.Printf("Inavlid answer. Please give %s.\n", p.keys)
			continue
		}
		if answer := proc(line); answer != nil {
			return *answer
		}
	}
}

// Askf : プロンプトを出してy/nの入力を迫る
func (p *Prompt) Askf(format string, args ...interface{}) bool {
	return p.Ask(fmt.Sprintf(format, args...))
}
