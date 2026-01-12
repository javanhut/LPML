.PHONY: install uninstall clean build

# Detect OS
UNAME_S := $(shell uname -s 2>/dev/null || echo Windows)
ifeq ($(UNAME_S),Linux)
    OS := linux
    INSTALL_DIR := /usr/local/bin
    BINARY := lpml
    SUDO := $(shell if [ -w $(INSTALL_DIR) ]; then echo ""; else echo "sudo"; fi)
else ifeq ($(UNAME_S),Darwin)
    OS := macos
    INSTALL_DIR := /usr/local/bin
    BINARY := lpml
    SUDO := $(shell if [ -w $(INSTALL_DIR) ]; then echo ""; else echo "sudo"; fi)
else
    OS := windows
    INSTALL_DIR := $(USERPROFILE)/bin
    BINARY := lpml.exe
    SUDO :=
endif

build:
	@echo "Building LPML binary for $(OS)..."
	go build -o $(BINARY)

install: build
	@echo "INSTALLING LAZY PAGE MAKER LANGUAGE!"
ifeq ($(OS),windows)
	@echo "Installing to $(INSTALL_DIR)..."
	@if not exist "$(INSTALL_DIR)" mkdir "$(INSTALL_DIR)"
	@copy /Y $(BINARY) "$(INSTALL_DIR)\$(BINARY)"
	@echo ""
	@echo "NOTE: Make sure $(INSTALL_DIR) is in your PATH"
	@echo "You can add it via: setx PATH \"%PATH%;$(INSTALL_DIR)\""
else
	@echo "Installing to $(INSTALL_DIR)..."
	@if [ ! -d "$(INSTALL_DIR)" ]; then $(SUDO) mkdir -p $(INSTALL_DIR); fi
	$(SUDO) mv ./$(BINARY) $(INSTALL_DIR)/$(BINARY)
endif
	@echo ""
	@echo "Installation complete! Try compiling a .lpml file using:"
	@echo "  lpml testfile.lpml"

uninstall:
ifeq ($(OS),windows)
	@if exist "$(INSTALL_DIR)\$(BINARY)" ( \
		echo "LPML binary exists, removing..." && \
		del "$(INSTALL_DIR)\$(BINARY)" \
	) else ( \
		echo "LPML binary doesn't exist. Skipping..." \
	)
else
	@if [ -f "$(INSTALL_DIR)/$(BINARY)" ]; then \
		echo "LPML binary exists, removing..."; \
		$(SUDO) rm -f $(INSTALL_DIR)/$(BINARY); \
	else \
		echo "LPML binary doesn't exist. Skipping..."; \
	fi
endif

clean:
	rm -f ./lpml ./lpml.exe
