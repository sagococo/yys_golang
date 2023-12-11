package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func capture(rec image.Rectangle) {
	img, err := screenshot.CaptureRect(rec)
	if err != nil {
		return
	}
	fileName := fmt.Sprintf("source.png")
	err = savePng(fileName, img)
	if err != nil {
		panic(err)
	}
}

func savePng(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("创建文件失败, filename:%s err:%s", filename, err))
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("关闭文件失败, filename:%s err:%s", filename, err))
		}
	}(file)
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(fmt.Sprintf("编码PNG失败, filename:%s err:%s", filename, err))
		return err
	}
	return nil
}
