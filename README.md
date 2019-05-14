# Tradelens Integrator

Code for integration on Tradelens platform, 
provides fetch and post code and a listener for events consuming.

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

```
go run main.go fetch
```

## Listening for Events

Events subscription listening and Redis storage

https://docs.tradelens.com/how_to/setting_up_subscriptions/

```
go run main.go listen
```

This endpoint offer a proxy for Redis consuming and third party integrations