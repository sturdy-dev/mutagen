package sturdy

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

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

type fileLock struct {
	mu      sync.RWMutex // lock within this process
	lock    *flock.Flock // lock between processes
	countMx sync.Mutex

	count uint
}

func New(filename string) *fileLock {
	return &fileLock{
		lock: flock.New(filename),
	}
}

func (fl *fileLock) Lock() error {
	fl.mu.Lock()
	if err := fl.lock.Lock(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return nil
}

func (fl *fileLock) Unlock() error {
	defer fl.mu.Unlock()
	if err := fl.lock.Unlock(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return nil
}

func (fl *fileLock) RLock() error {
	fl.countMx.Lock()
	fl.count++
	fl.countMx.Unlock()

	fl.mu.RLock()
	if err := fl.lock.RLock(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return nil
}

func (fl *fileLock) RUnlock() error {
	fl.countMx.Lock()
	fl.count--
	if fl.count == 0 {
		if err := fl.lock.Unlock(); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fl.mu.RUnlock()
				fl.countMx.Unlock()
				return nil
			}
			fl.mu.RUnlock()
			fl.countMx.Unlock()
			return err

		}
	}
	fl.mu.RUnlock()
	fl.countMx.Unlock()
	return nil
}

func CreateMutex(roots string, labels map[string]string) Mutex {
	if labels == nil {
		return noMutex{}
	}
	if !useLock(labels) {
		return noMutex{}
	}
	return &fileLock{
		lock: flock.New(filepath.Join(roots, lockFile)),
	}
}
