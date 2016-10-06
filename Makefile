DS_TARGETS= cp rm clone monitor stat

all: $(DS_TARGETS) $(TM_TARGETS)

$(DS_TARGETS):
	cd datastore/$@; go build $@.go

$(TM_TARGETS):
