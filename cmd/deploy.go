/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys the worker service in Knative",
	Long:  `Deploys the worker service in Knative`,
	Run:   FuncDeploy,
}

func FuncDeploy(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatalf("required 2 arguments\n")
	}

	svcName := args[0]
	image := args[1]

	log.Printf("Running %s, service: %s\n", image, svcName)

	outBuf := &bytes.Buffer{}

	deploySvcCmd := exec.Command("kn", "service", "create", svcName, fmt.Sprintf("--image=%s", image), "--port=h2c:8080")
	deploySvcCmd.Stdout = outBuf
	deploySvcCmd.Stderr = os.Stderr

	if err := deploySvcCmd.Run(); err != nil {
		log.Fatalf("failed to deploy service: %v\n", err)
	}

	outStr := outBuf.String()
	httpEndpoint := strings.Split(outStr, "\n")[len(strings.Split(outStr, "\n"))-2]
	grpcEndpoint := strings.ReplaceAll(httpEndpoint, "http://", "")
	grpcEndpoint += ":80"

	fmt.Printf("\nDeployed service %s at %s\n", svcName, httpEndpoint)
	fmt.Printf("You can now run:\n\n\tinvoker run %s\n\n", grpcEndpoint)
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
