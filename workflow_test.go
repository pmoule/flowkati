package flowkati

import (
	"testing"

	"github.com/pmoule/flowkati/status"
)

type data struct {
	IsProcessed bool
}

func (d data) ToDto() any {
	return ""
}

func (d data) String() string {
	return ""
}

func TestNewWorkflow(t *testing.T) {
	id := "Flow"
	wfData := data{}
	wf := NewWorkflow[data](id, wfData)
	expectedID := id

	if wf.ID() != expectedID {
		t.Errorf("ID is %s, expected: %s", wf.ID(), expectedID)
	}

	expectedStepCount := 0

	if len(wf.Steps()) != expectedStepCount {
		t.Errorf("step count is %d, expected: %d", len(wf.Steps()), expectedStepCount)
	}

	if wf.ActiveStep() != nil {
		t.Errorf("ActiveStep is %v, expected: %v", wf.ActiveStep(), nil)
	}

	expectedStatus := status.Finished

	if wf.Status() != expectedStatus {
		t.Errorf("Status is %s, expected: %v", wf.Status(), expectedStatus)
	}
}

func TestAddStep(t *testing.T) {
	id := "Flow"
	wfData := data{}
	wf := NewWorkflow[data](id, wfData)
	stepTitle := "StepTitle"
	stepDescription := "StepDescription"
	step1 := createWorkflowstep[data](stepTitle, stepDescription)
	wf.AddStep(step1)
	expectedStepCount := 1

	if len(wf.Steps()) != expectedStepCount {
		t.Errorf("step count is %d, expected: %d", len(wf.Steps()), expectedStepCount)
	}
}

func TestRun(t *testing.T) {
	id := "Flow"
	wfData := data{}
	wf := NewWorkflow[data](id, wfData)
	stepTitle := "StepTitle"
	stepDescription := "StepDescription"
	step1 := createWorkflowstep[data](stepTitle, stepDescription)
	wf.AddStep(step1)
	wf.Run()
	expectedStepIndex := 1

	if wf.ActiveStep().index != expectedStepIndex {
		t.Errorf("ActiveStep index is %d, expected: %d", wf.ActiveStep().index, expectedStepIndex)
	}
}
