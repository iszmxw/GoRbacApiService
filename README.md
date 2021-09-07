# GoRbacApiService
一个可以自动构建CURD控制器的go-api服务，并预设RBAC权限功能

## 引入的库

> [gin-gonic/gin](https://github.com/gin-gonic/gin)   【Gin框架】
>
> [thedevsaddam/govalidator](https://github.com/thedevsaddam/govalidator)   【govalidator表单验证器】
>
> [gorm.io/gorm](https://gorm.io/gorm)   【Gorm数据查询工具】
>
> [spf13/viper](https://github.com/spf13/viper)   【配置读取工具】
>
> [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)   【JSON Web 令牌（JWT）的 Golang 实现】
>
>



## 接口文档

### 公共请求字段

- 可以放在header头部

| 字段名      | 类型   | 说明                  |
| ----------- | ------ | --------------------- |
| Admin-Token | string | 用户登录时获取的token |
### 1、登录模块
#### 1.1、登录接口
- 请求地址
```url
/v1/admin/user/login
```
- 请求参数

| 字段名   | 类型   | 说明   |
| -------- | ------ | ------ |
| username | string | 用户名 |
| password | string | 密码   |

```json
{
    "username": "admin",
    "password": "123456"
}
```
- 返回参数

| 字段名     | 类型     | 说明      |
| ---------- | -------- | --------- |
| code       | int      | 错误代码  |
| data       | object{} |           |
| data.token | string   | token钥匙 |
| data.msg   | string   | 提示信息  |


```json
{
    "code": 20000,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsImV4cCI6ODgyOTk4MjExOCwiaXNzIjoiaXN6bXh3In0.yZyhhtDY5cRSypIhkbhqehAYV61cs6Zixbn7y7ZIEdw"
    },
    "msg": "登录成功！"
}
```
---

#### 1.2、退出接口
- 请求地址
```url
/v1/admin/user/logout
```
- 请求参数

| 字段名 | 类型   | 说明      |
| ------ | ------ | --------- |
| token  | string | token钥匙 |

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsImV4cCI6ODgyOTk4MjExOCwiaXNzIjoiaXN6bXh3In0.yZyhhtDY5cRSypIhkbhqehAYV61cs6Zixbn7y7ZIEdw"
}
```
- 返回参数

| 字段名 | 类型   | 说明     |
| ------ | ------ | -------- |
| code   | int    | 错误代码 |
| msg    | string | 提示信息 |


```json
{
    "code": 20000,
    "msg": "退出成功！"
}
```
---
#### 1.3、获取用户信息
- 请求地址
```url
/v1/admin/user/info
```
- 请求参数

| 字段名      | 类型   | 说明   |
| ----------- | ------ | ------ |
| Admin-Token | string | 用户名 |

```json
{
    "Admin-Token":"{{token}}"
}
```
- 返回参数

| 字段名            | 类型     | 说明     |
| ----------------- | -------- | -------- |
| code              | int      | 错误代码 |
| data              | object{} |          |
| data.avatar       | string   | 用户头像 |
| data.introduction | string   | 介绍描述 |
| data.name         | string   | 用户名   |
| data.roles        | string   | 拥有角色 |
| msg               | string   | 提示信息 |


```json
{
    "code": 20000,
    "data": {
        "avatar": "https://blog.54zm.com/style/web/iszmxw_simple_pro/static/images/head.jpg",
        "introduction": "admin",
        "name": "admin",
        "roles": "[admin]"
    },
    "msg": "登录成功"
}
```
---

### 2、系统模块
#### 2.1、登录日志
- 请求地址
```url
/v1/admin/dashboard/login_log
```
- 请求参数

| 字段名 | 类型 | 说明             |
| ------ | ---- | ---------------- |
| page   | int  | 页数             |
| limit  | int  | 每页获取数据条数 |

```json
{
    "page": 1,
    "limit": 10
}
```
- 返回参数

| 字段名            | 类型       | 说明       |
| ----------------- | ---------- | ---------- |
| code              | int        | 错误代码   |
| data              | object{}   |            |
| data.current_page | int        | 当前页数   |
| data.first_page   | int        | 第一页     |
| data.last_page    | int        | 最后一页   |
| data.page_size    | int        | 每页条数   |
| data.total        | int        | 总数据条数 |
| data.data         | [object{}] | 列表数据   |
| msg               | string     | 提示信息   |


```json
{
    "code": 20000,
    "data": {
        "current_page": 1,
        "first_page": 1,
        "last_page": 154,
        "page_size": 3,
        "total": 460,
        "data": [
            {
                "id": 460,
                "account_id": 1,
                "type": 0,
                "username": "admin",
                "role_id": 1,
                "role_name": "超级管理员",
                "ip": "127.0.0.1",
                "address": "本地开发",
                "created_at": "2021-08-26T20:48:38.045+08:00",
                "updated_at": "2021-08-26T20:48:38.045+08:00",
                "deleted_at": null
            },
            {
                "id": 459,
                "account_id": 1,
                "type": 0,
                "username": "admin",
                "role_id": 1,
                "role_name": "超级管理员",
                "ip": "127.0.0.1",
                "address": "本地开发",
                "created_at": "2021-08-26T20:48:21.941+08:00",
                "updated_at": "2021-08-26T20:48:21.941+08:00",
                "deleted_at": null
            },
            {
                "id": 458,
                "account_id": 1,
                "type": 0,
                "username": "admin",
                "role_id": 1,
                "role_name": "超级管理员",
                "ip": "127.0.0.1",
                "address": "本地开发",
                "created_at": "2021-08-25T22:40:54.39+08:00",
                "updated_at": "2021-08-25T22:40:54.39+08:00",
                "deleted_at": null
            }
        ]
    },
    "msg": "查询成功！"
}
```
---

#### 2.1、操作日志
- 请求地址
```url
/v1/admin/dashboard/operation_log
```
- 请求参数

| 字段名  | 类型   | 说明             |
| ------- | ------ | ---------------- |
| page    | int    | 页数             |
| limit   | int    | 每页获取数据条数 |
| orderBy | string | 排序             |

```json
{
    "page": 1,
    "limit": 3,
    "orderBy": "id asc"
}
```
- 返回参数

| 字段名            | 类型       | 说明       |
| ----------------- | ---------- | ---------- |
| code              | int        | 错误代码   |
| data              | object{}   |            |
| data.current_page | int        | 当前页数   |
| data.first_page   | int        | 第一页     |
| data.last_page    | int        | 最后一页   |
| data.page_size    | int        | 每页条数   |
| data.total        | int        | 总数据条数 |
| data.data         | [object{}] | 列表数据   |
| msg               | string     | 提示信息   |


```json
{
    "code": 20000,
    "data": {
        "current_page": 1,
        "first_page": 1,
        "last_page": 81,
        "page_size": 3,
        "total": 242,
        "data": [
            {
                "id": 1,
                "type": 1,
                "account_id": 1,
                "username": "admin",
                "role_name": "超级管理员",
                "content": "修改了登录密码！",
                "ip": "127.0.0.1",
                "address": "",
                "created_at": "2021-01-20T21:24:42.11+08:00",
                "updated_at": "2021-01-20T21:24:42.11+08:00",
                "deleted_at": null
            },
            {
                "id": 3,
                "type": 1,
                "account_id": 1,
                "username": "admin",
                "role_name": "超级管理员",
                "content": "修改了商户ID为【3】的合作商户名称！",
                "ip": "127.0.0.1",
                "address": "",
                "created_at": "2021-01-20T21:24:42.11+08:00",
                "updated_at": "2021-01-20T21:24:42.11+08:00",
                "deleted_at": null
            },
            {
                "id": 4,
                "type": 1,
                "account_id": 1,
                "username": "admin",
                "role_name": "超级管理员",
                "content": "修改了商户ID为【3】的合作商户名称！",
                "ip": "127.0.0.1",
                "address": "",
                "created_at": "2021-01-20T21:24:42.11+08:00",
                "updated_at": "2021-01-20T21:24:42.11+08:00",
                "deleted_at": null
            }
        ]
    },
    "msg": "查询成功！"
}
```
---

### 3、菜单角色

#### 3.1、角色列表
- 请求地址
```url
/v1/admin/roles/list
```
- 请求参数

| 字段名  | 类型   | 说明             |
| ------- | ------ | ---------------- |
| page    | int    | 页数             |
| limit   | int    | 每页获取数据条数 |
| orderBy | string | 排序             |

```json
{
    "page": 1,
    "limit": 3,
    "orderBy": "id asc"
}
```
- 返回参数

| 字段名            | 类型       | 说明       |
| ----------------- | ---------- | ---------- |
| code              | int        | 错误代码   |
| data              | object{}   |            |
| data.current_page | int        | 当前页数   |
| data.first_page   | int        | 第一页     |
| data.last_page    | int        | 最后一页   |
| data.page_size    | int        | 每页条数   |
| data.total        | int        | 总数据条数 |
| data.data         | [object{}] | 列表数据   |
| msg               | string     | 提示信息   |


```json
{
    "code": 20000,
    "data": {
        "current_page": 1,
        "first_page": 1,
        "last_page": 2,
        "page_size": 3,
        "total": 4,
        "data": [
            {
                "id": 4,
                "name": "运营部",
                "routes": "1,37,6,7,19,20,21,10,13,14,4,22,23,24,25,26,5,27,31,32,28,33,29,34,35,30,15,16,17,18,36",
                "desc": "运营部有用的权限",
                "created_at": "2021-01-20T21:24:42.11+08:00",
                "updated_at": "2021-01-20T21:24:42.11+08:00",
                "deleted_at": null
            },
            {
                "id": 3,
                "name": "技术部",
                "routes": "1,37,2,6,7,8,19,20,21,3,9,10,11,12,13,14,4,22,23,24,25,26,5,27,31,32,28,33,29,34,35,30,15,16,17,18,36",
                "desc": "技术部拥有的权限",
                "created_at": "2021-01-20T21:24:42.11+08:00",
                "updated_at": "2021-01-20T21:24:42.11+08:00",
                "deleted_at": null
            },
            {
                "id": 2,
                "name": "公司财务",
                "routes": "1,37,19,20,21,10,13,14,22,23,24,25,5,27,31,32,28,33,41,29,34,35,30,15,16,17,18,36,42,43,38,39,40",
                "desc": "财务拥有的权限",
                "created_at": "2021-01-20T21:24:42.11+08:00",
                "updated_at": "2021-01-20T21:24:42.11+08:00",
                "deleted_at": null
            }
        ]
    },
    "msg": "查询成功！"
}
```
---

#### 3.2、角色详情

- 请求地址
```url
/v1/admin/roles/routes
```
- 请求参数

| 字段名 | 类型 | 说明   |
| ------ | ---- | ------ |
| id     | int  | 角色id |

```json
{
    "id": 1
}
```
- 返回参数

| 字段名              | 类型       | 说明               |
| ------------------- | ---------- | ------------------ |
| code                | int        | 错误代码           |
| data                | object{}   |                    |
| data.all_route_list | [object{}] | 所有路由           |
| data.defaultChecked | [object{}] | 当前角色拥有的路由 |
| msg                 | string     | 提示信息           |


```json
{
	"code": 20000,
	"data": {
		"all_route_list": [{
			"id": 1,
			"is_menu": "1",
			"name": "系统首页",
			"parent_id": 0,
			"disabled": false,
			"children": [{
				"id": 8,
				"is_menu": "1",
				"name": "首页统计",
				"parent_id": 1,
				"disabled": false,
				"children": null
			}]
		}, {
			"id": 2,
			"is_menu": "1",
			"name": "系统管理",
			"parent_id": 0,
			"disabled": false,
			"children": [{
				"id": 3,
				"is_menu": "0",
				"name": "登录日志",
				"parent_id": 2,
				"disabled": false,
				"children": [{
					"id": 7,
					"is_menu": "1",
					"name": "日志列表",
					"parent_id": 3,
					"disabled": false,
					"children": null
				}]
			}, {
				"id": 4,
				"is_menu": "0",
				"name": "操作日志",
				"parent_id": 2,
				"disabled": false,
				"children": null
			}, {
				"id": 5,
				"is_menu": "0",
				"name": "角色权限",
				"parent_id": 2,
				"disabled": false,
				"children": null
			}, {
				"id": 6,
				"is_menu": "0",
				"name": "菜单配置",
				"parent_id": 2,
				"disabled": false,
				"children": null
			}, {
				"id": 12,
				"is_menu": "1",
				"name": "日志管理",
				"parent_id": 2,
				"disabled": false,
				"children": null
			}, {
				"id": 13,
				"is_menu": "1",
				"name": "用户管理",
				"parent_id": 2,
				"disabled": false,
				"children": null
			}]
		}, {
			"id": 9,
			"is_menu": "1",
			"name": "公用页面",
			"parent_id": 0,
			"disabled": false,
			"children": [{
				"id": 10,
				"is_menu": "1",
				"name": "404",
				"parent_id": 9,
				"disabled": false,
				"children": null
			}, {
				"id": 11,
				"is_menu": "1",
				"name": "用户信息",
				"parent_id": 9,
				"disabled": false,
				"children": null
			}, {
				"id": 14,
				"is_menu": "1",
				"name": "系统设置",
				"parent_id": 9,
				"disabled": false,
				"children": null
			}]
		}],
		"defaultChecked": ["1", "6", "7", "8", "19", "20", "3", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18"]
	},
	"msg": "查询成功！"
}
```
---

#### 3.3、菜单列表

### 4、公共模块

#### 4.1、上传图片

