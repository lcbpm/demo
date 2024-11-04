package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	numGoroutines = 1         // 增加协程数量
	batchSize     = 1000000   // 增加每个批次插入的数据量
	totalRecords  = 200000000 // 总插入数据量
	targetRecords = 400000    // 符合特定条件的数据量
)

func Test(t *testing.T) {
	// 数据库连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/mysql"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	recordsPerGoroutine := totalRecords / numGoroutines
	targetPerGoroutine := targetRecords / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(startID, targetCount int) {
			defer wg.Done()
			insertData(db, startID, recordsPerGoroutine, targetCount)
		}(i*recordsPerGoroutine, targetPerGoroutine)
	}

	wg.Wait()
}

func insertData(db *sql.DB, startID, count, targetCount int) {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	rand.Seed(time.Now().UnixNano())

	for i := startID; i < startID+count; i += batchSize {
		tx, err := db.Begin()
		if err != nil {
			log.Println(err)
			continue
		}

		stmt, err := tx.Prepare(`
            INSERT INTO talker_generator_message 
            (session_id, tmp_id, message_id, type, message, status, io, is_summary, created_at, updated_at, deleted_at) 
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        `)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			continue
		}

		records := make([][]interface{}, 0, batchSize)
		for j := i; j < i+batchSize && j < startID+count; j++ {
			sessionID := generateSessionID()
			tmpID := fmt.Sprintf("tmp_%d", j)
			status := 1
			var deletedAt *time.Time
			if rand.Float64() < 0.02 { // 随机判断
				sessionID = "2_83671:0_1292564"
				tmpID = ""
				status = 1
				deletedAt = nil
			} else {
				sessionID = generateSessionID()
				tmpID = ""
				status = 1
				deletedAt = nil
			}
			createdAt := randomDate(startDate, endDate)
			message := fmt.Sprintf("%v", "test" /*randomChineseString(100)*/)
			messageType := "type_value"
			io := rand.Intn(2) + 1
			isSummary := rand.Intn(2)
			record := []interface{}{sessionID, tmpID, j, messageType, message, status, io, isSummary, createdAt, createdAt, deletedAt}
			records = append(records, record)
		}

		// 打乱插入数据的顺序
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })

		for _, record := range records {
			_, err := stmt.Exec(record...)
			if err != nil {
				log.Println(err)
			}
		}

		stmt.Close()
		err = tx.Commit()
		if err != nil {
			log.Println(err)
		}

		log.Printf("Inserted batch starting at %d", i)
	}
}

func randomDate(start, end time.Time) *time.Time {
	delta := end.Sub(start).Seconds()
	random := rand.Int63n(int64(delta))
	randomDate := start.Add(time.Duration(random) * time.Second)
	return &randomDate
}

func generateSessionID() string {
	return fmt.Sprintf("2_%d:0_%d", rand.Intn(1000000), rand.Intn(1000000))
}

// 生成随机中文字符
func randomChineseCharacter() string {
	base := 0x4E00 // 汉字的 Unicode 编码起始点
	end := 0x9FA5  // 汉字的 Unicode 编码结束点
	r := rand.Intn(end-base+1) + base
	return string(rune(r))
}

// 生成指定长度的随机中文字符串
func randomChineseString(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		s += randomChineseCharacter()
	}
	return s
}
