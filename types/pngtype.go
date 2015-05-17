package types

type PNGType struct {
}

// this will return the header signature for the PNG type
func (t PNGType) HeaderBytesSignature() []byte{
	aSignature := []byte {0x89, 0x50, 0x4E ,0x47}
	return aSignature
}

// set the minimum size required for the validation here
// note, this should be less than or equal to the size of the signature set above
func (t PNGType) HeaderSize() uint8{
	return 4;
}