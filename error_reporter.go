package hydrocarbon

import "context"

type ErrorReporter interface {
	Report(ctx context.Context, err error)
}
