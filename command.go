package main

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute()
	Undo()
	GetDescription() string
}

type TextEditor struct {
	content string
	history []string
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		content: "",
		history: make([]string, 0),
	}
}

func (te *TextEditor) GetContent() string {
	return te.content
}

func (te *TextEditor) SetContent(content string) {
	te.history = append(te.history, te.content)
	te.content = content
}

func (te *TextEditor) RestorePreviousContent() {
	if len(te.history) > 0 {
		te.content = te.history[len(te.history)-1]
		te.history = te.history[:len(te.history)-1]
	}
}

func (te *TextEditor) AppendText(text string) {
	te.history = append(te.history, te.content)
	te.content += text
}

func (te *TextEditor) DeleteText(n int) {
	te.history = append(te.history, te.content)
	if n >= len(te.content) {
		te.content = ""
	} else {
		te.content = te.content[:len(te.content)-n]
	}
}

func (te *TextEditor) ReplaceText(old, new string) {
	te.history = append(te.history, te.content)
	te.content = strings.ReplaceAll(te.content, old, new)
}

func (te *TextEditor) Clear() {
	te.history = append(te.history, te.content)
	te.content = ""
}

func (te *TextEditor) Print() {
	fmt.Printf("内容: \"%s\"\n", te.content)
}

type WriteCommand struct {
	editor *TextEditor
	text   string
}

func NewWriteCommand(editor *TextEditor, text string) *WriteCommand {
	return &WriteCommand{
		editor: editor,
		text:   text,
	}
}

func (wc *WriteCommand) Execute() {
	wc.editor.AppendText(wc.text)
}

func (wc *WriteCommand) Undo() {
	wc.editor.RestorePreviousContent()
}

func (wc *WriteCommand) GetDescription() string {
	return fmt.Sprintf("Write: \"%s\"", wc.text)
}

type DeleteCommand struct {
	editor *TextEditor
	length int
}

func NewDeleteCommand(editor *TextEditor, length int) *DeleteCommand {
	return &DeleteCommand{
		editor: editor,
		length: length,
	}
}

func (dc *DeleteCommand) Execute() {
	dc.editor.DeleteText(dc.length)
}

func (dc *DeleteCommand) Undo() {
	dc.editor.RestorePreviousContent()
}

func (dc *DeleteCommand) GetDescription() string {
	return fmt.Sprintf("Delete: %d文字", dc.length)
}

type ReplaceCommand struct {
	editor *TextEditor
	old    string
	new    string
}

func NewReplaceCommand(editor *TextEditor, old, new string) *ReplaceCommand {
	return &ReplaceCommand{
		editor: editor,
		old:    old,
		new:    new,
	}
}

func (rc *ReplaceCommand) Execute() {
	rc.editor.ReplaceText(rc.old, rc.new)
}

func (rc *ReplaceCommand) Undo() {
	rc.editor.RestorePreviousContent()
}

func (rc *ReplaceCommand) GetDescription() string {
	return fmt.Sprintf("Replace: \"%s\" -> \"%s\"", rc.old, rc.new)
}

type ClearCommand struct {
	editor *TextEditor
}

func NewClearCommand(editor *TextEditor) *ClearCommand {
	return &ClearCommand{
		editor: editor,
	}
}

func (cc *ClearCommand) Execute() {
	cc.editor.Clear()
}

func (cc *ClearCommand) Undo() {
	cc.editor.RestorePreviousContent()
}

func (cc *ClearCommand) GetDescription() string {
	return "Clear: 全削除"
}

type MacroCommand struct {
	commands    []Command
	description string
}

func NewMacroCommand(description string) *MacroCommand {
	return &MacroCommand{
		commands:    make([]Command, 0),
		description: description,
	}
}

func (mc *MacroCommand) AddCommand(command Command) {
	mc.commands = append(mc.commands, command)
}

func (mc *MacroCommand) Execute() {
	for _, command := range mc.commands {
		command.Execute()
	}
}

func (mc *MacroCommand) Undo() {
	for i := len(mc.commands) - 1; i >= 0; i-- {
		mc.commands[i].Undo()
	}
}

func (mc *MacroCommand) GetDescription() string {
	return fmt.Sprintf("Macro: %s (%d commands)", mc.description, len(mc.commands))
}

type CommandInvoker struct {
	history []Command
}

func NewCommandInvoker() *CommandInvoker {
	return &CommandInvoker{
		history: make([]Command, 0),
	}
}

func (ci *CommandInvoker) ExecuteCommand(command Command) {
	fmt.Printf("実行: %s\n", command.GetDescription())
	command.Execute()
	ci.history = append(ci.history, command)
}

func (ci *CommandInvoker) UndoLastCommand() {
	if len(ci.history) > 0 {
		lastCommand := ci.history[len(ci.history)-1]
		fmt.Printf("取り消し: %s\n", lastCommand.GetDescription())
		lastCommand.Undo()
		ci.history = ci.history[:len(ci.history)-1]
	} else {
		fmt.Println("取り消すコマンドがありません。")
	}
}

func (ci *CommandInvoker) ShowHistory() {
	fmt.Println("\n--- コマンド履歴 ---")
	if len(ci.history) == 0 {
		fmt.Println("（履歴なし）")
	} else {
		for i, command := range ci.history {
			fmt.Printf("%d. %s\n", i+1, command.GetDescription())
		}
	}
	fmt.Println("---")
}

func ExecCommand() {
	fmt.Println("=== Command Pattern Demo ===")

	editor := NewTextEditor()
	invoker := NewCommandInvoker()

	fmt.Println("初期状態:")
	editor.Print()

	fmt.Println("\n--- 基本的なコマンド実行 ---")

	writeCmd1 := NewWriteCommand(editor, "Hello ")
	invoker.ExecuteCommand(writeCmd1)
	editor.Print()

	writeCmd2 := NewWriteCommand(editor, "World!")
	invoker.ExecuteCommand(writeCmd2)
	editor.Print()

	replaceCmd := NewReplaceCommand(editor, "World", "Go")
	invoker.ExecuteCommand(replaceCmd)
	editor.Print()

	deleteCmd := NewDeleteCommand(editor, 3)
	invoker.ExecuteCommand(deleteCmd)
	editor.Print()

	invoker.ShowHistory()

	fmt.Println("\n--- Undo テスト ---")
	fmt.Println("最後のコマンドを取り消し:")
	invoker.UndoLastCommand()
	editor.Print()

	fmt.Println("さらに取り消し:")
	invoker.UndoLastCommand()
	editor.Print()

	fmt.Println("\n--- Macro コマンドテスト ---")

	macro := NewMacroCommand("挨拶文作成")
	macro.AddCommand(NewClearCommand(editor))
	macro.AddCommand(NewWriteCommand(editor, "こんにちは、"))
	macro.AddCommand(NewWriteCommand(editor, "Goの世界へ！"))
	macro.AddCommand(NewReplaceCommand(editor, "Go", "素晴らしいGo"))

	invoker.ExecuteCommand(macro)
	editor.Print()

	fmt.Println("\nマクロを取り消し:")
	invoker.UndoLastCommand()
	editor.Print()

	invoker.ShowHistory()

	fmt.Println("\n=== Demo completed ===")
}
