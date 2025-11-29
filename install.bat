@echo off
setlocal enabledelayedexpansion

REM ------------------------------
REM Detect OS e ARCH
REM ------------------------------

for /f "tokens=2 delims==" %%a in ('wmic os get osarchitecture /value ^| find "="') do set ARCH=%%a

REM Normalizar arquitetura
if "%ARCH%"=="64-bit" set ARCH=x86_64
if "%ARCH%"=="32-bit" set ARCH=i386
if "%ARCH%"=="ARM64-based PC" set ARCH=arm64

echo Detected Architecture: %ARCH%

REM ------------------------------
REM Define version
REM ------------------------------
set VERSION=%1
if "%VERSION%"=="" set VERSION=v0.1.5

set REPO=KlangLang/loom
set FILE=loom_Windows_%ARCH%.zip
set URL=https://github.com/%REPO%/releases/download/%VERSION%/%FILE%

echo Downloading Loom %VERSION%
echo URL: %URL%
curl -L "%URL%" -o "%FILE%"

if not exist "%FILE%" (
    echo ❌ Failed to download %FILE%
    exit /b 1
)

REM ------------------------------
REM Extract ZIP
REM Requires Windows 10+ (tar builtin)
REM ------------------------------
echo Extracting...
tar -xf "%FILE%"

REM ------------------------------
REM Locate loom.exe
REM (procura dentro da pasta extraída)
REM ------------------------------

set LOOMBIN=
for /r %%f in (loom.exe) do (
    set LOOMBIN=%%f
)

if "%LOOMBIN%"=="" (
    echo ❌ loom.exe not found inside archive.
    exit /b 1
)

echo Found binary: %LOOMBIN%

REM ------------------------------
REM Install bin
REM ------------------------------

set TARGET=%USERPROFILE%\bin

if not exist "%TARGET%" (
    mkdir "%TARGET%"
)

echo Installing to %TARGET%\loom.exe
copy /y "%LOOMBIN%" "%TARGET%\loom.exe" >nul

REM ------------------------------
REM Check PATH
REM ------------------------------
echo.
echo Checking PATH...

echo %PATH% | find /i "%TARGET%" >nul
if errorlevel 1 (
    echo ⚠ "%TARGET%" is NOT in PATH.
    echo Add this to your PATH manually:
    echo.
    echo     setx PATH "%%PATH%%;%TARGET%%"
    echo.
) else (
    echo ✔ PATH OK
)

echo.
echo ✔ Loom installed!
echo Run: loom --version

endlocal
