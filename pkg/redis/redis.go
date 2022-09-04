package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorbac/pkg/config"
	"gorbac/pkg/logger"
	"reflect"
	"sync"
	"time"
)

// Client *Client 对象
var Client *Redis

// Redis 对象
type Redis struct {
	mutex      sync.Mutex    // 锁
	Client     *redis.Client // 客户端
	expireTime int           // 过期时间
	DefaultDB  int           // 默认数据库，用于切换操作后，在还原回来
}

// ConnectDB 初始化模型
func (rdb *Redis) ConnectDB(selectDB int) *Redis {
	rdb.expireTime = 600     // 过期时间
	rdb.DefaultDB = selectDB // 选择库
	// 初始化 Redis 连接信息
	var (
		err       error
		RedisIp   = config.GetString("redis.host")
		RedisPort = config.GetString("redis.port")
		Pw        = config.GetString("redis.password")
	)
	if len(Pw) > 0 {
		rdb.Client = redis.NewClient(&redis.Options{
			Addr:     RedisIp + ":" + RedisPort,
			DB:       rdb.DefaultDB, // use default DB
			Password: Pw,            // no password set
		})
	} else {
		rdb.Client = redis.NewClient(&redis.Options{
			Addr: RedisIp + ":" + RedisPort,
			DB:   rdb.DefaultDB, // use default DB
		})
	}
	_, err = rdb.Client.Ping().Result()
	if err != nil {
		fmt.Println("redis连接错误" + err.Error())
	}
	return rdb
}

func Set(key string) string {
	return fmt.Sprintf("%v:%v", config.GetString("app.name"), key)
}

func (rdb *Redis) CheckExist(key string) bool {
	a, err := rdb.Client.Exists(Set(key)).Result()
	if err != nil {
		fmt.Println("判断key存在失败")
		return false
	}
	if a == 1 {
		fmt.Println("key存在")
		return true
	}
	return false
}

func (rdb *Redis) Add(key string, value interface{}, exTime int) (bool, error) {
	if exTime >= 0 {
		rdb.expireTime = exTime
	}
	err := rdb.Client.Set(Set(key), value, time.Duration(rdb.expireTime)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return false, err
	}
	return true, nil
}

func (rdb *Redis) Incr(key string) int64 {
	return rdb.Client.Incr(Set(key)).Val()
}

func (rdb *Redis) Decr(key string) int64 {
	return rdb.Client.Decr(Set(key)).Val()
}

func (rdb *Redis) ZAdd(key string, members ...redis.Z) int64 {
	return rdb.Client.ZAdd(Set(key), members...).Val()
}

func (rdb *Redis) Expire(key string, exTime int) *redis.BoolCmd {
	return rdb.Client.Expire(Set(key), time.Duration(exTime)*time.Second)
}

func (rdb *Redis) Lpush(key string, values ...interface{}) (bool, error) {
	_, err := rdb.Client.LPush(Set(key), values...).Result()
	if err != nil {
		fmt.Println("Lpush 失败")
		return false, err
	}
	return true, nil
}

func (rdb *Redis) RPop(key string) string {
	return rdb.Client.RPop(Set(key)).Val()
}

