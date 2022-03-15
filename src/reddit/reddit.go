package reddit

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

const redditSub = "https://www.reddit.com/r/"

type subredditPage struct {
	Thing []string `json:"page"`
}

func RedditSearch(subreddit string, m *discordgo.MessageCreate) {

	res, err := http.Get(redditSub + subreddit)
	if err != nil {
		fmt.Println("Error 1:", err)
	}

	//res.Header.Set("User-Agent", "I am a project bot: https://github.com/amontg/GoSpicyRamen")

	body, readErr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if readErr != nil {
		fmt.Println("ReadErr: ", readErr)
	}

	fmt.Printf("%s", body)

}
