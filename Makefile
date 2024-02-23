# Makefile

# Directories
LIB_DIR = ./lib
BUILD_DIR = ./build
GO_STATIC_BIN = $(BUILD_DIR)/cgo_static
GO_DYNAMIC_BIN = $(BUILD_DIR)/cgo_dynamic
#LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:./lib

.PHONY: all static dynamic clean

all: static dynamic

# Build static binary
static:
	CGO_ENABLED=1 go build -ldflags '-extldflags "-static"' -o $(GO_STATIC_BIN) .

# Build dynamic binary
dynamic: $(C_SHARED_LIB)
	CGO_ENABLED=1 \
	CGO_LDFLAGS="-L$(LIB_DIR) -ladd -Wl,-rpath,$(LIB_DIR)" \
	go build -o $(GO_DYNAMIC_BIN) .
	@echo "Remember to set LD_LIBRARY_PATH to include $(LIB_DIR) when running the dynamic binary."

clean:
	rm -f $(GO_STATIC_BIN) $(GO_DYNAMIC_BIN) $(LIB_DIR)/addlib.o $(LIB_DIR)/libadd.a $(LIB_DIR)/libadd.so

gcc:
	gcc -shared -o $(LIB_DIR)/libadd.so -fPIC $(LIB_DIR)/addlib.c # build shared library
	gcc -c -o  $(LIB_DIR)/addlib.o  $(LIB_DIR)/addlib.c && ar rcs $(LIB_DIR)/libadd.a $(LIB_DIR)/addlib.o # build static library

ldd:
	ldd $(GO_DYNAMIC_BIN)
	ldd $(GO_STATIC_BIN)
