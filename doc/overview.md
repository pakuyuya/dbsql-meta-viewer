# 構成メモ

## モジュールの種類

-   client

    -   SPA
    -   内容をビルドして、serverのtemplatesに配置する

-   server
    -   Webサーバー
    -   静的ページの提供とWebAPIを提供

## モジュールの関連性

    client -[api call]-> server
      '`-[respond as html]-'

## APIリスト

| path               | Method | Description                   |
| ------------------ | ------ | ----------------------------- |
| /api/server/status | GET    | Respose server status AS JSON |
| /api/textdata/:id  | GET    | get textdata. signiture       |
| /api/searchtext    | GET    | search textdata.              |
