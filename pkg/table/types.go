package table

type Type interface {
	// Size in bytes need to hold instance of the type
	Size() uint16
	// Name of the type
	Name() string
}

type Value struct {
	t     Type
	bytes []byte
}

// type strict interface {
// 	// Type of the value
// 	Type() Type
// 	// Bytes holds the byte representation of the value
// 	Bytes() []byte
// }
