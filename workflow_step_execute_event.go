package slack

type WorkflowStepExecuteEvent struct {
	Type         string       `json:"type"`
	CallbackId   string       `json:"callback_id"`
	WorkflowStep WorkflowStep `json:"workflow_step"`
	EventTS      string       `json:"event_ts"`
}

type WorkflowStep struct {
	WorkflowStepExecuteId string                           `json:"workflow_step_execute_id"`
	WorkflowId            string                           `json:"workflow_id"`
	WorkflowInstanceId    string                           `json:"workflow_instance_id"`
	StepId                string                           `json:"step_id"`
	Inputs                map[string]WorkflowStepInputItem `json:"inputs"`
	Outputs               []WorkflowStepInputItem          `json:"outputs"`
}

type WorkflowStepInputItem struct {
	Value string `json:"value"`
}
type WorkflowStepOutputItem struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Label string `json:"label"`
}
