package main

import (
	"os"

	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"demo1/pkg/cmd/exec"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-pod", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := exec.NewDebugCmd(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}