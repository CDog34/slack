package slack

import (
	"context"
	"encoding/json"
	"log"
)

type UpdateWorkflowStepReq struct {
	EditId       string                       `json:"workflow_step_edit_id,omitempty"`
	Inputs       map[string]map[string]string `json:"inputs,omitempty"`
	Outputs      []map[string]string          `json:"outputs,omitempty"`
	StepImageUrl string                       `json:"step_image_url,omitempty"`
	StepName     string                       `json:"step_name,omitempty"`
}

type CompleteWorkflowStepReq struct {
	WorkflowStepExecuteId string            `json:"workflow_step_execute_id"`
	Outputs               map[string]string `json:"outputs,omitempty"`
}

type FailWorkflowStepReq struct {
	WorkflowStepExecuteId string `json:"workflow_step_execute_id"`
	Error                 struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (api *Client) UpdateWorkflowStep(req UpdateWorkflowStepReq) (err error) {
	encoded, err := json.Marshal(req)
	if err != nil {
		return
	}
	endpoint := api.endpoint + "workflows.updateStep"
	resp := map[string]interface{}{}
	err = postJSON(context.Background(), api.httpclient, endpoint, api.token, encoded, &resp, api)
	log.Println(resp)
	return
}

func (api *Client) CompleteWorkflowStep(req CompleteWorkflowStepReq) (err error) {
	encoded, err := json.Marshal(req)
	if err != nil {
		return
	}
	endpoint := api.endpoint + "workflows.stepCompleted"
	resp := map[string]interface{}{}
	err = postJSON(context.Background(), api.httpclient, endpoint, api.token, encoded, &resp, api)
	return
}
func (api *Client) FailWorkflowStep(req FailWorkflowStepReq) (err error) {
	encoded, err := json.Marshal(req)
	if err != nil {
		return
	}
	endpoint := api.endpoint + "workflows.stepFailed"
	resp := map[string]interface{}{}
	err = postJSON(context.Background(), api.httpclient, endpoint, api.token, encoded, &resp, api)
	return
}
