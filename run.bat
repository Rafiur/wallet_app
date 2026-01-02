@echo off
echo Starting Wallet App...

REM Start backend
start cmd /k "cd /d D:\Projects\wallet_app\backend && wallet_app.exe"

REM Start frontend
start cmd /k "cd /d D:\Projects\wallet_app\frontend && npm run dev"

echo Both frontend and backend are starting...
pause