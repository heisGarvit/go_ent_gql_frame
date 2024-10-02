package modules

import (
	"project/src/modules/auth"
	"project/src/modules/healthcheck"
)

// MUST add to gqlgen.yml autobind as well

type Resolver struct {
	*healthcheck.Resolver
	*auth.AuthGqlResolver
}
