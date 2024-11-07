package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

type RedisLeaderboard struct {
	client *redis.Client
	ctx    context.Context
}

// 初始化 RedisLeaderboard 实例
func NewRedisLeaderboard(addr, password string, db int) *RedisLeaderboard {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisLeaderboard{
		client: client,
		ctx:    context.Background(),
	}
}

// 获取当前周数
func getCurrentWeek() string {
	week := MondayTime(time.Now()).Format("2006-01-02")
	return week
}

func MondayTime(now time.Time) time.Time {
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	addDate := now.AddDate(0, 0, offset)

	return GetZeroTimeInLoc(addDate, now.Location())
}

// GetZeroTimeInLoc 获取0点0时0分的时间（指定时区的时间）
func GetZeroTimeInLoc(now time.Time, loc *time.Location) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

// 生成键名
func generateRoleKey(week string, roleID int) string {
	return fmt.Sprintf("demo:%s:%d", week, roleID)
}

func generateWeekKey(week string) string {
	return fmt.Sprintf("demo:%s", week)
}

// 为角色和用户添加分数，并记录时间
func (rl *RedisLeaderboard) AddScore(roleID, userID, score int) error {
	week := getCurrentWeek()
	roleKey := generateRoleKey(week, roleID)
	weekKey := generateWeekKey(week)
	timestamp := float64(time.Now().UnixNano()) / 1e18 // 使用更小的时间戳

	// 使用综合得分：主分数+时间作为次级排序
	compositeScore := float64(score) + timestamp
	_, err := rl.client.ZAdd(rl.ctx, roleKey, &redis.Z{Score: compositeScore, Member: fmt.Sprintf("%d", userID)}).Result()
	if err != nil {
		return err
	}

	// 将角色ID添加到周榜单键中，并使用综合得分进行排序
	_, err = rl.client.ZAdd(rl.ctx, weekKey, &redis.Z{Score: compositeScore, Member: fmt.Sprintf("%d", roleID)}).Result()
	return err
}

// 获取指定周的角色榜单，按总分数排序
func (rl *RedisLeaderboard) GetTopRoles(week string, limit int) ([]int, error) {
	weekKey := generateWeekKey(week)
	roleIDs, err := rl.client.ZRevRange(rl.ctx, weekKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	// 将角色ID从字符串转换为整数
	topRoles := make([]int, len(roleIDs))
	for i, id := range roleIDs {
		fmt.Sscanf(id, "%d", &topRoles[i])
	}

	return topRoles, nil
}

// 获取指定周和角色的用户榜单
func (rl *RedisLeaderboard) GetTopUsers(roleID int, week string, limit int) ([]int, error) {
	key := generateRoleKey(week, roleID)
	userIDs, err := rl.client.ZRevRange(rl.ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	// 将用户ID从字符串转换为整数
	topUsers := make([]int, len(userIDs))
	for i, id := range userIDs {
		fmt.Sscanf(id, "%d", &topUsers[i])
	}

	return topUsers, nil
}

func TestRedisLeaderboard(t *testing.T) {
	r2 := NewRedisLeaderboard("192.168.3.112:16378", "beta-redis-0", 85)

	// 测试添加分数
	if err := r2.AddScore(1, 1001, 100); err != nil {
		t.Fatalf("Failed to add score: %v", err)
	}
	if err := r2.AddScore(1, 1002, 100); err != nil {
		t.Fatalf("Failed to add score: %v", err)
	}
	if err := r2.AddScore(2, 1003, 200); err != nil {
		t.Fatalf("Failed to add score: %v", err)
	}

	// 获取当前周的周一日期
	currentWeek := getCurrentWeek()

	// 测试获取当前周的角色榜单
	topRoles, err := r2.GetTopRoles(currentWeek, 2)
	if err != nil || len(topRoles) != 2 {
		t.Fatalf("Expected 2 top roles, got: %v, err: %v", topRoles, err)
	}

	// 测试获取当前周指定角色的用户榜单
	topUsers, err := r2.GetTopUsers(1, currentWeek, 2)
	if err != nil || len(topUsers) != 2 {
		t.Fatalf("Expected 2 top users in role 1, got: %v, err: %v", topUsers, err)
	}
	fmt.Printf("%v,%v", topUsers, topRoles)
}
