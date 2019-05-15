# Tradelens Integrator

TradeLens integration Golang library, provides a consigment example usage

## Configuration

Create a configuration file ~/.hodl.yaml and fill the fields:

```
apiKey: apikeygeneratedoniam
cloudIAMUrl: https://iam.ng.bluemix.net/oidc/token
solutionId: gtd-sandbox
organizationId: yourorganizationid
smRootUrl: https://platform-sandbox.tradelens.com
```

## API

Examples for webhook X-Gtd-Signature events validation can be find at:

```
examples/example_1.go
```

Usage example for the API access and token generation can be found at:

```
examples/example_2.go
```

Documentation: https://platform-sandbox.tradelens.com/documentation/swagger/
