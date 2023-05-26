package youtube

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

	//"os/exec"

	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/bwmarrin/discordgo"

	//"github.com/bwmarrin/discordgo"
	paginator "github.com/TopiSenpai/dgo-paginator"
)

// youtubeSearchEndpoint contains YouTube endpoint for searching after a video
const youtubeSearchEndpoint string = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&key="

const ytVideoUrl string = "https://www.youtube.com/watch?v="

// these structs are for doing a Youtube search
// type structs create new type
type ytPageSearch struct {
	Items []itemsSearch `json:"items"`
}

type itemsSearch struct {
	Id      id      `json:"id"`
	Snippet snippet `json:"snippet"`
}

type id struct {
	VideoId string `json:"videoId"`
}

type snippet struct {
	Title     string    `json:"title"`
	Author    string    `json:"channelTitle"`
	Thumbnail thumbnail `json:"thumbnails"`
	Desc      string    `json:"description"`
}

type size struct {
	Url string `json:"url"`
}

type thumbnail struct {
	Size size `json:"medium"`
}

// t for triggering message, s for sending message
func YtSearch(query string, m *discordgo.MessageCreate) *paginator.Paginator { // *paginator.Paginator

	// this line accesses youtube api using our youtube key and searches using our keyword. confirmed working
	res, err := http.Get(youtubeSearchEndpoint + config.GetYoutubeKey() + "&q=" + query)
	//fmt.Println(youtubeSearchEndpoint + config.GetYoutubeKey() + "&q=" + query)
	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		//return utils.EmptyPaginator()
	}

	res.Header.Set("User-Agent", "I am a project bot: https://github.com/amontg/GoSpicyRamen")

	var page ytPageSearch

	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		fmt.Println(err)
		//return utils.EmptyPaginator()
	}

	// reads the response, decodes it onto a struct (ytPageSearch struct in this case)

	res.Body.Close()

	if len(page.Items) < 1 {
		fmt.Println("No results!")
		//return ""
	}

	return &paginator.Paginator{ // sm
		PageFunc: func(pageIndex int, embed *discordgo.MessageEmbed) {
			embed.Title = html.UnescapeString(page.Items[pageIndex].Snippet.Title)
			embed.Author = &discordgo.MessageEmbedAuthor{
				Name: page.Items[pageIndex].Snippet.Author,
			}
			embed.Color = 16777215
			embed.URL = ytVideoUrl + page.Items[pageIndex].Id.VideoId
			embed.Description = page.Items[pageIndex].Snippet.Desc
			embed.Image = &discordgo.MessageEmbedImage{ // thumbnail
				URL: page.Items[pageIndex].Snippet.Thumbnail.Size.Url,
			} // thumbnail
			/* For when I can figure out how videos work...
			embed.Video = &discordgo.MessageEmbedVideo{ // thumbnail
				URL: embed.URL + ".mp4",
			}
			*/
		},
		MaxPages:        len(page.Items),
		Expiry:          time.Now(),
		ExpiryLastUsage: true,
		ID:              m.ID + query,
	}
}
