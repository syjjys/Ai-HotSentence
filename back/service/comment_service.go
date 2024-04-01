package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"syj/hope/entity"
	"syj/hope/infra"
	"time"

	"github.com/go-redis/redis/v8"
)

func InsertComment(comment entity.Comment) {
	// 创建一个上下文对象
	ctx := context.Background()
	jsonString, _ := json.Marshal(comment)
	// 使用HSet命令将JSON字符串插入Redis哈希
	var err error = infra.RS.HSet(ctx, "comment:"+comment.SayId, comment.RecordTime, jsonString).Err()
	if err != nil {
		panic(err)
	}
}

func GetComments(sayId string) []entity.Comment {
	ctx := context.Background()
	result, err := infra.RS.HGetAll(ctx, "comment:"+sayId).Result()
	if err == redis.Nil {
		return nil
	}
	ls := []entity.Comment{}
	for _, value := range result {
		var com entity.Comment
		json.Unmarshal([]byte(value), &com)
		ls = append(ls, com)
	}
	fmt.Println(ls)
	sort.Slice(ls, func(i, j int) bool {
		// 解析时间字符串为time.Time类型
		timeI, errI := time.Parse("2006-01-02 15:04:05", ls[i].RecordTime)
		if errI != nil {
			fmt.Println("Error parsing time:", errI)
		}
		timeJ, errJ := time.Parse("2006-01-02 15:04:05", ls[j].RecordTime)
		if errJ != nil {
			fmt.Println("Error parsing time:", errJ)
		}

		// 从最新到最旧排序
		return timeI.After(timeJ)
	})

	return ls
}
