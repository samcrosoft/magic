package magic
import (
	"testing"
	"io/ioutil"
	"github.com/samcrosoft/magic/types"
	"path/filepath"
)

// this will test the png file and it should pass
func TestPngData(t *testing.T){
	sImagePath := "./corpus/image/png-sample.png"
	if oBytes, err := ioutil.ReadFile(sImagePath); err == nil{
		oType := types.PNGType{}
		if bValid, oErr := IsBytesContentAValidType(oBytes, oType); bValid == false{
			if oErr != nil{
				t.Error("Image Has Less Than The Required Header Bytes")
			}
		}
	}
}


// this will test the jpeg file
func TestJpegData(t *testing.T){
	sImagePath := "./corpus/image/jpeg-sample.jpg"
	sImagePath,_ = filepath.Abs(sImagePath)
	if oBytes, err := ioutil.ReadFile(sImagePath); err == nil{
		if bValid, oErr := IsBytesContentAValidType(oBytes, &types.JPEGType{}); bValid == false{
			if oErr != nil{
				t.Error("Image Has Less Than The Required Header Bytes")
			}else{
				t.Error("Image Is Not A Jpeg File")
			}
		}
	}
}