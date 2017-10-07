package hydrocarbon

import "context"

type ReadStatusStore interface {
	MarkRead(ctx context.Context, postID, userID string)
}
