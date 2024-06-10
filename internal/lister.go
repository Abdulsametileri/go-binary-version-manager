package internal

import "context"

type Lister interface {
	List(ctx context.Context, lib string) error
}
