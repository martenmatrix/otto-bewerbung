package parseArgs

import "strings"

type Arguments struct {
	userID string
	filter string
}

/*
ParseArgs takes an array of strings, which is usually returned by os.Args
and extracts and returns the values specified for userID and filter.
*/
func ParseArgs(args []string) Arguments {
	extractedArgs := Arguments{}

	for index, arg := range args {
		lowercaseArg := strings.ToLower(arg)

		if lowercaseArg == "-userid" {
			extractedArgs.userID = args[index+1]
		} else if lowercaseArg == "-filter" {
			extractedArgs.filter = args[index+1]
		}
	}

	return extractedArgs
}
