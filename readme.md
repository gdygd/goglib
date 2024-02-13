# 'Go general library'

`goglib` is a Go general lib in pure Go. For now, it only includes log lib.

## Installation

    go get -u github.com/gdygd/goglib

#### 1.Log library
- thread safe

##### log example

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
- output
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

#### 2.general library
 - func CheckElapsedTime(timer *time.Time, msDuration int) bool   // check elapsedtime
 - func MinInt(a int, b int) int  // check min value
 - func MaxInt(a int, b int) int  // check max value

 - func GetNumber(src []byte, pos int, length int, endian int) int  // Get Big-endian or little-endian value
 - func SetNumber(buf []byte, pos int, value int, length int, endian int) // Set Big-endian or little-endian value (32bit)
 - func SetNumber2(buf []byte, pos int, value int, length int, endian int)  // Set Big-endian or little-endian value (64bit)
 - func GenLRC(buf []byte, pos int, lastIdx int) byte

##### CheckElapsedTime example

  ```go
package main

import (
	"fmt"
	"goglib"
	"time"
)

func main() {
	const CHECK_INTERVAL1 = 100  // 100 millisecond
	const CHECK_INTERVAL2 = 1000 // 1000 millisecond

	var markTime1 time.Time = time.Now()

	if goglib.CheckElapsedTime(&markTime1, CHECK_INTERVAL1) {
		fmt.Println("100 millisecond was elapsed")
	} else {
		fmt.Println("100 millisecond wasn't elapsed")
	}

	var markTime2 time.Time = time.Now()
	time.Sleep(time.Second * 2)
	if goglib.CheckElapsedTime(&markTime2, CHECK_INTERVAL2) {
		fmt.Println("1000 millisecond was elapsed")
	} else {
		fmt.Println("1000 millisecond wasn't elapsed")
	}

}

  ```
- output
  ```
  100 millisecond wasn't elapsed
  1000 millisecond was elapsed

  ```

##### Endian example

  ```go

package main

import (
	"goglib"
)

var Applog *goglib.OLog2 = goglib.InitLogEnv("./logdir", "process1", 2) // "path, stamp string, initial debug lv"
func main() {
	var buf1 []byte = make([]byte, 2)
	var buf2 []byte = make([]byte, 4)
	var buf3 []byte = make([]byte, 8)

	goglib.SetNumber(buf1, 0, 1, 2, goglib.ED_BIG)
	goglib.SetNumber(buf2, 0, 2, 4, goglib.ED_BIG)
	goglib.SetNumber(buf3, 0, 1152921504606846977, 8, goglib.ED_LITTLE)

	Applog.Print(2, "buf1 (big endian)")
	Applog.Dump(2, "buf1 rawdata", buf1, len(buf1))

	Applog.Print(2, "buf2 (big endian)")
	Applog.Dump(2, "buf2 rawdata", buf2, len(buf2))

	Applog.Print(2, "buf3 (little endian)")
	Applog.Dump(2, "buf3 rawdata", buf3, len(buf3))

	val1 := goglib.GetNumber(buf1, 0, 2, goglib.ED_BIG)
	val2 := goglib.GetNumber(buf2, 0, 4, goglib.ED_BIG)
	val3 := goglib.GetNumber(buf3, 0, 8, goglib.ED_LITTLE)

	Applog.Print(2, "buf1 val : %d", val1)
	Applog.Print(2, "buf2 val : %d", val2)
	Applog.Print(2, "buf3 val : %d", val3)

	Applog.Fileclose()
}
  ```

- output
  ```
  PRINT  : [process1]2024/02/13 14:54:40.775770 buf1 (big endian)
  DUMP   : [process1]2024/02/13 14:54:40.775823   buf1 rawdata  [2]
          00 01

  PRINT  : [process1]2024/02/13 14:54:40.775861 buf2 (big endian)
  DUMP   : [process1]2024/02/13 14:54:40.775897   buf2 rawdata  [4]
          00 00 00 02

  PRINT  : [process1]2024/02/13 14:54:40.775931 buf3 (little endian)
  DUMP   : [process1]2024/02/13 14:54:40.775970   buf3 rawdata  [8]
          01 00 00 00 00 00 00 10

  PRINT  : [process1]2024/02/13 14:54:40.776011 buf1 val : 1
  PRINT  : [process1]2024/02/13 14:54:40.776043 buf2 val : 2
  PRINT  : [process1]2024/02/13 14:54:40.776071 buf3 val : 1152921504606846977

  ```