# 'dev library'
`goglib` is a backend dev library for linux and windows
1. it includes byte ordering, user error, time mangement, process and thread management
2. it includes sse(server send event), svg parser, 
3. it includes serial library, 

#### Byte ordering
- **func GetNumber(src []byte, pos int, length int, endian int) int**
```
A function that reads in Big Endian or Little Endian
input
 src : byte arrary
 pos : start index of array
 length : read length
 enddian : ED_BIG or ED_LITTLE

output :  Big or Little endian value

*example
pMessage => [0,0,0,1,0,0,0,2]
num1 := GetNumber(*pMessage, 0, 4, general.ED_BIG) // num : 1
num2 := GetNumber(*pMessage, 4, 4, general.ED_BIG) // num : 2
```

- **SetNumber(buf []byte, pos int, value int, length int, endian int)**
```
A function that writes in Big Endian or Little Endian
input
 buf : byte arrary
 pos : start index of array
 length : read length
 enddian : ED_BIG or ED_LITTLE

output :  Big or Little endian value

*example
info := make([]byte, 10)
idx:=0
SetNumber(info, idx, 1, 2, general.ED_BIG)
info => [0,2,0,0,0,0,0,0,0,0]

```

#### Time management
- **func CheckElapsedTime(timer *time.Time, msDuration int) bool**
```
A function that reads in Big Endian or Little Endian
input
 timer : 
 msDurations : 
 

output :  true > over duration, false not over duration

*example
sendtime := time.Now()
// CONNECT_INTERVAL : 500 (500ms)
isElapsed := CheckElapsedTime(&sendtime, CONNECT_INTERVAL)
```

#### Thread management
```
*example

package main

import (
	"log"
	"time"

	"gitlab.com/theroadlib/goglib"
)

var thrfunc *goglib.Thread = nil // msg process thread

func Threadf(t *goglib.Thread, chThrStop chan bool, arg1, arg2, arg3 interface{}) {

	var terminate = false

	//------------------------------------
	// thread loop
	//------------------------------------
	for {
		select {
		case <-chThrStop:
			// quit thread
			terminate = true
			break
		default:
			//
			//
			// code...
		}

		if terminate {
			break
		}

		log.Println("run thread..")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	thrfunc = goglib.NewThread() // make instance
	thrfunc.Init(Threadf, 10)    // init thread object
	thrfunc.Start()              // run thread
	time.Sleep(time.Second * 2) 
	thrfunc.Kill()              // quit thread

}


```