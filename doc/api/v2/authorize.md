# Authorize API

## Authorize - Code Mode

The authorization code method means that a third-party application first applies for an authorization code, and then uses the code to obtain a token.

This method is the most commonly used process and has the highest security. It is suitable for web applications with backends. The authorization code is transmitted through the front end, and the token is stored in the back end, and all communication with the resource server is completed in the back end. This separation of front and back ends can avoid token leakage.

### EndPoint

```http request
GET /api/v2/authorize HTTP/1.1
Host: ${server_host_address}
```

### Request-Params

+ `response_type`: **required**, must be "code".
+ `client_id`: **required**, the app_id of client.
+ `scope_filed`: **required**, the scope definition of this request.
+ `verify_key`: **required**, the key used to identify the request and response.
+ `verify_secret`: **required**, the key used to encrypt callback data.
+ `redirect_uri`: the redirect_uri of the response.

### Request-Example

```http request
GET /api/v2/authorize?response_type=code&client_id=nsr78mqpbx74b6bn1df34s93sad2v528&scope_field=nick_name.email_address.avatar_address.real_name.sub_id&verify_key=02502aafe77c03e647ea3128a3ff4236&verify_secret=d5d1360633ce35ee5e943b2f29dadc54&redirect_uri=test.domain.com HTTP/1.1
Host: ${server_host_address}
```

### Response-Params

+ `authorize_response`: the encrypted raw-data and token.
+ `verify_key`: the key used to identify the request and response, the same as `verify_key` field in the request.
+ `redirect_uri`: the redirect_uri of the response, if there is no `redirect_uri` field in the request, it will be the default redirect_uri left during client registration.

### Response-Example

> This response would be redirected to `redirect_uri`.

```go
import "fmt"
func main() {
	fmt.Print("")
}
```

```http request
POST /callback?verify_key=${verify_key} HTTP/1.1
Host: ${redirect_uri}
User-Agent: oauth-server/SchwarzOAuth
Content-Type: application/json
Content-Length: 238

{
    "authorize_response": {
        "algorithm": "HS256",
        "code": "1pmLEQjAnMynvO_WiOBX2km-8JCBISoGX6nXymqQgL4",
        "scope": {
            "nick_name": true,
            "email_address": true,
            "avatar_address": true,
            "real_name": false,
            "sub_id": false
        }
    }
}
```

## Authorize - Password Mode

If you trust an application highly, RFC 6749 also allows users to directly tell the application their username and password. The application uses your password to apply for a token. This method is called "password".

### EndPoint

```http request
GET /api/v2/authorize HTTP/1.1
Host: ${server_host_address}
```

### Request-Params

+ `response_type`: **required**, must be "password".
+ 

### Response-Params

### Response-Example

## Authorize - Implicit Mode

Some web applications are purely front-end applications and have no back-end. At this time, the above method cannot be used, and the token must be stored in the front end. RFC 6749 specifies the second method, allowing tokens to be issued directly to the front end. This method does not have the intermediate step of authorization code, so it is called (authorization code) "implicit".

### EndPoint

```http request
GET /api/v2/authorize HTTP/1.1
Host: ${server_host_address}
```

### Request-Params

+ `response_type`: **required**, must be "token".

### Response-Params

### Response-Example

## Authorize - Client Mode

The last method is client credentials, which is suitable for command-line applications without a front-end, that is, request tokens on the command line.

### EndPoint

```http request
GET /api/v2/authorize HTTP/1.1
Host: ${server_host_address}
```

### Request-Params

### Response-Params

### Response-Example
