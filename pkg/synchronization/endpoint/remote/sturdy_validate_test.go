package remote

import (
	"context"
	"fmt"
	"testing"
)

func TestSturdyValidateRoot(t *testing.T) {
	cases := []struct {
		in       string
		expected error
	}{
		{"/repos/foo/bar", nil},
		{"/repos/foo/not-bar", fmt.Errorf("invalid from api")}, // not valid by api lookup

		{"/repos/../foo/bar", fmt.Errorf("invalid path")},
		{"/repos/../foo/bar/xoxox", fmt.Errorf("invalid path")},
		{"/repos/../../../../foo/bar", fmt.Errorf("invalid path")},
		{"/etc/foo/bar", fmt.Errorf("invalid path")},
	}

	validateCodebaseView := func(_ context.Context, codebaseID, viewID string, isNewConnection bool) error {
		if codebaseID == "foo" && viewID == "bar" {
			return nil
		}
		return fmt.Errorf("invalid from api")
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			res := sturdyValidateRoot(context.TODO(), tc.in, validateCodebaseView)
			if res == nil && tc.expected == nil {
				return
			}
			if res != nil && tc.expected != nil && res.Error() == tc.expected.Error() {
				return
			}

			t.Errorf("expected=%v got=%v", tc.expected, res)
		})
	}
}
