package context

import "context"

type labelsContextKey string

var (
	labelsKey = labelsContextKey("labels")
)

func WithLabels(ctx context.Context, labels map[string]string) context.Context {
	return context.WithValue(ctx, labelsKey, labels)
}

func Labels(ctx context.Context) (map[string]string, bool) {
	labels, ok := ctx.Value(labelsKey).(map[string]string)
	return labels, ok
}
