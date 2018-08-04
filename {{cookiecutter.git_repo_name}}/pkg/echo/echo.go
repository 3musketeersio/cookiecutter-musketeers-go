package echo

import (
	"os"
)

// Echo returns the value of ECHO_MESSAGE
func Echo() string {
	return os.Getenv("ECHO_MESSAGE")
}
