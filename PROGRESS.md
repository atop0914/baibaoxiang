# 百宝箱 — 开发进度追踪

**项目启动日期：** 2026-06-04
**预计完成日期：** 2026-06-18
**当前阶段：** Day 1 完成

## 进度概览

| 天数 | 日期 | 任务 | 状态 |
|------|------|------|------|
| Day 1 | 06-04 | Go 后端骨架搭建 | ✅ 已完成 |
| Day 2 | 06-05 | uni-app 前端骨架 + 首页 | ⬜ 待开始 |
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
