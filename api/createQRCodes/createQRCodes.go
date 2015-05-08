package createQRCodes

import (
	"bytes"
	"code.google.com/p/rsc/qr"
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
)

func randStr(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func GetQrCode() {
	//maximize CUP usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	//generate a random string - perferably 6 or 8 characters
	randomStr := randStr(6, "alphanum")
	fmt.Println(randomStr)
	stuff2buy := "https://www.baidu.com/"
	code, err := qr.Encode(stuff2buy, qr.H)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imgByte := code.PNG()
	img, _, _ := image.Decode(bytes.NewReader(imgByte))
	out, err := os.Create("./QRImg.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("QR code generated and saved to QRimg1.png")
}