func (rdb *Redis) LLen(key string) (int64, error) {
	count, err := rdb.Client.LLen(Set(key)).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (rdb *Redis) Delete(key string) bool {
	err := rdb.Client.Del(Set(key)).Err()
	if err != nil {
		fmt.Println("删除key失败" + err.Error())
		logger.Error(err)
		return false
	}
	return true
}

func (rdb *Redis) Get(key string) (string, error) {
	defer rdb.mutex.Unlock()
	rdb.mutex.Lock()
	value, err := rdb.Client.Get(Set(key)).Result()
	return value, err
}

func (rdb *Redis) DefaultGet(key string) (string, error) {
	defer rdb.mutex.Unlock()
	rdb.mutex.Lock()
	value, err := rdb.Client.Get(key).Result()
	return value, err
}

func (rdb *Redis) WebGet(key string) (string, error) {
	defer rdb.mutex.Unlock()
	rdb.mutex.Lock()
	value, err := rdb.Client.Get(key).Result()
	return value, err
}

func (rdb *Redis) SelectDbGet(db int, key string) (string, error) {
	defer rdb.mutex.Unlock()
	rdb.mutex.Lock()
	pipe := rdb.Client.Pipeline()
	pipe.Do("select", db)
	_, _ = pipe.Get(Set(key)).Result()
	pipe.Do("select", rdb.DefaultDB)
	cmderArr, err := pipe.Exec()
	strMap := GetCmdResult(cmderArr)
	value, _ := strMap[1].(string)
	return value, err
}

func (rdb *Redis) SelectDbAdd(db int, key string, value interface{}, exTime int) (bool, error) {
	defer rdb.mutex.Unlock()
	rdb.mutex.Lock()
	pipe := rdb.Client.Pipeline()
	pipe.Do("select", db)
	if exTime >= 0 {
		rdb.expireTime = exTime
	}
	err := pipe.Set(Set(key), value, time.Duration(rdb.expireTime)*time.Second).Err()
	if err != nil {
		pipe.Do("select", rdb.DefaultDB)
		fmt.Println("设置key失败")
		return false, err
	}
	pipe.Do("select", rdb.DefaultDB)
	cmderArr, err := pipe.Exec()
	strMap := GetCmdResult(cmderArr)
	_, _ = strMap[1].(string)
	return true, err
}

func (rdb *Redis) GetKeys(db int, key string) (keys []string, cursor uint64, err error) {
	newRdb := *rdb.Client
	newRdb.Do("select", db)
	keys, cursor, err = newRdb.Scan(0, Set(key), 10000).Result()
	newRdb.Do("select", rdb.DefaultDB)
	return
}

// Redis Keys 命令 - 查找所有符合给定模式( pattern)的 key
// https://www.redis.net.cn/order/3535.html

func (rdb *Redis) Keys(key string) ([]string, error) {
	value, err := rdb.Client.Keys(Set(key)).Result()
	return value, err
}

func (rdb *Redis) SubExpireEvent(channels string) *redis.PubSub {
	// 订阅key过期事件
	//sub := Redis.Subscribe("__keyevent@0__:expired")
	return rdb.Client.Subscribe(channels)
	// 这里通过一个for循环监听redis-server发来的消息。
	// 当客户端接收到redis-server发送的事件通知时，
	// 客户端会通过一个channel告知我们。我们再根据
	// msg的channel字段来判断是不是我们期望收到的消息，
	// 然后再进行业务处理。
	//for {
	//	msg := <-sub.Channel()
	//	fmt.Println("Channel ", msg.Channel)
	//	fmt.Println("pattern ", msg.Pattern)
	//	fmt.Println("pattern ", msg.Payload)
	//}
}

func (rdb *Redis) Close() {
	err := rdb.Client.Close()
	if err != nil {
		fmt.Println("RedisCloseError", err.Error())
	}
}

func GetCmdResult(cmderArr []redis.Cmder) map[int]interface{} {
	strMap := make(map[int]interface{}, len(cmderArr))
	for idx, cmder := range cmderArr {
		StringType := reflect.TypeOf(cmder).String()
		// *ClusterSlotsCmd 未实现
		switch StringType {
		case "*redis.Cmd":
			cmd := cmder.(*redis.Cmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringCmd":
			cmd := cmder.(*redis.StringCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.SliceCmd":
			cmd := cmder.(*redis.SliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringSliceCmd":
			cmd := cmder.(*redis.StringSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringStringMapCmd":
			cmd := cmder.(*redis.StringStringMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringIntMapCmd":
			cmd := cmder.(*redis.StringIntMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.BoolCmd":
			cmd := cmder.(*redis.BoolCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.BoolSliceCmd":
			cmd := cmder.(*redis.BoolSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.IntCmd":
			cmd := cmder.(*redis.IntCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.FloatCmd":
			cmd := cmder.(*redis.FloatCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StatusCmd":
			cmd := cmder.(*redis.StatusCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.TimeCmd":
			cmd := cmder.(*redis.TimeCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.DurationCmd":
			cmd := cmder.(*redis.DurationCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringStructMapCmd":
			cmd := cmder.(*redis.StringStructMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XMessageSliceCmd":
			cmd := cmder.(*redis.XMessageSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XStreamSliceCmd":
			cmd := cmder.(*redis.XStreamSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XPendingCmd":
			cmd := cmder.(*redis.XPendingCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XPendingExtCmd":
			cmd := cmder.(*redis.XPendingExtCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.ZSliceCmd":
			cmd := cmder.(*redis.ZSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.ZWithKeyCmd":
			cmd := cmder.(*redis.ZWithKeyCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.CommandsInfoCmd":
			cmd := cmder.(*redis.CommandsInfoCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.GeoLocationCmd":
			cmd := cmder.(*redis.GeoLocationCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.GeoPosCmd":
			cmd := cmder.(*redis.GeoPosCmd)
			strMap[idx], _ = cmd.Result()
			break
		default:
			logger.Info(StringType)
			break
		}
	}
	return strMap
}
