@ECHO OFF
SET GOPATH=%GOPATH%;%CD%
go env -w GO111MODULE=auto
go env -w GOBIN=%CD%\bin
go env -w GOPATH=%GOPATH%
SET PATH=%PATH%;%CD%\bin
SET PROJ_PATH=%CD%
code .