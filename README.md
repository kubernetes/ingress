# NGINX Ingress Controller

[![Go Report Card](https://goreportcard.com/badge/github.com/kubernetes/ingress-nginx)](https://goreportcard.com/report/github.com/kubernetes/ingress-nginx)
[![GitHub license](https://img.shields.io/github/license/kubernetes/ingress-nginx.svg)](https://github.com/kubernetes/ingress-nginx/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/kubernetes/ingress-nginx.svg)](https://github.com/kubernetes/ingress-nginx/stargazers)
[![GitHub stars](https://img.shields.io/badge/contributions-welcome-orange.svg)](https://github.com/kubernetes/ingress-nginx/blob/main/CONTRIBUTING.md)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkubernetes%2Fingress-nginx.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkubernetes%2Fingress-nginx?ref=badge_shield)

## Overview

ingress-nginx is an Ingress controller for Kubernetes using [NGINX](https://www.nginx.org/) as a reverse proxy and load balancer.

[Learn more about Ingress on the main Kubernetes documentation site](https://kubernetes.io/docs/concepts/services-networking/ingress/).

## Get started

See the [Getting Started](https://kubernetes.github.io/ingress-nginx/deploy/) document.

## Troubleshooting

If you encounter issues, review the [troubleshooting docs](docs/troubleshooting.md), [file an issue](https://github.com/kubernetes/ingress-nginx/issues), or talk to us on the [#ingress-nginx channel](https://kubernetes.slack.com/messages/ingress-nginx) on the Kubernetes Slack server.

## Contributing

Thanks for taking the time to join our community and start contributing!

- This project adheres to the [Kubernetes Community Code of Conduct](https://git.k8s.io/community/code-of-conduct.md). By participating in this project, you agree to abide by its terms.
- See [CONTRIBUTING.md](CONTRIBUTING.md) for information about setting up your environment, the workflow that we expect, and instructions on the developer certificate of origin that we require.
- Check out the [open issues](https://github.com/kubernetes/ingress-nginx).
- Join our Kubernetes Slack channel for developer discussion : [#ingress-nginx-dev](https://kubernetes.slack.com/archives/C021E147ZA4).

## Changelog

See [the list of releases](https://github.com/kubernetes/ingress-nginx/releases) to find out about feature changes.
For detailed changes for each release; please check the [Changelog.md](Changelog.md) file.
For detailed changes on the `ingress-nginx` helm chart, please check the following [CHANGELOG.md](charts/ingress-nginx/CHANGELOG.md) file.

### Support Versions table 

| Ingress-nginx version | k8s supported version  | Alpine Version | Nginx Version |
|-----------------------|-------------           |----------------|---------------|
| v1.0.0-alpha.2        | 1.22, 1.21, 1.20, 1.19 | 3.13.5         |  1.20.1       |
| v1.0.0-alpha.1        | 1.21, 1.20, 1.19       | 3.13.5         |  1.20.1       |
| v0.48.1               | 1.21, 1.20, 1.19       | 3.13.5         |  1.20.1       |
| v0.47.0               | 1.21, 1.20, 1.19       | 3.13.5         |  1.20.1       |
| v0.46.0               | 1.21, 1.20, 1.19       | 3.13.2         |  1.19.6       |

# Get Involved

- **Contributing**: Pull requests are welcome!
  - Read [`CONTRIBUTING.md`](CONTRIBUTING.md) and check out [help-wanted](https://github.com/kubernetes/ingress-nginx/labels/help%20wanted) issues.
  - Submit github issues for any feature enhancements, bugs or documentation problems.
- **Support**: Join to [Kubernetes Slack](http://slack.kubernetes.io/) in the [#ingress-nginx-users](https://kubernetes.slack.com/messages/CANQGM8BA/) channel to ask questions to get support from the maintainers and other users.
  - The [github issues](https://github.com/kubernetes/ingress-nginx/issues) in the repository are **exclusively** for bug reports and feature requests.
- **Discuss**: Tweet using the `#IngressNginx` hashtag.

## Issues

Please make sure to read the [Issue Reporting Checklist](https://github.com/kubernetes/ingress-nginx/blob/main/CONTRIBUTING.md#issue-reporting-guidelines) before opening an issue. Issues not conforming to the guidelines **may be closed immediately**.

## License

[Apache License 2.0](https://github.com/kubernetes/ingress-nginx/blob/main/LICENSE)
