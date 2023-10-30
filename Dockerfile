# Gunakan image golang sebagai base image
FROM golang:1.20

# Setel direktori kerja ke direktori program Go
WORKDIR /my-restaurant-app

# Salin file Go program Anda ke dalam image
COPY . .

# Kompilasi program Go
RUN go build -o main .

# Expose port yang akan digunakan oleh server Anda
EXPOSE 8080

# Perintah untuk menjalankan aplikasi saat Docker container dijalankan
CMD ["./main"]
