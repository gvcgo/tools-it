package cloc

import (
	"strings"
	"testing"
)

func TestCloc(t *testing.T) {
	vmrDir := `D:\projects\go\src\gogvc\version-manager`
	result := Cloc(
		false,
		true,
		"files",
		"",
		"",
		"",
		"",
		"",
		"",
		vmrDir,
	)
	if !strings.Contains(result, "Go") {
		t.Error("Cloc failed to count Go files")
	} else {
		t.Log(result)
	}

	result = Cloc(
		true,
		true,
		"files",
		"",
		"",
		"",
		"",
		"",
		"",
		vmrDir,
	)

	if !strings.Contains(result, "D:") {
		t.Error("Cloc failed to count by files")
	} else {
		t.Log(result)
	}
}
