package controller

import (
	"app/controller/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/cache"
)

type Handler func(c *gin.Context) (response.Data, error)

func Wrapper(handler Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handler(c)
		buf, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		c.Data(http.StatusOK, "application/json; charset=utf-8", buf)
		c.Abort()
	}
}

type ctr struct {
	cache cache.ICache
}

//func (ctr *ctr) store(context *gin.Context, key string, expiration time.Duration, handler Handler) (response.Data, error) {
//
//	funWithErr := func() (response.Data, error) {
//		return handler(context)
//	}
//
//	return cache.StoreWithErr(context.RequestContext(), ctr.cache, key, expiration, funWithErr)
//
//	// if data, err := ctr.cache.Get(key); err == nil {
//	// 	return data, nil
//	// }
//
//	// data, err := handler(context)
//
//	// if err == nil {
//	// 	ctr.cache.Set(key, data, 10*time.Second)
//	// }
//	// return data, err
//}

//func (ctr *ctr) cacheDefaultExpiration() time.Duration {
//	return 60 * time.Minute
//}
