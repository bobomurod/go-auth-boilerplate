package ports

import "context"

type DatabaseClient interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
}
