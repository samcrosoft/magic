package bytescontentvalidator
import (
	"github.com/samcrosoft/bytescontentvalidator/types"
	"errors"
	"reflect"
)


func IsBytesContentAValidType(oBytes []byte, oType types.ValidFileType) (bool, error){
	iMinHeader := int(oType.GetMinHeaderSize())

	// check length of bytes
	if iLen := len(oBytes); iLen < iMinHeader {
		return false, errors.New("Header Too Small")
	}

	// check the header
	aMinimumHeader := oBytes[:iMinHeader]
	oTypeSignature := types.GetSignature(oType)
	bValid := reflect.DeepEqual(aMinimumHeader, oTypeSignature[:iMinHeader] )
	return bValid, nil
}