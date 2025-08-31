# Mugi-DC Bot

![Mugi-DC Bot](https://i.ytimg.com/vi/x7wM7zx7onQ/maxresdefault.jpg)

My kisah tuh Mugi🥰😘😍😚🥵.

## Fitur

- Mengirim pesan otomatis ke channel Discord yang ditentukan.
- Konfigurasi mudah melalui file `.env`.
- Bot akan keluar setelah mengirim pesan.

## Prasyarat

- Go versi 1.16 atau lebih baru (unduh dari [golang.org](https://golang.org/dl/)).
- Token bot Discord (dapatkan dari [Discord Developer Portal](https://discord.com/developers/applications)).
- ID channel Discord tempat pesan akan dikirim.

## Instalasi

1. **Clone atau download proyek ini** ke direktori lokal Anda.

2. **Install dependencies**:
   Buka terminal di direktori proyek dan jalankan:

   ```bash
   go mod tidy
   ```

   Ini akan mengunduh dan menginstall semua dependensi yang diperlukan, seperti `discordgo` dan `godotenv`.

3. **Buat file konfigurasi `.env`**:
   Di root direktori proyek, buat file bernama `.env` dengan konten berikut:
   ```
   DISCORD_TOKEN=your_discord_bot_token_here
   CHANNEL_ID=your_channel_id_here
   MESSAGE=your_message_here
   ```
   - Ganti `your_discord_bot_token_here` dengan token bot Discord Anda.
   - Ganti `your_channel_id_here` dengan ID channel Discord (klik kanan channel di Discord, pilih "Copy ID" – pastikan Developer Mode aktif).
   - Ganti `your_message_here` dengan pesan yang ingin dikirim. Jika dibiarkan kosong, bot akan menggunakan pesan default: "Halo semua! Pesan otomatis dari bot Go."

## Penggunaan

1. **Jalankan bot**:
   Di terminal, jalankan perintah berikut:

   ```bash
   go run main.go
   ```

2. Bot akan:

   - Terhubung ke Discord.
   - Mengirim pesan ke channel yang ditentukan.
   - Keluar secara otomatis setelah pengiriman.

3. **Output di terminal**:
   - "Bot sudah online sebagai [nama bot]"
   - "Pesan terkirim" jika berhasil, atau pesan error jika gagal.
   - "OTW BOSKU" sebelum menunggu sinyal untuk keluar.

## Konfigurasi

File `.env` digunakan untuk menyimpan konfigurasi sensitif:

- `DISCORD_TOKEN`: Token autentikasi bot Discord. Jangan bagikan token ini.
- `CHANNEL_ID`: ID unik channel Discord.
- `MESSAGE`: Pesan yang akan dikirim. Mendukung format Markdown Discord (misalnya, **bold**, _italic_, dll.).

Pastikan file `.env` tidak di-commit ke repository publik untuk alasan keamanan.

## Troubleshooting

- **Error "Gagal memuat file .env"**: Pastikan file `.env` ada di root direktori proyek dan formatnya benar.
- **Error "Pastikan DISCORD_TOKEN dan CHANNEL_ID ada di environment variable"**: Periksa apakah token dan channel ID diisi dengan benar di `.env`.
- **Error "Gagal membuat Discord session" atau "Tidak bisa membuka koneksi"**: Periksa token bot apakah valid dan bot memiliki izin yang cukup di server Discord.
- **Error "Gagal kirim pesan"**: Pastikan channel ID benar dan bot memiliki izin untuk mengirim pesan di channel tersebut.

Jika bot tidak berfungsi, periksa log di terminal untuk detail error.

## Dependensi

- [discordgo](https://github.com/bwmarrin/discordgo): Library Go untuk interaksi dengan Discord API.
- [godotenv](https://github.com/joho/godotenv): Library untuk memuat variabel environment dari file `.env`.

## Lisensi

Proyek ini menggunakan lisensi MIT. Lihat file `LICENSE` untuk detail lebih lanjut (jika ada).

## Kontribusi

Jika Anda ingin berkontribusi, silakan fork repository ini dan buat pull request dengan perubahan Anda.

## Dukungan

Untuk pertanyaan atau dukungan, buat issue di repository ini atau hubungi pengembang.
