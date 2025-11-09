package gotemplates

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
	"dragonbytes02/scaff/internal/scripts"
)

var goOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoHexa}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.GoMC}},
	{Description: "ELM", Msg: msg.RunningScriptMsg{Script: scripts.GoEML}},
}

func InitialJsPooModel() menu.Model {

	return menu.InitialModel(goOptions)
}
