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

#### SSE (Server Send Event) management
```
*example

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/theroadlib/goglib"
)

var Applog *goglib.OLog2 = goglib.InitLogEnv("./log", "test", 0)

func sendSse(data goglib.EventData) {

	for _, actSession := range goglib.ActivesseSessionList {
		goglib.CheckSSEMsgChannel(actSession.Key)

		goglib.SseMsgChan[actSession.Key] <- data
	}
}

// ------------------------------------------------------------------------------
// processEventMsg
// ------------------------------------------------------------------------------
func ProcessEventMsg() {

	for {
		select {
		case event := <-goglib.ChEvent:
			Applog.Print(1, "Get Event message [%s]", event.Msgtype)

			if len(event.Msgtype) > 0 {
				msg := &event
				sendSse(*msg)
			} else {
				Applog.Error("undefined sse..[%s](%d)", event.Msgtype, event.Id)
			}

		}
	}
}

// ------------------------------------------------------------------------------
// handleSSE
// ------------------------------------------------------------------------------
func handleSSE() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get sse session key
		sessionKey := goglib.GetSSeSessionKey()
		defer func() {
			Applog.Print(3, "Close sse.. [%d]", sessionKey)
			goglib.ClearSSeSessionKey(sessionKey)
		}()

		if sessionKey == 0 {
			// invalid key...
			Applog.Error("Access handleSSE invalid key.. [%d]", sessionKey)
			<-r.Context().Done()
			return
		}

		// prepare the header
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// prepare the flusher
		flusher, _ := w.(http.Flusher)

		// trap the request under loop forever
		for {
			select {

			case <-r.Context().Done():
				return
			default:
				sseMsg, ok := goglib.PopSSEMsgChannel(sessionKey)
				if ok {
					btData := sseMsg.PrepareMessage()
					//am.Applog.Print(2, "PRepareMessage : %v", string(btData[:]))
					fmt.Fprintf(w, "%s\n", btData)

					flusher.Flush()
				}
			}
			time.Sleep(time.Millisecond * 5)
		}
	}
}

func makeSseMessage() {

	// send sse

	for {
		type UserInfo struct {
			Name string
			Age  int
		}
		var Info UserInfo
		Info.Age = 1
		Info.Name = "Hello"

		b, _ := json.Marshal(Info)
		var evdata goglib.EventData = goglib.EventData{}
		evdata.Msgtype = "message_type"
		evdata.Data = string(b)
		evdata.Id = "1"

		goglib.SendSSE(evdata)

		time.Sleep(time.Second * 2)
	}
}

// ------------------------------------------------------------------------------
// GetHelloWorld, get
// ------------------------------------------------------------------------------
func GetHelloWorld(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(string("hello world"))
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", GetHelloWorld).Methods("GET")
	r.HandleFunc("/sse", handleSSE())

	// sse msg routine
	go ProcessEventMsg()
	go makeSseMessage()

	http.ListenAndServe(":5000", r)
}



```