package common

import "regexp"

const (
	TimeFmt    = "2006-01-02 15:04:05"
	TimeFmts   = "2006-01-02"
	TimeFmtm   = "2006-01"
	TimeFmtpck = "2006-01-02T15:04:05.999999999Z"
	TimeFmtVer = "20060102150405.999999999"
)

var (
	RegHost1 = regexp.MustCompile(`^([\w\.]+)?:(\d+)?$`)
	RegHost2 = regexp.MustCompile(`^([\w\.]+)(:\d+)?$`)
	RegUrl   = regexp.MustCompile(`^(https?:)\/\/([\w\.]+)(:\d+)?`)
	RegNum   = regexp.MustCompile(`^\d+$`)
	RegVar   = regexp.MustCompile(`\${{\s*([^}\s]+)\s*}}`)
)
