package cmd

import (
    "context"
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCheckCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "check",
        Short: "Check status of pods in all namespaces",
        Run: func(cmd *cobra.Command, args []string) {
            kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
            config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
            if err != nil {
                fmt.Println("Failed to load kubeconfig:", err)
                return
            }

            clientset, err := kubernetes.NewForConfig(config)
            if err != nil {
                fmt.Println("Failed to create Kubernetes client:", err)
                return
            }

            namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
            if err != nil {
                fmt.Println("Failed to list namespaces:", err)
                return
            }

            for _, ns := range namespaces.Items {
                pods, err := clientset.CoreV1().Pods(ns.Name).List(context.TODO(), metav1.ListOptions{})
                if err != nil {
                    fmt.Printf("Failed to list pods in namespace %s: %v\n", ns.Name, err)
                    continue
                }

                running, pending, failed := 0, 0, 0
                for _, pod := range pods.Items {
                    switch pod.Status.Phase {
                    case "Running":
                        running++
                    case "Pending":
                        pending++
                    case "Failed":
                        failed++
                    }
                }

                fmt.Printf("Namespace: %s\nRunning: %d, Pending: %d, Failed: %d\n\n", ns.Name, running, pending, failed)
            }
        },
    }
}