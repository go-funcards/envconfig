package envconfig

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

func ReadConfig(path string, cfg any) error {
	log.Printf("read application config from path %s\n", path)

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		help, _ := cleanenv.GetDescription(cfg, nil)
		log.Println(help)
		return err
	}
	return nil
}
