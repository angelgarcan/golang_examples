package main

import (
	"flag"
	"io/ioutil"

	"k8s.io/klog/v2"
	// use "k8s.io/klog/v2" instead "k8s.io/klog"
)

/* klog main */
// https://github.com/kubernetes/klog/blob/main/examples/set_output/usage_set_output.go
func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false") // Nedded to not mix outputs.
	flag.Set("v", "1")
	// flag.Set("alsologtostderr", "false")
	flag.Set("stderrthreshold", "INFO") // INFO | WARNING | ERROR | FATAL
	flag.Parse()

	/* All output to a new buffer */
	// buf := new(bytes.Buffer)
	// klog.SetOutput(buf)
	// klog.SetOutput(os.Stdout)
	klog.SetOutput(ioutil.Discard) // Files output discarded.

	/* Separate each output by severity */
	// infoBuf := new(bytes.Buffer)
	// klog.SetOutputBySeverity("INFO", infoBuf)
	// warnBuf := new(bytes.Buffer)
	// klog.SetOutputBySeverity("WARNING", warnBuf)
	// errBuf := new(bytes.Buffer)
	// klog.SetOutputBySeverity("ERROR", errBuf)

	klog.Info("Info message")
	klog.V(2).Info("Info V2 message")
	klog.Warning("Warning message")
	klog.Error("Error message")
	// klog.Fatal("Fatal message")
	klog.Info("DONE!!!")
	klog.Flush()

	// fmt.Printf("-----INFO:\n >>%s<<\n", infoBuf.String())
	// fmt.Printf("-----WARN:\n >>%s<<\n", warnBuf.String())
	// fmt.Printf("-----ERROR:\n >>%s<<\n", errBuf.String())

}

// /* glog main */
// func main() {
// 	flag.Parse()
// 	glog.Info("Info message")
// 	glog.Warning("Warning message")
// 	glog.Error("Error message")
// 	// glog.Fatal("Fatal message")
// 	glog.Info("DONE!!!")
// 	glog.Flush()
// }
