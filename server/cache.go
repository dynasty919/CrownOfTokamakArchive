package main

import (
	"CrownOfTokamak/persist"
	"CrownOfTokamak/util"
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
)

func Put(client *redis.Client, ch chan util.AnsInfo) {

	// 确保连接正常
	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)

	for {
		select {
		case info := <-ch:
			// 将 AnsInfo 结构体编码为 JSON
			jsonData, err := json.Marshal(info)
			if err != nil {
				log.Fatal(err)
			}

			// 存储 JSON 数据到 Redis，使用 ID 作为键
			err = client.Set(info.Id, jsonData, 0).Err()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("AnsInfo with ID %s title %s stored in Redis.\n", info.Id, info.Title)

			persist.Store(info)
		}
	}

}

func Get(client *redis.Client, id string) (string, error) {
	return client.Get(id).Result()
}
