package types


type ValidFileType interface {
	HeaderBytesSignature() []byte
	HeaderSize() uint8
}



func Signature(v ValidFileType) []byte{
	return v.HeaderBytesSignature()
}