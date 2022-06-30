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
	# Fetch the next problem in chronological order
	# Can deduce the next in line, then call `get-1234`
	echo "Fetching the next problem"

.PHONY: run-%
run-%:
	echo "Running $*"

.PHONY: test-%
test-%:
	echo "Testing $*"
