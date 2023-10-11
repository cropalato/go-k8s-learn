//
// main.go
// Copyright (C) 2023 rmelo <Ricardo Melo <rmelo@ludia.com>>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"context"
	"fmt"
  "log"
  "flag"
  "os"
	"path/filepath"

	//"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	// There are muliples plugins in more complex contexts
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
)

var (
  flags = log.LstdFlags | log.Lshortfile
  al = aggregatedLogger{
    infoLogger: log.New(os.Stderr, "INFO: ", flags),
    warnLogger: log.New(os.Stderr, "WARN: ", flags),
    errorLogger: log.New(os.Stderr, "ERROR: ", flags),
  }


)

func main() {
  var k8sCfg *string

  // We could check the environment variable KUBERNETES_SERVICE_HOST to be sure we are inside a k8s cluster.

  // Trying to load config from pod mount point.
  config, err := rest.InClusterConfig()
  if err != nil {

    al.warn("Unable to load a service account token from inside the Pod, trying to load from kubeconfig.")
    // Trying to load local kubeconfig file.
    k8sCfg = flag.String("kubeconfig", "", "absolute path to kubeconfig file")
    flag.Parse()

    if *k8sCfg != "" {
      config, err = clientcmd.BuildConfigFromFlags("", *k8sCfg)
    } else {
      home := homedir.HomeDir()
      config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
    }
    if err != nil {
      panic(err.Error())
    }
  }

  // Create the clientset
  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  // Get number of pods
  pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
