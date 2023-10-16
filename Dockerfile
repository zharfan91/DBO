
# Menggunakan image golang:1.x sebagai base image
FROM golang:1.21.1

# Menentukan direktori kerja di dalam container
WORKDIR /app

# Menyalin seluruh file aplikasi Go ke dalam direktori kerja di dalam container
COPY . .


# Menyalin file .env-example ke dalam direktori kerja dan mengganti namanya menjadi .env
COPY .env .env


# Menjalankan perintah build aplikasi Go
RUN go build -o main .

# Mengungkapkan port 80 yang akan digunakan oleh aplikasi
EXPOSE 80

# Perintah yang akan dijalankan ketika container dijalankan
CMD ["./main"]
