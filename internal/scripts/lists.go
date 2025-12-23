package scripts

type ScriptApp int

type IScript interface {
	Execute() error
}

const (
	// JavaScript Funcional
	JsFuntionalClean ScriptApp = iota // 0 (clean arquitecture)
	JsFuntionalHexa                   // 1 (Arquitectura Hexagonal)
	JsFuntionalMcs                    // 2 (model controller services)
	JsFuntionalMC                     // 3 (Modular clean)

	// JavaScript Orientado a Objetos (POO)
	JsPooClean // 4
	JsPooHexa  // 5
	JsPooMcs   // 6
	JsPooMC    // 7

	// Go
	GoHexa // 8
	GoMcs  // 9
	GoMC   // 10
)

var TemplateName = map[ScriptApp]string{
	JsFuntionalClean: "js-clean-func-template",
	JsFuntionalHexa:  "js-hexagonal-func-template",
	JsFuntionalMC:    "js-modularclean-func-template",
	JsFuntionalMcs:   "js-mcs-func-template",

	JsPooClean: "js-clean-poo-template",
	JsPooHexa:  "js-hexagonal-poo-template",
	JsPooMC:    "js-modularclean-poo-template",
	JsPooMcs:   "js-mcs-poo-template",

	GoHexa: "go-hexagonal-template",
	GoMcs:  "go-mcs-template",
	GoMC:   "go-modularclean-template",
}
