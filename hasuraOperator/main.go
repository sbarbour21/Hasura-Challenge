package main

import (
    "context"
    "fmt"
    "os"
    "time"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/cache"
)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    configMapInformer := cache.NewSharedIndexInformer(
        &cache.ListWatch{
            ListFunc: func(options metav1.ListOptions) (object runtime.Object, e error) {
                return clientset.CoreV1().ConfigMaps("default").List(context.TODO(), options)
            },
            WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
                return clientset.CoreV1().ConfigMaps("default").Watch(context.TODO(), options)
            },
            },
            &corev1.ConfigMap{},
            0,
            cache.Indexers{},
            )

    deploymentController := cache.NewSharedIndexInformer(
        &cache.ListWatch{
            ListFunc: func(options metav1.ListOptions) (object runtime.Object, e error) {
                return clientset.AppsV1().Deployments("default").List(context.TODO(), options)
            },
            WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
                return clientset.AppsV1().ConfigMaps("default").Watch(context.TODO(), options)
            },
            },
            &corev1.Deployment{},
            0,
            cache.Indexers{},
            )
}