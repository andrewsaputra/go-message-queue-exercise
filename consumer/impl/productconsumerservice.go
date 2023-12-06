package impl

import (
	"andrewsaputra/go-message-queue-exercise/consumer/api"
	"errors"
	"fmt"
	"strconv"

	_ "image/jpeg"

	"github.com/disintegration/imaging"
)

func ConstructProductConsumerService(imageDir string, compressedImageDir string, targetHeight int, productAccessor api.ProductDataAccessor) api.ConsumerService {
	return &ProductConsumerService{
		ImageDir:           imageDir,
		CompressedImageDir: compressedImageDir,
		TargetHeight:       targetHeight,
		ProductAccessor:    productAccessor,
	}
}

type ProductConsumerService struct {
	ImageDir           string
	CompressedImageDir string
	TargetHeight       int
	ProductAccessor    api.ProductDataAccessor
}

func (this *ProductConsumerService) OnConsumed(body []byte) error {
	id, err := strconv.Atoi(string(body))
	if err != nil {
		return err
	}

	product, err := this.ProductAccessor.Get(id)
	if err != nil || product == nil {
		if err == nil {
			err = errors.New(fmt.Sprintf("Product not found, id: %d", id))
		}
		return err
	}

	compressedImages := []string{}
	for _, url := range product.Images {
		compressResult, err := this.compressAndSave(url)
		if err != nil {
			return err
		}

		compressedImages = append(compressedImages, compressResult)
	}

	_, err = this.ProductAccessor.SetCompressedImages(id, compressedImages)
	if err != nil {
		return err
	}

	fmt.Println("ProductConsumerService.OnConsumed :", product.Id)
	return nil
}

func (this *ProductConsumerService) compressAndSave(imageName string) (string, error) {
	inputPath := fmt.Sprintf("%s/%s", this.ImageDir, imageName)
	img, err := imaging.Open(inputPath)
	if err != nil {
		return "", err
	}

	scaleFactor := float64(this.TargetHeight) / float64(img.Bounds().Dy())
	targetWidth := int(scaleFactor * float64(img.Bounds().Dx()))
	imgResized := imaging.Resize(img, targetWidth, this.TargetHeight, imaging.Box)
	compressedImageName := fmt.Sprintf("compressed_%s", imageName)
	outputPath := fmt.Sprintf("%s/%s", this.CompressedImageDir, compressedImageName)
	err = imaging.Save(imgResized, outputPath)
	if err != nil {
		return "", err
	}

	return compressedImageName, nil
}
