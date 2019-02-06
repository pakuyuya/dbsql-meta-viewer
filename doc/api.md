
## GET /api/searver/status

### Query parameter

none

### returns
```
{
    "status" : "ok"
}
```

## GET /api/textdata/:id


### Query parameter

none

### returns

```

```

## GET /api/searchtext

### Query parameter

<dl>
  <dt>query</dt>
  <dd>required. Query string encoded UTF-8 for search.</dd>
  <dt>wordOnly</dt>
  <dd>true | false (default: false)</dd>
  <dt>ignoreCase</dt>
  <dd>true | false (default: false)</dd>
  <dt>regex</dt>
  <dd>true | false (default: false)</dd>
  <dt>bodyOnly</dt>
  <dd>true | false (default: false)</dd>
  <dt>readBody</dt>
  <dd>true | false (default: false)</dd>
</dl>

### returns

```
{
    "count": 3,
    "datas": [
        { idx: 0, caption: "1st entry", body: "body", since: "2019-01-03T14:00:32" },
        { idx: 1, caption: "2nd entry", body: "body", since: "2019-01-03T14:00:32" },
        { idx: 2, caption: "3rd entry", body: "body", since: "2019-01-03T14:00:32" }
    ]
}
```
