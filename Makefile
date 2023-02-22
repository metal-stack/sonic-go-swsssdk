.PHONY: all
all:
	go test -v ./...

.PHONY: generate
generate:
	./generator \
		-output_file=api/sonic.go \
		-package_name=api \
		yang/sonic-interface.yang \
		yang/sonic-vrf.yang \
		yang/sonic-vlan.yang 