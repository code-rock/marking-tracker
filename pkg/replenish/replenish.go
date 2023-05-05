package replenish

import (
	"image/color"
	"math/rand"
	"log"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"time"
)

func changeImg(id int, count int, folder string) {
	src, err := imaging.Open(fmt.Sprintf("images/%s/%d.png", folder, id ))
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

	err = imaging.Save(img1, fmt.Sprintf("images/%s/out_%d%d.png", folder, id, count))
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
		for id, _ := range files {
			for j:= 1; j <= count; j++ {
				go func (i, j int, folder string) {
					changeImg(i + 1, j, folder)
				}(id, j, folder.Name())
			}
		}
	}

	time.Sleep(1000* time.Millisecond)
}

// folders, err := ioutil.ReadDir("./images")
    // if err != nil {
    //     log.Fatal(err)
    // }

    // for _, folder := range folders {
	// 	files, err_files := ioutil.ReadDir(fmt.Sprintf("./images/%s", folder.Name()))
	// 	if err_files != nil {
	// 		log.Fatal(err_files)
	// 	}

	// 	for _, file := range files {
	// 		fmt.Println(file.Name(), file.IsDir())
	// 	}
    // }