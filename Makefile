CWDIR := $(dir $(lastword $(MAKEFILE_LIST)))

include $(CWDIR)/common.conf
-include $(CWDIR)/.env

COMMON_ENV := \
	PORT=${PORT} \

run-local:
	@env $(COMMON_ENV) go run cmd/server/main.go

run-migrate:
	@env $(COMMON_ENV) go run cmd/db/main.go
