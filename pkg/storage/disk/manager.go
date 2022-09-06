package disk

import (
	"errors"
	"os"
)

var ErrDBFileNotFound = errors.New("DB File Not Found")
var ErrLogFileNotFound = errors.New("Log File Not Found")

// Manager struct abstracts the persistance of pages on disk
type Manager struct {
	db  *os.File
	log *os.File
}

// Read()
// Write()
// Close()

// ManagerFromFiles constructs a Manager based on the provided log and db files
func ManagerFromFiles(db, log *os.File) (*Manager, error) {
	if db == nil {
		return nil, ErrDBFileNotFound
	}
	if log == nil {
		return nil, ErrLogFileNotFound
	}
	if _, err := os.Stat(db.Name()); errors.Is(err, os.ErrNotExist) {
		return nil, ErrDBFileNotFound
	}
	if _, err := os.Stat(log.Name()); errors.Is(err, os.ErrNotExist) {
		return nil, ErrLogFileNotFound
	}
	return &Manager{
		db:  db,
		log: log,
	}, nil
}
