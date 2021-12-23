package sturdy

import (
	"path/filepath"

	"github.com/gofrs/flock"
)

const (
	// lockFile is the name of the lock file used to prevent concurrent
	// synchronize mutagen with sturdy server.
	lockFile = ".git/sturdy.lock"
)

type Mutex interface {
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}

type noMutex struct{}

func (noMutex) Lock() error {
	return nil
}

func (noMutex) Unlock() error {
	return nil
}

func (noMutex) RLock() error {
	return nil
}

func (noMutex) RUnlock() error {
	return nil
}

type flockMutex struct {
	flck *flock.Flock
}

func (f *flockMutex) Lock() error {
	return f.flck.Lock()
}

func (f *flockMutex) Unlock() error {
	return f.flck.Unlock()
}

func (f *flockMutex) RLock() error {
	return f.flck.RLock()
}

func (f *flockMutex) RUnlock() error {
	return f.flck.Unlock()
}

func CreateMutex(roots string, labels map[string]string) Mutex {
	if labels == nil {
		return noMutex{}
	}
	if !useLock(labels) {
		return noMutex{}
	}
	return &flockMutex{
		flck: flock.New(filepath.Join(roots, lockFile)),
	}
}
