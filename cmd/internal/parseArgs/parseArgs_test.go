package parseArgs

import "testing"

func TestIfArgsAreParsedCorrect(t *testing.T) {
	input := []string{"./cmd", "-uSeRId", "1", "-FilTEr", "magnam"}

	parsed := ParseArgs(input)

	if parsed.userID != "1" {
		t.Errorf("parsed user id should be 1, was %s", parsed.userID)
	}

	if parsed.filter != "magnam" {
		t.Errorf("parsed filter should be magnam, was: %s", parsed.filter)
	}
}
