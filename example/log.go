package main

import (
	"goglib"
)

//---------------------------------------------------------------------------
// Log
//---------------------------------------------------------------------------
var Applog *goglib.OLog2 = goglib.InitLogEnv("./logdir", "process1", 2) // "path, stamp string, initial debug lv"
func main() {

	Applog.Print(1, "Debug #1")
	Applog.Print(2, "Debug #%d", 2)
	Applog.SetLevel(1)
	Applog.Print(1, "Debug #1")

	Applog.Always("Debug Always")
	Applog.Info("Debug Info")
	Applog.Warn("Debug Warn")
	Applog.Error("Debug Err")

	var arr []byte = []byte(string("123123123abcd12312312asdfasdfaosdfhaskdjf"))
	Applog.Dump(2, "rawdata", arr, len(arr))

	Applog.Fileclose()
}
