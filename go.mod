module github.com/omniboost/go-sage-intacct

go 1.25.1

require (
	github.com/cydev/zero v0.0.0-20160322155811-4a4535dd56e7
	github.com/elliotchance/pie/v2 v2.9.1
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/schema v1.2.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/pkg/errors v0.9.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	golang.org/x/exp v0.0.0-20220321173239-a90fa8a75705 // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
