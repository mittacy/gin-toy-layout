package bizUtil

import "context"

const (
	RequestIdKey = "requestId"
)

func GetRequestId(c context.Context) string {
	if v, ok := c.Value(RequestIdKey).(string); ok {
		return v
	}
	return ""
}
