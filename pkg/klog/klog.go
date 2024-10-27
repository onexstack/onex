package klog

import (
	"k8s.io/klog/v2"
)

func C(ctx context.Context) {
	return klog.From
}
