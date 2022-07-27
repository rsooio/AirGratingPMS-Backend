package customworkflow

import (
	"crypto/rand"

	"github.com/dtm-labs/dtmgrpc/workflow"
)

func MustRefister(handler workflow.WfFunc, custom ...func(wf *workflow.Workflow)) string {
	b := make([]byte, 32)
	rand.Read(b)
	name := string(b)
	err := workflow.Register(name, handler, custom...)
	if err != nil {
		return MustRefister(handler, custom...)
	} else {
		return name
	}
}
