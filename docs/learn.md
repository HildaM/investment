# Investment AI 项目架构分析

## 项目概述

Investment AI 是一个基于 Go 语言开发的智能投资分析系统，集成了爬虫技术、AI 代理和数据分析功能。系统主要用于分析美股和A股市场，通过爬取财联社、雅虎财经等数据源，结合多个AI模型进行智能分析，并生成投资研报。

## 系统架构概览

项目采用 DDD（领域驱动设计）架构模式，基于 Freedom 框架构建，具有清晰的分层结构：

```mermaid
graph TB
    subgraph "Presentation Layer"
        A[Controller Layer]
        B[HTTP API]
    end
    
    subgraph "Application Layer"
        C[Domain Services]
        D[AI Agents]
        E[Spider Services]
    end
    
    subgraph "Infrastructure Layer"
        F[Repository]
        G[Database]
        H[External APIs]
        I[Email Service]
    end
    
    subgraph "Domain Layer"
        J[Domain Models]
        K[Value Objects]
        L[Domain Events]
    end
    
    A --> C
    C --> D
    C --> E
    C --> F
    F --> G
    D --> H
    E --> H
    C --> I
    C --> J
    J --> K
    J --> L
```

## 核心模块分析

### 1. 适配器层 (Adapter)

适配器层负责与外部系统的交互，包含以下子模块：

#### 1.1 AI 模块 (`adapter/ai`)

**AI 代理系统架构：**

```mermaid
classDiagram
    class Agent {
        <<interface>>
        +Run(ctx) Message
        +GetInputMessage() []Message
    }
    
    class YahooAnalyst {
        -inputMessage []Message
        -agent *react.Agent
        +NewYahooAnalyst()
        +Run(ctx) Message
    }
    
    class ClsAnalyst {
        -inputMessage []Message
        -agent *react.Agent
        +NewClsAnalyst()
        +Run(ctx) Message
    }
    
    class SelectAnalyst {
        -inputMessage []Message
        -agent *react.Agent
        +NewSelectAnalyst()
        +Run(ctx) Message
    }
    
    Agent <|-- YahooAnalyst
    Agent <|-- ClsAnalyst
    Agent <|-- SelectAnalyst
```

**AI 工具系统：**

```mermaid
classDiagram
    class BaseTool {
        <<interface>>
        +Invoke(ctx, input) output
    }
    
    class ClsTelegramSearch {
        +GetContent(ctx, request) response
    }
    
    class ClsDepthSearch {
        +GetContent(ctx, request) response
    }
    
    class ClsDetail {
        +GetContent(ctx, request) response
    }
    
    BaseTool <|-- ClsTelegramSearch
    BaseTool <|-- ClsDepthSearch
    BaseTool <|-- ClsDetail
```

#### 1.2 爬虫模块 (`adapter/spider`)

爬虫系统负责从各种数据源获取金融信息：

```mermaid
sequenceDiagram
    participant Main as 主程序
    participant Spider as 爬虫服务
    participant Browser as 浏览器引擎
    participant DataSource as 数据源
    participant DB as 数据库
    
    Main->>Spider: 调用爬虫服务
    Spider->>Browser: 启动浏览器实例
    Browser->>DataSource: 访问目标网站
    DataSource-->>Browser: 返回页面数据
    Browser->>Spider: 解析并提取数据
    Spider->>DB: 存储数据
    Spider-->>Main: 返回处理结果
```

**主要爬虫功能：**
- 财联社新闻爬取 (`GetClsNews`)
- 财联社深度文章 (`GetClsDepthList`, `GetClsDepthDetail`)
- 财联社指数数据 (`GetClsQuotation`)
- 雅虎财经新闻 (`GetYahooNews`)

#### 1.3 控制器模块 (`adapter/controller`)

提供 RESTful API 接口，处理 HTTP 请求：

```mermaid
sequenceDiagram
    participant Client as 客户端
    participant Controller as 控制器
    participant Service as 领域服务
    participant Repository as 仓储
    
    Client->>Controller: HTTP 请求
    Controller->>Service: 调用业务逻辑
    Service->>Repository: 数据操作
    Repository-->>Service: 返回数据
    Service-->>Controller: 返回结果
    Controller-->>Client: HTTP 响应
```

#### 1.4 仓储模块 (`adapter/repository`)

