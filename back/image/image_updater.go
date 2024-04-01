package image

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"os"
	"github.com/google/uuid"
)

func UpdateImage(imageName string) string {
	inputFile, err := os.Open(imageName)
    if err != nil {
        panic(err)
    }
    defer inputFile.Close()

    // 解码图片
    img, _, err := image.Decode(inputFile)
    if err != nil {
        panic(err)
    }

    // 获取原图片的边界
    bounds := img.Bounds()
    width := bounds.Dx()
    height := bounds.Dy()

    // 计算新的高度，裁剪掉下方10%
    newHeight := int(float64(height) * 0.9)
	newWidth := int(float64(width) * 1)

    // 创建裁剪区域，这里裁剪掉下方10%
    rect := image.Rect(0, 0, newWidth, newHeight)
    croppedImg := img.(interface {
        SubImage(r image.Rectangle) image.Image
    }).SubImage(rect)

    // 创建输出文件
    fname := uuid.New().String() + ".jpg"
    outputFile, err := os.Create(fname)
    if err != nil {
        panic(err)
    }
    defer outputFile.Close()

    // 将裁剪后的图片写入文件
    jpeg.Encode(outputFile, croppedImg, nil)
    return fname
}