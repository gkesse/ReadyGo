GSRC = ..\code\GProject\src
GBIN = bin
GBUILD = build
GTARGET = $(GBIN)\gp_go.exe

all: clean compile run
test: clean compile run_test

compile:
	@if not exist $(GBIN) ( @mkdir $(GBIN) )
	go build -o $(GTARGET) $(GSRC)
run: 
	@echo.
	@$(GTARGET)
	@echo.
run_test: 
	@echo.
	@$(GTARGET) test
	@echo.
clean: 
	@if not exist $(GBIN) ( @mkdir $(GBIN) )
	del /q $(GBIN)\*
sqlite: 
	go get github.com/mattn/go-sqlite3
qt: 
	go get github.com/therecipe/qt
qt_setup: 
	go get -u -v github.com/therecipe/qt/cmd/qtsetup
