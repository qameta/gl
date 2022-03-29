package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type cmdOption func(*cobra.Command)

func CommandBuilder(parent *cobra.Command, cr Runner, cliText, shortDesc string, longDesc string, options ...cmdOption) *cobra.Command {
	c := &cobra.Command{
		Use:   cliText,
		Short: shortDesc,
		Long:  longDesc,
		PreRun: func(cmd *cobra.Command, args []string) {
			if bindErr := viper.BindPFlags(cmd.Flags()); bindErr != nil {
				log.Fatalln(bindErr)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			conf := NewCmdConfig(cmd, args)
			crErr := cr(conf)
			if crErr != nil {
				log.Fatalln(crErr)
			}
		},
	}
	if parent != nil {
		parent.AddCommand(c)
	}
	for _, co := range options {
		co(c)
	}
	return c
}

func AddStringFlag(cmd *cobra.Command, name, shorthand, dflt, desc string, opts ...flagOpt) {
	cmd.Flags().StringP(name, shorthand, dflt, desc)
	for _, opt := range opts {
		opt(cmd, name)
	}
}

func AddStringSliceFlag(cmd *cobra.Command, name, shorthand string, def []string, desc string, opts ...flagOpt) {
	cmd.Flags().StringSliceP(name, shorthand, def, desc)
	for _, o := range opts {
		o(cmd, name)
	}
}

type flagOpt func(c *cobra.Command, name string)
