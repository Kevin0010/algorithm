package tool

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var num = os.Args[1]

func Run(fn interface{}, args ...interface{}) {
	_, file, _, _ := runtime.Caller(1)
	fileNameIdx := strings.LastIndex(file, "/") + 1
	dotIdx := strings.Index(file, ".")
	if num != file[fileNameIdx:dotIdx] {
		return
	}
	printResult(file[fileNameIdx:], doRun(fn, args...)[0].Interface())
	os.Exit(0)
}

func RunAndPrintSpec(p interface{}, fn interface{}, args ...interface{}) {
	_, file, _, _ := runtime.Caller(1)
	fileNameIdx := strings.LastIndex(file, "/") + 1
	dotIdx := strings.Index(file, ".")
	if num != file[fileNameIdx:dotIdx] {
		return
	}
	doRun(fn, args...)
	printResult(file[fileNameIdx:], reflect.ValueOf(p).Elem().Interface())
	os.Exit(0)
}

func doRun(fn interface{}, args ...interface{}) []reflect.Value {
	var in []reflect.Value
	for i := range args {
		in = append(in, reflect.ValueOf(args[i]))
	}
	return reflect.ValueOf(fn).Call(in)
}

func printResult(title string, i interface{}) {
	fmt.Println("-----", title, "-----")
	switch x := i.(type) {
	case bool, int, int32, int64, string, []int, []string:
		fmt.Println(x)
	case [][]int:
		for i2 := range x {
			row := x[i2]
			fmt.Println(row)
		}
	default:
		marshal, e := json.Marshal(x)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(string(marshal))
	}
}
