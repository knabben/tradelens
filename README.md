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

## API hit

POST, GET and PATCH API for:

https://platform-sandbox.tradelens.com/documentation/swagger/

Usage example for the API can be found at:

```
examples/main.go
```