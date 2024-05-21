package main

import (
	"flag"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/khulnasoft-lab/kubernetes-scanner/v2/scanner/compliance"
	"github.com/khulnasoft-lab/kubernetes-scanner/v2/util"
	"github.com/sirupsen/logrus"
)

var (
	managementConsoleUrl  = flag.String("mgmt-console-url", "", "Khulnasoft Management Console URL")
	managementConsolePort = flag.Int("mgmt-console-port", 443, "Khulnasoft Management Console Port")
	clusterName           = flag.String("cluster-name", "", "Cluster Name")
	debug                 = flag.Bool("debug", false, "set log level to debug")
)

func main() {
	flag.Parse()

	// setup logrus
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		PadLevelText:           true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// return funcName(f.Func.Name()) + "()", " " + path.Base(f.File) + ":" + strconv.Itoa(f.Line)
			return "", path.Base(f.File) + ":" + strconv.Itoa(f.Line)
		},
	})

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	nodeId := util.GetKubernetesClusterId()
	if nodeId == "" {
		nodeId = *clusterName
	}
	config := util.Config{
		ManagementConsoleUrl:  *managementConsoleUrl,
		ManagementConsolePort: strconv.Itoa(*managementConsolePort),
		KhulnasoftKey:          os.Getenv("KHULNASOFT_KEY"),
		NodeName:              *clusterName,
		NodeId:                nodeId,
	}

	complianceScanner, err := compliance.NewComplianceScanner(config)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	err = complianceScanner.RunComplianceScan()
	if err != nil {
		logrus.Error(err.Error())
	}
	// read results from file
}
