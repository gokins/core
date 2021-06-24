package common

import "regexp"

var (
	RegHost1 = regexp.MustCompile(`^([\w\.]+)?:(\d+)?$`)
	RegHost2 = regexp.MustCompile(`^([\w\.]+)(:\d+)?$`)
	RegUrl   = regexp.MustCompile(`^(https?:)\/\/([\w\.]+)(:\d+)?`)
	RegNum   = regexp.MustCompile(`^\d+$`)
)
