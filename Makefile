SHELL := /bin/bash
arklibgo := ~/ProjectsGo/arkAlias.sh
version = ~/ProjectsGo/arkAlias.sh getlastversion
PROJECTNAME=$(shell basename `pwd`)
.PHONY: check

.SILENT: build run getlasttag buildzip buildwin buildlinux buildwasm www buildandroid

wasm:
# скопировать из каталога GO куда надо ,   cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
	@echo $$($(version))
	$(info +Компиляция Wasm)
	GOOS=js GOARCH=wasm go build -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/wasm/html/$(PROJECTNAME).wasm cmd/wasm/main.go

linux:
	@echo $$($(version))
	$(info +Компиляция Linux)
	go build -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/main/$(PROJECTNAME) cmd/main/main.go
	./bin/main/$(PROJECTNAME)


buildlinux:
	@echo $$($(version))
	$(info +Компиляция Linux)
	go build -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/main/$(PROJECTNAME) cmd/main/main.go
buildzip:
	$(info +Компиляция с сжатием)
	go build -ldflags "-s -w" -o ./bin/main/$(PROJECTNAME) cmd/main/main.go
buildwin:
	$(info +Компиляция windows)
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o ./bin/main/$(PROJECTNAME).exe -tags static -ldflags "-s -w -X 'main.versionProg=$$($(version))'" cmd/main/main.go

gio:
	$(info +Компиляция WASM)
	PROJECTNAME="xxxx" /home/arkadii/go/bin/gogio -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/main/html/gio -target js cmd/gio/main.go 
	cp ./bin/main/html/gio/main.wasm ./bin/main/html/main.wasm 
	cp ./bin/main/html/gio/wasm.js ./bin/main/html/wasm.js 

buildandroid:
	$(info +Компиляция Android)
	ANDROID_SDK_ROOT=/home/arkadii/Android/Sdk/ go run gioui.org/cmd/gogio -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/main/$(PROJECTNAME).apk -target android -icon appicon.png -arch arm64 -appid Go.arkiv cmd/main/main.go

run: buildlinux buildwin 
	$(info +Запуск)
	./bin/main/$(PROJECTNAME)

build: buildlinux buildwin #buildwasm buildandroid
	$(info +Сборка)

#www: build
#	$(info +Старт сервера http://172.16.172.10:8080)
#	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("wasm")))'

