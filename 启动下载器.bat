@echo off
chcp 65001 >nul
echo Starting WX Channels Downloader...

:: Open browser after 3 seconds
start /b cmd /c "timeout /t 3 /nobreak >nul && start http://127.0.0.1:2022"

go run main.go
pause