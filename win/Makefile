GSRC = ..\code\GProject\src
GBIN = bin
GBUILD = build
GTARGET = $(GBIN)\gp_go.exe

all: clean compile run

compile:
	@if not exist $(GBIN) ( @mkdir $(GBIN) )
	go build -o $(GTARGET) $(GSRC)
run: 
	@echo.
	@$(GTARGET) $(args)
	@echo.
clean: 
	@if not exist $(GBIN) ( @mkdir $(GBIN) )
	del /q $(GBIN)\*
sqlite: 
	go get github.com/mattn/go-sqlite3
qt: 
	go get github.com/therecipe/qt
qt_module: 
	go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...
qt_setup: 
	$(GOPATH)\bin\qtsetup
qt_env: 
	$(GOPATH)\bin\qtenv
