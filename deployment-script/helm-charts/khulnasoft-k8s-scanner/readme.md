# Helm chart for Khulnasoft Kubernetes Scanner

### Install

```shell
helm repo add khulnasoft-k8s-scanner https://khulnasoft-helm-charts.s3.amazonaws.com/khulnasoft-k8s-scanner
```

```shell
helm show values khulnasoft-k8s-scanner/khulnasoft-k8s-scanner
helm show readme khulnasoft-k8s-scanner/khulnasoft-k8s-scanner
```

```shell
helm install khulnasoft-k8s-scanner khulnasoft-k8s-scanner/khulnasoft-k8s-scanner \
    --set managementConsoleUrl="40.40.40.40" \
    --set khulnasoftKey="xxxxx" \
    --set clusterName="prod-cluster" \
    --namespace khulnasoft-k8s-scanner \
    --create-namespace
```
