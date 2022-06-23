

goctl rpc new {name} --home ..\..\deploy\goctl\temp

goctl api go -api {name}.api -dir . -style goZero --home ..\..\deploy\goctl\temp

# model

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:33306)/test" -table="*"  -dir="./model" -c