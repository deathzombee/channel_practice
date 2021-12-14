// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//start := time.Now()
	var a bytes.Buffer
	var b bytes.Buffer
	var c bytes.Buffer
	var d bytes.Buffer
	var e bytes.Buffer
	//urls := []string{
	//"https://roambarcelona.com/clock-pt1?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D",
	//"https://roambarcelona.com/clock-pt2?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D",
	//"https://roambarcelona.com/clock-pt3?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D",
	//"https://roambarcelona.com/clock-pt4?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D",
	//"https://roambarcelona.com/clock-pt5?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D",

	//}
	ch0 := make(chan string)
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go fetch("https://roambarcelona.com/clock-pt1?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D", ch0) // start a goroutine
	go fetch1("https://roambarcelona.com/clock-pt2?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D", ch1)
	go fetch2("https://roambarcelona.com/clock-pt3?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D", ch2)
	go fetch3("https://roambarcelona.com/clock-pt4?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D", ch3)
	go fetch4("https://roambarcelona.com/clock-pt5?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D", ch4)

	a.WriteString(<-ch0)
	b.WriteString(<-ch1)
	c.WriteString(<-ch2)
	d.WriteString(<-ch3)
	e.WriteString(<-ch4)
	//fmt.Print(<-ch,"\n") // receive from channel ch

	fmt.Print(a.String())
	fmt.Print(b.String())
	fmt.Print(c.String())
	fmt.Print(d.String())
	fmt.Print(e.String())

	doop, error := http.Get("https://roambarcelona.com/get-flag?verify=Na2Q%2BeqhSP5hTRLDwpTNoA%3D%3D&string=" + a.String() + b.String() + c.String() + d.String() + e.String())
	if error != nil {

	}
	defer doop.Body.Close()
	dd, error := io.ReadAll(doop.Body)
	fmt.Println(" ")
	fmt.Println(string(dd))
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch0 chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch0 <- fmt.Sprint(err) // send to channel ch
		return
	}

	teebody, err := io.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch0 <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	ch0 <- fmt.Sprintf("%s", teebody)

}
func fetch1(url string, ch1 chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch1 <- fmt.Sprint(err) // send to channel ch
		return
	}
	teebody1, err := io.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch1 <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch1 <- fmt.Sprintf("%s", teebody1)
}
func fetch2(url string, ch2 chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch2 <- fmt.Sprint(err) // send to channel ch
		return
	}
	teebody2, err := io.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch2 <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch2 <- fmt.Sprintf("%s", teebody2)
}
func fetch3(url string, ch3 chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch3 <- fmt.Sprint(err) // send to channel ch
		return
	}
	teebody3, err := io.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch3 <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch3 <- fmt.Sprintf("%s", teebody3)
}
func fetch4(url string, ch4 chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch4 <- fmt.Sprint(err) // send to channel ch
		return
	}
	teebody4, err := io.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch4 <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch4 <- fmt.Sprintf("%s", teebody4)
}
