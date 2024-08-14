package main //TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
import (
	"flag"
	"fmt"
	"server-api/cmd/app"
	"server-api/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	fmt.Println(*configFlag)
	c := config.NewConfig(*configFlag)
	app.NewApp(c)
}
