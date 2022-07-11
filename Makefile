SHELL := /bin/bash -e

MAKEFILE_PATH := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

.PHONY: get-%
get-%:
	# 1) Check if dir already exists
	# 2) Scrape the webpage for the PE problem
	# 2a) Error out if doesn't exist
	# 3) Parse title
	# 4) Create new directory for problem
	# 5) Copy template files to new dir, and update
	#    the stub vars in the files
	echo "Fetching $*"

.PHONY: get-next
get-next:
	# 1) Get top level dir names in root folder
	# 2) Match on `####-*$` and ignore non-matches
	# 3) Split at `-` and keep numbers
	# 4) Sort numbers and check sequence until a gap is found
	# 5) Zero pad if necessary
	# 6) Call `get-####` with the found number
	echo "Fetching the next problem"

# DIRECTORIES := $(patsubst %-,%,$(notdir $(wildcard $(MAKEFILE_PATH)/*)))
DIRECTORIES := $(filter %-*,$(notdir $(wildcard $(MAKEFILE_PATH)/*)))
PROBLEMS := $(sort $(foreach dir,$(DIRECTORIES),$(firstword $(subst -, ,$(dir)))))

.PHONY: foo
foo:
	echo $(PROBLEMS)

.PHONY: bar
bar:
	@echo $(subst -, ,1234-abc123)
	@echo "************************************"
	@echo $(sort $(foreach dir,1234-abc123 5678-xyz567,$(firstword $(subst -, ,$(dir)))))

.PHONY: baz
baz:
	@echo $(call numonly,$(subst -,$(COMMA),1234-abc123))
	@echo "************************************"
	@echo $(filter %-%,1234-abc123 5678-xyz567)

letters:=a b c
numbers:=1 2 3 4

define GEN_RULE
$(letter).dat.$(number) : $(letter).rlt.$(number)
    ./rlt2dat $$< $$@
endef

spam:
	$(foreach number,$(numbers), $(foreach letter,$(letters), $(eval $(GEN_RULE)) ) )

.PHONY: run-%
run-%:
	echo "Running $*"

.PHONY: test-%
test-%:
	echo "Testing $*"
