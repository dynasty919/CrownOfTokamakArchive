package persist

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
 	"CrownOfTokamak/server"
)

var (
	once sync.Once
	chDb chan AnsInfo
)

func init() {
	once.Do(initDb)
}

func initDb() {
	// 数据库连接信息
	db, err := sql.Open("mysql", "root:fuckyou@tcp(127.0.0.1:3307)/tok")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		select {
		case info := <-chDb:
			// 插入记录
			_, err = db.Exec("INSERT INTO tbl_file (Author, Title, Content, PostTime, Counter, Id) VALUES (?, ?, ?, ?, ?, ?)",
				info.,
				"标题",
				"文章内容",
				time.Now(),
				42,
				"标题的Sha1哈希",
			)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("AnsInfo with ID %s title %s stored in Mysql.\n", info.Id, info.Title)
		}
	}


}

func Store(jsonData []byte) error {
 	go func() {
 		chDb <- jsonData
	}()
}
