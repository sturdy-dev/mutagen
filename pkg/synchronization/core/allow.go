package core

// allower is just an ignorer with a different interface for readability.
type allower struct {
	ignorer *ignorer
}

func newAllower(patterns []string) (*allower, error) {
	ignorer, err := newIgnorer(patterns)
	if err != nil {
		return nil, err
	}
	return &allower{
		ignorer: ignorer,
	}, nil
}

func (a *allower) allowed(path string, directory bool) bool {
	return a.ignorer.ignored(path, directory)
}
