package layout

import "fmt"

func (m AppModel) View() string {

	switch m.state {
	case RunningScript:
		return fmt.Sprintf("\n%s configuring project...", m.Spinner.View())

	case ScriptError:
		return fmt.Sprintf("\n%s ", m.ScriptError)

	case ScriptFinished:
		return fmt.Sprintf("\n\nfinished successfully\n%s  \n", m.ScriptFinished)

	case LanguageModel:
		return m.Language.View()

	case ParadigmModel:
		return m.Paradigm.View()

	case JsPoo:
		return m.JsPoo.View()

	case JsFuntional:
		return m.JsFuntional.View()

	default:
		return "unknown error"
	}
}
