# invoker

## Usage Guide

Follow [vHive's Quickstart Guide](https://github.com/vhive-serverless/vHive/blob/main/docs/quickstart_guide.md) to setup a single node cluster.

### Build

```
make build
sudo cp bin/invoker /usr/local/bin
```

### Deploy Worker Image

```
joshjms@node-000:~/invoker$ invoker deploy pikachu docker.io/joshjms/worker:v1
2025/08/29 01:57:10 Running docker.io/joshjms/worker:v1, service: pikachu
Warning: Kubernetes default value is insecure, Knative may default this to secure in a future release: spec.template.spec.containers[0].securityContext.allowPrivilegeEscalation, spec.template.spec.containers[0].securityContext.capabilities, spec.template.spec.containers[0].securityContext.runAsNonRoot, spec.template.spec.containers[0].securityContext.seccompProfile

Deployed service pikachu at http://pikachu.default.192.168.1.240.sslip.io
You can now run:

        invoker run pikachu.default.192.168.1.240.sslip.io:80
```

### Generate Reports

```
ENDPOINT=pikachu.default.192.168.1.240.sslip.io:80 scripts/generate_report.sh
```

The end-to-end time data and graphs can be viewed in `reports/`.

