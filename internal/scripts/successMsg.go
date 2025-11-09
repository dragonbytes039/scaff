package scripts

func SuccessMsg(selection ScriptApp) string {
	switch selection {
	case JsFuntionalClean, JsFuntionalHexa, JsFuntionalMcs, JsFuntionalMC, JsPooClean, JsPooHexa, JsPooMcs, JsPooMC:
		return "\nconfigure your environment:\nexecute:\n\npnpm i\n"
	}
	return ""
}
