
# 生成api
buildApi :
	goctl api go -api $(apis) -dir $(dir) --style=goZero --home=$(template)