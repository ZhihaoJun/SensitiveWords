# sensitive words filter
敏感词检测

## api
`GET /api/check`

* s 字符串

```
GET /api/check?s=text
```

返回

``` javascript
{
    "error": "ok",
    "msg":   "check sensitive success",
    "sensitives": [] // 一个字符串数据表示哪些词是敏感的
}
```
