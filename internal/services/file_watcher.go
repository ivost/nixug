package services

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"sync"
	"time"
)

type FileWatcher struct {
	Lock        sync.RWMutex
	Ch          chan bool
	fw          *fsnotify.Watcher
	path        string
	fileChanged bool
}

func NewFileWatcher(path string) (*FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if check(err) {
		return nil, err
	}
	if err = watcher.Add(path); check(err) {
		return nil, err
	}
	return &FileWatcher{fw: watcher, path: path, Ch: make(chan bool)}, nil
}

func (w *FileWatcher) HasChanged() bool {
	w.Lock.Lock()
	defer w.Lock.Unlock()
	return w.fileChanged
}

func (w *FileWatcher) SetDirty(value bool) bool {
	w.Lock.Lock()
	old := w.fileChanged
	w.fileChanged = value
	w.Lock.Unlock()
	return old
}

// Watch will block and check for file change events
func (w *FileWatcher) Watch() {
	var err error
	for {
		// set dirty flag on WRITE events
		event, ok := <-w.fw.Events
		log.Printf("event %+v, ok %v", event, ok)
		//&& event.Op&fsnotify.Write == fsnotify.Write
		if !ok {
			continue
		}
		w.SetDirty(true)
		// many editors do backup = RENAME - need new watcher
		if event.Op&fsnotify.Rename != fsnotify.Rename {
			continue
		}
		// got rename
		w.fw.Close()
		w.fw = nil
		// need small delay for rename/chmod to complete
		// or loop with stat
		for i := 0; i < 100; i++ {
			if _, err = os.Stat(w.path); err == nil {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
		if w.fw, err = fsnotify.NewWatcher(); check(err) {
			w.fw = nil
			break
		}
		if err = w.fw.Add(w.path); check(err) {
			break
		}
	}

}
