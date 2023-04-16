package cookiestore

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// client := redis.NewClient(&redis.Options{
// 	Addr: "localhost:6379",
// 	Password: "",
// 	DB: 0,
// })

// pong, err := client.Ping(client.Context()).Result()
// fmt.Println(pong, err)

type CookieRepository struct {
	store *CookieStore
}

func (r *CookieRepository) Get(req *http.Request, sessionName string) (*sessions.Session, error) {

	// Get a session.
	session, err := r.store.RediStore.Get(req, sessionName)
	if err != nil {
		return nil, fmt.Errorf("get session error: [%w]", err)
	}

	return session, nil
	// // Add a value.
	// session.Values["foo"] = "bar"

	// // Save.
	// if err = sessions.Save(req, rsp); err != nil {
	// 	t.Fatalf("Error saving session: %v", err)
	// }

	// // Delete session.
	// session.Options.MaxAge = -1
	// if err = sessions.Save(req, rsp); err != nil {
	// 	t.Fatalf("Error saving session: %v", err)
	// }

}

func (r *CookieRepository) Delete(rsp http.ResponseWriter, req *http.Request, sessionName string) error {

	session, err := r.store.RediStore.Get(req, sessionName)
	if err != nil {
		return fmt.Errorf("get session error: [%w]", err)
	}

	session.Options.MaxAge = -1
	if err = sessions.Save(req, rsp); err != nil {
		return fmt.Errorf("saving session error: [%w]", err)
	}

	return nil
}
