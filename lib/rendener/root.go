package rendener

import "context"

type Renderer interface {
	Render(ctx context.Context, data any, key string, addKeys ...string) (string, error)
}
