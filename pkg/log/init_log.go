package log

import "github.com/cloudwego/kitex/pkg/klog"

func InitLog() {
	klog.SetLogger(logger)
}
