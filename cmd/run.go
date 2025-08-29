/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joshjms/invoker/invoker"
	"github.com/spf13/cobra"
)

var (
	rps          int
	durationMs   int64
	distribution string
	outputPath   string
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
	var distributionOpt invoker.Distribution

	switch distribution {
	case "uniform":
		distributionOpt = invoker.DistributionUniform
	case "poisson":
		distributionOpt = invoker.DistributionPoisson
	default:
		log.Fatalf("invalid distribution\n")
	}

	opts := invoker.InvokerOptions{
		Rps:          rps,
		Distribution: distributionOpt,
		DurationMs:   durationMs,
		Endpoint:     args[0],
	}

	inv := invoker.NewInvoker(opts)
	reports, err := inv.Run()
	if err != nil {
		log.Fatalf("failed to run invoker: %v\n", err)
	}

	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed to open output file: %v\n", err)
	}
	defer f.Close()

	for _, report := range reports {
		if _, err := fmt.Fprintf(f, "%d,%d,%d,%d\n", report.RequestAt, report.WorkerStartAt, report.WorkerEndAt, report.ResponseAt); err != nil {
			log.Fatalf("failed to write report: %v\n", err)
		}
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&rps, "rps", "r", 10, "Requests per second")
	runCmd.Flags().Int64VarP(&durationMs, "duration", "d", 50, "Duration of the cpu-spin in ms")
	runCmd.Flags().StringVarP(&distribution, "distribution", "D", "uniform", "Distribution of the requests. Options: uniform, poisson")
	runCmd.Flags().StringVarP(&outputPath, "outputPath", "o", "invoke.log", "Path to the output file")
}
