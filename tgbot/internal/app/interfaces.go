package app

import "context"

type Searcher interface {
	GetSuitableComices(ctx context.Context, keywords string, limit int) (map[string]string, error)
}
