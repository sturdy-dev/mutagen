// +build !windows,!darwin

package filesystem

// probeExecutabilityPreservationFastByPath attempts to perform a fast
// executability preservation test by path, without probe files. The
// successfulness of the test is indicated by the second return parameter.
func probeExecutabilityPreservationFastByPath(path string) (bool, bool) {
	return false, false
}

// probeExecutabilityPreservationFast attempts to perform a fast executability
// preservation test, without probe files. The successfulness of the test is
// indicated by the second return parameter.
func probeExecutabilityPreservationFast(directory *Directory) (bool, bool) {
	return false, false
}
