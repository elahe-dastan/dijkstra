package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Road struct {
	RoadName    string `json:"road_name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Length      int    `json:"length"`
	SpeedLimit  int    `json:"speed_limit"`
}

type Roads struct {
	RoadsArr []Road `json:"road_details"`
}

func NewRoads(cfg string) Roads {
	file, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Fatal(err)
	}

	data := Roads{}

	err = json.Unmarshal(file, &data)
	if err != nil{
		log.Fatal(err)
	}

	return data
}