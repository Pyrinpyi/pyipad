@echo off

CALL TASKKILL /F /IM pyrin.exe
CALL TASKKILL /F /IM pyrinminer.exe
CALL TASKKILL /F /IM pyrinwallet.exe

CALL rmdir /q /s %localappdata%\Pyipad
CALL rmdir /q /s %localappdata%\Pyrinminer
CALL rmdir /q /s %localappdata%\Pyrinwallet

@REM CALL pyipad.exe /utxoindex /nodnsseed /testnet /a 192.168.1.169
CALL pyipad.exe /utxoindex /nodnsseed /a 192.168.1.169