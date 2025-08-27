/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: RunFunc,
}

func RunFunc(cmd *cobra.Command, args []string) {
	fmt.Println("run called")
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&rps, "rps", "r", 10, "Requests per second")
	runCmd.Flags().Int64VarP(&durationMs, "duration", "d", 100, "Duration of the cpu-spin in ms")
	runCmd.Flags().StringVarP(&distribution, "distribution", "D", "uniform", "Distribution of the requests. Options: uniform, poisson")
	runCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "Worker endpoint")
	runCmd.MarkFlagRequired("endpoint")
}
