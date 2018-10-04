package main

import (
	"os"

	"github.com/postfinance/kubectl-ctx/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-ctx", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCtxCmd(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
