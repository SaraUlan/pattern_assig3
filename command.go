package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type Editor struct {
	text string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (te *Editor) Copy() {
	te.text = "Has been copied"
}
func (te *Editor) Cut() {
	te.text = "Has been cut"
}
func (te *Editor) Paste() {
	te.text = "Pasted" + te.text
}

type CopyCommand struct {
	editor *Editor
}

func NewCopyCommand(editor *Editor) *CopyCommand {
	return &CopyCommand{editor}
}

func (cc *CopyCommand) Execute() {
	cc.editor.Copy()
}

type CutCommand struct {
	editor *Editor
}

func NewCutCommand(editor *Editor) *CutCommand {
	return &CutCommand{editor}
}

func (cc *CutCommand) Execute() {
	cc.editor.Cut()
}

type PasteCommand struct {
	editor *Editor
}

func NewPastCommand(editor *Editor) *PasteCommand {
	return &PasteCommand{editor}
}

func (pc *PasteCommand) Execute() {
	pc.editor.Paste()
}

type EditorInvoker struct {
	command Command
}

func (ei *EditorInvoker) SetCommand(command Command) {
	ei.command = command
}

func (ei *EditorInvoker) Execute() {
	ei.command.Execute()
}

func main() {
	editor := NewEditor()
	copyCommand := NewCopyCommand(editor)
	CutCommand := NewCutCommand(editor)
	PasteCommand := NewPastCommand(editor)

	invoker := &EditorInvoker{}

	invoker.SetCommand(copyCommand)
	invoker.Execute()
	fmt.Println(editor.text)

	invoker.SetCommand(CutCommand)
	invoker.Execute()
	fmt.Println(editor.text)

	invoker.SetCommand(PasteCommand)
	invoker.Execute()
	fmt.Println(editor.text)
}
