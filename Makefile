DS_TARGETS= cp rm monitor stat
TM_TARGETS= ln delete
TARGETS= clone

all: $(DS_TARGETS) $(TM_TARGETS) $(TARGETS)

$(DS_TARGETS):
	cd datastore/$@; go build $@.go

$(TM_TARGETS):
	cd tm/$@; go build $@_*.go

$(TARGETS):
	cd tm/$@; go build $@_*.go;
	cd datastore/$@; go build $@.go
