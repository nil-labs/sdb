package storage

// PageSize sets the page size in bytes
const PageSize = 4096

// Page is the main storage abstraction
// all data is represented as pages and pages are stored in persistent storage and also
// temp cached in main memory
type page struct {
	id   int64
	data []byte
}

func (p *page) ID() int64 {
	return p.id
}

func (p *page) Data() []byte {
	return p.data
}
func (p *page) Size() int {
	return PageSize
}

func NewPage() Page {
	return &page{
		data: make([]byte, PageSize),
	}
}

type Page interface {
	ID() int64
	Data() []byte
	Size() int
}
