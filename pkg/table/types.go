package table

type Type interface {
	// Size in bytes need to hold instance of the type
	Size() uint16
	// Name of the type
	Name() string
}

type Value interface {
	Type() Type
	Bytes() []byte
	Size() uint32
}
