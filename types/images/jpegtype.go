package images

type JpegType struct {

}

// this will return the header signature for the png type
func (t *JpegType) GetHeaderBytesSignature() []byte{
	aSignature := []byte {0xff, 0xd8, 0xff,0xe0}
	return aSignature
}


// this will return the minimum size required for the validation
func (t *JpegType) GetMinHeaderSize() uint{
	return 3;
}