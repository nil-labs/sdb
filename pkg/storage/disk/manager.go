package disk

import (
	"errors"
	"os"
)

var (
	ErrDBFileNotFound = errors.New("DB File Not Found")
)

// PAGE_SIZE defines the page size inn bites
// TODO potentially swtich from a fixed size page size to detecting the Page Size based on the OS/ARCH
const PAGE_SIZE = 4096

// Manager struct abstracts the persistance of pages on disk
type Manager struct {
	db  *os.File
	log *os.File
}

// Page smallest set of data readable/writable to the FS
type Page [PAGE_SIZE]byte

// PageId the offset at which the Page starts in the DB file
type PageId int

// Read()
// Write()
// Close()

// ManagerFromFile constructs a Manager based on the provided db file
func ManagerFromFile(db *os.File) (*Manager, error) {
	if db == nil {
		return nil, ErrDBFileNotFound
	}

	if _, err := os.Stat(db.Name()); errors.Is(err, os.ErrNotExist) {
		return nil, ErrDBFileNotFound
	}

	return &Manager{
		db: db,
	}, nil
}

// WritePage writes a page to the DB file
func WritePage(id PageId, page Page) error {
	return nil
}
