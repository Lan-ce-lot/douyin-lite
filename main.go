package main

import (
	"github.com/houqingying/douyin-lite/pkg/config"
	"github.com/houqingying/douyin-lite/repository"
	"k8s.io/klog"
)

func main() {
	// 1. Initialize configuration
	if err := config.Setup(); err != nil {
		klog.Fatalf("config.Setup() error: %s", err)
	}
	// 2. Initialize database
	if err := repository.Setup(&config.Config.Database); err != nil {
		klog.Fatalf("repository.Setup() error: %s", err)
	}

	for {

	}
}
