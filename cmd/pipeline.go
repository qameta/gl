package cmd

import (
	"github.com/spf13/cobra"
	glClient "qameta.io/gl/internal/gitlab-client"
	"strings"
)

func Pipeline() *cobra.Command {
	pipeCmd := &cobra.Command{
		Use:   "pipeline",
		Short: "Runs Pipeline in Gitlab",
		Long:  "Runs Pipeline in Gitlab, use credentials before running e.g. gl auth login --token <token> --project-id 1",
	}
	runCmd := CommandBuilder(pipeCmd, Run, "run", "Runs pipeline", "Runs gitlab pipeline")
	AddStringFlag(runCmd, "ref", "", "master", "Git branch of Gitlab Repo")
	AddStringSliceFlag(runCmd, "env", "e", []string{}, "Environment variables")
	return pipeCmd
}

func Run(c *CmdConfig) error {
	runOpts := RunOpts{}
	authOpts := AuthOpts{}
	cobra.CheckErr(c.PopulateOpts(&runOpts))
	cobra.CheckErr(c.PopulateOpts(&authOpts))
	envs := map[string]string{}
	for _, env := range runOpts.Envs {
		split := strings.Split(env, "=")
		envs[strings.ToUpper(split[0])] = strings.ToUpper(split[1])
	}
	return glClient.RunPipeline(authOpts.GitlabToken, authOpts.Endpoint, authOpts.ProjectId, runOpts.Ref, envs)
}

type RunOpts struct {
	Ref  string   `mapstructure:"ref"`
	Envs []string `mapstructure:"env"`
}
