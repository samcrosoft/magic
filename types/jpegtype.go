package types

type JPEGType struct {

}

// set the header signature for the jpeg type
func (t *JPEGType) HeaderBytesSignature() []byte{
	aSignature := []byte {0xff, 0xd8, 0xff,0xe0}
	return aSignature
}


// set the minimum size required for the validation
func (t *JPEGType) HeaderSize() uint8{
	return 3;
}