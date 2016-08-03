package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	// url to POST request to
	u = "http://m.investing.com/economic-calendar/services/filter/"

	// country codes to query. Not sure what these actually are.
	codes = []int{
		25, 32, 6, 37, 72, 22, 17, 39, 14, 10, 35, 43, 36, 110, 11, 26, 12,
		4, 5, 56,
	}
)

func getForm() url.Values {
	form := url.Values{
		"timeZone":   []string{"8"},
		"currentTab": []string{"today"},
	}

	for _, code := range codes {
		form.Add("country[]", strconv.Itoa(code))
	}

	return form
}

func main() {
	// build the request data
	form := getForm()
	buf := bytes.NewBuffer([]byte(form.Encode()))

	// build the request
	req, err := http.NewRequest("POST", u, buf)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// build the client and send the request
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Received HTTP status code %v\n", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		// TODO get date and combine date/time for each item
		time := strings.TrimSpace(s.Find(".time p").Text())
		title := strings.TrimSpace(s.Find(".rightSide p").Text())
		currency := s.Find("div.curr").Text()
		previous := s.Find("i.prev").Text()
		actual := s.Find("i.act").Text()
		forecast := s.Find("i.fore").Text()
		// fmt.Printf("time='%v'\n", time)
		fmt.Printf("%s (%s) : %v : %v vs %v (previous %v)\n", time, currency, title, forecast, actual, previous)
	})
}
