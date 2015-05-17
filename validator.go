package magic
import (
	"github.com/samcrosoft/magic/types"
	"errors"
	"bytes"
)

// check if the supplied byte array has the same file signature as the type supplied
// a result of true will mean the byte array contains the valid file format and false means otherwise
// note: an error may be returned in the likely case where the length of byte array contains less than the minimum
// required
func IsBytesContentAValidType(oBytes []byte, oType types.ValidFileType) (bool, error){
	iMinHeader := int(oType.HeaderSize())

	// check length of bytes
	if iLen := len(oBytes); iLen < iMinHeader {
		return false, errors.New("Header Too Small")
	}

	// check the header
	aMinimumHeader := oBytes[:iMinHeader]
	oTypeSignature := types.Signature(oType)
	bValid := bytes.Compare(aMinimumHeader, oTypeSignature[:iMinHeader] )
	return bValid == 0, nil
}