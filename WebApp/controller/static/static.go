// Package static serves static files like CSS, JavaScript, and images.
package static

import (
	"net/http"

	"github.com/huzhaer/qianxun/lib/flight"
	"github.com/huzhaer/teamlite_core/router"

	"strings"
	"os"
	"io"
	"path"
)

// Load the routes.
func Load() {
	// Serve static files
	router.Get("/public/*filepath", GetPublic)
}

func GetPublic(w http.ResponseWriter, r *http.Request)  {
	 c := flight.Context(w, r)
	 _ = c

	fullUri := r.URL.Path[1:]
	uriSlice := strings.Split(fullUri,"/")
	thepath := path.Join("./public",uriSlice[len(uriSlice) - 2 ], uriSlice[len(uriSlice) - 1 ])
	http.ServeFile(w , r, thepath)
}

// Index maps static files.
func Index(w http.ResponseWriter, r *http.Request) {
	// c := flight.Context(w, r)

	fullUri := r.URL.Path[1:]
	// File path
	// path := path.Join("./assets" ,r.URL.Path[1:])
	// res, getErr  := http.Get("http://localhost:3000/" + fullUri )
	// defer res.Body.Close()

	uriSlice := strings.Split(fullUri,"/")
	filename := uriSlice[len(uriSlice)-1]
	DownloadFile(filename,"http://localhost:3000/"+fullUri,w)
}

func DownloadFile(filepath string, url string,w http.ResponseWriter) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ctp := resp.Header.Get("Content-Type")
	cset := resp.Header.Get("charset")
	w.Header().Set("Content-Type",ctp)
	w.Header().Set("charset",cset)
	// Write the body to file
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
