package goglib

import (
	"sync"
)

// ---------------------------------------------------------------------------
// Mutex
// ---------------------------------------------------------------------------
var rwMSSeSession = new(sync.RWMutex) // read, write vds session mutex
// ---------------------------------------------------------------------------
// Mutex(SSE mutex)
// ---------------------------------------------------------------------------
var sseMsgMTX = &sync.Mutex{} // sse message mutex

// ---------------------------------------------------------------------------
// SSE session key list
// ---------------------------------------------------------------------------
var sseSessionKeyList []SessionObj = makeSSeSessionKey()
var ActivesseSessionList []SessionObj = []SessionObj{}

// ---------------------------------------------------------------------------
// SSE channel
// ---------------------------------------------------------------------------
var SseMsgChan map[int]chan EventData = initSSeMsgChannel()

// ---------------------------------------------------------------------------
// makeSSeSessionKey
// ---------------------------------------------------------------------------
func makeSSeSessionKey() []SessionObj {
	sseKeyList := make([]SessionObj, 0)

	for idx := 1; idx <= CHBUF_SSE; idx++ {
		var session = SessionObj{Key: idx}
		sseKeyList = append(sseKeyList, session)
	}

	return sseKeyList
}

// ---------------------------------------------------------------------------
// initSSeMsgChannel
// ---------------------------------------------------------------------------
func initSSeMsgChannel() map[int]chan EventData {
	mpChannel := make(map[int]chan EventData)
	//세션 key별 채널 생성
	for key := 1; key <= CHBUF_SSE; key++ {
		mpChannel[key] = make(chan EventData, CHBUF_SSE)
	}

	return mpChannel
}

// ---------------------------------------------------------------------------
// PopSSEMsgChannel
// ---------------------------------------------------------------------------
func PopSSEMsgChannel(key int) (EventData, bool) {
	var popData EventData
	sseMsgMTX.Lock()
	if len(SseMsgChan[key]) > 0 {
		popData = <-SseMsgChan[key]
		sseMsgMTX.Unlock()

		return popData, true
	}
	sseMsgMTX.Unlock()

	return popData, false
}

func clearSSEMsgChannel(key int) {
	for len(SseMsgChan[key]) > 0 {
		PopSSEMsgChannel(key)
	}
}

func CheckSSEMsgChannel(key int) {
	if len(SseMsgChan[key]) >= CHBUF_SSE {
		clearSSEMsgChannel(key)
	}
}

func removeActiveSSeKey(index int) {
	ActivesseSessionList = append(ActivesseSessionList[:index], ActivesseSessionList[index+1:]...)
}

// ---------------------------------------------------------------------------
// GetSSeSessionKey
// ---------------------------------------------------------------------------
func GetSSeSessionKey() int {
	rwMSSeSession.Lock()

	var frontKeyObj SessionObj = SessionObj{}

	if len(sseSessionKeyList) > 0 {

		frontKeyObj, sseSessionKeyList = sseSessionKeyList[0], sseSessionKeyList[1:]
		ActivesseSessionList = append(ActivesseSessionList, frontKeyObj)

	}

	rwMSSeSession.Unlock()

	return frontKeyObj.Key
}

// ---------------------------------------------------------------------------
// ClearSSeSessionKey
// ---------------------------------------------------------------------------
func ClearSSeSessionKey(key int) {
	rwMSSeSession.Lock()

	sseSessionKeyList = append(sseSessionKeyList, SessionObj{Key: key})
	for index, data := range ActivesseSessionList {
		if data.Key == key {
			removeActiveSSeKey(index)
			break
		}
	}

	rwMSSeSession.Unlock()
}
