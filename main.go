package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/eleven26/rds_exporter/exporter"
)

func parseConfig(path *string) exporter.Config {
	if _, err := os.Stat(*path); os.IsNotExist(err) {
		panic(err)
	}

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	var conf exporter.Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func main() {
	address := flag.String("web.listen-address", "127.0.0.1:9094", "监听地址，格式如 \"0.0.0.0:9094\"")
	configFile := flag.String("config.file", "config.yml", "配置文件路径")
	flag.Parse()

	conf := parseConfig(configFile)

	log.Printf("监听地址：%s", *address)
	log.Printf("rds 实例id：%s", conf.InstanceId)

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		response, err := exporter.FetchLastMinuteDBInstancePerformance(exporter.Config{
			InstanceId:      conf.InstanceId,
			AccessKeyId:     conf.AccessKeyId,
			AccessKeySecret: conf.AccessKeySecret,
		})
		if err != nil {
			panic(err)
		}

		// log.Println(response.Json())

		w.Write([]byte(response.Html()))
	})

	log.Fatal(http.ListenAndServe(*address, nil))
}
