package reddit

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"

	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/bwmarrin/discordgo"
)

const redditLink = "https://www.reddit.com"
const redditQuery = "https://www.reddit.com/search.json?q="

type Data struct {
	Data struct {
		Amount   int `json:"dist"`
		Children []struct {
			Data struct {
				Selftext              string  `json:"selftext"`
				Title                 string  `json:"title"`
				SubredditNamePrefixed string  `json:"subreddit_name_prefixed"`
				Author                string  `json:"author"`
				URL                   string  `json:"url_overridden_by_dest"`
				Thumbnail             string  `json:"thumbnail"`
				CreatedUtc            float64 `json:"created_utc"`
				Permalink             string  `json:"permalink"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func RedditSearch(query string, m *discordgo.MessageCreate) *paginator.Paginator {

	res, err := http.Get(redditQuery + query)
	fmt.Println(redditQuery + query)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//res.Header.Set("User-Agent", "I am a project bot: https://github.com/amontg/GoSpicyRamen")
	var page Data

	err = json.NewDecoder(res.Body).Decode(&page)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	result := &paginator.Paginator{ // sm
		PageFunc: func(pageIndex int, embed *discordgo.MessageEmbed) {
			embed.Title = html.UnescapeString(page.Data.Children[pageIndex].Data.Title)
			embed.Author = &discordgo.MessageEmbedAuthor{
				Name: page.Data.Children[pageIndex].Data.Author,
			}
			embed.Color = 16777215
			embed.URL = redditLink + page.Data.Children[pageIndex].Data.Permalink
			embed.Description = page.Data.Children[pageIndex].Data.Selftext
			embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
				URL: CheckURL(page.Data.Children[pageIndex].Data.Thumbnail),
			}
		},
		MaxPages:        page.Data.Amount,
		Expiry:          time.Now(),
		ExpiryLastUsage: true,
		ID:              m.ID + query,
	}

	return result
}

func CheckSize(s string, num int) string {
	if strings.Count(s, "")-1 > num {
		s = utils.truncate(s, num-3)
	}

	return s
}

func CheckURL(s string) string {
	if strings.Contains(s, "https://") {
		return ToGIF(s)
	} else if strings.Contains(s, "http://") {
		return ToGIF(s)
	}

	return ""
}

func ToGIF(s string) string {
	if strings.Contains(s, ".gifv") {
		strings.Replace(s, ".gifv", ".gif", 1)
	}

	return s
}
