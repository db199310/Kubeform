[![Build Status](https://github.com/kubeform/kubeform/workflows/CI/badge.svg)](https://github.com/kubeform/kubeform/actions?workflow=CI)
[![Docker Pulls](https://img.shields.io/docker/pulls/kubeform/kfc.svg)](https://hub.docker.com/r/kubeform/kfc/)
[![Slack](https://slack.appscode.com/badge.svg)](https://slack.appscode.com)
[![Twitter](https://img.shields.io/twitter/follow/kubeform.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=Kubeform)

# Kubeform

Kubeform by AppsCode is a Kubernetes operator for [Terraform](https://www.terraform.io/). Kubeform provides auto-generated Kubernetes CRDs for Terraform resources and modules so that you can manage any cloud infrastructure in a Kubernetes native way. You just write a CRD for a cloud infrastructure, apply it and Kubeform will create it for you! Kubeform currently supports 5 top cloud platforms. These are AWS, Google Cloud, Azure, Digitalocean and Linode.

## Features

- Native kubernetes support
- Built on Terraform
- Supports Terraform resources and modules
- Use cloud infrastructures as code
- Define & Manage cloud infrastructures as Kubernetes `CRD` (Custom Resource Definition)
- Supports multiple cloud platform
- 100% open source

## Installation

Kubeform pipeline is setup to automatically deploy the master branch in two stages. In the first stage, the deployment lands in staging (sandbox) cluster. The second stage, if approved on mgmtprd environment, deploys kubeform to production clusters that face users. It also tags master branch and increment minor version number.

On feature branch, kubeform is disabled to deploy automatically.  It can be enabled easily by changing kueform deployment manifest on the target cluster where you want the feature build to deploy to.

Assuming the feature is developed on a branch `feature/new-feature`. Then,

* If the target cluster is managed by flux then set value of flux annotation attribute in the deployment manifest to deploy image named after the feature branch. Clone the flux repository, change the kubeform deployment manifest, commit and push the change.
```yaml
kind: Deployment
metadata:
  name: kubeform
  annotations:
    fluxcd.io/automated: "true"
    fluxcd.io/tag.operator: glob: feature-new-feature-*
```
In `azure-pipeline.yml`, change variables `FLUX_REPO` and `FLUX_PATH` and set `UpdateFluxOnFeatureBranch` to `true` value.

* If the target cluster is not managed by flux then after each successfull build manually change value of `spec.template.spec.containers[0].image` attribute to value named after your branch by editing the deployment manifest.
```sh
kubectl -n kubeform edit deployment kubeform
```
```yaml
spec:
  template:
    spec:
      containers:
      - name: operator
        image: shellai.azurecr.io/sedp/kubeform/kfc:feature-new-feature-abcdef_linux_amd64
```
Also apply the built CRDs to the cluster `kubectl apply api/crds/modules*.yaml`.

Verify that the deployment has the image set,
```sh
kubectl -n kubeform get deployment apiversiontest-sedp-kubeform -o jsonpath='{.spec.template.spec.containers[0].image}{"\n"}'
shellai.azurecr.io/sedp/kubeform/kfc:new-feature_linux_amd64
```
And wait for the pod to restart,
```sh
kubectl -n kubeform get pods -w
NAME                        READY   STATUS       RESTARTS   AGE
kubeform-85655b597f-wxfpc   1/1     Terminating  0          20h
kubeform-85655b597f-9vhl4   1/1     Running      0          21s
```

## Using Kubeform

Want to learn how to use Kubeform? Please start [here](https://kubeform.com/docs/latest/guides/).

## Contribution guidelines

Want to help improve Kubeform? Please start [here](https://kubeform.com/docs/latest/welcome/contributing/).

## Acknowledgement

- Many thanks to [Hashicorp](https://www.hashicorp.com/) for [Terraform](https://www.terraform.io/) project.

## Support

We use Slack for public discussions. To chit chat with us or the rest of the community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C8NCX6N23/details/) channel `#kubeform`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

If you have found a bug with Kubeform or want to request for new features, please [file an issue](https://github.com/kubeform/project/issues/new)
