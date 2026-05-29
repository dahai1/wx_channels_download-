@echo off
chcp 65001 >nul
title WX Video Downloader

:: Open browser after 2 seconds
start /b cmd /c "timeout /t 2 /nobreak >nul && start http://127.0.0.1:2022"

:: Start service
"%~dp0WX_Video_Downloader.exe" server

pause
