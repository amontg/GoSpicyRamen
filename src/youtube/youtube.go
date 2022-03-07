package youtube

import (
	//"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	//"os/exec"

	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/bwmarrin/discordgo"

	"github.com/amontg/GoSpicyRamen/src/utils"
	//"github.com/bwmarrin/discordgo"
	paginator "github.com/TopiSenpai/dgo-paginator"
)

//	youtubeSearchEndpoint contains YouTube endpoint for searching after a video
const youtubeSearchEndpoint string = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&key="

//	youtubeFindEndpoint contains endpoint for finding more details about a video
const youtubeFindEndpoint string = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key="

//
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
	Thumbnail thumbnail `json:"thumbnails"`
	Desc      string    `json:"channelTitle"`
}

type size struct {
	Url string `json:"url"`
}

type thumbnail struct {
	Size size `json:"medium"`
}

type videoResponse struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
}

// these structs find a video on youtube
// page.Items[].Snippet.Title
//
type ytPageFind struct {
	Items []itemsFind `json:"items"`
}

type itemsFind struct {
	Snippet snippet `json:"snippet"`
}

// t for triggering message, s for sending message
func YtSearch(query string, m *discordgo.MessageCreate) *paginator.Paginator {

	// this line accesses youtube api using our youtube key and searches using our keyword. confirmed working
	res, err := http.Get(youtubeSearchEndpoint + config.GetYoutubeKey() + "&q=" + query)
	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		return utils.EmptyPaginator()
	}

	var page ytPageSearch

	// reads the response, decodes it onto a struct (ytPageSearch struct in this case)
	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		fmt.Println(err)
		return utils.EmptyPaginator()
	}

	//the following commented out code prints out the struct nicely
	s, _ := json.MarshalIndent(page, "", "\t")
	fmt.Print(string(s))

	res.Body.Close()

	if len(page.Items) < 1 {
		fmt.Println("No results!")
		err = errors.New("empty search result")
		//return ""
	}

	var embed *discordgo.MessageSend = new(discordgo.MessageSend)

	var msg *paginator.Paginator

	msg.PageFunc = func(num int, embed *discordgo.MessageEmbed) {
		// I have to create an embed for every page
		embed = ytEmbed(num, &page)
	}
	msg.MaxPages = len(page.Items)
	msg.Expiry = time.Now()
	msg.ExpiryLastUsage = true
	msg.ID = m.ID + query

	embed.Components = []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Left Arrow",
						ID:   "949946487335448636",
					},
					Style:    discordgo.PrimaryButton,
					CustomID: "PreviousResult",
					Disabled: true,
				},
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Right Arrow",
						ID:   "949946487616462848",
					},
					Style:    discordgo.PrimaryButton,
					CustomID: "NextResult",
					Disabled: false,
				},
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Shuffle Icon",
						ID:   "949946487410929714",
					},
					Style:    discordgo.PrimaryButton,
					CustomID: "Shuffle",
				},
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Trash",
						ID:   "949946686191587328",
					},
					Style:    discordgo.DangerButton,
					CustomID: "Dellete",
				},
			},
		},
	}

	// CreateMessage(*discordgo.Session, channelID string, paginator *Paginator) err

	return utils.EmptyPaginator()
}

func createYTEmbeds(page *ytPageSearch) []*discordgo.MessageEmbed {
	var embed *discordgo.MessageEmbed

	embed.Title = page.Items[num].Snippet.Title
	embed.Color = 16777215
	embed.URL = page.Items[num].Snippet.Title // page.Items[#].Snippet.Title
	embed.Description = page.Items[0].Snippet.Desc
	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      page.Items[0].Snippet.Thumbnail.Size.Url,
		ProxyURL: "",
		Width:    500,
		Height:   500,
	}

	// for-each loop to go over every page.Items[0-4], create []MessageEmbed

}
