# 疾书 

## 运行

```
go run .
```

## 数据库

数据库sql文件为 `Jishu.sql` ，请设置环境变量以指定数据库

## API文档

蒟蒻初次手写，请多关照


###用户管理

####用户注册 `POST {base_url}/api/register`

- Request
  ```json
  {
    "username":"balabala",
    "password":"abcd1234",   //是否需要规定密码格式？
  }
  ```

- Response
  ```json
  {
    "success":true,
    "data":{
      "userId":1,
    }
  }
  ```

- Error
  ```json
  {
    "success":false,
    "error":"message",
  }
  ```


#### 用户登录 `POST {base_url}/api/login`

- Request
  ```json
  {
    "username":"balabala",
    "password":"abcd1234"
  }
  ```

- Response
  ```json
  {
    "success":true,
    "data"{
      "userId":1,
    }
  }
  ```

- Error
  ```json
  {
    "success":false,
    "error":"message",
  }
  ```


#### 用户登出 `DELETE {base_url}/api/logout`

- Respuest
  ```json
  session获取userId
  ```

- Response
  ```json
  {
    "success":true,
    "data":"登出成功"
  }
  ```

- Error
  ```json
  {
    "success":false,
  }

