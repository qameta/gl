package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gl",
		Short: "A CLI util for interacting with Gitlab",
		Long:  `A CLI util for interacting with Gitlab`,
	}
	Globals = GlobalOptions{}
)

const (
	ArgConfig = "config"
)

type GlobalOptions struct {
	Config string `mapstructure:"config"`
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.AddCommand(Auth())
	rootCmd.AddCommand(Pipeline())
	globalFlags := rootCmd.PersistentFlags()

	globalFlags.StringVarP(&Globals.Config, ArgConfig, "", filepath.Join(SetUpHomeConfig(), "config.yaml"), "Specify a custom config file")
	cobra.CheckErr(viper.BindPFlag(ArgConfig, globalFlags.Lookup(ArgConfig)))

	cobra.OnInitialize(InitConfig)
}
