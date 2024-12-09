package main

import (
	"fmt"
	"net/http"
	"telegram-personal-trainer-bot/internal/config"
)

//func main() {
//	cfg := config.MustConfig()
//	_ = redis.MustRedis(cfg.Redis)
//
//	pref := tele.Settings{
//		Token:  cfg.Token,
//		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
//	}
//
//	b, err := tele.NewBot(pref)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	b.Handle("/hello", func(c tele.Context) error {
//		return c.Send("Hello!")
//	})
//
//	b.Start()
//}

func handle(w http.ResponseWriter, r *http.Request) {
	cfg := config.MustConfig()
	fmt.Fprintf(w, "Hello, World! %s", cfg.Token)
}

//

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
