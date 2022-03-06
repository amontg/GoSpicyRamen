package youtube

import (
	//"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	//"os/exec"

	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"

	//"github.com/amontg/GoSpicyRamen/src/utils"
	//embed "github.com/clinet/discordgo-embed"
	//"github.com/amontg/GoSpicyRamen/src/utils"
	"github.com/bwmarrin/discordgo"
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
	Title string `json:"title"`
}

type videoResponse struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
}

// these structs find a video on youtube
type ytPageFind struct {
	Items []itemsFind `json:"items"`
}

type itemsFind struct {
	Snippet snippet `json:"snippet"`
}

func YtSearch(query string, m *discordgo.MessageCreate) {

	// this line accesses youtube api using our youtube key and searches using our keyword. confirmed working
	res, err := http.Get(youtubeSearchEndpoint + config.GetYoutubeKey() + "&q=" + query)
	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		//return ""
	}

	var page ytPageSearch

	// reads the response from line #63 and decodes it onto a struct (ytPageSearch struct in this case)
	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		fmt.Println(err)
		//return ""
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
	//

	//videoId := page.Items[0].Id.VideoId
	//videoUrl := ytVideoUrl + videoId

	//utils.SendChannelMessage(channelID, videoUrl)
	//return videoUrl

	msg := discordgo.MessageSend{
		Content: "I'm a filthy test.",
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "1",
				Style:    discordgo.PrimaryButton,
				Disabled: false,
			},
		},
	}

	test, err := context.Dg.ChannelMessageSendComplex(m.ChannelID, &msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(test)
	fmt.Println("Attempted to search YouTube.")
}

// func name(para type) (return type, return type)
// this func will pull up the specific video's page, call using videoId from YtSearch(). not used for now
func ytFind(videoId string) {
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

	//return videoTitle, nil
}
