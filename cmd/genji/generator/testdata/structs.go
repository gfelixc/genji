package testdata

// Basic is the simplest struct that can be used with Genji: No tags, no methods, no comments.
// This must not be generated by the generator.
type Basic struct {
	A    string
	B    int
	C, D int32
}

// basic is like Basic except that it is unexported.
type basic struct {
	A    []byte
	B    uint16
	C, D float32
}
