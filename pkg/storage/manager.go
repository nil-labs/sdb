package storage

import (
	"errors"
	"os"
)

var (
	ErrDBFileNotFound = errors.New("DB File Not Found")
)

// Manager struct abstracts the persistance of pages on disk
type Manager struct {
	db *os.File
}

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

// WritePage to the DB file - thread safe as the underlying access to the store is sync-ed
func (m *Manager) WritePage(p Page) error {
	offset := p.ID() * int64(p.Size())
	n, err := m.db.WriteAt(p.Data(), offset)

	if err != nil {
		return err
	}
	if n != p.Size() {
		return errors.New("less than a full page written")
	}
	return m.db.Sync()
}

// ReadPage from the DB file
// thread safe - can be called by multiple go routines
func (m *Manager) ReadPage(p Page) error {
	offset := p.ID() * int64(p.Size())
	n, err := m.db.ReadAt(p.Data(), offset)
	if err != nil {
		return err
	}
	if n != p.Size() {
		return errors.New("Less than a full page red")
	}
	return nil
}

func (m *Manager) Close() error {
	return m.db.Close()
}
