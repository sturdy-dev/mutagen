package sturdy

import (
	"fmt"
	"path"
	"strings"
)

// path looks like:
//   ${optionalPrefix}/repos/${codebaseID}/${viewID}
func ParseCodebaseViewPath(root string) (codebaseID, viewID string, err error) {
	root = path.Clean(root)
	parts := strings.Split(root, "/")

	if len(parts) < 4 {
		return "", "", fmt.Errorf("invalid number of parts")
	}

	if parts[len(parts)-3] != "repos" {
		return "", "", fmt.Errorf("invalid path")
	}

	return parts[len(parts)-2], parts[len(parts)-1], nil
}
