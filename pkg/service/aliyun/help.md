
## 请求方式举例

```bash
# 这里传内网 ip，会提示参数传入不对
curl -XPOST 127.0.0.1:8000/s/addsgrule \
-H "content-type: application/json" \
-d '{
    "ip": "192.168.0.100",
    "sgid": "sg-2zeb1ux0h4683ehrocq0",
    "remark": "lxm"
}'
```