package urbandictionary

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

const urbandictQuery = "https://api.urbandictionary.com/v0/define?term="
const urbandictRandom = "https://api.urbandictionary.com/v0/random"

// search: .../v0/define?term=

type Data struct {
	List []struct {
		Definition  string    `json:"definition"`
		Permalink   string    `json:"permalink"`
		ThumbsUp    int       `json:"thumbs_up"`
		Author      string    `json:"author"`
		Word        string    `json:"word"`
		Defid       int       `json:"defid"`
		CurrentVote string    `json:"current_vote"`
		WrittenOn   time.Time `json:"written_on"`
		Example     string    `json:"example"`
		ThumbsDown  int       `json:"thumbs_down"`
	} `json:"list"`
}

func UrbanDictSearch(query string, m *discordgo.MessageCreate) *paginator.Paginator {
	var res *http.Response
	var err error
	if strings.Compare("random", query) == 0 {
		res, err = http.Get(urbandictRandom)
	} else {
		res, err = http.Get(urbandictQuery + query)
	}

	//fmt.Println(urbandictQuery + query)
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

	if len(page.List) <= 0 {
		utils.SimpleMessage(m.ChannelID, "I found no results.")
		return nil
	}

	return &paginator.Paginator{ // sm
		PageFunc: func(pageIndex int, embed *discordgo.MessageEmbed) {
			embed.Title = html.UnescapeString(page.List[pageIndex].Word)
			embed.Author = &discordgo.MessageEmbedAuthor{
				Name: page.List[pageIndex].Author,
			}
			embed.Color = 16777215
			embed.URL = page.List[pageIndex].Permalink
			embed.Description = utils.CheckSize(html.UnescapeString(page.List[pageIndex].Definition), 4096)
		},
		MaxPages:        len(page.List),
		Expiry:          time.Now(),
		ExpiryLastUsage: true,
		ID:              m.ID + query,
	}
}