数据访问层，封装数据库操作：

```mermaid
classDiagram
    class CLSRepository {
        -db *gorm.DB
        +SaveArticle(article) error
        +GetArticle(articleId) (*ClsDepthArticle, error)
        +db() *gorm.DB
    }
    
    class ClsDepthArticle {
        +ArticleID int
        +CreatedAt time.Time
        +Title string
        +Brief string
        +Content string
    }
    
    CLSRepository --> ClsDepthArticle
```

### 2. 领域层 (Domain)

#### 2.1 持久化对象 (`domain/po`)

定义数据库实体模型：

```mermaid
erDiagram
    ClsDepthArticle {
        int article_id PK
        datetime created_at
        string title
        string brief
        text content
    }
```

#### 2.2 值对象 (`domain/vo`)

定义数据传输对象和业务值对象：

```mermaid
classDiagram
    class ClsDepthArticle {
        +ArticleID int
        +Ctime int
        +SortScore int
        +Title string
        +Brief string
    }
    
    class ClsDepthArticleExt {
        +URL string
        +Brief string
        +Datetime string
    }
    
    class ClsDepthArticleContent {
        +ID int
        +Content string
    }
```

### 3. 基础设施层 (Infrastructure)

#### 3.1 请求处理 (`infra/request.go`)

统一处理 HTTP 请求的解析和验证：

```mermaid
flowchart TD
    A[HTTP 请求] --> B{请求类型}
    B -->|JSON| C[ReadJSON]
    B -->|Query| D[ReadQuery]
    B -->|Form| E[ReadForm]
    C --> F[数据验证]
    D --> F
    E --> F
    F --> G[返回解析结果]
```

#### 3.2 响应处理 (`infra/response.go`)

统一的 JSON 响应格式：

```json
{
  "code": 200,
  "error": "",
  "data": {}
}
```

### 4. 工具层 (Utility)

#### 4.1 AI 模型配置 (`utility/models.go`)

支持多种 AI 模型提供商：

```mermaid
graph LR
    A[AI 模型管理] --> B[硅基流动 DeepSeek]
    A --> C[腾讯云 DeepSeek]
    A --> D[阿里千问]
    A --> E[Google Gemini]
    
    B --> F[聊天模型]
    B --> G[嵌入模型]
    C --> F
    C --> G
    D --> F
    D --> G
    E --> F
    E --> G
```

#### 4.2 爬虫工具 (`utility/crwl.go`)

基于 Conda 环境的爬虫执行器：

```mermaid
sequenceDiagram
    participant Go as Go 程序
    participant Conda as Conda 环境
    participant Crwl as Crwl 工具
    participant Target as 目标网站
    
    Go->>Conda: 激活 investment 环境
    Conda->>Crwl: 执行 crwl 命令
    Crwl->>Target: 访问目标 URL
    Target-->>Crwl: 返回页面内容
    Crwl->>Crwl: 转换为 Markdown
    Crwl-->>Conda: 返回结果
    Conda-->>Go: 返回爬取内容
```

#### 4.3 邮件服务 (`utility/email.go`)

支持 HTML 格式的邮件发送功能。

### 5. 服务器配置 (Server)

#### 5.1 配置管理 (`server/conf/config.go`)

```mermaid
classDiagram
    class Configuration {
        +App freedom.Configuration
        +DB DBConf
        +Redis RedisConf
        +System SystemConf
        +Other map
    }
    
    class SystemConf {
        +SFAPIkey string
        +TTAPIkey string
        +QWENAPIkey string
        +GeminiAPIkey string
        +ToMailList array
        +FromMail string
        +FromMailPassword string
        +MailServer string
        +MailServerPort int
    }
    
    Configuration --> SystemConf
```

## 系统工作流程

### 美股分析流程

```mermaid
flowchart TD
    A[启动美股分析] --> B[爬取雅虎财经数据]
    B --> C[构建分析提示词]
    C --> D[创建 YahooAnalyst 代理]
    D --> E[AI 分析生成报告]
    E --> F[转换为 HTML 格式]
    F --> G[发送邮件报告]
    G --> H[等待1分钟]
    H --> I[开始A股分析]
```

### A股分析流程

