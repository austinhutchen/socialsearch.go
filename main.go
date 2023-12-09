package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// RedditData represents the structure of Reddit API response data
type RedditData struct {
	Data struct {
		Children []struct {
			Data struct {
				Title       string `json:"title"`
				Body        string `json:"body,omitempty"`
				Permalink   string `json:"permalink,omitempty"`
				Score       int    `json:"score,omitempty"`
				NumComments int    `json:"num_comments,omitempty"`
				Subreddit   string `json:"subreddit,omitempty"`
				Author      string `json:"author,omitempty"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func getPushshiftData(dataType string, params map[string]string) (*RedditData, error) {
	baseURL := fmt.Sprintf("https://api.pushshift.io/reddit/search/%s/", dataType)

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var redditData RedditData
	err = json.Unmarshal(body, &redditData)
	if err != nil {
		return nil, err
	}

	return &redditData, nil
}

func getReddit(ans, listing, limit, timeframe string) (*RedditData, error) {
	client := &http.Client{}
	baseURL := fmt.Sprintf("https://www.reddit.com/r/%s/%s.json?limit=%s&t=%s", ans, listing, limit, timeframe)
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "yourbot")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var redditData RedditData
	err = json.Unmarshal(body, &redditData)
	if err != nil {
		return nil, err
	}

	return &redditData, nil
}

func main() {
	limit := "50"
	timeframe := "week"
	listing := "hot"

	fmt.Println("---REDDIT SEARCH V.1---")
	fmt.Println("1 -> Search keyword")
	fmt.Println("2 -> View subreddit")
	fmt.Println("3 -> Show random post")

	var mainChoice int
	fmt.Print("Choose an option: ")
	fmt.Scanln(&mainChoice)

	if mainChoice == 1 {
		var exit int
		for exit != 1 {
			var ans string
			fmt.Print("Enter the keyword you want more information on: ")
			fmt.Scanln(&ans)
			if ans == "end" {
				return
			} else {
				fmt.Println("ENTER 1 FOR MOST COMMON COMMENTS")
				fmt.Println("ENTER 2 FOR MOST COMMON SUBREDDITS")
				fmt.Println("ENTER 3 FOR MOST COMMON USERS")

				var choice int
				fmt.Print("Would you prefer comments (1), users (2), or subreddits? (3): ")
				fmt.Scanln(&choice)

				if choice == 1 {
					// Implement logic for comments search
				} else if choice == 2 {
					// Implement logic for subreddits search
				} else if choice == 3 {
					// Implement logic for users search
				} else {
					fmt.Println("Fix yo input")
					continue
				}
			}
		}
	} else if mainChoice == 2 {
		var choice string
		fmt.Print("Enter the subreddit you would like to search for: ")
		fmt.Scanln(&choice)

		result, err := getReddit(choice, listing, limit, timeframe)
		if err != nil {
			fmt.Println("An error occurred:", err)
			return
		}

		// Display results
		for i := 0; i < 11; i++ {
			fmt.Println(result.Data.Children[i].Data.Author)
			fmt.Println("----end----")
		}
	} else if mainChoice == 3 {
		post := 0
		fmt.Printf("Your random post of the day is: %v\n", post)
	} else {
		return
	}
}

