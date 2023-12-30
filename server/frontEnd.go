package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"html/template"
	"net/http"
	"os"
	"sort"
	"strings"
)

func WriteContentToResponseWrite(w http.ResponseWriter, info AnsInfo) {

	tmpl, err := template.ParseFiles("./server/template/article.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := make(map[string]interface{})
	form["Title"] = info.Title
	form["Author"] = info.Author
	form["PostTime"] = info.PostTime
	form["Content"] = template.HTML(strings.ReplaceAll(info.Content, "\n", "<br>"))

	// 将数据填充到模板中
	err = tmpl.Execute(w, form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DrawMainPage(w http.ResponseWriter, r *http.Request, client *redis.Client) {
	// 获取所有 AnsInfo
	ansInfos, err := getAllAnsInfo(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 读取 HTML 模板文件
	htmlBytes, err := os.ReadFile("./server/template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	htmlTemplate := string(htmlBytes)

	// 解析 HTML 模板
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将数据填充到模板中
	err = tmpl.Execute(w, ansInfos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAllAnsInfo(client *redis.Client) ([]AnsInfo, error) {
	// 获取所有键值对
	keys, err := client.Keys("*").Result()
	if err != nil {
		return nil, err
	}

	var ansInfos []AnsInfo

	// 遍历所有键
	for _, key := range keys {
		// 获取值
		jsonData, err := client.Get(key).Result()
		if err != nil {
			return nil, err
		}

		// 解析为 AnsInfo 结构体
		var ansInfo AnsInfo
		err = json.Unmarshal([]byte(jsonData), &ansInfo)
		if err != nil {
			return nil, err
		}

		ansInfos = append(ansInfos, ansInfo)
	}

	// 按 Counter 从大到小排序
	sort.Slice(ansInfos, func(i, j int) bool {
		return ansInfos[i].Counter > ansInfos[j].Counter
	})

	return ansInfos, nil
}
