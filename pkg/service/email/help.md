
## 请求方法举例

```bash
curl -XPOST 127.0.0.1:8000/s/mailto \
-d '{
    "mailto": [
        "limingrui@trgroup.cn"
    ],
    "subject": "test subject",
    "body": "这是你的测试邮件了"
}'
```