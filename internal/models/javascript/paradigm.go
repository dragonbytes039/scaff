package javascript

import (
	"dragonbytes039/scaff/internal/components/menu"
	"dragonbytes039/scaff/internal/msg"
)

var LanguagesOptions = []menu.OptionsInfo{
	{Description: "object-oriented programming", Msg: msg.ChangeToJsPooMsg{}},
	{Description: "functional programming", Msg: msg.ChangeToJsFuntionalMsg{}},
}

func InitialParadigmModel() menu.Model {

	return menu.InitialModel(LanguagesOptions, "select paradigm")
}
