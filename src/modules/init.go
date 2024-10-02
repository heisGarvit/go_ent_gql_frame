package modules

import (
	"project/src/modules/auth"
	"project/src/modules/healthcheck"
)

type Resolver struct {
	*auth.AuthGqlResolver
	*healthcheck.Resolver
}

func init() {
	(Resolver{}).Ping(nil)
}
