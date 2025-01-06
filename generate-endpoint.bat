@echo off
set /p endpoint="Enter endpoint name: "
go run cmd/generator/generate.go %endpoint%
pause
