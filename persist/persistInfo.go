package persist

import (
	"CrownOfTokamak/util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var (
	once sync.Once
	chDb chan util.AnsInfo
)

func init() {
	chDb = make(chan util.AnsInfo)
	once.Do(initDb)
}

func initDb() {
	go persistData() // 数据库连接信息
}

func persistData() {
	log.Println("Persisting")
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
				info.Author,
				info.Title,
				info.Content,
				info.PostTime,
				info.Counter,
				info.Id,
			)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("AnsInfo with ID %s title %s stored in Mysql.\n", info.Id, info.Title)
		}
	}
}

func Store(jsonData util.AnsInfo) {
	go func() {
		chDb <- jsonData
	}()
}
