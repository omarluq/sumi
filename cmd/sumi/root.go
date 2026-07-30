package main

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	var configPath string

	cmd := &cobra.Command{
		Use:           "sumi",
		Short:         "sumi is a command line tool",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVar(&configPath, "config", "", "config file path")
	cmd.AddCommand(newConfigCmd(&configPath))
	cmd.AddCommand(newVersionCmd())

	return cmd
}
