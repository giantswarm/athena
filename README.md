[![CircleCI](https://circleci.com/gh/giantswarm/athena.svg?style=shield&circle-token=e86584296950c3be820856e802a37336c2c7d540)](https://circleci.com/gh/giantswarm/athena)
[![Docker Repository on Quay](https://quay.io/repository/giantswarm/athena/status?token=789fdfa0-2068-473d-9ab6-7bde4aaf46dc "Docker Repository on Quay")](https://quay.io/repository/giantswarm/athena)

# athena

athena is a service that knows everything about your installation. Its purpose is to provide metadata to clients, so they could easily establish a connection with the Kubernetes API, and identify the installation that they're talking to.

## Examples

You can find example queries in the [examples folder](./examples). You can execute these in the GraphQL playground app (at the `/` route).

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
