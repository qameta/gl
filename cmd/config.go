package cmd

import (
	"context"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Writer = os.Stdout

type CmdConfig struct {
	Args []string
	Out  io.Writer
	Ctx  context.Context
}

type Runner func(*CmdConfig) error

func NewCmdConfig(com *cobra.Command, args []string) *CmdConfig {
	return &CmdConfig{
		Args: args,
		Out:  Writer,
		Ctx:  com.Context(),
	}
}

func (c *CmdConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (c *CmdConfig) PopulateOpts(opts interface{}) error {
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}
	return nil
}

func InitConfig() {
	cfgFile := viper.GetString("config")
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("GL")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if _, err := os.Stat(cfgFile); err == nil {
		if readConfErr := viper.ReadInConfig(); readConfErr != nil {
			log.Fatalf("failed reading config: %v", readConfErr)
		}
	}
	unmarshalErr := viper.Unmarshal(&Globals)
	if unmarshalErr != nil {
		log.Fatalf("failed unmarshalling config: %v", unmarshalErr)
	}
}

func SetUpHomeConfig() string {
	cfgDir, homeDirErr := homedir.Dir()
	if homeDirErr != nil {
		log.Fatalf("failed getting user directory: %v", homeDirErr)
	}
	ch := filepath.Join(cfgDir, ".gl")
	createConfDir := os.MkdirAll(ch, 0755)
	if createConfDir != nil {
		log.Fatalf("gl failed creating config dir: %v", createConfDir)
	}
	return ch
}

type Defaults interface {
	Get() map[string]string
}
