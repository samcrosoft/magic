package types


type ValidFileType interface {
	GetHeaderBytesSignature() []byte
	GetMinHeaderSize() uint
}



func GetSignature(v ValidFileType) []byte{
	return v.GetHeaderBytesSignature()
}