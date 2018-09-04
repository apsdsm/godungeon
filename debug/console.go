package debug

import "fmt"

var (
	log []string
)

func Log(s string) {
	logLine := fmt.Sprintf("%d: %s", len(log)+1, s)
	log = append(log, logLine)
}

func Clear() {
	log = make([]string, 0)
}

func Tail(lines int) []string {
	len := len(log)
	var start int

	start = len - lines

	if start < 0 {
		start = 0
	}

	return log[start:]
}
