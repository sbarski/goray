package main

import (
	"flag"
	"fmt"
	_ "goray/geometry"
	"log"
	"net/http"
	"strconv"
)

func runRayTracer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

/*
 * Command line arguments allowed are: -port=xxxx
 * Example: go run goray.go -port=9234
 */
func main() {
	port := flag.Int("port", 9223, "Required port number")
	out := flag.String("output", "ppm", "Required outpyut must be defined")
	flag.Parse()

	switch *out {
	case "ppm":
		break
	case "web":
		http.HandleFunc("/", runRayTracer)
		err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)

		if err != nil {
			log.Fatal("Could not run web server: ", err)
		}
	}

}
