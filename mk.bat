@ECHO OFF

SETLOCAL ENABLEEXTENSIONS

SET build_dir=./build
SET exe_fname=life.exe

IF NOT EXIST %build_dir% (
	MKDIR "%build_dir%"
)

FOR /f "tokens=*" %%a IN (
	'git rev-list --abbrev-commit -1 HEAD'
) DO (
	SET version=%%a
)

FOR %%a IN ("./src", "./src/world", "./src/config", "./src/config/theme") DO (
	go test %%a
	
	IF ERRORLEVEL 1 (
        ECHO.
		ECHO Tests failed. Aborting build.
        ECHO.

        EXIT /B 1
	)
)

go build -trimpath -ldflags "-s -w -X main.Version=%version%" -o "%build_dir%/%exe_fname%" ./src

ENDLOCAL
