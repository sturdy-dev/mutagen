package remote

import (
	"fmt"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
	"log"
)

func sturdyValidateConfiguration(request *InitializeSynchronizationRequest) error {
	if request.Alpha {
		return fmt.Errorf("must be beta")
	}
	if request.Configuration.IgnoreVCSMode != core.IgnoreVCSMode_IgnoreVCSModeIgnore {
		return fmt.Errorf("must ignore vcs")
	}
	if request.Configuration.SynchronizationMode != core.SynchronizationMode_SynchronizationModeTwoWayResolved {
		return fmt.Errorf("unexpected SynchronizationMode")
	}
	log.Println("valid configuration")
	return nil
}
