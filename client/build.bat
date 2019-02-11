rem @echo off

setlocal

npm install --no-audit --prefix ./
npm run build

echo build client is successful

setlocal
