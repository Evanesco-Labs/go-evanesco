# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: eva android ios eva-cross evm all test clean
.PHONY: eva-linux eva-linux-386 eva-linux-amd64 eva-linux-mips64 eva-linux-mips64le
.PHONY: eva-linux-arm eva-linux-arm-5 eva-linux-arm-6 eva-linux-arm-7 eva-linux-arm64
.PHONY: eva-darwin eva-darwin-386 eva-darwin-amd64
.PHONY: eva-windows eva-windows-386 eva-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

eva:
	$(GORUN) build/ci.go install ./cmd/eva
	@echo "Done building."
	@echo "Run \"$(GOBIN)/eva\" to launch eva."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/eva.aar\" to use the library."
	@echo "Import \"$(GOBIN)/eva-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/kevinburke/go-bindata/go-bindata@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

eva-cross: eva-linux eva-darwin eva-windows eva-android eva-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/eva-*

eva-linux: eva-linux-386 eva-linux-amd64 eva-linux-arm eva-linux-mips64 eva-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-*

eva-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/eva
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep 386

eva-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/eva
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep amd64

eva-linux-arm: eva-linux-arm-5 eva-linux-arm-6 eva-linux-arm-7 eva-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep arm

eva-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/eva
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep arm-5

eva-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/eva
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep arm-6

eva-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/eva
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep arm-7

eva-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/eva
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep arm64

eva-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/eva
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep mips

eva-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/eva
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep mipsle

eva-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/eva
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep mips64

eva-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/eva
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/eva-linux-* | grep mips64le

eva-darwin: eva-darwin-386 eva-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/eva-darwin-*

eva-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/eva
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/eva-darwin-* | grep 386

eva-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/eva
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/eva-darwin-* | grep amd64

eva-windows: eva-windows-386 eva-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/eva-windows-*

eva-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/eva
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/eva-windows-* | grep 386

eva-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/eva
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/eva-windows-* | grep amd64
