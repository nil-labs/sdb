package storage

import "encoding/binary"

// PageSize sets the page size in bytes
const (
	PageSize = 4096
)

type Page struct {
	bytes [PageSize]byte
}

/**
 * Slotted page format:
 *  ---------------------------------------------------------
 *  | HEADER | ... FREE SPACE ... | ... INSERTED TUPLES ... |
 *  ---------------------------------------------------------
 *                                ^
 *                                free space pointer
 *
 *  Header format (size in bytes):
 *  ----------------------------------------------------------------------------
 *  | PageId (4)| LSN (4)| PrevPageId (4)| NextPageId (4)| FreeSpacePointer(4) |
 *  ----------------------------------------------------------------------------
 *  ----------------------------------------------------------------
 *  | TupleCount (4) | Tuple_1 offset (4) | Tuple_1 size (4) | ... |
 *  ----------------------------------------------------------------
 *
 */

func (p *Page) ID() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[0:4])
}

func (p *Page) LSN() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[4:9])
}
func (p *Page) PreviousPageID() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[9:14])
}
func (p *Page) NextPageID() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[14:19])
}
func (p *Page) FreeSpacePointer() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[19:24])
}
func (p *Page) TuplesCount() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[24:29])
}
func (p *Page) TuplePointer() uint32 {
	return binary.LittleEndian.Uint32(p.bytes[29:34])
}

func (p *Page) InsertTuple() error {
	return nil
}
func (p *Page) DeleteTuple() error {
	return nil
}

/**
 * Tuple format:
 * ---------------------------------------------------------------------
 * | FIXED-SIZE or VARIED-SIZED OFFSET | PAYLOAD OF VARIED-SIZED FIELD |
 * ---------------------------------------------------------------------
 */

type Tuple struct {
	bytes []byte
	size  uint32
	id    RID
}

type RID struct {
	pageID uint32
	slot   uint32
}
