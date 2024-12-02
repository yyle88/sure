COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/sortslice/blob/5f56c911501ffcef244e46d7c9f96b2ca60e5b16/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
