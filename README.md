[![CircleCI](https://circleci.com/gh/giantswarm/athena.svg?style=shield&circle-token=e86584296950c3be820856e802a37336c2c7d540)](https://circleci.com/gh/giantswarm/athena)
[![Docker Repository on Quay](https://quay.io/repository/giantswarm/athena/status?token=789fdfa0-2068-473d-9ab6-7bde4aaf46dc "Docker Repository on Quay")](https://quay.io/repository/giantswarm/athena)

# Athena

Athena is a service that knows some useful things about your cluster. Its purpose is to provide some non-sensitive data (e. g. the CA certificate of the Kubernetes API, the cluster identifier, the cloud provider) to public clients, so they could easily establish a connection with the Kubernetes API, and identify the cluster that they're talking to.

Athena is typically running in every Giant Swarm management cluster, but is also useful in workload clusters.

## Installing in a workload cluster

If [Dex](https://github.com/giantswarm/dex-app) is already configured in the workload cluster, Athena can be used to provide OIDC access information to [kubectl gs](https://github.com/giantswarm/kubectl-gs) for easy login via SSO.

The app is installed in workload clusters, via our [app platform](https://docs.giantswarm.io/app-platform/).

Other than the app itself, you will need to provide a `values.yaml` configuration.

The management cluster name is needed as minimal configuration.

```yaml
managementCluster:
  name: test
```

It is also possible to override the api and issuer addresses as well as the cluster name and provider in case it is needed:
```yaml
managementCluster:
  name: test
clusterID: example
provider:
  kind: aws
kubernetes:
  api:
    address: https://api.test.example.io
oidc:
  issuerAddress: https://dex.test.example.io
```
Access to athena can be restricted to certain CIDRs.
```yaml
security:
  subnet:
    customer:
      public: x.x.x.x/x,x.x.x.x/x
      private: x.x.x.x/x
    restrictAccess:
      gsAPI: true
```

## Examples

Athena provides a GraphQL service. You can find example queries in the [examples folder](https://github.com/giantswarm/athena/blob/main/examples). You can execute these in the GraphQL playground app (at the `/` route).

## How to add a new property?

Adding a new query property is relatively simple. We can illustrate this by adding a new `party` property.

1. Create a new schema for your new property

#### **`pkg/graph/graphql/party.graphql`**

```graphql
type Party {
  name: String!
}
```

2. Extend the `Query` by adding your new property to it.

#### **`pkg/graph/graphql/party.graphql`**

```graphql
type Party {
  name: String!
}
+
+   extend type Query {
+     party: Party!
+   }
```

3. Run the code generator

```nohighlight
$ go generate ./...
```

4. Add your resolver implementation (what to return when that parameter is queried).

#### **`pkg/graph/resolver/party.resolvers.go`**

```go
func (r *queryResolver) Party(ctx context.Context) (*model.Party, error) {
-  	 panic(fmt.Errorf("not implemented"))
+    p := &model.Party{
+		Name: "something",
+	 }
+
+    return p, nil
}
```

5. See it in action

You can run the app locally, and execute a query for this in the GraphQL playground app (at the `/` route).
