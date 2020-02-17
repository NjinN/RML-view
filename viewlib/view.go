package nativelib

import (
	. "github.com/NjinN/RML-view/core"

	"strings"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func View(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == BLOCK {
		var inTE, outTE *walk.TextEdit

		MainWindow{
			Title:   "xiaochuan测试",
			MinSize: Size{600, 400},
			Layout:  VBox{},
			Children: []Widget{
				HSplitter{
					Children: []Widget{
						TextEdit{AssignTo: &inTE, MaxLength: 10},
						TextEdit{AssignTo: &outTE, ReadOnly: true},
					},
				},
				PushButton{
					Text: "SCREAM",
					OnClicked: func() {
						outTE.SetText(strings.ToUpper(inTE.Text()))
					},
				},
			},
		}.Run()
		
		return &Token{NIL, nil}, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}
