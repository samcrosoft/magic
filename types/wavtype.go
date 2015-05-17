package types


type WAVType struct {}

func (w *WAVType) HeaderBytesSignature() []byte{
	return []byte{0x52, 0x49, 0x46,0x46};
}


func (w *WAVType) HeaderSize() uint8{
	return 4;
}

