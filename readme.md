# 'Go general library'

`goglib` is a Go general lib in pure Go. For now, it only includes log lib.

## Installation

    go get -u github.com/gdygd/goglib

#### 1.Log library
- thread safe

##### example

  ```go
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
  ```
- output:
  ```
PRINT  : [process1]2023/11/28 13:34:50.234569 Debug #2  
PRINT  : [process1]2023/11/28 13:34:50.234617 Debug #1  
ALWAYS : [process1]2023/11/28 13:34:50.234641 Debug Always  
WARN   : [process1]2023/11/28 13:34:50.234665 Debug Warn  
ERROR  : [process1]2023/11/28 13:34:50.234693 Debug Err  
DUMP   : [process1]2023/11/28 13:34:50.234775   rawdata  [41]  
         31 32 33 31 32 33 31 32 33 61 62 63 64 31 32 33 31 32 33 31  
         32 61 73 64 66 61 73 64 66 61 6F 73 64 66 68 61 73 6B 64 6A  
         66  
  ```