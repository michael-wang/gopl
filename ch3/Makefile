DIRS := $(wildcard */.)

clean:
	for dir in $(DIRS); do \
		cd $$dir; make clean; cd ../; \
	done
