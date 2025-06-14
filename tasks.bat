@echo off
REM tasks.bat - Usage: tasks <command> -os <windows|linux> -env <local|dev|prod>

setlocal ENABLEDELAYEDEXPANSION

echo [TRACE] Starting tasks.bat with args: %*

REM Check arguments
if "%1"=="" (
    echo Usage: tasks ^<command^> -os ^<windows|linux^> -env ^<local|dev|prod^>
    exit /b 1
)

echo [TRACE] Initial command: %1
set CMD=%1
shift

echo [TRACE] Parsing flags...
REM Initialize flags
set PLATFORM=
set ENV=

REM Parse flags using a loop
:parse_loop
echo [TRACE] Arg: %1
if "%1"=="" goto after_args
if "%1"=="-os" (
    shift
    set PLATFORM=%2
    echo [TRACE] Set PLATFORM=%PLATFORM%
    shift
    goto parse_loop
)
if "%1"=="-env" (
    shift
    set ENV=%2
    echo [TRACE] Set ENV=%ENV%
    shift
    goto parse_loop
)
shift
goto parse_loop

:after_args
echo [TRACE] Args parsed: PLATFORM=%PLATFORM% ENV=%ENV%
if "%PLATFORM%"=="" (
    echo Error: Platform not specified. Use -os windows or -os linux.
    exit /b 1
)
if "%ENV%"=="" (
    echo Error: Environment not specified. Use -env local, dev, or prod.
    exit /b 1
)

echo [TRACE] Building script path...
set SCRIPTDIR=commands\%PLATFORM%
set SCRIPTNAME=%CMD%.%ENV%

echo [TRACE] Looking for script...
set SCRIPT=
if exist "%SCRIPTDIR%\%SCRIPTNAME%.bat" set SCRIPT="%SCRIPTDIR%\%SCRIPTNAME%.bat"
if not defined SCRIPT if exist "%SCRIPTDIR%\%SCRIPTNAME%.cmd" set SCRIPT="%SCRIPTDIR%\%SCRIPTNAME%.cmd"
if not defined SCRIPT if exist "%SCRIPTDIR%\%SCRIPTNAME%.sh" set SCRIPT="%SCRIPTDIR%\%SCRIPTNAME%.sh"
if not defined SCRIPT if exist "%SCRIPTDIR%\%SCRIPTNAME%" set SCRIPT="%SCRIPTDIR%\%SCRIPTNAME%"

echo [TRACE] SCRIPT=%SCRIPT%

if not defined SCRIPT (
    echo Error: Command script not found: %SCRIPTDIR%\%SCRIPTNAME%[.bat|.cmd|.sh]
    exit /b 1
)

REM Always use .\ prefix for Windows scripts
set WIN_SCRIPT=%SCRIPT:~1,-1%
set WIN_SCRIPT=.\%WIN_SCRIPT%

echo [TRACE] WIN_SCRIPT=%WIN_SCRIPT%

echo [TRACE] Executing script...
if /I "%PLATFORM%"=="windows" (
    call %WIN_SCRIPT% %*
    echo [TRACE] Script returned %ERRORLEVEL%
) else if /I "%PLATFORM%"=="linux" (
    echo To run: bash %SCRIPT% %*
    exit /b 0
) else (
    echo Error: Unsupported platform "%PLATFORM%".
    exit /b 1
)

echo [TRACE] Done.
pause
endlocal