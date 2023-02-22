.PHONY: all
all:
	go test -v ./...

# generator can be installed from
# https://github.com/openconfig/ygot
.PHONY: generate
generate:
	rm -rf api
	mkdir api
	./generator \
		-output_file=api/sonic.go \
		-package_name=api \
		-generate_fakeroot \
		-annotations=false \
		-compress_paths=false \
		-exclude_state=true \
		-exclude_modules=ietf-interfaces \
		-include_schema=true \
		-fakeroot_name=device \
		-shorten_enum_leaf_names=true \
		-typedef_enum_with_defmod=true \
		yang/*.yang

		# yang/sonic-breakout_cfg.yang \
		# yang/sonic-interface.yang \
		# yang/sonic-port.yang \
		# yang/sonic-portchannel.yang \
		# yang/sonic-types.yang \
		# yang/sonic-vlan.yang \
		# yang/sonic-vrf.yang