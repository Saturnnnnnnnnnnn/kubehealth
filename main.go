package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "kubehealth/cmd"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "kubehealth",
        Short: "Check Kubernetes cluster health",
    }

    rootCmd.AddCommand(cmd.NewCheckCommand())

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}