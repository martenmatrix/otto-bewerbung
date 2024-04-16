package parseArgs

import "testing"

func TestIfArgsAreParsedCorrect(t *testing.T) {
	input := []string{"./cmd", "-uSeRId", "1", "-FilTEr", "magnam"}

	parsed, err := ParseArgs(input)

	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if parsed.UserID != 1 {
		t.Errorf("parsed user id should be 1, was %d", parsed.UserID)
	}

	if parsed.Filter != "magnam" {
		t.Errorf("parsed filter should be magnam, was: %s", parsed.Filter)
	}
}
