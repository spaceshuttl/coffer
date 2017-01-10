package cmd

import (
	"strings"

	"github.com/mitchellh/cli"
)

const (
	New    = "new"
	List   = "ls"
	Find   = "find"
	Remove = "remove"
)

type NoteCommand struct {
	ui *cli.BasicUi
}

func NewNote(ui *cli.BasicUi) *NoteCommand { return &NoteCommand{ui: ui} }

func (e *NoteCommand) Help() string {
	txt := `Usage: coffer note <option>
note manages all aspects of note management

Note Options:
	new
		Creates a new note.
	ls
		List all notes by Title
	find [query]
		Perform a fuzzy search against your notes
	rm [id]
		TBC - removes a note
	`
	return strings.TrimSpace(txt)
}

func (e *NoteCommand) Run(args []string) int {
	if len(args) == 0 {
		e.ui.Error(e.Help())
		e.Help()
		return 1
	}

	switch args[0] {
	case New:
		err := e.newNote()
		if err != nil {
			e.ui.Error(err.Error())
			return 1
		}
		return 0

	case List:
		err := e.listNote()
		if err != nil {
			e.ui.Error(err.Error())
			return 1
		}
		return 0

	case Find:
		err := e.findNote()
		if err != nil {
			e.ui.Error(err.Error())
			return 1
		}
		return 0

	case Remove:
		err := e.removeNote()
		if err != nil {
			e.ui.Error(err.Error())
			return 1
		}
		return 0

	default:
		e.ui.Error(e.Help())
		return 1
	}
}

func (e *NoteCommand) Synopsis() string {
	return "Makes a secret note and adds it in to the bucket"
}

func (e *NoteCommand) newNote() error {
	return nil
}

func (e *NoteCommand) listNote() error {
	return nil
}
func (e *NoteCommand) findNote() error {
	return nil
}
func (e *NoteCommand) removeNote() error {
	return nil
}
