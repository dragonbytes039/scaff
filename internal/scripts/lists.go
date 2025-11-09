package scripts

type ScriptApp int

type IScript interface {
	Execute() error
}

const (
	// JavaScript Funcional
	JsFuntionalClean ScriptApp = iota // 0 (creal arquitecture)
	JsFuntionalHexa                   // 1 (Arquitectura Hexagonal)
	JsFuntionalMcs                    // 2 (model controller services)
	JsFuntionalMC                     // 3 (Modular clean)

	// JavaScript Orientado a Objetos (POO)
	JsPooClean // 4
	JsPooHexa  // 5
	JsPooMcs   // 6
	JsPooMC    // 7

	// Go
	GoClean // 8
	GoHexa  // 9
	GoMC    // 11 (Microservices/Microcomponents)
	GoEML   // 10 (Model-View-Controller)
)

var ScriptsDetailsUrls = map[ScriptApp]string{
	JsFuntionalClean: "https://github.com/dragonbytes02/js-clean-func-template.git",
	JsFuntionalHexa:  "https://github.com/dragonbytes02/js-hexagonal-func-template.git",
	JsFuntionalMC:    "https://github.com/dragonbytes02/js-modularclean-func-template.git",
	JsFuntionalMcs:   "https://github.com/dragonbytes02/js-mcs-func-template.git",

	JsPooClean: "https://github.com/dragonbytes02/js-clean-poo-template.git",
	JsPooHexa:  "https://github.com/dragonbytes02/js-hexagonal-poo-template.git",
	JsPooMcs:   "https://github.com/dragonbytes02/js-mcs-poo-template.git",
	JsPooMC:    "in process...",

	GoClean: "in process...",
	GoHexa:  "in process...",
	GoMC:    "in process...",
	GoEML:   "in process...",
}
