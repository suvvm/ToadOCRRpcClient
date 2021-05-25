package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/png"
	"os"
	"path"
	"strconv"
	"strings"
)

var supportImageExtNames = []string{".jpg", ".png"}

func GetBasicToken(appSecret, verifyStr string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(appSecret + verifyStr))
	md5Token := hex.EncodeToString(hasher.Sum(nil))
	return md5Token, nil
}

func PixelHashStr(pxs []byte) string {
	var resp string
	for i := 0; i < len(pxs); i++ {
		pixelVal := int(pxs[i])
		resp += strconv.Itoa(pixelVal)
	}
	return resp
}

func GetImageGrayBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	extname := strings.ToLower(path.Ext(filename))	// 判断扩展名是否合法
	if !isImage(extname) {
		return nil, fmt.Errorf("file not image")
	}
	var m image.Image
	if extname == ".png" {
		// decode图片
		m, err = png.Decode(file)
	} else {
		m, _, err = image.Decode(file)
	}
	if err != nil {
		return nil, err
	}
	bounds := m.Bounds()
	// mgrey := image.NewGray(bounds)
	imgBytes := make([]byte, 0)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// fmt.Printf("\n")
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			v := byte((float32(r)*299 + float32(g)*587 + float32(b)*114) / 1000)
			if v < 10 {
				v = 0
			}
			if v > 100 {
				v = 255
			}
			// fmt.Printf("%3v ", v)
			// mgrey.Set(x, y, color.Gray{v})
			imgBytes = append(imgBytes, v)
		}
	}
	return imgBytes, nil
}

func isImage(extName string) bool {
	for i := 0; i < len(supportImageExtNames); i++ {
		if supportImageExtNames[i] == extName {
			return true
		}
	}
	return false
}
