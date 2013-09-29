package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	_ "math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringToInt(s string) int {
	r, err := strconv.Atoi(strings.Trim(s, " \n\r"))

	check(err)

	return r
}

func WriteToPPM(pixels []float64, width int, height int, depth int, filename string) {
	file, err := os.Create(filename)
	check(err)

	defer file.Close()

	w := bufio.NewWriter(file)

	header := fmt.Sprintf("P6\n%d %d\n%d\n", width, height, depth)

	_, err = w.Write([]byte(header))
	check(err)

	buf := make([]byte, width*height*3)

	for i := 0; i < width*height*3; i += 3 {
		buf[i] = byte(pixels[i] * 255.0)
		buf[i+1] = byte(pixels[i+1] * 255.0)
		buf[i+2] = byte(pixels[i+2] * 255.0)
	}

	_, err = w.Write(buf)
	check(err)

	// for i := 0; i < width; i++ {
	// 	for j := 0; j < height; j++ {
	// 		r := math.Max(0.0, math.Min(255.0, math.Pow(pixels[0], 1/2.2)*255+0.5))
	// 		g := math.Max(0.0, math.Min(255.0, math.Pow(pixels[1], 1/2.2)*255+0.5))
	// 		b := math.Max(0.0, math.Min(255.0, math.Pow(pixels[2], 1/2.2)*255+0.5))

	// 		w := []byte{byte(r), byte(g), byte(b)}
	// 		_, err = file.Write(w)

	// 		pixels = pixels[3:]

	// 		if i == 0 && j == 0 {
	// 			fmt.Println(w)
	// 		}

	// 		check(err)
	// 	}
	// }

	err = w.Flush()
	check(err)
}

func ReadPPM(filename string) (data []float64, w int, h int, c int, err error) {
	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	r := bufio.NewReader(file)

	//read first line -- PPM Type

	for {
		s, _, err := r.ReadLine()

		check(err)

		if string(s) == "P6" { //correct format
			break
		}

		if !strings.HasPrefix(string(s), "#") { //if not a comment
			return nil, 0, 0, 0, errors.New("Wrong format PPM or file is corrupt")
		}
	}

	width, err := r.ReadString(' ')
	check(err)

	height, err := r.ReadString('\n')
	check(err)

	maxval, err := r.ReadString('\n')
	check(err)

	cw := stringToInt(width)
	ch := stringToInt(height)
	cm := stringToInt(maxval)

	d := make([]byte, cw*ch*3)

	for {
		n, err := r.Read(d)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}
	}

	raw := make([]float64, cw*ch*3)

	for i, val := range d {
		raw[i] = ((float64)(val)) / 255.0
	}

	fmt.Println(raw[0], raw[1], raw[2])

	return raw, cw, ch, cm, nil
}
