# JWT

- 认证（Authentication，英文缩写 authn）: 用来验证某个用户是否具有访问系统的权限。如果认证通过，该用户就可以访问系统，从而创建、修改、删除、查询平台支持的资源。
- 授权（Authorization，英文缩写 authz）: 用来验证某个用户是否具有访问某个资源的权限，如果授权通过，该用户就能对资源做增删改查等操作。

> 认证证明了你是谁，授权决定了你能做什么。不要在请求参数中使用明文密码，也不要在任何存储中保存明文密码。

JWT 是 Bearer Token 的一个具体实现，由 JSON 数据格式组成，通过 HASH 散列算法生成一个字符串。该字符串可以用来进行授权和信息交换。

## jwt格式

Header.Payload.Signature格式

```shell
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NDI4NTY2MzcsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MzUwODA2MzcsInN1YiI6ImFkbWluIn0.Shw27RKENE_2MVBq7-c8OmgYdF92UmdwS8xE-Fts2FM
```

Signature 是 Token 的签名部分，通过如下方式生成：将 Header 和 Payload 分别 base64 编码后，用 . 连接。然后再使用 Header 中声明的加密方式，利用 secretKey 对连接后的字符串进行加密，加密后的字符串即为最终的 Signature。

secretKey 是密钥，保存在服务器中，一般通过配置文件来保存，例如：

```yaml
jwt:
    realm: iam jwt # jwt标识
    key: ksmjmasksmmlljjmmm, # 服务端秘钥
    timeout: 24h # token过期时间 （小时）
    max-refresh: 24h # token更新时间（小时）
```

**注意：密钥一定不能泄露。密钥泄露后，入侵者可以使用该密钥来签发 JWT Token，从而入侵系统。**

- 不要存放敏感信息在 Token 里；
- Payload 中的 exp 值不要设置得太大，一般开发版本 7 天，线上版本 2 小时。当然，你也可以根据需要自行设置。

