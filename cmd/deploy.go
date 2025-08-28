/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys the worker service in Knative",
	Long:  `Deploys the worker service in Knative`,
	Run:   FuncDeploy,
}

func FuncDeploy(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalf("required 1 argument\n")
	}

	svcName := uuid.NewString()
	image := args[0]

	log.Printf("Running %s, service: %s\n", image, svcName)

	deploySvcCmd := exec.Command("kn", "service", "create", svcName, fmt.Sprintf("--image=%s", image))
	deploySvcCmd.Stdout = os.Stdout
	deploySvcCmd.Stderr = os.Stderr

	if err := deploySvcCmd.Run(); err != nil {
		log.Fatalf("failed to deploy service: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
