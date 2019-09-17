package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
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
	rawConfig   clientcmdapi.Config
	args        []string

	userSpecifiedContext string
	availableContexts    []string // contains a sorted list of all contexts

	genericclioptions.IOStreams
}

// NewCtxOptions provides an instance of CtxOptions with default values
func NewCtxOptions(streams genericclioptions.IOStreams) *CtxOptions {
	return &CtxOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
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
			return opt.Run()
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

	for ctx := range o.rawConfig.Contexts {
		o.availableContexts = append(o.availableContexts, ctx)
	}
	sort.Strings(o.availableContexts)

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

// Run lists all available contexts in a user's KUBECONFIG or updates the curren context if
// the user passed one.
func (o *CtxOptions) Run() error {
	selected := []string{}
	for _, ctx := range o.availableContexts {
		if ctx == o.userSpecifiedContext {
			selected = []string{ctx}
			break
		}
		if strings.Contains(ctx, o.userSpecifiedContext) {
			selected = append(selected, ctx)
		}
	}

	switch len(selected) {
	case 0:
		return fmt.Errorf("can't change context to %q, context not found in KUBECONFIG", o.userSpecifiedContext)
	case 1:
		return o.changeCurrentCtx(selected[0])
	default:
		o.printContexts(selected)
		return nil
	}
}

func (o *CtxOptions) changeCurrentCtx(newCtx string) error {
	currentCtx := o.rawConfig.CurrentContext
	if currentCtx != newCtx {
		o.rawConfig.CurrentContext = newCtx
		if err := clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), o.rawConfig, true); err != nil {
			return err
		}
		fmt.Fprintf(o.Out, "current context set to %q\n", newCtx)
	}
	return nil
}

// prints each context in a user's KUBECONFIG, the current context is printed in red.
func (o *CtxOptions) printContexts(contexts []string) {
	red := color.New(color.FgRed)
	currentCtx := o.rawConfig.CurrentContext
	for _, ctx := range contexts {
		if ctx == currentCtx {
			red.Fprintf(o.Out, "%s\n", currentCtx)
		} else {
			fmt.Fprintf(o.Out, "%s\n", ctx)
		}
	}
}
