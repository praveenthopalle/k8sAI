# serve

The serve commands allow you to run k8sAI in a grpc server mode.
This would be enabled typically through `k8sAI serve` and is how the in-cluster k8sAI deployment functions when managed by the [k8sAI-operator](https://github.com/k8sAI-ai/k8sAI-operator)

The grpc interface that is served is hosted on [buf](https://buf.build/k8sAI-ai/schemas) and the repository for this is [here](https://github.com/k8sAI-ai/schemas)

## grpcurl

A fantastic tool for local debugging and development is `grpcurl`
It allows you to form curl like requests that are http2
e.g.

```
grpcurl -plaintext -d '{"namespace": "k8sAI", "explain" : "true"}' localhost:8080 schema.v1.ServiceAnalyzeService/Analyze
```

```
grpcurl -plaintext  localhost:8080 schema.v1.ServiceConfigService/ListIntegrations
{
  "integrations": [
    "trivy"
  ]
}

```

```
grpcurl -plaintext -d '{"integrations":{"trivy":{"enabled":"true","namespace":"default","skipInstall":"false"}}}' localhost:8080 schema.v1.ServiceConfigService/AddConfig
```
