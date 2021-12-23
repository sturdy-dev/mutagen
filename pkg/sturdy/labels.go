package sturdy

func WithLock(labels map[string]string) map[string]string {
	if labels == nil {
		labels = make(map[string]string)
	}
	labels["sturdy.lock"] = "true"
	return labels
}

func useLock(labels map[string]string) bool {
	if labels == nil {
		return false
	}
	_, ok := labels["sturdy.lock"]
	return ok
}
