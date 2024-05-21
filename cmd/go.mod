module github.com/khulnasoft-lab/kubernetes-scanner/cmd/v2

go 1.22.2

replace github.com/khulnasoft-lab/kubernetes-scanner/v2 => ../

require (
	github.com/khulnasoft-lab/kubernetes-scanner/v2 v2.0.0
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
