package sessions

import (
	"log"
	"sync"
)

var (
	frontSessions     = make(map[uint64]*FrontSession)
	frontSessionMutex sync.Mutex
)

func AddFrontSession(session *FrontSession) {
	frontSessionMutex.Lock()
	defer frontSessionMutex.Unlock()

	frontSessions[session.RId] = session
	log.Println("AddFrontSession", session.RId)
}

func RemoveFrontSession(rid uint64) {
	frontSessionMutex.Lock()
	defer frontSessionMutex.Unlock()

	log.Println("removeFrontSession")
	delete(frontSessions, rid)
}

func GetFrontSession(rid uint64) *FrontSession {
	frontSessionMutex.Lock()
	defer frontSessionMutex.Unlock()
	session, _ := frontSessions[rid]
	if session == nil {
		log.Println("GetFrontSession : error", rid, frontSessions)
	}
	return session
}

func FrontSessionLen() int {
	frontSessionMutex.Lock()
	defer frontSessionMutex.Unlock()

	return len(frontSessions)
}

func FetchFrontSession(callback func(*FrontSession)) {
	frontSessionMutex.Lock()
	defer frontSessionMutex.Unlock()

	for _, session := range frontSessions {
		callback(session)
	}
}
