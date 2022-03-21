package unique

import "context"

func GetUUID(ctx context.Context) (uuid string) {
	if val := ctx.Value("id"); val != nil {
		uuid = val.(string)
	}
	return
}
