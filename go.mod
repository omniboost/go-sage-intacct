module github.com/omniboost/go-sage-intacct

go 1.15

require (
	github.com/Azure/go-ntlmssp v0.0.0-20191115210519-2b2be6cc8ed4
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/schema v1.2.0
	github.com/omniboost/go-sageone-za v0.0.0-20200615105129-947664bd04a9
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gopkg.in/guregu/null.v3 v3.5.0
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
