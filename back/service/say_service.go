package service

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"syj/hope/entity"
	"syj/hope/infra"
	"time"

	"github.com/go-redis/redis/v8"
)

func InsertSay(say entity.Say) {
	// 创建一个上下文对象
	ctx := context.Background()
	say.RecordTime = time.Now().Format("2006-01-02 15:04:05")
	jsonString, _ := json.Marshal(say)
	// 使用HSet命令将JSON字符串插入Redis哈希
	var err error = infra.RS.HSet(ctx, "say", say.Id, jsonString).Err()
	if err != nil {
		panic(err)
	}
}

func GetSayById(id string) *entity.Say {
	var say entity.Say
	ctx := context.Background()
	value, err := infra.RS.HGet(ctx, "say", id).Result()
	if err == redis.Nil {
		return nil
	}
	print(value)
	json.Unmarshal([]byte(value), &say)
	return &say
}

func GetSays() []entity.Say {
	ctx := context.Background()
	result, err := infra.RS.HGetAll(ctx, "say").Result()
	if err == redis.Nil {
		return nil
	}
	ls := []entity.Say{}
	for _, value := range result {
		var say entity.Say
		json.Unmarshal([]byte(value), &say)
		say.CommentNum = strconv.Itoa(len(GetComments(say.Id)))
		say.LikeNum = GetLike(say.Id)
		ls = append(ls, say)
	}
	sort.Slice(ls, func(i, j int) bool {
		// 将LikeNum从string转换为int进行比较
		likeNumI, errI := strconv.Atoi(ls[i].LikeNum)
		if errI != nil {
			panic(errI) // 处理转换错误
		}
		likeNumJ, errJ := strconv.Atoi(ls[j].LikeNum)
		if errJ != nil {
			panic(errJ) // 处理转换错误
		}
		return likeNumI > likeNumJ // 从大到小排序
	})
	return ls
}

func Like(sayId string) {
	ctx := context.Background()
	infra.RS.Incr(ctx, "like:"+sayId)
}

func GetLike(sayId string) string {
	ctx := context.Background()
	val, err := infra.RS.Get(ctx, "like:"+sayId).Result()
	if err == redis.Nil {
		return "0"
	}
	return val
}
