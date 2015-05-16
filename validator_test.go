package bytescontentvalidator
import (
	"testing"
	"io/ioutil"
	"github.com/samcrosoft/bytescontentvalidator/types/images"
	"path/filepath"
)

// this will test the png file and it should pass
func TestPngData(t *testing.T){
	sImagePath := "./corpus/image/png-sample.png"
	if oBytes, err := ioutil.ReadFile(sImagePath); err == nil{
		oType := &images.PngType{}
		if bValid, oErr := IsBytesContentAValidType(oBytes, oType); bValid == false{
			if oErr != nil{
				t.Error("Image Has Less Than The Required Header Bytes")
			}
		}
	}
}


// this will test the jpeg file
func TestInvalidPngData(t *testing.T){
	sImagePath := "./corpus/image/jpeg-sample.jpg"
	sImagePath,_ = filepath.Abs(sImagePath)
	if oBytes, err := ioutil.ReadFile(sImagePath); err == nil{
		if bValid, oErr := IsBytesContentAValidType(oBytes, &images.JpegType{}); bValid == false{
			if oErr != nil{
				t.Error("Image Has Less Than The Required Header Bytes")
			}else{
				t.Error("Image Is Not A Jpeg File")
			}
		}
	}
}