package javascript

import (
	"dragonbytes039/scaff/internal/components/menu"
	"dragonbytes039/scaff/internal/msg"
	"dragonbytes039/scaff/internal/scripts"
)

var jsPooOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsPooClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsPooHexa}},
	{Description: "Model Controller Services", Msg: msg.RunningScriptMsg{Script: scripts.JsPooMcs}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.JsPooMC}},
}

func InitialJsPooModel() menu.Model {

	return menu.InitialModel(jsPooOptions, "select template")
}
