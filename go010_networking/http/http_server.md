#

## `HTTP Methods`

- `GET` - yêu cần tài nguyên từ server, phản hồi qua server response.
- `HEAD` - tương tự GET nhưng không yêu cầu server phản hồi tài nguyên qua response. Dùng để truy xuất các thông tin về tài nguyên như _kích thước_, ...
- `POST` - đẩy dữ liệu lên server trong request body.
- `PUT` - tương tự POST, dùng để tải dữ liệu lên server. Thực tế thường dùng để cập nhât/thay thế toàn bộ tài nguyên nào đó đã có => _cần gửi cả object lên server_
- `PATCH` - giống PUT nhưng thay vì cập nhật toàn bộ, PATCH cập nhật một phần tài nguyên hiện có.
- `DELETE` - yêu cầu server xóa tài nguyên chỉ định.
- `OPTIONS`, `CONNECT`, `TRACE`

#

## `HTTP status code`

- `200 OK`
- `201 Created`
- `202 Accepted`
- `204 No Content`
- `304 Not Modified`
- `400 Bad Request`
- `403 Forbidden`
- `404 Not Found`
- `405 Method Not Allowed`
- `426 Upgrade Required`
- `500 Internal Server Error`
- `502 Bad Gateway`
- `503 Service Unavailable`
- `504 Gateway Timeout`

#

## HTTP Server

- Cấu trúc `http.Server`

  ```go
  type Server struct {
    Addr              string // "host:port"
    Handler           Handler // router
    TLSConfig         *tls.Config

    ReadTimeout       time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout      time.Duration
    IdleTimeout       time.Duration

    MaxHeaderBytes    int

    ConnState         func(net.Conn, ConnState)
    ErrorLog          *log.Logger
    // other options
  }
  ```

- HTTP Server gồm 3 thành phần chính:

  - _`Handler Function`_ xử lý request.

    ```go
    func handler(w http.ResponseWriter, r *http.Request) {
        // truy xuất request
        r.Method() // GET, POST
        r.URL.Path() // Kiểm tra đường dẫn URL hiện tại.
        r.URL.Query.get("key") // Lấy giá trị của tham số từ chuỗi truy vấn (query string).

        // Đọc header "Content-Type"
        contentType := r.Header.Get("Content-Type")
        // Đọc header "Authorization"
        authToken := r.Header.Get("Authorization")

        // đọc body
        // ... trong handler ...
        var data Type
        // Giải mã Body (r.Body là một io.Reader)
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            // Xử lý lỗi nếu JSON không hợp lệ
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        //application/x-www-form-urlencoded
        // Phân tích dữ liệu form từ cả Body và URL Query
        r.ParseForm()
        // Lấy giá trị từ form
        name := r.FormValue("name")


        // Responnse
        // 1. Thiết lập Content-Type (Quan trọng nhất)
        w.Header().Set("Content-Type", "application/json")
        // 2. Thiết lập Status Code
        // Nếu bạn chỉ dùng w.Write, Go sẽ mặc định dùng 200 OK.
        // Để dùng mã khác (ví dụ: 201 Created), bạn phải gọi:
        w.WriteHeader(http.StatusCreated) // Status Code 201

        // Text/HTML
        fmt.Fprintf(w, "Đã xử lý thành công yêu cầu của bạn.")
        // HOẶC: w.Write([]byte("Đã xử lý thành công"))
        // JSON
        response := map[string]interface{}{ // response body
            "status": "success",
            "data":   "Tài nguyên mới đã được tạo",
        }
        // Mã hóa và gửi phản hồi
        err := json.NewEncoder(w).Encode(response)
        if err != nil {
            log.Println("Lỗi gửi phản hồi JSON:", err)
        }
    }
    ```

  - _`Route Registration`_ đăng ký định tuyến, liên kết một route (ex: `/`, `/api`, ..) với một handler function.

    ```go
    // tạo router
    router := &http.NewServerMux
    // đăng ký route:
    router.HandleFunc("/route", RouteHandler)
    ```

  - _`Listen and Serve`_ : Lắng nghe và phục vụ => Khởi động máy chủ trên cổng cụ thể.

    ```go
    server := &http.Server{
        // config
    }

    err := server.ListenAndServe() // http/1.1
    //or
    err := server.ListenAndServeTLS(CertFile, KeyFile) // http/1.1 , http/2 tùy client
    ```

**`Note`**: Để có thể chạy HTTP/2, HTTP/3 cần có chứng chỉ TLS.  
Để `demo`, có thể sử dụng chứng chỉ TLS tự ký:

```powershell
# Mở cmd, powershell
mkdir certs
cd certs
openssl req -x509 -newkey rsa:2048 -nodes -keyout server.key -out server.crt -days 365 -subj "/CN=localhost"
```

Để chạy `HTTP/3` server:

```powershell
# powershell, cmd
go mod init example.com/http3demo
go get github.com/quic-go/quic-go/http3
```

Các bước thiết lập `http/3 server` tương đối giống `http/1.1` và `http/2` server, chỉ khác chỗ: dùng `http3.Server` thay cho `http.Server` để tạo server.

TLSConfig cho `HTTP/2`, `HTTP/3`:

```go
tlsConf := &tls.Config{
    NextProtos: []string{"h3", "h2", "http/1.1"},
    MinVersion: tls.VersionTLS13,
}

// Load chứng chỉ
cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
if err != nil {
    log.Fatalf("failed to load cert: %v", err)
}
tlsConf.Certificates = []tls.Certificate{cert}

// sau đó
/*
server := http.Server{
    //config
    TLSConfig : tlsConf
}
*/
```

- có thể dùng chứng chỉ tự ký - `self-signed certificate` để thử:

  1. tạo `server.cnf`:

  ```
  [ req ]
  default_bits       = 2048
  prompt             = no
  default_md         = sha256
  distinguished_name = dn
  x509_extensions    = v3_req

  [ dn ]
  C = VN
  ST = Hanoi
  L = Hanoi
  O = Demo
  OU = Dev
  CN = localhost

  [ v3_req ]
  subjectAltName = @alt_names

  [ alt_names ]
  DNS.1 = localhost
  IP.1 = 127.0.0.1

  ```

  2. chạy lệnh: `openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt -config server.cnf`
