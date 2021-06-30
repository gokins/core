package core

import (
	"fmt"
	"github.com/gokins-main/core/utils"
	"testing"
)

type testBase1 struct {
	I int
	K string
}
type testBase2 struct {
	I int
	K string
}
type test1 struct {
	a  int
	B  int64
	C  int `convs:"Cs"`
	Ts testBase1
	Mp *map[string]interface{}
}
type test2 struct {
	a  int64
	B  int64
	Cs int
	Ts *testBase2
	Mp *map[string]interface{}
}

func Test1(t *testing.T) {
	t1 := test1{1, 2, 3, testBase1{
		I: 123,
		K: "hellos",
	}, &map[string]interface{}{
		"abcd": "12345",
	}}
	t2 := &test2{}
	err := utils.Struct2Struct(t2, t1)
	if err != nil {
		println("err:" + err.Error())
		return
	}
	println(fmt.Sprintf("t1:%v", t1))
	println(fmt.Sprintf("t2:%v", t2))
	println(fmt.Sprintf("t2s:%v", t2.Ts))
}
