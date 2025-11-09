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
	JsFuntionalClean: "prueba",
}
