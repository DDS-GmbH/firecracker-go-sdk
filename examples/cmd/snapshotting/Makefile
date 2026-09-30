# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

RELEASE_URL=https://github.com/firecracker-microvm/firecracker/releases
VER=v1.17.0

ARCH=$(shell uname -m)
GID = $(shell id -g)

PWD=$(shell pwd)
GOBIN=$(PWD)/bin
FC_TEST_DATA_PATH?=$(PWD)

# --location is needed to follow redirects on github.com
curl = curl --location

GO_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1,2)
define install_go
	GOBIN=$(GOBIN) go install $(1)@$(2)
endef

all: snapshot-example plugins root-drive vmlinux firecracker

plugins: bin/tc-redirect-tap bin/ptp bin/host-local | bin

bin:
	mkdir -p bin

bin/tc-redirect-tap: bin
	$(call install_go,github.com/awslabs/tc-redirect-tap/cmd/tc-redirect-tap,v0.0.0-20250516183331-34bf829e9a5c)

bin/ptp: bin
	$(call install_go,github.com/containernetworking/plugins/plugins/main/ptp,v1.9.1)

bin/host-local: bin
	$(call install_go,github.com/containernetworking/plugins/plugins/ipam/host-local,v1.9.1)

vmlinux:
	$(curl) -o vmlinux https://s3.amazonaws.com/spec.ccfc.min/img/quickstart_guide/${ARCH}/kernels/vmlinux.bin

firecracker:
	$(curl) ${RELEASE_URL}/download/${VER}/firecracker-${VER}-${ARCH}.tgz | tar -xz
	mv release-${VER}-${ARCH}/firecracker-${VER}-${ARCH} firecracker
	rm -rf release-${VER}-${ARCH}

# Use rootfs from S3 following firecracker getting-started guide
# https://github.com/firecracker-microvm/firecracker/blob/main/docs/getting-started.md
root-drive:
	$(curl) -o root-drive.img https://s3.amazonaws.com/spec.ccfc.min/firecracker-ci/20260916-dcfc69b625d0-0/$(arch)/ubuntu-24.04.squashfs
	
	
snapshot-example:
	go build -o $@

run: snapshot-example
	./snapshot-example

clean:
	rm -rf bin firecracker root-drive.img snapshot-example vmlinux  

.PHONY: all clean plugins run
