# Ai-HotSentence
## 自动生成AI网络热句和名言，sdk是调用的讯飞大模型
## 代码包含前后端，前端是用React，后端用的golang
## 代码比较臃肿，轻喷
## 如何在本地启动：
- 启动前需要去讯飞大模型官网获取免费试用的apikey,appsecret,appid，并且替换掉`image.go`和`spark_chat.go`两个文件中对应的apikey,appsecret,appid。
- 需要一个可以访问图片文件的nginx服务配置，并替换掉`task.go`中的 imagePre和reNamePre
- 需要一个启动好的Redis,并替换掉`redis.go`中的Addr,Password。
- 进入golang项目文件夹中执行`go mod tidy`自动获取并安装所有项目依赖，然后执行`go run start.go`启动后端。
- 进入react项目文件夹中执行`npm install`同样安装所有前端项目依赖，然后执行`npm run start`启动前端。
- 在网页打开`localhost:`+ 前端启动的端口就可以访问成功了。
