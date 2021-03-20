package guid

import (
	"sync"
	"time"
)

type Guid struct {
	serverId      uint16
	lock          sync.RWMutex
	lastTimestamp int64
	sequence      int32
}

func (this *Guid) NewID() uint64 {
	this.lock.Lock()
	defer this.lock.Unlock()

	timestamp := time.Now().Unix()

	if timestamp == this.lastTimestamp {
		this.sequence += 1
	} else {
		this.sequence = 0
	}
	this.lastTimestamp = timestamp

	return uint64(timestamp<<40) | (uint64(this.serverId) << 12) | uint64(this.sequence)
}

func NewGuid(serverid uint16) *Guid {
	return &Guid{
		serverId:      serverid,
	}
}
