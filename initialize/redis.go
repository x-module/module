/**
 * Created by GoLand
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 16:05
 * @desc   redis.go
 */

package system

import (
	"github.com/x-module/utils/dirver"
	"github.com/x-module/utils/handler"
)

// InitializeRedisPool 初始化redis连接池
func InitializeRedisPool(config dirver.RedisConfig) {
	handler.NewRedis(config)
}
