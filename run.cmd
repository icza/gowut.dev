@echo off
call setgopath.cmd

:cycle

go run src/code.google.com/p/gowut/examples/showcase.go

rem ERRORLEVEL tests greater than or equal condition (not equal)

if ERRORLEVEL 2 goto goerror
if ERRORLEVEL 1 goto restart

goto exit



:restart
echo .
echo .
echo .
echo Restarting...
goto cycle

:goerror
echo .
echo .
pause >nul
goto restart

:exit
echo Exiting... Good bye!
