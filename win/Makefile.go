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
	@$(GTARGET)
	@echo.
clean: 
	@if not exist $(GBIN) ( @mkdir $(GBIN) )
	del /q $(GBIN)\*
