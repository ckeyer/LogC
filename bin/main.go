package main

import (
	"fmt"

	// "github.com/ckeyer/LogC/routers"
	logpkg "github.com/ckeyer/go-log"
	"io"
	"net/http"
	"os"
)

type Logger interface {
	io.Writer
	io.Reader
}

var (
	logfile = "test.log"
	pool    Logger
	log     = logpkg.GetDefaultLogger()
)

type Content struct {
	w   http.ResponseWriter
	req *http.Request
}

func getFilePooler(name string) Logger {
	f, err := os.OpenFile(name, os.O_APPEND, 0600)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return f
}

func main() {
	// c, err := conf.GetConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// r := routers.Init()

	// _, _ = c, r
	f, err := os.OpenFile(logfile, os.O_CREATE, 0666)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer f.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	addr := fmt.Sprintf(":%d", 8087)
	log.Notice("Http is running at ", addr)
	err = http.ListenAndServe(addr, mux)
	panic(err)
}

// Index
func Index(w http.ResponseWriter, req *http.Request) {
	c := &Content{
		w:   w,
		req: req,
	}
	switch req.Method {
	case "POST":
		c.Post()
	case "GET":
		c.Get()
	}
}

// Get ...
func (c *Content) Get() {
	c.w.Header().Add("Content-type", "text/plain")
	f, err := os.OpenFile(logfile, os.O_RDONLY, 0400)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(c.w, f)
	if err != nil {
		log.Error(err.Error())
		return
	}
	c.w.WriteHeader(http.StatusOK)
}

// Post ...
func (c *Content) Post() {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(f, c.req.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	f.Write([]byte("\n"))
	c.w.WriteHeader(http.StatusCreated)
}
