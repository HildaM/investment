# Investment AI Agent

一个基于 AI 的智能投资Agent，专注于 A 股和美股的实时市场分析。系统通过爬取多个数据源，结合 ReAct模式反向调用爬虫进行深度分析，生成专业的市场分析报告。

### 效果预览
<img align="center" src="https://github.com/8treenet/investment/blob/master/img/1.png?raw=true">
<img align="center" src="https://github.com/8treenet/investment/blob/master/img/2.png?raw=true">

## 功能特点

### 1. 美股市场分析 (Yahoo Finance)
- 实时获取雅虎财经新闻数据
- 多维度市场分析：
  - 资金流向与市场情绪
  - 美股财报季追踪
  - 宏观经济与货币政策
  - 行业与政策催化
  - 全球扰动与地缘政治
  - 大宗商品与外汇联动
  - 科技创新与市场趋势
  - 系统性风险监测

### 2. A股市场分析 (财联社)
- 实时数据获取：
  - 深度文章分析
  - 实时电报追踪
  - 市场行情监控
- 专业分析维度：
  - 资金流向分析
  - 重要事件追踪
  - 宏观经济数据解读
  - 政策信号解析
  - 全球市场联动
  - 高频行业数据
  - 风险提示
  - 地缘政治分析

### 3. ETF 投资分析
- 智能 ETF 推荐
- 投资信号分析
- 风险评分系统
- 投资逻辑验证

## 技术架构
### 主要依赖
```go
require (
    github.com/cloudwego/eino v0.3.19
    github.com/8treenet/freedom v1.9.6
    github.com/go-rod/rod v0.116.2
    gorm.io/gorm v1.25.12
    // ... 其他依赖
)
```


## 直接使用
```bash
# 克隆项目
git clone https://github.com/8treenet/investment
cd investment

# 安装依赖
go mod tidy

# 配置文件，查看配置说明
vi server/conf/config.toml

# 分析美股(需开启梯子)和A股后发送邮件。默认使用qwen模型，代码可改，就不配置了。
go run main.go
```

### 配置说明
```env
# AI 模型配置

#硅基流动key,推荐使用
sf_api_key = "your_openai_api_key"
#腾讯云key, 腾讯云的DeepSeek 不支持 tools。
tt_api_key = "your_openai_api_key"
#qwen key,推荐使用
qwen_api_key = "your_openai_api_key"
#gemini_api_key
gemini_api_key = "your_openai_api_key"


# 邮件通知配置
#接收人
to_mail_list = ["112233@gmail.com","445566@qq.com"]
#邮件发送者
from_mail = "44123@qq.net"
#邮件发送者密码
from_mail_password = "123123"
#邮件服务器地址
mail_server = "smtp.qq.com"
#邮件服务器端口
mail_server_port = 465
```

## 二次开发

### 项目结构
```
.
├── adapter/       # 适配器层
│   ├── ai/        # AI 相关实现
│   └── spider/    # 爬虫实现
├── domain/        # 领域模型
├── infra/         # 基础设施
├── server/        # 服务器
├── utility/       # 工具函数
├── main.go        # 主程序
└── go.mod         # 依赖管理
```

### 1. 美股分析
```go
// 创建雅虎分析师实例
agent, err := NewYahooAnalyst(ctx, roles, inputMsg)
if err != nil {
    panic(err)
}

// 运行分析
output, err := agent.Run(ctx)
if err != nil {
    panic(err)
}

// 发送分析报告
htmldata, _ := utility.ConvertHtml(output.Content)
utility.SendHtmlMail("美股分析报告", htmldata)
```

### 2. A股分析
```go
// 创建财联社分析师实例
agent, err := NewClsAnalyst(ctx, roles, inputMsg)
if err != nil {
    panic(err)
}

// 运行分析
output, err := agent.Run(ctx)
if err != nil {
    panic(err)
}

// 发送分析报告
htmldata, _ := utility.ConvertHtml(output.Content)
utility.SendHtmlMail("A股分析报告", htmldata)
```

## 注意事项

1. API 密钥安全
   - 不要将 API 密钥提交到代码仓库
   - 使用环境变量或配置文件管理密钥
   - 定期轮换密钥

2. 数据使用
   - 遵守数据源的使用条款
   - 控制爬虫频率
   - 注意数据存储安全

3. 系统维护
   - 定期更新依赖
   - 监控系统资源使用
   - 备份重要数据


