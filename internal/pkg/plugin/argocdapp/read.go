package argocdapp

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	. "github.com/devstream-io/devstream/internal/pkg/plugin/common"
	"github.com/devstream-io/devstream/internal/pkg/plugin/installer"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
)

func Read(options configmanager.RawOptions) (statemanager.ResourceStatus, error) {
	var err error
	defer func() {
		HandleErrLogsWithPlugin(err, Name)
	}()

	// Initialize Operator with Operations
	operator := &installer.Operator{
		PreExecuteOperations: installer.PreExecuteOperations{
			validate,
		},
		GetStatusOperation: getDynamicStatus,
	}

	// Execute all Operations in Operator
	status, err := operator.Execute(options)
	if err != nil {
		return nil, err
	}
	return status, nil
}
