check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check install
	GO111MODULE=off swagger generate spec -0 ./swagger.yaml --scan-models