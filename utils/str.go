package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gokins/core/common"
	"github.com/gokins/core/runtime"
)

// RandomString 随机生成字符串
func RandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bts := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bts[r.Intn(len(bts))])
	}
	return string(result)
}

func RepVarString(origin string, vars map[string]*runtime.Variables, mustShow ...bool) string {
	ms := false
	if len(mustShow) == 1 && mustShow[0] {
		ms = true
	}
	all := common.RegVar.FindAllStringSubmatch(origin, -1)
	for _, v := range all {
		rv, ok := vars[v[1]]
		va := ""
		st := false
		if ok {
			st = rv.Secret
			va = rv.Value
		}
		if !ms && st {
			origin = strings.ReplaceAll(origin, v[0], "***")
		} else {
			origin = strings.ReplaceAll(origin, v[0], va)
		}
	}
	return origin
}
