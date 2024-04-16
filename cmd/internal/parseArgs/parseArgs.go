package parseArgs

import (
	"strconv"
	"strings"
)

type Arguments struct {
	UserID int
	Filter string
}

/*
ParseArgs takes an array of strings, which is usually returned by os.Args
and extracts and returns the values specified for userID and filter.
*/
func ParseArgs(args []string) (Arguments, error) {
	extractedArgs := Arguments{}

	for index, arg := range args {
		lowercaseArg := strings.ToLower(arg)

		if lowercaseArg == "-userid" {
			userIdString := args[index+1]

			userId, err := strconv.Atoi(userIdString)
			if err != nil {
				return extractedArgs, err
			}

			extractedArgs.UserID = userId
		} else if lowercaseArg == "-filter" {
			extractedArgs.Filter = args[index+1]
		}
	}

	return extractedArgs, nil
}
