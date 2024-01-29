package flowkati

import (
	"testing"

	"github.com/pmoule/flowkati/status"
)

func createWorkflowstep[T StepData](title string, description string) Workflowstep[T] {
	action := func(step *Workflowstep[T], data T) {
	}
	step := NewWorkflowstep[T](title, description, action)

	return step
}

func TestNewWorkflowstep(t *testing.T) {
	stepTitle := "StepTitle"
	stepDescription := "StepDescription"
	ws := createWorkflowstep[data](stepTitle, stepDescription)

	expectedTitle := stepTitle

	if ws.Title != expectedTitle {
		t.Errorf("Title is %s, expected: %s", ws.Title, expectedTitle)
	}

	expectedDescription := stepDescription

	if ws.Description != expectedDescription {
		t.Errorf("Description is %s, expected: %s", ws.Description, expectedDescription)
	}

	expectedStatus := status.Running

	if ws.Status != expectedStatus {
		t.Errorf("Status is %s, expected: %v", ws.Status, expectedStatus)
	}
}

func TestExecute(t *testing.T) {
	stepTitle := "StepTitle"
	stepDescription := "StepDescription"
	action := func(step *Workflowstep[*data], data *data) {
		data.IsProcessed = true
	}
	ws := NewWorkflowstep[*data](stepTitle, stepDescription, action)
	ws.Data = &data{}
	ws.Execute()

	if !ws.Data.IsProcessed {
		t.Errorf("data isProcessed is %t, expected: %t", ws.Data.IsProcessed, true)
	}

	expectedDescription := stepDescription

	if ws.Description != expectedDescription {
		t.Errorf("Description is %s, expected: %s", ws.Description, expectedDescription)
	}
}

func TestFinish(t *testing.T) {
	stepTitle := "StepTitle"
	stepDescription := "StepDescription"
	action := func(step *Workflowstep[*data], data *data) {
		data.IsProcessed = true
	}
	ws := NewWorkflowstep[*data](stepTitle, stepDescription, action)
	ws.Data = &data{}
	ws.Finish()

	expectedStatus := status.Finished

	if ws.Status != expectedStatus {
		t.Errorf("Status is %s, expected: %v", ws.Status, expectedStatus)
	}
}
