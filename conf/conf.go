package conf

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var gConfig = &Config{}

type Config struct {
	ApiPrefix []string
	Articles  map[int32]*Article
}

type Article struct {
	DlLinks []string
}

func init() {
	var configFile string
	flag.StringVar(&configFile, "c", "conf/config.yaml", "conf file path")

	reloadConf(configFile)
	rand.Seed(time.Now().Unix())
}

func GetArticleAdLinks(articleId int32) []string {
	if len(gConfig.Articles) == 0 {
		return nil
	}

	article, ok := gConfig.Articles[articleId]
	if !ok || article == nil || len(article.DlLinks) == 0 {
		return nil
	}

	adLinks := make([]string, 0, len(article.DlLinks)*len(gConfig.ApiPrefix))
	for _, apiPrefix := range gConfig.ApiPrefix {
		for _, dlLink := range article.DlLinks {
			adLinks = append(adLinks, apiPrefix+dlLink)
		}
	}

	return adLinks
}

func reloadConf(configFile string) {
	err := ParseYaml(configFile, gConfig)
	if err != nil {
		log.Panicf("ParseYaml err:%v\n", err)
	}

	log.Printf("config:%+v", gConfig)
	go reloadYamlFile(configFile, time.Minute, gConfig)
}

func reloadYamlFile(configFile string, duration time.Duration, serverConf *Config) {
	var lastMtime = getFileMtime(configFile)
	for {
		time.Sleep(duration)
		if curMtime := getFileMtime(configFile); curMtime > lastMtime {
			lastMtime = curMtime
			err := ParseYaml(configFile, serverConf)
			if err != nil {
				log.Panicf("ParseYaml err:%v\n", err)
			}
			log.Printf("config:%+v", serverConf)
		}
	}
}

func getFileMtime(file string) int64 {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatalf("file stat err:%v\n", err)
		return 0
	}
	return fileInfo.ModTime().Unix()
}
