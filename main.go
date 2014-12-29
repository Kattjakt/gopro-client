package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const BASEPATH string = "http://10.5.5.9/videos/DCIM/100GOPRO/"
const DIR string = "GoPro"

type Entry struct {
	filename string
	size     string // to be used in the future
}

func getEntries() ([]Entry, error) {
	doc, err := goquery.NewDocument(BASEPATH)
	if err != nil {
		log.Fatal(err)
	}
	items := make([]Entry, doc.Find("tbody tr").Length())
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		items[i].filename = s.Find(".link").Text()
		items[i].size = s.Find("span").Text()
	})
	return items, err // do proper error handling maybe?
}

func worker(done chan bool, entry Entry) {
	out, _ := os.Create(DIR + "/" + entry.filename)
	resp, _ := http.Get(BASEPATH + entry.filename)
	defer out.Close()
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
	fmt.Println("Successfully downloaded", entry.filename)
	done <- true
}

func main() {
	fmt.Print("Checking for GoPro device... ")
	for {
		res, err := http.Get(BASEPATH)
		if err == nil && res.StatusCode == 200 {
			fmt.Println("OK")
			// we should probably do some error checking here
			break
		}
		time.Sleep(time.Second * 5)
	}
	items, err := getEntries()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found", len(items), "items, downloading... ")
	done := make([]chan bool, len(items))
	os.Mkdir(DIR, 'd')
	for i, filename := range items {
		go worker(done[i], filename)
	}

	for i, _ := range done {
		<-done[i]
	}
	fmt.Println("Finished!")
}
