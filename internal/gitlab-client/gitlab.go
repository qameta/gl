package gitlab_client

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func RunPipeline(token, endpoint, projectId, ref string, vars map[string]string) error {
	gl, clientErr := gitlab.NewClient(token, gitlab.WithBaseURL(endpoint))
	if clientErr != nil {
		return clientErr
	}
	var pipelineEnvs []*gitlab.PipelineVariable
	for k, v := range vars {
		pipelineEnvs = append(pipelineEnvs, &gitlab.PipelineVariable{
			Key:          k,
			Value:        v,
			VariableType: "env_var",
		})
	}
	pl, _, pipeErr := gl.Pipelines.CreatePipeline(projectId, &gitlab.CreatePipelineOptions{Ref: &ref, Variables: &pipelineEnvs})
	if pipeErr != nil {
		return pipeErr
	}
	fmt.Printf("Started pipeline %s", pl.WebURL)
	return nil
}
