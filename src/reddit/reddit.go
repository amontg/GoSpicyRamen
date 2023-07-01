package reddit

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"time"

	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/amontg/GoSpicyRamen/src/utils"
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
		log.Panic(err)
	}

	if len(page.Data.Children) <= 0 {
		utils.SimpleMessage(m.ChannelID, "I found no results.")
		return nil
	}

	return &paginator.Paginator{ // sm
		PageFunc: func(pageIndex int, embed *discordgo.MessageEmbed) {
			embed.Title = CheckSize(html.UnescapeString(page.Data.Children[pageIndex].Data.Title), 250)
			embed.Author = &discordgo.MessageEmbedAuthor{
				Name: page.Data.Children[pageIndex].Data.Author + " in " + page.Data.Children[pageIndex].Data.SubredditNamePrefixed,
			}
			embed.Color = 16777215
			embed.URL = redditLink + page.Data.Children[pageIndex].Data.Permalink
			embed.Description = CheckSize(html.UnescapeString(page.Data.Children[pageIndex].Data.Selftext), 4096)
			embed.Image = &discordgo.MessageEmbedImage{
				URL: CheckURL(page.Data.Children[pageIndex].Data.Thumbnail),
			}
		},
		MaxPages:        len(page.Data.Children),
		Expiry:          time.Now(),
		ExpiryLastUsage: true,
		ID:              m.ID + query,
	}
}

func CheckSize(s string, num int) string {
	if strings.Count(s, "")-1 > num {
		s = utils.Truncate(s, num-3)
	}

	return s
}

func CheckURL(s string) string {
	if strings.Contains(s, "https://") || strings.Contains(s, "http://") {
		if strings.Contains(s, ".gifv") {
			return ToGIF(s)
		} else if strings.Contains(s, ".jpg") || strings.Contains(s, ".jpeg") || strings.Contains(s, ".png") {
			fmt.Println(s)
			return s
		}
	}

	return ""
}

func ToGIF(s string) string {
	s = strings.Replace(s, ".gifv", ".gif", 1)
	fmt.Println(s)
	return s
}
