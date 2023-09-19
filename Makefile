# Global variables.
# Notice both 'bin' and 'out' are git ignored.
BIN = ./bin
OUT = ./out

# Clean up for each program.
clean:
	rm -fr $(BIN)
	rm -fr $(OUT)