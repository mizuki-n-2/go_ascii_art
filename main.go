package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 読み込みファイル名
	readFileName := "snoopy.JPG"
	reader, err := os.Open("img/" + readFileName)
	if err != nil {
		fmt.Println("ファイルの読み込みエラー", err)
		os.Exit(1)
	}

	img, fileFormat, err := image.Decode(reader)
	if err != nil {
		fmt.Println("画像の変換エラー", err)
		os.Exit(1)
	}

	fmt.Printf("%s形式データを得ました\n", fileFormat)

	defer reader.Close()

	marks := []string{"*", "+", "-"}
	var marksStr string

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayness := c.Y / (255 / 3)
			if c.Y == 255 {
				grayness = 2
			}
			marksStr += marks[grayness]
		}
		marksStr += "\n"
	}

	data := []byte(marksStr)
	// 書き込みファイル名
	writeFileName := strings.Split(readFileName, ".")[0] + ".txt"
	err = ioutil.WriteFile("ascii_art/"+writeFileName, data, 0777)
	if err != nil {
		fmt.Println("ファイルの書き込みエラー", err)
		os.Exit(1)
	}

	fmt.Println("txtファイルを保存しました")
}
