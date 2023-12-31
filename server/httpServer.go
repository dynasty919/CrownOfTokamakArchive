package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"strings"
)

func httpServer(ch chan AnsInfo) {

	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6390", // 这里使用新的 Redis 容器的地址和端口
		Password: "",               // 如果有密码，填写密码
		DB:       0,                // 默认数据库
	})
	defer client.Close()

	go Put(client, ch)

	// 注册动态路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 解析请求路径
		path := strings.Trim(r.URL.Path, "/")

		if path == "" {
			// 如果路径为空，提供 主页 页面
			DrawMainPage(w, r, client)
			return
		}

		parts := strings.Split(path, "/")
		//log.Printf("url has %d parts %s:", len(parts), parts)  // 添加这行用于调试

		if len(parts) == 0 {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// 获取 Id
		id := parts[0]

		// 从 Redis 中获取对应 Id 的数据
		jsonData, err := Get(client, id)
		if err == redis.Nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// 解码 JSON 数据到 AnsInfo 结构体
		var ansInfo AnsInfo
		if err := json.Unmarshal([]byte(jsonData), &ansInfo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 返回 JSON 数据
		WriteContentToResponseWrite(w, ansInfo)

	})

	log.Println("Server is listening on :2222...")

	http.ListenAndServe(":2222", nil)
}
