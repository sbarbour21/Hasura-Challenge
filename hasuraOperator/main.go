package main

import (
    "fmt"
    "os"
    "time"

)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        fmt.Fprintf(e)
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        fmt.Fprintf(e)
    }

    configMapInformer := cache.NewSharedIndexInformer(
        &cache.ListWatch{
            ListFunc: func(options metav1),
        }
        )
}