package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	message := os.Getenv("MESSAGE")

	if token == "" || channelID == "" {
		log.Fatal("Pastikan DISCORD_TOKEN dan CHANNEL_ID ada di environment variable")
	}
	if message == "" {
		message = "Halo semua! Pesan otomatis dari bot Go."
	}

	// Session bot
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Gagal membuat Discord session:", err)
	}

	// Berhasil langsung kirim pesan
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot sudah online sebagai", s.State.User.Username)

		_, err := s.ChannelMessageSend(channelID, message)
		if err != nil {
			fmt.Println("Gagal kirim pesan:", err)
		} else {
			fmt.Println("Pesan terkirim")
		}

		// Setelah kirim → keluar
		s.Close()
	})

	// Buka koneksi
	err = dg.Open()
	if err != nil {
		log.Fatal("Tidak bisa membuka koneksi:", err)
	}

	// Tunggu sampai bot ditutup
	fmt.Println("OTW BOSKU")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
}