
package webcam

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
    "gocv.io/x/gocv"
	"github.com/hybridgroup/mjpeg"
)

var (
    deviceID int
    err      error
    cam   *gocv.VideoCapture
    stream   *mjpeg.Stream
)

func Connect() {
    /*  if len(os.Args) < 3 {
            fmt.Println("How to run:\n\tmjpeg-streamer [camera ID] [host:port]")
            return
        }
    */
    // parse args
    deviceID := 0   // os.Args[1]
    host := ":3005" //os.Args[2]

    // open webcam
    cam, err = gocv.OpenVideoCapture(deviceID)
    if err != nil {
        fmt.Printf("Error opening capture device: %v\n", deviceID)
        return
    }
    defer cam.Close()

    // create the mjpeg stream
    stream = mjpeg.NewStream()

    // start capturing
    go mjpegCapture()

    fmt.Println("Capturing. Point your browser to " + host)

    //start http server
	//return stream
    http.Handle("/", stream)
    log.Fatal(http.ListenAndServe(host, nil))
}

func mjpegCapture() {
    img := gocv.NewMat()
    defer img.Close()

    for {
        if ok := cam.Read(&img); !ok {
            fmt.Printf("Device closed: %v\n", deviceID)
            return
        }
        if img.Empty() {
            continue
        }

        buf, _ := gocv.IMEncode(".jpg", img)
        stream.UpdateJPEG(buf.GetBytes())
        buf.Close()
    }
}

// package webcam

// import (
// 	// "os"
// 	"log"
// 	"fmt"
// 	//"net/http"
// 	// "image/color"
// 	"gocv.io/x/gocv"
// 	"github.com/hybridgroup/mjpeg"
// )

// var (
// 	deviceID int
// 	err      error
// 	webcam   *gocv.VideoCapture
// 	stream   *mjpeg.Stream
// )


// func mjpegCapture() {
// 	img := gocv.NewMat()
// 	defer img.Close()

// 	for {
// 		if ok := webcam.Read(&img); !ok {
// 			fmt.Printf("Device closed: %v\n", deviceID)
// 			return
// 		}
// 		// if img.Empty() {
// 		// 	continue
// 		// }

// 		// buf, _ := gocv.IMEncode(".jpg", img)
// 		// stream.UpdateJPEG(buf.GetBytes())
// 		// buf.Close()
// 	}
// }

// func Connect() {
// 	// if len(os.Args) < 3 {
// 	// 	fmt.Println("How to run:\n\tmjpeg-streamer [camera ID] [host:port]")
// 	// 	return
// 	// }

// 	deviceID := 0 // os.Args[1]
// 	host := ":3001" //os.Args[2]
// 	webcam, err := gocv.OpenVideoCapture(deviceID)
// 	if err != nil {    
// 		log.Fatalf("Error opening web cam: %v", err)
// 	}
// 	defer webcam.Close()
	
// 	stream = mjpeg.NewStream()
// 	go mjpegCapture()

// 	fmt.Println("Capturing. Point your browser to " + host)
// 	fmt.Println(stream)
// 	// start http server
// 	// http.Handle("/", stream)
// 	// log.Fatal(http.ListenAndServe(host, nil))
// 	// window := gocv.NewWindow("Find out the composition")
// 	// defer window.Close()

// 	// img := gocv.NewMat()
// 	// defer img.Close()

// 	// blue := color.RGBA{0, 255, 255, 0}

// 	// classifier := gocv.NewCascadeClassifier()
// 	// defer classifier.Close()

// 	// if !classifier.Load("data/haarcascade_frontalface_default.xml") {
// 	// 	fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
// 	// 	return
// 	// }

// 	// for {
// 	// 	if ok := webcam.Read(&img); !ok {
// 	// 		fmt.Printf("Cannot read device %v\n", deviceID)
// 	// 		return
// 	// 	}

// 	// 	if img.Empty() {
// 	// 		continue
// 	// 	}

// 	// 	rects := classifier.DetectMultiScale(img)
// 	// 	fmt.Printf("Found %d faces\n", len(rects))
			
// 	// 	for _, r := range rects {
// 	// 		gocv.Rectangle(&img, r, blue, 3)
// 	// 	}

// 	// 	window.IMShow(img)
// 	// 	window.WaitKey(1)
// 	// }
// }
