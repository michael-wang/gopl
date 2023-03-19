package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"time"
)

var save = flag.Bool("save", false, "Exercise 1.10: save response body to file ($hostname.html)")

func main() {
	flag.Parse()

	start := time.Now()
	ch := make(chan string)
	for _, arg := range flag.Args() {
		fmt.Println("fetching: ", arg)
		go fetch(arg, ch)
	}
	for range flag.Args() {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	fmt.Printf("fetching: %s, save: %t\n", url, save)
	start := time.Now()
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	fmt.Println("url: ", url)

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	out := ioutil.Discard
	if *save {
		u, err := neturl.Parse(url)
		if err != nil {
			ch <- fmt.Sprintf("Failed to parse url: %s\nerr: %v\n", url, err)
			return
		}

		filename := fmt.Sprintf("./out/%s.out", u.Host)
		file, err := os.Create(filename)
		if err != nil {
			ch <- fmt.Sprintf("Failed to create file: %s\nerr: %v\n", filename, err)
			return
		}

		out = file
		defer file.Close()
	}

	nbytes, err := io.Copy(out, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
