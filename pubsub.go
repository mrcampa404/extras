package pubsub

import (
	"encoding/json"
	"gopkg.in/redis.v2"
)

type PubSub struct {
	client *redis.Client
}

var Service *PubSub

func init() {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     "redis-18582.c12281.us-central1-mz.gcp.cloud.rlrcp.com:18582",
		Password: "WKJeO6VNC7tZX6ae672rbXNJUUXCy3Wm",
		DB:       0,
		PoolSize: 10,
	})
	Service = &PubSub{client}
}

func (ps *PubSub) PublishString(channel, message string) *redis.IntCmd {
	return ps.client.Publish(channel, message)
}

func (ps *PubSub) Publish(channel string, message interface{}) *redis.IntCmd {
	// TODO reflect if interface{} type is string, Publish as-is
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	return ps.client.Publish(channel, messageString)
}
