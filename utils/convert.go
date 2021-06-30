package utils

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime/debug"
	"strings"
)

func Struct2Struct(dst, src interface{}) error {
	if dst == nil || src == nil {
		return errors.New("param err")
	}
	dstVlf := reflect.ValueOf(dst)
	srcVlf := reflect.ValueOf(src)
	if dstVlf.Type().Kind() != reflect.Ptr {
		//dstVlf=dstVlf.Elem()
		return errors.New("dst must ptr")
	}
	dstVlf = dstVlf.Elem()
	if srcVlf.Type().Kind() == reflect.Ptr {
		srcVlf = srcVlf.Elem()
	}
	if dstVlf.Type().Kind() != reflect.Struct || srcVlf.Type().Kind() != reflect.Struct {
		return errors.New("des/src is not a struct")
	}
	dstFln := dstVlf.Type().NumField()
	for i := 0; i < dstFln; i++ {
		fdTyp := dstVlf.Type().Field(i)
		fdVlf := dstVlf.FieldByIndex(fdTyp.Index)

		tags := strings.Split(fdTyp.Tag.Get("convs"), ",")
		srcFnm := fdTyp.Name
		if tags[0] != "" {
			srcFnm = tags[0]
		}
		//srcfdTyp:=dstVlf.Type()
		srcfdTyp, ok := srcVlf.Type().FieldByName(srcFnm)
		if !ok {
			continue
		}
		err := saveSet(fdVlf, srcVlf.FieldByIndex(srcfdTyp.Index))
		if err == nil {
			continue
		}

		dstFdVlf := fdVlf
		srcFdVlf := srcVlf.FieldByIndex(srcfdTyp.Index)
		if srcFdVlf.Type().Kind() == reflect.Ptr {
			if srcFdVlf.IsZero() {
				continue
			}
			srcFdVlf = srcFdVlf.Elem()
		}
		if dstFdVlf.Type().Kind() == reflect.Ptr {
			ne := reflect.New(dstFdVlf.Type().Elem())
			err = saveSet(dstFdVlf, ne)
			if err != nil {
				return err
			}
			dstFdVlf = ne.Elem()
		}
		err = saveSet(dstFdVlf, srcFdVlf)
		if err == nil {
			continue
		}
		if dstFdVlf.Type().Kind() == reflect.Struct && srcFdVlf.Type().Kind() == reflect.Struct {
			err = Struct2Struct(dstFdVlf.Addr().Interface(), srcFdVlf.Interface())
			if err != nil {
				return err
			}
		}
		/*if dstFdVlf.Type().Kind() == reflect.Map&&srcFdVlf.Type().Kind()==reflect.Struct {
			Struct2Struct(dstVlf.Interface(),srcVlf.Interface())
		}*/
	}
	return nil
}
func saveSet(dst, src reflect.Value) (rterr error) {
	defer func() {
		if err := recover(); err != nil {
			rterr = fmt.Errorf("recover:%v", err)
			logrus.Warnf("convert saveSet recover:%v", err)
			logrus.Warnf("saveSet stack:%s", string(debug.Stack()))
		}
	}()
	if dst.Type() != src.Type() {
		return errors.New("saveSet type err")
	}
	dst.Set(src)
	return nil
}

/*func Map2Struct() {

}
*/
