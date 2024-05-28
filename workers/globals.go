package main

import "sync"

var (
	orphan      bool
	orphanMutex sync.Mutex
)

func SetOrphan(value bool) {
	orphanMutex.Lock()
	if value != orphan {
		orphan = value
		if orphan { // If orphan is set to true, trigger reassignment
			handleOrphanedStatus()
		}
	}
	orphanMutex.Unlock()
}

func GetOrphan() bool {
	orphanMutex.Lock()
	defer orphanMutex.Unlock()
	return orphan
}
