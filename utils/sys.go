package utils

import (
	"os"
	"runtime"
	"strings"
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

type EnvVal map[string]string

func AllEnv() EnvVal {
	rt := EnvVal{}
	envs := os.Environ()
	for _, v := range envs {
		i := strings.Index(v, "=")
		if i > 0 {
			k := v[:i]
			val := os.Getenv(k)
			if val != "" {
				rt[k] = val
			}
		}
	}
	return rt
}
func (c *EnvVal) SetOs() {
	os.Clearenv()
	for k, v := range *c {
		os.Setenv(k, v)
	}
}
