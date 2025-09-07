package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/william1nguyen/ping-battle/internal/config"
	"github.com/william1nguyen/ping-battle/internal/model"
	"github.com/william1nguyen/ping-battle/internal/redis"
)

var (
	topKey         = "ping:top"
	leaderboardKey = "ping:unique"
)

type LeaderBoardEntry struct {
	Username string
	Score    float64
}

func CreateSession(username string) *model.User {
	sessionID := uuid.New().String()
	key := fmt.Sprintf("session:%s", sessionID)
	redis.Rdb.Set(key, username, time.Duration(config.Cfg.SessionTTL)*time.Second)
	return &model.User{
		Username:  username,
		SessionID: sessionID,
	}
}

func validateLock(username string) error {
	lockKey := fmt.Sprintf("ping:lock:%s", username)

	if ok, _ := redis.Rdb.SetNX(lockKey, 1, 5*time.Second).Result(); !ok {
		return fmt.Errorf("already processing")
	}

	defer redis.Rdb.Del(lockKey)
	return nil
}

func validateRateLimit(username string) error {
	rateKey := fmt.Sprintf("ping:rate:%s", username)
	cnt, _ := redis.Rdb.Incr(rateKey).Result()

	if cnt == 1 {
		redis.Rdb.Expire(rateKey, 60*time.Second)
	}

	if cnt > 2 {
		return fmt.Errorf("rate limit exceed")
	}

	return nil
}

func updateStatus(username string) {
	countKey := fmt.Sprintf("ping:count:%s", username)
	redis.Rdb.Incr(countKey)

	redis.Rdb.ZIncrBy(topKey, 1, username)
	redis.Rdb.PFAdd(leaderboardKey, username)

	time.Sleep(5 * time.Second)
}

func ProcessPing(username string) (string, error) {
	if err := validateLock(username); err != nil {
		return "", err
	}

	if err := validateRateLimit(username); err != nil {
		return "", err
	}

	updateStatus(username)

	return "ping", nil
}

func GetTopUsers() ([]LeaderBoardEntry, error) {
	zset, err := redis.Rdb.ZRevRangeWithScores(topKey, 0, 9).Result()
	if err != nil {
		return nil, err
	}

	var top []LeaderBoardEntry
	for _, z := range zset {
		top = append(top, LeaderBoardEntry{
			Username: z.Member.(string),
			Score:    z.Score,
		})
	}

	return top, nil
}

func GetUniqueCount() (int64, error) {
	return redis.Rdb.PFCount(leaderboardKey).Result()
}
