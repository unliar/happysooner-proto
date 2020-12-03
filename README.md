![happysooner-proto](https://socialify.git.ci/unliar/happysooner-proto/image?description=1&font=Rokkitt&forks=1&issues=1&language=1&owner=1&pattern=Floating%20Cogs&stargazers=1&theme=Dark)
![protoc](https://github.com/unliar/happysooner-proto/workflows/protoc/badge.svg)

## 错误码规则

错误码定义 累计 7 位 第一位表示主模块，第二位表示子模块 ，后四位自由发挥。

100\*\*\*\* 表示 账户/account 服务错误

200\*\*\*\* 表示 支付/pay 服务错误

300\*\*\*\* 表示 写作/writing 服务错误

400\*\*\*\* 表示 推送/push 服务错误

500\*\*\*\* 表示 计数/counter 服务错误

600\*\*\*\* 表示 文件/file 服务错误

700\*\*\*\* 表示 搜索/search 服务错误
