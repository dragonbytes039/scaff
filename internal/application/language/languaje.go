package language

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
)

var LanguagesOptions = []menu.OptionsInfo{
	{Description: "Go", Msg: msg.ChangeToGoOptionsMsg{}},
	{Description: "Typescript", Msg: msg.ChangeToParadigmMsg{}},
}

func InitialLanguageModel() menu.Model {

	return menu.InitialModel(LanguagesOptions)
}
