package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Unit struct {
	Type   string `json:type`
	Player string `json:player`
	Region string `json:region`
}

func main() {

	state := make(map[string][]Unit)
	state["state"] = mockState(20)
	stateJson, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	redisClient.Set("game-state", stateJson, 0).Result()

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			enableCors(&w)
			gameState := redisClient.Get("game-state").Val()
			fmt.Fprintln(w, gameState)
		},
	)
	http.ListenAndServe(":8080", nil)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// This is just for the purpose of erly stage experimenting with the front-end
type RegionFromJSON struct {
	Name string `json:name`
	Type string `json:type`
}

type Board map[string]RegionFromJSON

func mockState(max int) []Unit {
	rand.Seed(time.Now().Unix())
	var units []Unit
	jsonFile, err := os.Open("regions.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var board Board
	json.Unmarshal(byteValue, &board)
	counter := 0
	players := [7]string{"Russia", "Germany", "France", "Italy", "England", "Turkey", "Austria"}
	for id, region := range board {
		if max < counter {
			break
		}
		var unitType string
		if region.Type == "w" {
			unitType = "Fleet"
		} else if region.Type == "l" || region.Type == "x" || region.Type == "A" || region.Type == "R" || region.Type == "I" || region.Type == "E" || region.Type == "T" || region.Type == "F" || region.Type == "G" {
			unitType = "Army"
		}
		newUnit := Unit{Region: id, Type: unitType, Player: players[rand.Intn(len(players))]}
		units = append(units, newUnit)
	}
	return units
}
