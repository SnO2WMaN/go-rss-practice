package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	/*
		url := "https://jser.info/rss/"

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		byteArray, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		f := Feed{}
		xml.Unmarshal(byteArray, &f)
		fmt.Println(f)
	*/

	type Item struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
	}

	type Channel struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Items       []Item `xml:"items"`
	}

	type Feed struct {
		XMLName xml.Name `xml:"rss"`

		Version string `xml:"version,attr"`
		Channel Channel
	}

	feed := Feed{
		Version: "2.0",
		Channel: Channel{
			Title: "JSer.info",
			Items: []Item{
				{Title: "2022-05-17のJS: Solid v1.4.0、playwright v1.22.0、Safari 15.5"},
				{Title: "2022-05-11のJS: Node v18.1.0(node --test)、State of CSS 2022、Performance insights"},
			},
		},
	}
	buf, _ := xml.MarshalIndent(feed, "", " ")
	fmt.Println(string(buf))
}
