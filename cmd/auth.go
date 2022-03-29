package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gopkg.in/yaml.v2"
)

var cfgFileWriter = defaultConfigFileWriter

func Auth() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Display commands for authenticating with an account",
		Long:  "The subcommands of `gl auth` allow you to login.",
	}
	loginCmd := CommandBuilder(authCmd, RunAuthLogin, "login", "Login to gitlab", "Login to use a specific account.")
	AddStringFlag(loginCmd, "token", "t", "", "Gitlab Token")
	AddStringFlag(loginCmd, "project-id", "p", "", "Gitlab Project Id")
	AddStringFlag(loginCmd, "endpoint", "e", "https://gitlab.com", "Gitlab endpoint")
	CommandBuilder(authCmd, Print, "print", "Prints login details", "Prints login details")
	return authCmd
}

func RunAuthLogin(c *CmdConfig) error {
	authOpts := AuthOpts{}
	cobra.CheckErr(c.PopulateOpts(&authOpts))
	return writeConfig()
}

func Print(c *CmdConfig) error {
	authOpts := AuthOpts{}
	cobra.CheckErr(c.PopulateOpts(&authOpts))
	fmt.Println(fmt.Sprintf("Logged in with Token: %s, Project: %s", authOpts.GitlabToken, authOpts.ProjectId))
	return nil
}

type AuthOpts struct {
	Endpoint    string `mapstructure:"endpoint"`
	GitlabToken string `mapstructure:"token"`
	ProjectId   string `mapstructure:"project-id"`
}

func writeConfig() error {
	f, fwErr := cfgFileWriter()
	defer cobra.CheckErr(fwErr)
	b, marshallErr := yaml.Marshal(filter(viper.AllSettings()))
	if marshallErr != nil {
		return fmt.Errorf("unable to encode configuration to YAML format")
	}
	_, writeErr := f.Write(b)
	if writeErr != nil {
		return fmt.Errorf("unable to write configuration")
	}
	return nil
}

func defaultConfigFileWriter() (io.WriteCloser, error) {
	cfgFile := viper.GetString("config")
	f, createErr := os.Create(cfgFile)
	if createErr != nil {
		return nil, createErr
	}
	if chmodErr := os.Chmod(cfgFile, 0600); chmodErr != nil {
		return nil, chmodErr
	}
	return f, nil
}

func filter(data map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for key, value := range data {
		result[key] = value
	}
	return result
}
