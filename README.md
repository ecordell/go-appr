# go-appr
This package is a collection of `openapi-generator` generated client bindings for `App Registry` which can be used to talk to `quay.io` or [appr](https://github.com/app-registry/appr) to pull down application packages.

Swagger spec file has been obtained from here -  [appr-api-swagger.yaml](https://github.com/app-registry/appr/blob/master/Documentation/server/appr-api-swagger.yaml)

### Example
Below is an example of how you can use the client bindings to connect to an `app registry` server and invoke its APIs. This example assumes that an instance of `appr` is running on your machine (`localhost:5000`). For more information on how to run `appr` server on local environment visit https://github.com/app-registry/appr

```go
package main

import (
	"context"
	"log"

	"github.com/antihax/optional"

	apprclient "github.com/ecordell/go-appr/client"
)

func main() {
	client := apprclient.NewAPIClient(&apprclient.Configuration{
		BasePath:      "https://quay.io/cnr",
		DefaultHeader: make(map[string]string),
		UserAgent:     "go-appr/1.0.0",
		Scheme:        "https",
	})

	packages, err := listPackages(client)
	if err != nil {
		log.Fatalf("error - %v", err)
	}

	log.Printf("success - found [%d] package(s)\n", len(packages))
}

func listPackages(client *apprclient.APIClient) ([]apprclient.Package, error) {
	packages, _, err := client.PackageApi.ListPackages(context.TODO(),
		&apprclient.ListPackagesOpts{Namespace: optional.NewString("community-operators")})
	if err != nil {
		return nil, err
	}

	return packages, nil
}
```
