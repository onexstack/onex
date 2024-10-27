---
title: EasyAI
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# EasyAI

Base URLs:

# Authentication

# Conversation

## POST NewChat

POST /api/v1/conversations

> Body 请求参数

```json
{
  "query": "how to speak english",
  "files": [
    "test.pdf",
    "song.mp3"
  ],
  "response_mode": "blocking",
  "app_name": "base-chat-english-teacher",
  "app_namespace": "easyai",
  "debug": false
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|X-Namespace|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

