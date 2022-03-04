package bot

import (
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"

	//"github.com/amontg/GoSpicyRamen/src/youtube"
	"github.com/amontg/GoSpicyRamen/src/handlers"
)

func Start() {
	context.Initialize(config.GetBotToken())
	//fmt.Println(config.GetBotToken())
	handlers.AddHandlers()
	context.OpenConnection()
	//youtube.InitializeRoutine()

}

func Stop() {
	context.Dg.Close()
}
