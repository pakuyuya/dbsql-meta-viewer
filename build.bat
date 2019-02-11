@echo off

echo building client..
call .\client\build.bat

echo building server..
call .\server\build.bat

echo copy dists...
rd /Q /S .\dist

mkdir dist

xcopy .\server\dist .\dist /S /E
copy  .\client\dist\*.html .\dist\templates
xcopy .\client\dist\static .\dist\static /S /E
