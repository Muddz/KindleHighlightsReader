package clean

import "strings"

func WhiteSpaces(text string) string {
	return strings.TrimSpace(text)
}

func Prefixes() {
	//#PreFixes
	//itself. A male runs the risk of losing far more women 	(\w+\S\s)
	//s. A male runs the risk of losing far more women			(\w+\S\s)
	//. A male runs the risk of losing far more women			(^\S\s)
}

func PostFixes() {

	//#PostFixes
	//A male runs the risk of losing far more women. itself 	(\.\s\w*)
	//A male runs the risk of losing far more women. i			(\.\s\w*)
}
