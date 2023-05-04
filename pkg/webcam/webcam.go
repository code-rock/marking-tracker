package webcam

import (
	"log"
	"fmt"
	"image/color"
	"gocv.io/x/gocv"
)

func Connect() {
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {    
		log.Fatalf("Error opening web cam: %v", err)
	}
	defer webcam.Close()
	
	window := gocv.NewWindow("Find out the composition")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	blue := color.RGBA{0, 255, 255, 0}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("data/haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Cannot read device %v\n", deviceID)
			return
		}

		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("Found %d faces\n", len(rects))
			
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}