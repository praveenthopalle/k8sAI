![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/k8sAI-ai/k8sAI)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/k8sAI-ai/k8sAI/release.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/k8sAI-ai/k8sAI)
[![OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/projects/7272/badge)](https://bestpractices.coreinfrastructure.org/projects/7272)
[![Link to documentation](https://img.shields.io/static/v1?label=%F0%9F%93%96&message=Documentation&color=blue)](https://docs.k8sAI.ai/)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go version](https://img.shields.io/github/go-mod/go-version/k8sAI-ai/k8sAI.svg)](https://github.com/k8sAI-ai/k8sAI)
[![codecov](https://codecov.io/github/k8sAI-ai/k8sAI/graph/badge.svg?token=ZLR7NG8URE)](https://codecov.io/github/k8sAI-ai/k8sAI)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/k8sAI-ai/k8sAI/main)

`k8sAI` is a tool for scanning your Kubernetes clusters, diagnosing, and triaging issues in simple English.

It has SRE experience codified into its analyzers and helps to pull out the most relevant information to enrich it with AI.

_Out of the box integration with OpenAI, Azure, Cohere, Amazon Bedrock, Google Gemini and local models._

<a href="https://www.producthunt.com/posts/k8sAI?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-k8sAI" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=389489&theme=light" alt="k8sAI - k8sAI&#0032;gives&#0032;Kubernetes&#0032;Superpowers&#0032;to&#0032;everyone | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a> <a href="https://hellogithub.com/repository/9dfe44c18dfb4d6fa0181baf8b2cf2e1" target="_blank"><img src="https://abroad.hellogithub.com/v1/widgets/recommend.svg?rid=9dfe44c18dfb4d6fa0181baf8b2cf2e1&claim_uid=gqG4wmzkMrP0eFy" alt="Featured｜HelloGitHub" style="width: 250px; height: 54px;" width="250" height="54" /></a>


<img src="images/demo4.gif" width=650px; />

# CLI Installation

### Linux/Mac via brew

```sh
$ brew install k8sAI
```

or

```sh
brew tap k8sAI-ai/k8sAI
brew install k8sAI
```

<details>
  <summary>RPM-based installation (RedHat/CentOS/Fedora)</summary>

**32 bit:**

  <!---x-release-please-start-version-->

  ```
  sudo rpm -ivh https://github.com/k8sAI-ai/k8sAI/releases/download/v0.3.48/k8sAI_386.rpm
  ```
  <!---x-release-please-end-->

**64 bit:**

  <!---x-release-please-start-version-->
  ```
  sudo rpm -ivh https://github.com/k8sAI-ai/k8sAI/releases/download/v0.3.48/k8sAI_amd64.rpm
  ```
  <!---x-release-please-end-->
</details>

<details>
  <summary>DEB-based installation (Ubuntu/Debian)</summary>

**32 bit:**

  <!---x-release-please-start-version-->

```
curl -LO https://github.com/k8sAI-ai/k8sAI/releases/download/v0.3.48/k8sAI_386.deb
sudo dpkg -i k8sAI_386.deb
```

  <!---x-release-please-end-->

**64 bit:**

  <!---x-release-please-start-version-->

```
curl -LO https://github.com/k8sAI-ai/k8sAI/releases/download/v0.3.48/k8sAI_amd64.deb
sudo dpkg -i k8sAI_amd64.deb
```

  <!---x-release-please-end-->
</details>

<details>

  <summary>APK-based installation (Alpine)</summary>

**32 bit:**

  <!---x-release-please-start-version-->
  ```
  wget https://github.com/k8sAI-ai/k8sAI/releases/download/v0.3.48/k8sAI_386.apk
  apk add --allow-untrusted k8sAI_386.apk
  ```
  <!---x-release-please-end-->

**64 bit:**

  <!---x-release-please-start-version-->
  ```
  wget https://github.com/k8sAI-ai/k8sAI/releases/download/v0.3.48/k8sAI_amd64.apk
  apk add --allow-untrusted k8sAI_amd64.apk
  ```
  <!---x-release-please-end-->
</details>

<details>
  <summary>Failing Installation on WSL or Linux (missing gcc)</summary>
  When installing Homebrew on WSL or Linux, you may encounter the following error:

```
==> Installing k8sAI from k8sAI-ai/k8sAI Error: The following formula cannot be installed from a bottle and must be
built from the source. k8sAI Install Clang or run brew install gcc.
```

If you install gcc as suggested, the problem will persist. Therefore, you need to install the build-essential package.

```
   sudo apt-get update
   sudo apt-get install build-essential
```

</details>

### Windows

- Download the latest Windows binaries of **k8sAI** from the [Release](https://github.com/k8sAI-ai/k8sAI/releases)
  tab based on your system architecture.
- Extract the downloaded package to your desired location. Configure the system _path_ variable with the binary location

## Operator Installation

To install within a Kubernetes cluster please use our `k8sAI-operator` with installation instructions available [here](https://github.com/k8sAI-ai/k8sAI-operator)

_This mode of operation is ideal for continuous monitoring of your cluster and can integrate with your existing monitoring such as Prometheus and Alertmanager._

## Quick Start

- Currently, the default AI provider is OpenAI, you will need to generate an API key from [OpenAI](https://openai.com)
  - You can do this by running `k8sAI generate` to open a browser link to generate it
- Run `k8sAI auth add` to set it in k8sAI.
  - You can provide the password directly using the `--password` flag.
- Run `k8sAI filters` to manage the active filters used by the analyzer. By default, all filters are executed during analysis.
- Run `k8sAI analyze` to run a scan.
- And use `k8sAI analyze --explain` to get a more detailed explanation of the issues.
- You also run `k8sAI analyze --with-doc` (with or without the explain flag) to get the official documentation from Kubernetes.

## Analyzers

k8sAI uses analyzers to triage and diagnose issues in your cluster. It has a set of analyzers that are built in, but
you will be able to write your own analyzers.

### Built in analyzers

#### Enabled by default

- [x] podAnalyzer
- [x] pvcAnalyzer
- [x] rsAnalyzer
- [x] serviceAnalyzer
- [x] eventAnalyzer
- [x] ingressAnalyzer
- [x] statefulSetAnalyzer
- [x] deploymentAnalyzer
- [x] cronJobAnalyzer
- [x] nodeAnalyzer
- [x] mutatingWebhookAnalyzer
- [x] validatingWebhookAnalyzer

#### Optional

- [x] hpaAnalyzer
- [x] pdbAnalyzer
- [x] networkPolicyAnalyzer
- [x] gatewayClass
- [x] gateway
- [x] httproute
- [x] logAnalyzer

## Examples

_Run a scan with the default analyzers_

```
k8sAI generate
k8sAI auth add
k8sAI analyze --explain
k8sAI analyze --explain --with-doc
```

_Filter on resource_

```
k8sAI analyze --explain --filter=Service
```

_Filter by namespace_

```
k8sAI analyze --explain --filter=Pod --namespace=default
```

_Output to JSON_

```
k8sAI analyze --explain --filter=Service --output=json
```

_Anonymize during explain_

```
k8sAI analyze --explain --filter=Service --output=json --anonymize
```

<details>
<summary> Using filters </summary>

_List filters_

```
k8sAI filters list
```

_Add default filters_

```
k8sAI filters add [filter(s)]
```

### Examples :

- Simple filter : `k8sAI filters add Service`
- Multiple filters : `k8sAI filters add Ingress,Pod`

_Remove default filters_

```
k8sAI filters remove [filter(s)]
```

### Examples :

- Simple filter : `k8sAI filters remove Service`
- Multiple filters : `k8sAI filters remove Ingress,Pod`

</details>

<details>

<summary> Additional commands </summary>

_List configured backends_

```
k8sAI auth list
```

_Update configured backends_

```
k8sAI auth update $MY_BACKEND1,$MY_BACKEND2..
```

_Remove configured backends_

```
k8sAI auth remove -b $MY_BACKEND1,$MY_BACKEND2..
```

_List integrations_

```
k8sAI integrations list
```

_Activate integrations_

```
k8sAI integrations activate [integration(s)]
```

_Use integration_

```
k8sAI analyze --filter=[integration(s)]
```

_Deactivate integrations_

```
k8sAI integrations deactivate [integration(s)]
```

_Serve mode_

```
k8sAI serve
```

_Analysis with serve mode_

```
grpcurl -plaintext -d '{"namespace": "k8sAI", "explain" : "true"}' localhost:8080 schema.v1.ServerAnalyzerService/Analyze
{
  "status": "OK"
}
```

_Analysis with custom headers_

```
k8sAI analyze --explain --custom-headers CustomHeaderKey:CustomHeaderValue
```

_Print analysis stats_

```
k8sAI analyze -s
The stats mode allows for debugging and understanding the time taken by an analysis by displaying the statistics of each analyzer.
- Analyzer Ingress took 47.125583ms 
- Analyzer PersistentVolumeClaim took 53.009167ms 
- Analyzer CronJob took 57.517792ms 
- Analyzer Deployment took 156.6205ms 
- Analyzer Node took 160.109833ms 
- Analyzer ReplicaSet took 245.938333ms 
- Analyzer StatefulSet took 448.0455ms 
- Analyzer Pod took 5.662594708s 
- Analyzer Service took 38.583359166s
```

_Diagnostic information_

To collect diagnostic information use the following command to create a `dump_<timestamp>_json` in your local directory.
```
k8sAI dump
```

</details>

## LLM AI Backends

k8sAI uses the chosen LLM, generative AI provider when you want to explain the analysis results using --explain flag e.g. `k8sAI analyze --explain`. You can use `--backend` flag to specify a configured provider (it's `openai` by default).

You can list available providers using `k8sAI auth list`:

```
Default:
> openai
Active:
Unused:
> openai
> localai
> ollama
> azureopenai
> cohere
> amazonbedrock
> amazonsagemaker
> google
> huggingface
> noopai
> googlevertexai
> ibmwatsonxai
```

For detailed documentation on how to configure and use each provider see [here](https://docs.k8sAI.ai/reference/providers/backend/).

_To set a new default provider_

```
k8sAI auth default -p azureopenai
Default provider set to azureopenai
```

## Key Features

<details>

With this option, the data is anonymized before being sent to the AI Backend. During the analysis execution, `k8sAI` retrieves sensitive data (Kubernetes object names, labels, etc.). This data is masked when sent to the AI backend and replaced by a key that can be used to de-anonymize the data when the solution is returned to the user.

<summary> Anonymization </summary>

1. Error reported during analysis:

```bash
Error: HorizontalPodAutoscaler uses StatefulSet/fake-deployment as ScaleTargetRef which does not exist.
```

2. Payload sent to the AI backend:

```bash
Error: HorizontalPodAutoscaler uses StatefulSet/tGLcCRcHa1Ce5Rs as ScaleTargetRef which does not exist.
```

3. Payload returned by the AI:

```bash
The Kubernetes system is trying to scale a StatefulSet named tGLcCRcHa1Ce5Rs using the HorizontalPodAutoscaler, but it cannot find the StatefulSet. The solution is to verify that the StatefulSet name is spelled correctly and exists in the same namespace as the HorizontalPodAutoscaler.
```

4. Payload returned to the user:

```bash
The Kubernetes system is trying to scale a StatefulSet named fake-deployment using the HorizontalPodAutoscaler, but it cannot find the StatefulSet. The solution is to verify that the StatefulSet name is spelled correctly and exists in the same namespace as the HorizontalPodAutoscaler.
```

Note: **Anonymization does not currently apply to events.**

### Further Details

**Anonymization does not currently apply to events.**

_In a few analysers like Pod, we feed to the AI backend the event messages which are not known beforehand thus we are not masking them for the **time being**._

- The following is the list of analysers in which data is **being masked**:-

  - Statefulset
  - Service
  - PodDisruptionBudget
  - Node
  - NetworkPolicy
  - Ingress
  - HPA
  - Deployment
  - Cronjob

- The following is the list of analysers in which data is **not being masked**:-

  - RepicaSet
  - PersistentVolumeClaim
  - Pod
  - Log
  - **_\*Events_**

**\*Note**:

- k8gpt will not mask the above analysers because they do not send any identifying information except **Events** analyser.
- Masking for **Events** analyzer is scheduled in the near future as seen in this [issue](https://github.com/k8sAI-ai/k8sAI/issues/560). _Further research has to be made to understand the patterns and be able to mask the sensitive parts of an event like pod name, namespace etc._

- The following is the list of fields which are not **being masked**:-

  - Describe
  - ObjectStatus
  - Replicas
  - ContainerStatus
  - **_\*Event Message_**
  - ReplicaStatus
  - Count (Pod)

**\*Note**:

- It is quite possible the payload of the event message might have something like "super-secret-project-pod-X crashed" which we don't currently redact _(scheduled in the near future as seen in this [issue](https://github.com/k8sAI-ai/k8sAI/issues/560))_.

### Proceed with care

- The K8gpt team recommends using an entirely different backend **(a local model) in critical production environments**. By using a local model, you can rest assured that everything stays within your DMZ, and nothing is leaked.
- If there is any uncertainty about the possibility of sending data to a public LLM (open AI, Azure AI) and it poses a risk to business-critical operations, then, in such cases, the use of public LLM should be avoided based on personal assessment and the jurisdiction of risks involved.

</details>

<details>
<summary> Configuration management</summary>

`k8sAI` stores config data in the `$XDG_CONFIG_HOME/k8sAI/k8sAI.yaml` file. The data is stored in plain text, including your OpenAI key.

Config file locations:
| OS | Path |
| ------- | ------------------------------------------------ |
| MacOS | ~/Library/Application Support/k8sAI/k8sAI.yaml |
| Linux | ~/.config/k8sAI/k8sAI.yaml |
| Windows | %LOCALAPPDATA%/k8sAI/k8sAI.yaml |

</details>

<details>
There may be scenarios where caching remotely is preferred.
In these scenarios k8sAI supports AWS S3 or Azure Blob storage Integration.

<summary> Remote caching </summary>
<em>Note: You can only configure and use only one remote cache at a time</em>

_Adding a remote cache_

- AWS S3
  - _As a prerequisite `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` are required as environmental variables._
  - Configuration, `k8sAI cache add s3 --region <aws region> --bucket <name>`
  - Minio Configuration with HTTP endpoint ` k8sAI cache add s3 --bucket <name> --endpoint <http://localhost:9000>`
  - Minio Configuration with HTTPs endpoint, skipping TLS verification ` k8sAI cache add s3 --bucket <name> --endpoint <https://localhost:9000> --insecure`
    - k8sAI will create the bucket if it does not exist
- Azure Storage
  - We support a number of [techniques](https://learn.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication?tabs=bash#2-authenticate-with-azure) to authenticate against Azure
  - Configuration, `k8sAI cache add azure --storageacc <storage account name> --container <container name>`
    - k8sAI assumes that the storage account already exist and it will create the container if it does not exist
    - It is the **user** responsibility have to grant specific permissions to their identity in order to be able to upload blob files and create SA containers (e.g Storage Blob Data Contributor)
- Google Cloud Storage
  - _As a prerequisite `GOOGLE_APPLICATION_CREDENTIALS` are required as environmental variables._
  - Configuration, ` k8sAI cache add gcs --region <gcp region> --bucket <name> --projectid <project id>`
    - k8sAI will create the bucket if it does not exist

_Listing cache items_

```
k8sAI cache list
```

_Purging an object from the cache_
Note: purging an object using this command will delete upstream files, so it requires appropriate permissions.

```
k8sAI cache purge $OBJECT_NAME
```

_Removing the remote cache_
Note: this will not delete the upstream S3 bucket or Azure storage container

```
k8sAI cache remove
```

</details>

<details>
<summary> Custom Analyzers</summary>

There may be scenarios where you wish to write your own analyzer in a language of your choice.
k8sAI now supports the ability to do so by abiding by the [schema](https://github.com/k8sAI-ai/schemas/blob/main/protobuf/schema/v1/analyzer.proto) and serving the analyzer for consumption.
To do so, define the analyzer within the k8sAI configuration and it will add it into the scanning process.
In addition to this you will need to enable the following flag on analysis:

```
k8sAI analyze --custom-analysis
```

Here is an example local host analyzer in [Rust](https://github.com/k8sAI-ai/host-analyzer)
When this is run on `localhost:8080` the k8sAI config can pick it up with the following additions:

```
custom_analyzers:
  - name: host-analyzer
    connection:
      url: localhost
      port: 8080
```

This now gives the ability to pass through hostOS information ( from this analyzer example ) to k8sAI to use as context with normal analysis.

_See the docs on how to write a custom analyzer_

_Listing custom analyzers configured_
```
k8sAI custom-analyzer list
```

_Adding custom analyzer without install_
```
k8sAI custom-analyzer add --name my-custom-analyzer --port 8085
```

_Removing custom analyzer_
```
k8sAI custom-analyzer remove --names "my-custom-analyzer,my-custom-analyzer-2"
```

</details>

## Documentation

Find our official documentation available [here](https://docs.k8sAI.ai)

## Contributing

Please read our [contributing guide](./CONTRIBUTING.md).
