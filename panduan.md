============================================
VekoLoad - Panduan Pengguna (Bahasa Indonesia)
============================================

1. Instalasi
------------
Unduh binary sesuai OS:
  Windows: https://vekoload.io/bin/windows/vekoload.exe
  Linux  : curl -sL https://vekoload.io/linux | bash
  MacOS  : brew install vekoload

2. Pengujian HTTP Dasar
-----------------------
vekoload http --url https://situs-anda.com --rps 1000 --duration 30s

• --url   : Target URL
• --rps   : Request per detik
• --method: GET/POST/PUT (default: GET)
• --data  : Data untuk POST (contoh: '{"user":"test"}')

3. Pengujian WebSocket
----------------------
vekoload ws --url ws://situs-anda.com/ws --connections 5000 --messages 20

• --connections : Jumlah koneksi simultan
• --messages   : Pesan per koneksi
• --interval   : Jeda antar pesan (default: 1s)
• --data       : Pesan yang dikirim (default: "ping")

4. Konfigurasi File
-------------------
Buat file `test.toml`:
```toml
[target]
url = "https://api.saya.com"
protocol = "http"

[load]
rps = 2000
duration = "5m"

[auth]
token = "Bearer xyz"
