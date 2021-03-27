package sturdy

import (
	"fmt"
	"path"
	"strings"
)

func ParseCodebaseViewPath(root string) (codebaseID, viewID string, err error) {
	root = path.Clean(root)
	parts := strings.Split(root, "/")

	if len(parts) != 4 {
		return "", "", fmt.Errorf("unknown number of paths")
	}

	if parts[0] != "" || parts[1] != "repos" {
		return "", "", fmt.Errorf("invalid path")
	}

	return parts[2], parts[3], nil
}
