package ch1

import (
	"fmt"
	"image"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func Serve1() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echoes the Path component of the request URL
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url.Path = %q\n", r.URL.Path)
}

var mu sync.Mutex
var count int

func Serve2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// handler2 echoes the Path componeent of the requested URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func Serve3() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q", k, v)
	}
}

func LissajousServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
// For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
// number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the
// string parameter into an integer. You can see its documentation with go doc strconv.Atoi.

func lissajous(out io.Writer, c float64) {
	var (
		cycles  float64 = 5     // number of complete x oscillator revolutions
		res             = 0.001 // angular resolution
		size            = 100   // image canvas cover [-size..+size]
		nframes         = 64    // number of animation frames
		delay           = 8     // delay between frames in 10ms units
	)

	if c != 0 {
		cycles = c
	}

	freq := rand.Float64() * 3.0 // realtive frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	cycles := *new(float64)

	q := r.URL.Query()

	if q.Has("cycles") {
		tmp, err := strconv.ParseFloat(q.Get("cycles"), 64)
		cycles = tmp // tmp avoids shadowing cycles
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't convert value %v int\n", q.Get("cycles"))
		}
	}

	lissajous(w, cycles)
}

func Server() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
