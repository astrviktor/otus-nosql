package service

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"project/internal/config"
	"project/internal/model"
	"time"
)

type Handler struct {
	log *zap.Logger
	cfg *config.Config

	client *fasthttp.Client
	rdb    *redis.Client

	ctx  context.Context
	done context.CancelFunc
}

func NewHandler(log *zap.Logger, cfg *config.Config) (*Handler, error) {
	client := &fasthttp.Client{}
	client.ReadTimeout = 30 * time.Second
	client.WriteTimeout = 30 * time.Second
	client.MaxConnsPerHost = 1024

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Service.Redis.Host, cfg.Service.Redis.Port),
		Password: cfg.Service.Redis.Password,
		DB:       cfg.Service.Redis.DB,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Debug("redis ping", zap.String("pong", pong))

	ctx, done := context.WithCancel(context.Background())

	h := &Handler{
		log:    log,
		cfg:    cfg,
		client: client,
		rdb:    rdb,
		ctx:    ctx,
		done:   done,
	}

	return h, nil
}

func (h *Handler) GetRedis() *redis.Client {
	return h.rdb
}

func randomString(n int) string {
	length := n / 2
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

const testCount = 10

func (h *Handler) makeTestData() []model.Record {
	var testData []model.Record

	for i := 0; i < testCount; i++ {
		record := model.Record{
			Id:   uuid.New().String(),
			Data: randomString(h.cfg.Service.DataSize),
		}

		bytes, _ := json.Marshal(record)
		record.Binary = string(bytes)

		testData = append(testData, record)
	}

	return testData
}

func (h *Handler) TestRedisString(ctx *fasthttp.RequestCtx) {
	testData := h.makeTestData()
	h.rdb.FlushDB(h.ctx)

	result := model.Result{
		Name: "redis string test, write / read, sec",
		Size: h.cfg.Service.DataSize,
	}

	var err error

	now := time.Now()
	for _, data := range testData {

		err = h.rdb.Set(h.ctx, data.Id, data.Binary, 0).Err()
		if err != nil {
			h.log.Error("redis string set error", zap.Error(err))
		}
	}
	result.WriteDuration = time.Since(now).Seconds() / testCount

	now = time.Now()
	for _, data := range testData {
		_, err = h.rdb.Get(h.ctx, data.Id).Result()
		if err != nil {
			h.log.Error("redis string get error", zap.Error(err))
		}
	}
	result.ReadDuration = time.Since(now).Seconds() / testCount

	body, err := json.Marshal(result)
	if err != nil {
		h.log.Error("fail to marshal data", zap.Error(err))
		ctx.Error("fail to marshal data", http.StatusInternalServerError)
		return
	}

	ctx.Success("application/json", body)
}

func (h *Handler) TestRedisHset(ctx *fasthttp.RequestCtx) {
	testData := h.makeTestData()
	h.rdb.FlushDB(h.ctx)

	result := model.Result{
		Name: "redis hset test, write / read, sec",
		Size: h.cfg.Service.DataSize,
	}

	var err error

	now := time.Now()
	for _, data := range testData {
		err = h.rdb.HSet(h.ctx, data.Id, data).Err()
		if err != nil {
			h.log.Error("redis hset set error", zap.Error(err))
		}
	}
	result.WriteDuration = time.Since(now).Seconds() / testCount

	now = time.Now()
	for _, data := range testData {
		_, err = h.rdb.HGet(h.ctx, data.Id, "data").Result()
		if err != nil {
			h.log.Error("redis hset get error", zap.Error(err))
		}
	}
	result.ReadDuration = time.Since(now).Seconds() / testCount

	body, err := json.Marshal(result)
	if err != nil {
		h.log.Error("fail to marshal data", zap.Error(err))
		ctx.Error("fail to marshal data", http.StatusInternalServerError)
		return
	}

	ctx.Success("application/json", body)
}

func (h *Handler) TestRedisZset(ctx *fasthttp.RequestCtx) {
	testData := h.makeTestData()
	h.rdb.FlushDB(h.ctx)

	result := model.Result{
		Name: "redis zset test, write / read, sec",
		Size: h.cfg.Service.DataSize,
	}

	var err error

	now := time.Now()
	for _, data := range testData {

		err = h.rdb.ZAdd(h.ctx, data.Id, redis.Z{Score: 10, Member: data.Binary}).Err()
		if err != nil {
			h.log.Error("redis zset set error", zap.Error(err))
		}
	}
	result.WriteDuration = time.Since(now).Seconds() / testCount

	now = time.Now()
	for _, data := range testData {
		_, err = h.rdb.ZRem(h.ctx, data.Id, "data").Result()
		if err != nil {
			h.log.Error("redis zset get error", zap.Error(err))
		}
	}
	result.ReadDuration = time.Since(now).Seconds() / testCount

	body, err := json.Marshal(result)
	if err != nil {
		h.log.Error("fail to marshal data", zap.Error(err))
		ctx.Error("fail to marshal data", http.StatusInternalServerError)
		return
	}

	ctx.Success("application/json", body)
}

func (h *Handler) TestRedisList(ctx *fasthttp.RequestCtx) {
	testData := h.makeTestData()
	h.rdb.FlushDB(h.ctx)

	result := model.Result{
		Name: "redis list test, write / read, sec",
		Size: h.cfg.Service.DataSize,
	}

	var err error

	now := time.Now()
	for _, data := range testData {

		err = h.rdb.RPush(h.ctx, data.Id, data).Err()
		if err != nil {
			h.log.Error("redis list set error", zap.Error(err))
		}
	}
	result.WriteDuration = time.Since(now).Seconds() / testCount

	now = time.Now()
	for _, data := range testData {
		_, err = h.rdb.LPop(h.ctx, data.Id).Result()
		if err != nil {
			h.log.Error("redis list get error", zap.Error(err))
		}
	}
	result.ReadDuration = time.Since(now).Seconds() / testCount

	body, err := json.Marshal(result)
	if err != nil {
		h.log.Error("fail to marshal data", zap.Error(err))
		ctx.Error("fail to marshal data", http.StatusInternalServerError)
		return
	}

	ctx.Success("application/json", body)
}

func (h *Handler) Run() {}

func (h *Handler) Stop() {}
