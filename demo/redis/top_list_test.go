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

// 为角色和用户添加分数，并记录时间
func (rl *RedisLeaderboard) AddScore(roleID, userID string, score int) error {
	key := "role:" + roleID
	timestamp := float64(time.Now().UnixNano()) / 1e9 // 使用秒级时间戳

	// 使用综合得分：主分数+时间作为次级排序
	compositeScore := float64(score*1000000) + (1e9 - timestamp)
	_, err := rl.client.ZAdd(rl.ctx, key, &redis.Z{Score: compositeScore, Member: userID}).Result()
	return err
}

// 获取角色榜单，按总分数排序
func (rl *RedisLeaderboard) GetTopRoles(limit int) ([]string, error) {
	roleKeys, err := rl.client.Keys(rl.ctx, "role:*").Result()
	if err != nil {
		return nil, err
	}

	roleScores := make([]struct {
		RoleID string
		Score  float64
	}, 0)

	for _, key := range roleKeys {
		score, err := rl.client.ZScore(rl.ctx, key, "*").Result()
		if err != nil {
			continue
		}
		roleID := key[len("role:"):]
		roleScores = append(roleScores, struct {
			RoleID string
			Score  float64
		}{RoleID: roleID, Score: score})
	}

	// 排序按总分数，获取前 N 个角色
	topRoles := make([]string, 0, limit)
	for i := 0; i < limit && i < len(roleScores); i++ {
		topRoles = append(topRoles, roleScores[i].RoleID)
	}
	return topRoles, nil
}

// 获取指定角色的用户榜单
func (rl *RedisLeaderboard) GetTopUsers(roleID string, limit int) ([]string, error) {
	key := "role:" + roleID
	users, err := rl.client.ZRevRange(rl.ctx, key, 0, int64(limit-1)).Result()
	return users, err
}

func main() {
	rl := NewRedisLeaderboard("localhost:6379", "", 0)

	// 添加分数示例
	rl.AddScore("warrior", "user1", 100)
	rl.AddScore("warrior", "user2", 100) // 相同分数，时间早的优先
	rl.AddScore("mage", "user3", 200)

	// 获取角色榜单
	topRoles, _ := rl.GetTopRoles(2)
	fmt.Println("Top Roles:", topRoles)

	// 获取指定角色下的用户榜单
	topUsers, _ := rl.GetTopUsers("warrior", 3)
	fmt.Println("Top Users in warrior:", topUsers)
}

func TestRedisLeaderboard(t *testing.T) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	rl := &RedisLeaderboard{client: client, ctx: ctx}

	// 测试添加分数
	if err := rl.AddScore("warrior", "user1", 100); err != nil {
		t.Fatalf("Failed to add score: %v", err)
	}
	if err := rl.AddScore("warrior", "user2", 100); err != nil {
		t.Fatalf("Failed to add score: %v", err)
	}
	if err := rl.AddScore("mage", "user3", 200); err != nil {
		t.Fatalf("Failed to add score: %v", err)
	}

	// 测试获取角色榜单
	topRoles, err := rl.GetTopRoles(2)
	if err != nil || len(topRoles) != 2 {
		t.Fatalf("Expected 2 top roles, got: %v, err: %v", topRoles, err)
	}

	// 测试获取指定角色的用户榜单
	topUsers, err := rl.GetTopUsers("warrior", 2)
	if err != nil || len(topUsers) != 2 {
		t.Fatalf("Expected 2 top users in warrior, got: %v, err: %v", topUsers, err)
	}
}
