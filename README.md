# Go Say Hello

Module Go sederhana yang menyediakan fungsi untuk menampilkan pesan "Hello".

## Instalasi

Untuk menggunakan module ini di project Go Anda, jalankan:

```bash
go get github.com/Pharseus/go-say-hello
```

## Penggunaan

```go
package main

import (
    "fmt"
    gosayhello "github.com/Pharseus/go-say-hello"
)

func main() {
    message := gosayhello.SayHello()
    fmt.Println(message) // Output: Hello
}
```

## Cara Membuat Module Go dan Memberikan Tag

### 1. Inisialisasi Module

Buat module baru dengan menjalankan:

```bash
go mod init github.com/Pharseus/go-say-hello
```

Ini akan membuat file `go.mod` yang berisi informasi tentang module Anda.

### 2. Buat Kode Module

Buat file `say_hello.go` dengan kode fungsi Anda:

```go
package gosayhello

func SayHello() string {
    return "Hello"
}
```

### 3. Push ke GitHub

Commit dan push kode Anda ke repository GitHub:

```bash
git add .
git commit -m "Initial commit: Add SayHello function"
git push origin master
```

### 4. Buat Tag Versi

Untuk membuat tag versi (misalnya v1.0.0):

```bash
# Buat tag lokal
git tag v1.0.0

# Atau buat tag dengan pesan
git tag -a v1.0.0 -m "Release version 1.0.0"

# Push tag ke remote repository
git push origin v1.0.0
```

### 5. Verifikasi Tag

Untuk melihat semua tag yang ada:

```bash
git tag
```

Untuk melihat detail tag tertentu:

```bash
git show v1.0.0
```

### 6. Menggunakan Versi Spesifik

Pengguna dapat menginstall versi spesifik dari module Anda:

```bash
go get github.com/Pharseus/go-say-hello@v1.0.0
```

## Semantic Versioning

Module ini mengikuti [Semantic Versioning](https://semver.org/):

- **MAJOR** version (v1.x.x): Perubahan yang tidak kompatibel dengan versi sebelumnya
- **MINOR** version (vx.1.x): Penambahan fitur yang kompatibel dengan versi sebelumnya
- **PATCH** version (vx.x.1): Perbaikan bug yang kompatibel dengan versi sebelumnya

## Lisensi

MIT License

## Kontributor

- [Pharseus](https://github.com/Pharseus)
