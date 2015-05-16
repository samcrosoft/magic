package images

type PngType struct {

}

// this will return the header signature for the png type
func (t PngType) GetHeaderBytesSignature() []byte{
	aSignature := []byte {0x89, 0x50, 0x4E ,0x47}
	return aSignature
}

// this will return the minimum size required for the validation
func (t PngType) GetMinHeaderSize() uint{
	return 4;
}