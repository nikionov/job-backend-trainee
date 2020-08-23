package awesomeProject1

import (
	"awesomeProject1/internal/app/apiserver"
	"flag"
	"github.com/burntsushi/toml"
	"log"
)

var (
	configPath string
)

func init(){
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main(){
	flag.Parse()
	config := apiserver.New()
	_, err := toml.DecodeFile(configPath, config)

	if err != nill{
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
}