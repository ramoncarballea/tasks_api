@echo off
setlocal
REM Get the root directory (the parent of this script's parent)
set SCRIPT_DIR=%~dp0
cd /d %SCRIPT_DIR%\..\..

REM Step 1: Download dependencies
echo Fetching Go dependencies...
go mod download
if errorlevel 1 (
    echo Failed to fetch dependencies.
    exit /b 1
)

REM Step 2: Vendor dependencies
echo Vendoring Go dependencies...
go mod vendor
if errorlevel 1 (
    echo Vendoring failed.
    exit /b 1
)

REM Step 3: Build the app
echo Building the app...
go build -o app.exe cmd\main.go
if errorlevel 1 (
    echo Build failed.
    exit /b 1
)

REM Step 4: Run the app
echo Running the app...
start cmd /k app.exe
endlocal
