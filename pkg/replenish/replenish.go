package replenish

import (
	"image/color"
	"math/rand"
	"log"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"time"
	"strings"
	"os"
)

func changeImg(file string, count int, folder string) {
	src, err := imaging.Open(fmt.Sprintf("images/%s/%s", folder, file))
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	fmt.Println(rand.Float64() * 360)
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)
	img1 := imaging.Blur(src, rand.Float64() * 5)
	img1 = imaging.Rotate(src, rand.Float64() * 360,  color.White)
	img1 = imaging.Resize(img1, 360, 360, imaging.Lanczos)
	img1 = imaging.PasteCenter(src, img1) 
	img1 = imaging.AdjustSigmoid(img1, rand.Float64(), rand.Float64() * 10)
	img1 = imaging.Grayscale(img1)

	err = imaging.Save(img1, fmt.Sprintf("images/%s/out_%d%s", folder,  count, file))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func ReplenishImageSet(count int) {
	folders, err := ioutil.ReadDir("./images")
    if err != nil {
        log.Fatal(err)
    }
	for _, folder := range folders {
		files, err := ioutil.ReadDir(fmt.Sprintf("./images/%s", folder.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if !strings.HasPrefix(file.Name(), "out_") {
				for j:= 0; j < count; j++ {
					go func (j int, file, folder string) {
						changeImg(file, j, folder)
					}(j, file.Name(), folder.Name())
				}
			}
		}
	}

	time.Sleep(5000* time.Millisecond)
}

func RemoveAdditional() {
	folders, err := ioutil.ReadDir("./images")
    if err != nil {
        log.Fatal(err)
    }
	for _, folder := range folders {
		files, err := ioutil.ReadDir(fmt.Sprintf("./images/%s", folder.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
			fmt.Println(strings.HasPrefix(file.Name(), "out") )
			if strings.HasPrefix(file.Name(), "out") {
				go func (folder, file string) {
					os.Remove(fmt.Sprintf("./images/%s/%s", folder, file))
				} (folder.Name(), file.Name())
			}
		}
	}

	time.Sleep(10000* time.Millisecond)
}
