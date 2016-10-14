DS= cp clone stat rm monitor
TM= ln_remote delete_remote clone_remote
TARGETS_DS= $(addprefix remotes/datastore/ipfs/,$(DS))
TARGETS_TM= $(addprefix remotes/tm/ipfs/,$(TM))
TM_SCRIPTS= $(addprefix remotes/tm/ipfs/,delete ln clone)
ifndef ONE_LOCATION
	ONE_LOCATION= /var/lib/one
endif

all: build

build: $(TARGETS_DS) $(TARGETS_TM)

check:
	golint helpers
	for module in $(DS); do golint src/datastore/$(module); done
	for module in $(TM); do golint src/tm/$(module); done

install: build
	install -m 0755 -d $(ONE_LOCATION)/remotes/datastore/ipfs
	install -m 0755 -d $(ONE_LOCATION)/remotes/tm/ipfs
	install -m 0755 -t $(ONE_LOCATION)/remotes/datastore/ipfs $(TARGETS_DS)
	install -m 0755 -t $(ONE_LOCATION)/remotes/tm/ipfs $(TARGETS_TM)
	install -m 0755 -t $(ONE_LOCATION)/remotes/tm/ipfs $(TM_SCRIPTS)

uninstall:
	rm -rf $(ONE_LOCATION)/remotes/datastore/ipfs
	rm -rf $(ONE_LOCATION)/remotes/tm/ipfs

$(TARGETS_DS):
	go build -o $@ src/datastore/$(@F)/$(@F).go
$(TARGETS_TM):
	go build -o $@ src/tm/$(@F)/$(@F).go

deps:
	go get -u github.com/hsanjuan/go-ipfs-api
	go get -u github.com/hsanjuan/one-ipfs/helpers
	go get -u github.com/golang/lint/golint

clean:
	rm -fv remotes/datastore/ipfs/* remotes/tm/ipfs/*_remote
