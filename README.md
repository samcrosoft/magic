# bytescontentvalidator

This is a simple GoLang package that validates the content of a byte array to a file type. The need for this package became evident when I needed to validate a byte array as a png file 
to then forwarded to a waveform point extractor.

------

Installation
============
To install bytescontentvalidator, you can `go get` this like you do any other go package as follows

    go get github.com/samcrosoft/bytecontentsvalidator
------

Example
============

Want to see how simple it is to use this package,
```Go
    sImagePath := "./corpus/image/png-sample.png"
	if oBytes, err := ioutil.ReadFile(sImagePath); err == nil{
		oType := &images.PngType{}
		if bValid, oErr := IsBytesContentAValidType(oBytes, oType); bValid == true{
			// then image is a valid png
		}
	}
```

At the moment, there are a couple of image type formats

1. images.PngType
2. images.JpegType

## Tests
 
To run the tests, just run go tests
 
## Contributors

To Contribute to this project is just simple, fork the project, create a new filetype with a unique magic number identifying the file type and commit it.

## License

This is licensed in MIT, you can do everything you want with it, just give a shoutout if you can.