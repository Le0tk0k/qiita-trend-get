package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// type Qiita struct {
// 	Trend `json:"trend"`
// 	Scope string `json:"scope"`
// }

// type Trend struct {
// 	Edges `json:"edges"`
// }

// type Edges []struct {
// 	FollowingLikers []string `json:"followingLikers"`
// 	IsLikeByViewer  bool     `json:"isLikeByViewer"`
// 	IsNewArrival    bool     `json:"isNewArrival"`
// 	HasCodeBlock    bool     `json:"hasCodeBlock"`
// 	Node            `json:"node"`
// }

// type Node struct {
// 	CreatedAt  time.Time `json:"createdAt"`
// 	LikesCount int       `json:"likesCount"`
// 	Title      string    `json:"title"`
// 	Uuid       string    `json:"uuid"`
// 	Author     `json:"author"`
// }

// type Author struct {
// 	ProfileImageURL string `json:"profileImageUrl"`
// 	UrlName         string `json:"urlName"`
// }

type Qiita struct {
	Trend `json:"trend"`
}

type Trend struct {
	Edges `json:"edges"`
}

type Edges []struct {
	Node `json:"node"`
}

type Node struct {
	Title  string `json:"title"`
	Uuid   string `json:"uuid"`
	Author `json:"author"`
}

type Author struct {
	UrlName string `json:"urlName"`
}

var trend Qiita

func main() {
	resp, err := http.Get("http://localhost:8080/trend")
	if err != nil {
		log.Fatal()
		fmt.Fprintln(os.Stderr, "http request err: ", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal()
		fmt.Fprintln(os.Stderr, "read err: ", err)
	}

	if err := json.Unmarshal(bytes, &trend); err != nil {
		fmt.Fprintln(os.Stderr, "unmarshal err:", err)
	}

	for i, v := range trend.Trend.Edges {
		fmt.Println("No.", i+1)
		fmt.Printf("\t%s\n\t%s\n", v.Node.Title, "https://qiita.com/"+v.Node.Author.UrlName+"/items/"+v.Node.Uuid)
	}
}
