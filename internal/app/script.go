package app

import (
	"dragonbytes039/scaff/fst"
	"dragonbytes039/scaff/internal/msg"
	"dragonbytes039/scaff/internal/scripts"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *AppModel) runScript(selection scripts.ScriptApp, name_project string) tea.Cmd {

	return func() tea.Msg {

		template := scripts.TemplateName[selection]

		err := fst.SpreadTemplate(template, name_project)

		if err != nil {
			return msg.ScriptErrorMsg{Err: string("output") + ": " + err.Error()}
		}
		messagess := scripts.SuccessMsg(selection)
		return msg.ScriptFinishedMsg{Msg: messagess}
	}
}
