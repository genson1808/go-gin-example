Để tạo một web server, Go cung cấp đúng 1 hàm thuộc gói `net/http` để xử lý đó là 
`ListenAndServe(addr string, handler Handler)`. `Addr` là địa chỉ server và cổng nhưng
thường chỉ có cổng vì chúng ta mở server tại cùng ứng dụng web nên địa chỉ là 
localhost có thể bỏ. Tham số thứ 2 là kiểu `interface http.Handler`, nơi đảm nhận 
việc nhận và xử lý yêu cầu từ client. Interface này có duy nhất phương thức 
`ServeHTTP(ResponseWriter, *Request)`. Để có thể tạo một đối tượng xử lý yêu cầu 
từ client, chúng ta cần khai báo một kiểu dữ liệu thỏa mãn `interface http.Handler`, 
tức có phương thức `ServeHTTP(ResponseWriter, *Request)`. Phương thức `ServeHTTP` có 
2 tham số: `interface http.ResponseWriter` đảm nhận việc ghi phần đầu và thân vào 
gói phản hồi và con trỏ thực thể cấu trúc `http.Request` chứa các thông tin phần 
đầu và thân của gói yêu cầu. 
