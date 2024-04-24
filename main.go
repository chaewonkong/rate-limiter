package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func SlidingWindowLimiter(ctx context.Context, service string, ip string, limit int, window time.Duration) {
	key := fmt.Sprintf("rate-limit:%s:%s", service, ip)
	currentTime := time.Now()

	// ZSet 이용
	// 현재시간 - window 이전 시간 이후의 데이터만 남기고 삭제

	// window 범위 내 갯수 count

	// count >= limit; return false

	// ip를 ZSet에 Add
	// score: 현재시간 in unix, member: score와 동일

	// Expire 업데이트
}

// redis hash를 이용하는 케이스
// key: ip, field: timestamp, value: count

// redis timeseries를 이용하는 케이스

func main() {
	redisClient := NewRedisClient("localhost:6379")

}

// question
// rate-limiter를 어떻게 WAF에 적용할 것인가?
// 차단 IPSet은 어떻게 만들지?
