package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func templateFile(filename string) bool {
	// 读入源图像和模板图像
	filename = fmt.Sprintf("images/%v.png", filename)
	src := gocv.IMRead("source.png", gocv.IMReadGrayScale)
	template := gocv.IMRead(filename, gocv.IMReadGrayScale)

	// 计算模板在源图像中的匹配程度
	result := gocv.NewMat()
	defer result.Close()

	gocv.MatchTemplate(src, template, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	// 寻找最大匹配值和位置
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	if maxVal > 0.85 {
		clickArea := image.Rect(
			shotSquare.Min.X+maxLoc.X,
			shotSquare.Min.Y+maxLoc.Y,
			shotSquare.Min.X+maxLoc.X+template.Cols(),
			shotSquare.Min.Y+maxLoc.Y+template.Rows(),
		)
		randomClick(clickArea)
		return true
	}
	return false
}
