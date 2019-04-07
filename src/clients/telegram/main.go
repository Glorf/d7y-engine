package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type OrderContent struct {
	From     string `json:from`
	To       string `json:to`
	UnitType string `json:unitType`
}

type Order struct {
	OrderType    string `json:orderType`
	Player       string `json:player`
	Location     string `json:location`
	UnitType     string `json:unitType`
	OrderContent `json:orderContent`
}

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	bot, err := tgbotapi.NewBotAPI("Enter your api key here")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		playerFrom := update.Message.From.UserName
		orders := parseOrders(update.Message.Text, playerFrom)
		ordersToSend := make(map[string][]Order)
		ordersToSend["content"] = orders
		orderJson, err := json.Marshal(ordersToSend)
		if err != nil {
			log.Panic(err)
		}
		_, err = redisClient.SAdd("orders", orderJson, 0).Result()
		if err != nil {
			log.Panic(err)
		}
		for range orders {
			replyText := "thanks I got the message!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			bot.Send(msg)
		}
	}
}

func parseOrders(message string, username string) []Order {
	ordersStringList := strings.Split(strings.ReplaceAll(message, "-", " "), "\n")
	var ordersList []Order
	for index := range ordersStringList {
		splittedOrder := strings.Fields(ordersStringList[index])
		if len(splittedOrder) == 3 {
			order := Order{
				Player:    username,
				OrderType: "M",
				Location:  splittedOrder[1],
				UnitType:  splittedOrder[0],
				OrderContent: OrderContent{
					From: splittedOrder[1],
					To:   splittedOrder[2],
				},
			}
			ordersList = append(ordersList, order)
		} else if len(splittedOrder) == 6 {
			order := Order{
				Player:    username,
				OrderType: splittedOrder[2],
				Location:  splittedOrder[1],
				UnitType:  splittedOrder[0],
				OrderContent: OrderContent{
					UnitType: splittedOrder[3],
					From:     splittedOrder[4],
					To:       splittedOrder[5],
				},
			}
			ordersList = append(ordersList, order)
		}
	}
	return ordersList
}
