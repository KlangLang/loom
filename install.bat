@echo off
setlocal enabledelayedexpansion

REM ------------------------------
REM Detect OS e ARCH
REM ------------------------------

for /f "tokens=2 delims==" %%a in ('wmic os get osarchitecture /value ^| find "="') do set ARCH=%%a

if "%ARCH%"=="64-bit" set ARCH=x86_64
if "%ARCH%"=="32-bit" set ARCH=i386
if "%ARCH%"=="ARM64-based PC" set ARCH=arm64

echo Detected -^> Windows-%ARCH%

REM ------------------------------
REM Get latest version from GitHub API
REM ------------------------------
set REPO=KlangLang/loom

if "%1"=="" (
    echo Fetching latest version from GitHub...
    
    REM Cria arquivo temporário para JSON
    set TMPJSON=%TEMP%\loom_releases.json
    
    curl -s "https://api.github.com/repos/%REPO%/releases/latest" -o "!TMPJSON!"
    
    REM Extrai tag_name do JSON usando PowerShell
    for /f "delims=" %%i in ('powershell -Command "(Get-Content '!TMPJSON!' | ConvertFrom-Json).tag_name"') do set VERSION=%%i
    
    del "!TMPJSON!"
    
    if "!VERSION!"=="" (
        echo ❌ Failed to fetch latest version, using fallback v0.1.6
        set VERSION=v0.1.6
    ) else (
        echo Found latest version: !VERSION!
    )
) else (
    set VERSION=%1
    echo Using specified version: !VERSION!
)

set FILE=loom_Windows_%ARCH%.tar.gz
set URL=https://github.com/%REPO%/releases/download/!VERSION!/%FILE%

echo Downloading: !URL!

REM ------------------------------
REM Create temp directory
REM ------------------------------
set TMPDIR=%TEMP%\loom_install_%RANDOM%
mkdir "%TMPDIR%"

REM ------------------------------
REM Download
REM ------------------------------
curl -L --fail "!URL!" -o "%TMPDIR%\%FILE%"

if not exist "%TMPDIR%\%FILE%" (
    echo ❌ Failed to download %FILE%
    rmdir /s /q "%TMPDIR%"
    exit /b 1
)

REM ------------------------------
REM Extract tar.gz
REM ------------------------------
echo Extracting...
tar -xzf "%TMPDIR%\%FILE%" -C "%TMPDIR%"

if errorlevel 1 (
    echo ❌ Failed to extract %FILE%
    rmdir /s /q "%TMPDIR%"
    exit /b 1
)

REM ------------------------------
REM Locate loom.exe
REM ------------------------------
set LOOMBIN=
for /r "%TMPDIR%" %%f in (loom.exe) do (
    set LOOMBIN=%%f
)

if "!LOOMBIN!"=="" (
    echo ❌ loom.exe not found inside archive.
    rmdir /s /q "%TMPDIR%"
    exit /b 1
)

echo Found binary: !LOOMBIN!

REM ------------------------------
REM Install to %USERPROFILE%\bin
REM ------------------------------
set TARGET=%USERPROFILE%\bin

if not exist "%TARGET%" (
    mkdir "%TARGET%"
)

echo Installing to %TARGET%\loom.exe
copy /y "!LOOMBIN!" "%TARGET%\loom.exe" >nul

if errorlevel 1 (
    echo ❌ Failed to copy loom.exe
    rmdir /s /q "%TMPDIR%"
    exit /b 1
)

REM ------------------------------
REM Cleanup
REM ------------------------------
rmdir /s /q "%TMPDIR%"

REM ------------------------------
REM Add to PATH
REM ------------------------------
echo.
echo Adding to PATH...

powershell -Command "$target = '%TARGET%'; $currentPath = [Environment]::GetEnvironmentVariable('Path', 'User'); if ($currentPath -notlike \"*$target*\") { [Environment]::SetEnvironmentVariable('Path', \"$currentPath;$target\", 'User'); Write-Host '✔ Added to PATH!' -ForegroundColor Green; Write-Host '' } else { Write-Host '✔ Already in PATH' -ForegroundColor Green; Write-Host '' }"

set PATH=%PATH%;%TARGET%

echo.
echo ✔ Loom !VERSION! installed!
echo → Run: loom --version

endlocal