# 百宝箱 — 开发进度追踪

**项目启动日期：** 2026-06-04
**预计完成日期：** 2026-06-18
**当前阶段：** Day 2 完成

## 进度概览

| 天数 | 日期 | 任务 | 状态 |
|------|------|------|------|
| Day 1 | 06-04 | Go 后端骨架搭建 | ✅ 已完成 |
| Day 2 | 06-05 | uni-app 前端骨架 + 首页 | ✅ 已完成 |
| Day 3 | 06-06 | JSON 格式化工具 | ⬜ 待开始 |
| Day 4 | 06-07 | Base64 + 时间戳 | ⬜ 待开始 |
| Day 5 | 06-08 | 二维码生成 | ⬜ 待开始 |
| Day 6 | 06-09 | 颜色转换 | ⬜ 待开始 |
| Day 7 | 06-10 | 文本处理 | ⬜ 待开始 |
| Day 8-9 | 06-11~12 | 单位换算 | ⬜ 待开始 |
| Day 10 | 06-13 | UUID 生成 | ⬜ 待开始 |
| Day 11-12 | 06-14~15 | 正则测试 | ⬜ 待开始 |
| Day 13 | 06-16 | Markdown 预览 | ⬜ 待开始 |
| Day 14 | 06-17 | 收尾优化 + 上线 | ⬜ 待开始 |

状态标记：⬜ 待开始 | 🔄 进行中 | ✅ 已完成 | ❌ 阻塞

## 每日详细记录

### 2026-06-04 — Day 1：Go 后端骨架搭建

**完成内容：**
- ✅ 初始化 Go module (`go mod init baibaoxiang`)
- ✅ 安装依赖：gin v1.12.0, go-sqlite3 v1.14.44
- ✅ 编写 `config/config.go`：环境变量配置（PORT/DB_PATH/DEBUG）
- ✅ 编写 `database/sqlite.go`：SQLite 初始化 + tool_usage 建表 + RecordUsage 函数
- ✅ 编写 `middleware/cors.go`：跨域支持，OPTIONS 预检返回 204
- ✅ 编写 `middleware/logger.go`：请求日志中间件
- ✅ 编写 `handlers/health_handler.go`：GET /api/ping 健康检查
- ✅ 编写 `main.go`：Gin 服务入口，注册中间件和路由，优雅关闭
- ✅ 编译通过，服务启动成功
- ✅ curl /api/ping 返回 `{"code":0,"message":"success","data":{"status":"ok","go":"go1.25.0","service":"baibaoxiang"}}`
- ✅ CORS OPTIONS 预检正常返回 204 + 正确的 Allow-* 头
- ✅ SQLite 数据库文件 `backend/data/baibaoxiang.db` 自动创建

**遇到的问题：**
- go-sqlite3 需要 CGO_ENABLED=1（已知约束，正常处理）
- gin v1.12.0 要求 Go 1.25+，go mod tidy 自动升级了 toolchain 到 go1.25.0

**文件清单：**
```
backend/
├── go.mod / go.sum
├── main.go
├── config/config.go
├── database/sqlite.go
├── middleware/cors.go
├── middleware/logger.go
├── handlers/health_handler.go
└── data/baibaoxiang.db (运行时生成)
```

### 2026-06-05 — Day 2：uni-app 前端骨架 + 首页

**完成内容：**
- ✅ 使用 `npx degit dcloudio/uni-preset-vue#vite-ts` 初始化 uni-app Vue3 项目
- ✅ 精简 package.json：只保留微信小程序和 H5 相关依赖
- ✅ 配置 `pages.json`：11 个页面路由（首页 + 10 个工具页）
- ✅ 配置 `manifest.json`：AppID、微信小程序设置、H5 开发代理
- ✅ 编写首页 `pages/index/index.vue`：
  - 自定义导航栏（蓝色渐变背景 + 状态栏适配）
  - 10 个工具卡片网格列表，每个卡片有图标、名称、描述
  - 点击卡片跳转到对应工具页面
- ✅ 编写 `components/ToolCard.vue`：图标+名称+描述+箭头，点击跳转
- ✅ 编写 `components/NavBar.vue`：自定义导航栏组件，支持返回按钮
- ✅ 编写 `components/ResultBox.vue`：结果展示组件，带复制功能
- ✅ 编写 `api/index.ts`：封装 uni.request，统一错误处理，预定义所有工具 API
- ✅ 编写 `utils/index.ts`：防抖、节流、复制、日期格式化等工具函数
- ✅ 创建 10 个工具页面的占位文件（待后续开发）
- ✅ 更新 `App.vue`：全局样式（配色、字体、卡片、按钮）
- ✅ H5 构建成功
- ✅ 微信小程序构建成功

**文件清单：**
```
frontend/
├── package.json / package-lock.json
├── vite.config.ts
├── tsconfig.json
├── index.html
├── .gitignore
└── src/
    ├── App.vue (全局样式)
    ├── main.ts (入口)
    ├── pages.json (路由配置)
    ├── manifest.json (应用配置)
    ├── api/index.ts (API 封装)
    ├── utils/index.ts (工具函数)
    ├── components/
    │   ├── NavBar.vue (导航栏)
    │   ├── ToolCard.vue (工具卡片)
    │   └── ResultBox.vue (结果展示)
    ├── pages/
    │   ├── index/index.vue (首页)
    │   ├── json-formatter/index.vue (占位)
    │   ├── base64/index.vue (占位)
    │   ├── timestamp/index.vue (占位)
    │   ├── qrcode/index.vue (占位)
    │   ├── color/index.vue (占位)
    │   ├── text/index.vue (占位)
    │   ├── unit/index.vue (占位)
    │   ├── uuid/index.vue (占位)
    │   ├── regex/index.vue (占位)
    │   └── markdown/index.vue (占位)
    └── static/logo.png
```
