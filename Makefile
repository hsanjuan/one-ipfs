DS= cp clone stat rm monitor
TM= ln_remote delete_remote clone_remote
TARGETS_DS= $(addprefix remotes/datastore/ipfs/,$(DS))
TARGETS_TM= $(addprefix remotes/tm/ipfs/,$(TM))

build: $(TARGETS_DS) $(TARGETS_TM)

$(TARGETS_DS):
	go build -o $@ src/datastore/$(@F)/$(@F).go
$(TARGETS_TM):
	go build -o $@ src/tm/$(@F)/$(@F).go

.PHONY: clean
clean:
	rm -fv remotes/datastore/ipfs/* remotes/tm/ipfs/*_remote
