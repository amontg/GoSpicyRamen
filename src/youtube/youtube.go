package youtube

import (
	//"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	//"os/exec"

	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/bwmarrin/discordgo"

	//"github.com/amontg/GoSpicyRamen/src/context"
	"github.com/amontg/GoSpicyRamen/src/utils"
	//embed "github.com/clinet/discordgo-embed"
	//"github.com/amontg/GoSpicyRamen/src/utils"
	//"github.com/bwmarrin/discordgo"
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
func YtSearch(query string) *discordgo.MessageSend {

	// this line accesses youtube api using our youtube key and searches using our keyword. confirmed working
	res, err := http.Get(youtubeSearchEndpoint + config.GetYoutubeKey() + "&q=" + query)
	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		return utils.EmptyComplex()
	}

	var page ytPageSearch

	// reads the response from line #63 and decodes it onto a struct (ytPageSearch struct in this case)
	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		fmt.Println(err)
		return utils.EmptyComplex()
	}

	// fmt.Print(page) -- gonna change to print in readable format
	/*
			the following commented out code prints out the struct nicely
		s, _ := json.MarshalIndent(page, "", "\t")
		fmt.Print(string(s))
	*/

	//
	res.Body.Close()

	if len(page.Items) < 1 {
		fmt.Println("No results!")
		err = errors.New("empty search result")
		//return ""
	}

	displayedResult := 0

	videoId := page.Items[displayedResult].Id.VideoId
	videoUrl := ytVideoUrl + videoId

	var msg *discordgo.MessageSend = new(discordgo.MessageSend)
	var video = ytFind(videoId)

	//msg.Content = videoUrl
	//fmt.Println(video.Items[0])

	//fmt.Println(video)
	msg.Embeds = []*discordgo.MessageEmbed{
		{
			Title:       video.Items[displayedResult].Snippet.Title,
			Color:       16777215,
			URL:         videoUrl,
			Description: video.Items[displayedResult].Snippet.Desc,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    video.Items[displayedResult].Snippet.Thumbnail.Size.Url,
				Width:  500,
				Height: 500,
			},
		},
	}

	msg.Components = []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Left Arrow",
						ID:   "949946487335448636",
					},
					Style:    discordgo.PrimaryButton,
					CustomID: "ytPreviousResult",
					Disabled: true,
				},
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Right Arrow",
						ID:   "949946487616462848",
					},
					Style:    discordgo.PrimaryButton,
					CustomID: "ytNextResult",
					Disabled: false,
				},
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Shuffle Icon",
						ID:   "949946487410929714",
					},
					Style:    discordgo.PrimaryButton,
					CustomID: "ytShuffle",
				},
				discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "Trash",
						ID:   "949946686191587328",
					},
					Style:    discordgo.DangerButton,
					CustomID: "ytDellete",
				},
			},
		},
	}

	return msg
}

// func name(para type) (return type, return type)
// this func will pull up the specific video's page, call using videoId from YtSearch(). not used for now
func ytFind(videoId string) ytPageFind {
	res, err := http.Get(youtubeFindEndpoint + config.GetYoutubeKey() + "&id=" + videoId)
	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		//return "", err
	}

	var page ytPageFind

	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		fmt.Println(err)
		//return "", err
	}

	res.Body.Close()

	if len(page.Items) < 1 {
		fmt.Println("Empty results!")
		err = errors.New("empty search result")
		//return "", err
	}

	//videoTitle := page.Items[0].Snippet.Title
	//fmt.Println(page)

	return page
}
