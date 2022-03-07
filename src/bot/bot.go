package bot

import (
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"

	//"github.com/amontg/GoSpicyRamen/src/youtube"
	"github.com/amontg/GoSpicyRamen/src/handlers"
)

func Start() {
	context.Initialize(config.GetBotToken())
	handlers.AddHandlers()
	context.OpenConnection()

}

func Stop() {
	context.Dg.Close()
}
