#

### Thao tác với luồng I/O với thư viện `io`

Các interfaces `io.Reader` và `io.Writer` là các lớp cơ sở cho mọi hoạt động I/O trong Go. Hầu hết các nguồn/đích I/O đều triển khai một hoặc cả hai interfaces này, bao gồm: `net.Conn`, tệp `*os.File`, bộ nhớ `*bytes.Buffer`, ..v.vv..

- Phương thức chính của `io.Reader`:

```go
// @params
// - b : raw bytes -> buffer để chứa dữ liệu đọc được
// @return
// - n : số byte đã đọc
// - err : lỗi, đặc biệt là io.EOF
Read(b []byte) (n int, err error)
```

- Phương thức chính của `io.Writer`:

```go
// @params
// - b : raw bytes
// @return
// - n : số byte write được
// - err
Write(b []byte) (n int, err error)
```

Ngoài ra, `io` còn có các method:

```go
io.Copy(dst io.Writer, src io.Reader) // hiệu qua cho việc chuyển dữ liệu lớn

io.ReadAll(r io.Reader) // đọc toàn bộ r cho tới khi EOF

io.LimitReader(r io.Reader, n int64) // tạo 1 Reader chỉ cho phép đọc tối đa n byte từ r.
```

#

### Thao tác với luồng I/O với thư viện `bufio - Buffered I/O`

Tương đương `BufferedInputStream - bufio.Reader`/`BufferedOutputStream - bufio.WriterX` trong Java.

- Các phương thức của `bufio.Reader`:

```go
// Tạo đối tượng bufid.Reader
bufio.NewReader(rd io.Reader) *bufio.Reader
// or
// NewReaderSize(rd io.Reader, size int)

// các methods của reader:


ReadBytes(delim byte) (b []byte, err error) // read cho tới khi gặp ký tự phân cách

ReadString(delim byte) (s string, err error) // đọc tới khi gặp phân cách và trả về chuỗi

Peek(n int) (b []byte, err error) // trả về tối đa n byte tiếp theo nhưng không di chuyển vị trí con trỏ  // dùng để xem trước dữ liệu mà không đọc nó

UnreadByte() (err error) // đẩy byte cuối cùng vừa đọc trở lại bộ đệm // di chuyển con trỏ ngược lại

ReadByte() (c byte, err error) // read a byte
ReadRune() (r rune, size int, err error) // read rune
Read(b byte[]) (n int, err error)  // read multi-bytes
```

- Các phương thức của `bufio.Reader`:

```go
// Tạo đối tượng bufio.Writer
bufio.NewWriter(wr io.Writer) *bufio.Writer
// hoặc
// NewWriterSize(wr io.Writer, size int)

// Các phương thức của Writer

WriteString(s string) (n int, err error)

Write(b []byte) (n int, err error)
WriteByte(c byte) err error

Flush() err error // gửi dữ liệu đi ngay khi gọi lệnh này

Buffered() length int// trả về độ dài buffer (số byte trong bộ đệm đang có)
```

#

### Thao tác với luồng I/O với thư viện `encoding/binary`

Đọc ghi dữ liệu có cấu trúc, tương đương `DataInputStream - binary.Read()` và `DataOutputStream - binary.Write()` trong Java, làm việc với type/struct có kích thước cố định (int, float, ... )

2 hàm chính:

- Đọc dữ liệu:

  ```go
    // @params
    // - rd: nguồn byte
    // - order: binary.BigEndian (thường sử dụng) / binary.LittleEndian
    // - data: biến để chứa dữ liệu
    binary.Read(rd io.Reader, order ByteOrder, &data any)

  ```

- Ghi dữ liệu:

  ```go
  binary.Write(wr io.Writer, order ByteOrder, data any) // chuyển data-> binary sau đó ghi vào wr
  ```

- order = `binary.BigEndian` (thường sử dụng) / binary.`LittleEndian`:
  - BE: byte có trọng số lớn nhất được lưu/truyền trước.
  - LE: byte có trọng số nhỏ nhất được lưu/truyền trước.