```mermaid
flowchart TD
    A[启动A股分析] --> B[爬取财联社深度文章]
    B --> C[爬取财联社电报新闻]
    C --> D[爬取财联社指数数据]
    D --> E[构建分析提示词]
    E --> F[创建 ClsAnalyst 代理]
    F --> G[AI 工具链分析]
    G --> H[生成投资报告]
    H --> I[转换为 HTML 格式]
    I --> J[发送邮件报告]
```

### AI 代理工作流程

```mermaid
sequenceDiagram
    participant User as 用户输入
    participant Agent as AI 代理
    participant Tools as AI 工具
    participant Model as AI 模型
    participant Data as 数据源
    
    User->>Agent: 分析请求
    Agent->>Model: 理解任务
    Model->>Agent: 制定分析计划
    
    loop ReAct 循环
        Agent->>Tools: 调用搜索工具
        Tools->>Data: 获取数据
        Data-->>Tools: 返回数据
        Tools-->>Agent: 返回结果
        Agent->>Model: 分析数据
        Model->>Agent: 生成思考
    end
    
    Agent->>Model: 生成最终报告
    Model-->>Agent: 返回报告
    Agent-->>User: 输出分析结果
```

## 技术栈

### 后端技术
- **框架**: Freedom (基于 Iris 的 DDD 框架)
- **语言**: Go 1.21+
- **数据库**: MySQL (通过 GORM)
- **缓存**: Redis
- **AI 框架**: Eino (字节跳动开源)
- **爬虫**: Rod + Python Crwl

### AI 模型
- **DeepSeek R1**: 主要分析模型
- **千问 Plus**: 财联社分析
- **Gemini 2.0**: 备用模型
- **文本嵌入**: 多种嵌入模型支持

### 外部服务
- **数据源**: 财联社、雅虎财经
- **邮件**: SMTP 邮件服务
- **AI API**: 多厂商 API 支持

## 部署架构

```mermaid
graph TB
    subgraph "运行环境"
        A[Go 应用服务]
        B[MySQL 数据库]
        C[Redis 缓存]
        D[Conda 环境]
    end
    
    subgraph "外部依赖"
        E[财联社 API]
        F[雅虎财经 API]
        G[AI 模型 API]
        H[SMTP 邮件服务]
    end
    
    A --> B
    A --> C
    A --> D
    A --> E
    A --> F
    A --> G
    A --> H
```

## 配置文件结构

```toml
[db]
addr = "数据库连接地址"
max_open_conns = 100
max_idle_conns = 10

[redis]
addr = "Redis连接地址"
password = "密码"
db = 0

[system]
sf_api_key = "硅基流动API密钥"
tt_api_key = "腾讯云API密钥"
qwen_api_key = "千问API密钥"
gemini_api_key = "Gemini API密钥"
to_mail_list = ["接收邮箱列表"]
from_mail = "发送邮箱"
from_mail_password = "邮箱密码"
mail_server = "SMTP服务器"
mail_server_port = 587
```

## 项目特点

### 优势
1. **模块化设计**: 清晰的 DDD 分层架构
2. **多模型支持**: 集成多个 AI 模型提供商
3. **智能分析**: ReAct 模式的 AI 代理
4. **数据丰富**: 多数据源整合
5. **自动化**: 定时分析和邮件推送

### 技术亮点
1. **AI 工具链**: 自定义 AI 工具系统
2. **爬虫集成**: Go + Python 混合爬虫方案
3. **配置灵活**: 支持多环境配置
4. **错误处理**: 完善的错误处理机制
5. **日志系统**: 结构化日志记录

## 扩展建议

### 功能扩展
1. **实时监控**: 添加市场实时监控功能
2. **策略回测**: 集成投资策略回测系统
3. **风险评估**: 增强风险评估模型
4. **用户系统**: 添加多用户支持
5. **API 网关**: 提供标准化 API 接口

### 技术优化
1. **缓存策略**: 优化数据缓存机制
2. **并发处理**: 提升并发处理能力
3. **监控告警**: 添加系统监控和告警
4. **容器化**: Docker 容器化部署
5. **微服务**: 拆分为微服务架构

## 总结

Investment AI 项目是一个设计良好的智能投资分析系统，采用现代化的技术栈和架构模式。系统通过集成多种数据源和 AI 模型，实现了自动化的市场分析和报告生成。项目具有良好的可扩展性和维护性，为进一步的功能扩展奠定了坚实的基础。