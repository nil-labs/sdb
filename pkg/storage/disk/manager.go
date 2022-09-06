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
type Page []byte

// PageId the offset at which the Page starts in the DB file
type PageId int64

// Read()
// Write()
// Close()

// ManagerFromFile constructs a Manager based on the provided db file
func ManagerFromFile(file string) (*Manager, error) {

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return nil, ErrDBFileNotFound
	}
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &Manager{
		db: f,
	}, nil
}

// WritePage to the DB file
func (m *Manager) WritePage(id PageId, page Page) error {
	offset := int64(id * PAGE_SIZE)
	n, err := m.db.WriteAt(page[:], offset)

	if err != nil {
		return err
	}
	if n != PAGE_SIZE {
		return errors.New("Less than a full page written")
	}
	return m.db.Sync()
}

// ReadPage from the DB file
func (m *Manager) ReadPage(id PageId, page Page) error {
	offset := int64(id * PAGE_SIZE)
	n, err := m.db.ReadAt(page, offset)
	if err != nil {
		return err
	}
	if n != PAGE_SIZE {
		return errors.New("Less than a full page readed")
	}
	return nil
}

func (m *Manager) Close() error {
	return m.db.Close()
}

func NewPage() Page {
	return make(Page, PAGE_SIZE, PAGE_SIZE)
}
