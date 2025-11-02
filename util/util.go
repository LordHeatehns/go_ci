package util

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

const (
	FORMAT_YYYYMMDD   = "20060102"
	FORMAT_YYYY_MM_DD = "2006-01-02"
	FORMAT_DD_MM_YYYY = "02-01-2006"
)

func StringIsEmpty(text string) bool {
	return len(strings.TrimSpace(text)) == 0

}

func StringIsNotEmpty(text string) bool {
	return len(strings.TrimSpace(text)) > 0
}

func StringTofloat64(text string) (float64, error) {
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, err
	}

	return f, nil

}

func Obj2Map[T any](obj T) map[string]any {
	mapResult := map[string]any{}
	jsonStr, _ := json.Marshal(&obj)
	err := json.Unmarshal(jsonStr, &mapResult)
	if err != nil {
		return nil
	}
	return mapResult
}

func Json[T any](obj T) string {
	jsonRaw, err := json.Marshal(&obj)
	if err != nil {
		return ""
	}
	return string(jsonRaw)
}

func NewUUID(prefix string) string {
	if len(prefix) < 3 {
		prefix = prefix + strings.Repeat("A", 3-len(prefix))
	}
	prefix = strings.ToUpper(prefix)

	id := ksuid.New()
	return prefix + id.String()
}

func CurrentTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
func GetDateNowFormat(format string) string {

	return time.Now().Format(format)
}

func GetDateFormat(timestamp int64, format string) string {
	return time.UnixMilli(timestamp).Format(format)
}

func CeliNumber(numerator int64, denominator int64) int64 {
	result := math.Ceil(float64(numerator) / float64(denominator))
	return int64(result)
}

func GetToken(token string) (string, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return "", errors.New("invalid token")
	}

	// Remove "Bearer " prefix and return the token
	token = strings.TrimPrefix(token, "Bearer ")
	return token, nil
}

// func Key(s string) string {
// 	conf, err := configuration.LoadConfigFile()
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	if StringIsEmpty(conf.Prefix_redis) {
// 		return s
// 	}
// 	return conf.Prefix_redis + s
// }

// func NewUUID(prefix string) string {
// 	if len(prefix) < 3 {
// 		prefix = prefix + strings.Repeat("A", 3-len(prefix))
// 	}
// 	prefix = strings.ToUpper(prefix)

// 	id := ksuid.New()
// 	return prefix + id.String()
// }

// func GetCurrentEpochTime() int64 {
// 	return time.Now().UnixNano() / int64(time.Millisecond)
// }

// func GenerateCacheKey(c *fiber.Ctx) string {
// 	conf, err := configuration.LoadConfigFile()
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	cacheKey := conf.Redis.Prefix
// 	cacheKey = strings.ToUpper(cacheKey)

// 	path := c.Path()
// 	cacheKey += path

// 	rawQuery := c.Request().URI().QueryString()

// 	queryParams, err := url.ParseQuery(string(rawQuery))
// 	if err != nil {
// 		log.Error(err)
// 		return cacheKey
// 	}
// 	if len(queryParams) > 0 {
// 		var keys []string
// 		for key := range queryParams {
// 			keys = append(keys, key)
// 		}
// 		sort.Strings(keys)
// 		var keyValuePairs []string
// 		for _, key := range keys {
// 			values := queryParams[key]
// 			for _, value := range values {
// 				keyValuePairs = append(keyValuePairs, key+"="+value)
// 			}
// 		}
// 		if len(keyValuePairs) > 0 {
// 			cacheKey += "?" + strings.Join(keyValuePairs, "&")
// 		}
// 	}
// 	return cacheKey
// }
