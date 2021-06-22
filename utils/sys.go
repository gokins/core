package utils

import (
	"os"
	"runtime"
)

func HomePath() string {
	home := ""
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}
	return home
}

func EnvDefault(nm string, defs ...string) string {
	s := os.Getenv(nm)
	if s == "" && len(defs) > 0 {
		s = defs[0]
	}
	return s
}
