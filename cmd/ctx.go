package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

var (
	ctxExample = `
  # view all contexts in your KUBECONFIG
  kubectl ctx

  # switch current context to foo
  kubectl ctx foo
`
)

// CtxOptions provides information required to update the current context
// in a user's KUBECONFIG
type CtxOptions struct {
	configFlags *genericclioptions.ConfigFlags
	rawConfig   api.Config
	args        []string

	userSpecifiedContext string

	genericclioptions.IOStreams
}

// NewCtxOptions provides an instance of CtxOptions with default values
func NewCtxOptions(streams genericclioptions.IOStreams) *CtxOptions {
	return &CtxOptions{
		configFlags: genericclioptions.NewConfigFlags(),
		IOStreams:   streams,
	}
}

// NewCtxCmd provides a cobra command wrapping CtxOptions
func NewCtxCmd(streams genericclioptions.IOStreams) *cobra.Command {
	opt := NewCtxOptions(streams)

	cmd := &cobra.Command{
		Use:          "ctx [new-context]",
		Short:        "Display/Switch current context in your KUBECONFIG",
		Example:      ctxExample,
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := opt.Complete(c, args); err != nil {
				return err
			}

			if err := opt.Validate(); err != nil {
				return err
			}

			if err := opt.Run(); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}

// Complete sets all information required for updating the current context
func (o *CtxOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return err
	}

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *CtxOptions) Validate() error {
	if len(o.args) > 1 {
		return fmt.Errorf("either one or no arguments are allowed")
	}

	if len(o.args) > 0 {
		o.userSpecifiedContext = o.args[0]
	}

	return nil
}

// Run lists all available contexts in a user's KUBECONFIG or updates the current
// context if the user passed one.
func (o *CtxOptions) Run() error {
	if len(o.userSpecifiedContext) > 0 {
		if err := o.changeCurrentCtx(); err != nil {
			return err
		}
	} else {
		o.printContexts()
	}
	return nil
}

func (o *CtxOptions) changeCurrentCtx() error {
	currentCtx := o.rawConfig.CurrentContext
	newCtx := o.userSpecifiedContext

	// check if the user provided context exists
	if _, ok := o.rawConfig.Contexts[newCtx]; !ok {
		return fmt.Errorf("can't change context to \"%s\", context not found in KUBECONFIG", newCtx)
	}

	if currentCtx != newCtx {
		o.rawConfig.CurrentContext = newCtx
		if err := clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), o.rawConfig, true); err != nil {
			return err
		}
		fmt.Fprintf(o.Out, "current context set to \"%s\"\n", newCtx)
	}
	return nil
}

// prints each context in a user's KUBECONFIG, the current context is printed
// in red
func (o *CtxOptions) printContexts() {
	red := color.New(color.FgRed)
	for cluster := range o.rawConfig.Clusters {
		if cluster == o.rawConfig.CurrentContext {
			red.Fprintf(o.Out, "%s\n", cluster)
		} else {
			fmt.Fprintf(o.Out, "%s\n", cluster)
		}
	}
}
