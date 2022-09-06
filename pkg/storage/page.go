package storage

// PAGE_SIZE sets the page size in bytes
const PAGE_SIZE = 4096

// Page is the main storage abstraction
// all data is represented as pages and pages are stored in persistent storage and also
// temp cached in main memory
type page struct {
	id   int64
	data []byte
}

func (p *page) Id() int64 {
	return p.id
}

func (p *page) Data() []byte {
	return p.data
}
func (p *page) Size() int {
	return PAGE_SIZE
}

func NewPage() *page {
	return &page{
		data: make([]byte, PAGE_SIZE, PAGE_SIZE),
	}
}

type Page interface {
	Id() int64
	Data() []byte
	Size() int
}
