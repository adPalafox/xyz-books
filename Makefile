CWDIR := $(dir $(lastword $(MAKEFILE_LIST)))

include $(CWDIR)/common.conf

COMMON_ENV := \
	PORT=${PORT} \

run-local:
	@env $(COMMON_ENV) go run cmd/server/main.go
