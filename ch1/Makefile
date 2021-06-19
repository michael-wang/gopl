DIRS := $(wildcard */.)

clean:
	for dir in $(DIRS); do \
		cd $$dir; go clean; cd ../; \
	done
