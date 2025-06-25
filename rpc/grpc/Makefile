# ==============================================================================

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
PROJ_ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# hello worldProtobuf 文件存放路径
HELLOWORLD_ROOT=$(PROJ_ROOT_DIR)/helloworld/pb

.PHONY: test
test:
	@echo "MAKEFILE_LIST = $(MAKEFILE_LIST)"
	@echo "COMMON_SELF_DIR = $(COMMON_SELF_DIR)"
	@echo "PROJ_ROOT_DIR = $(PROJ_ROOT_DIR)"
	@echo "HELLOWORLD_ROOT = $(HELLOWORLD_ROOT)"

.PHONY: helloworld
helloworld: # 编译 protobuf 文件.
	@echo "===========> Generate protobuf files"
	@protoc                                              \
		--proto_path=$(HELLOWORLD_ROOT)                          \
		--go_out=paths=source_relative:$(HELLOWORLD_ROOT)        \
		--go-grpc_out=paths=source_relative:$(HELLOWORLD_ROOT)   \
		$(shell find $(HELLOWORLD_ROOT) -name *.proto)