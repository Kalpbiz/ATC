package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/mux"
    "time"
)

var jwtKey = []byte("super-secret")

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

var users = map[string]string{
    "admin": "password123",
}

type Flight struct {
    ID     int    `json:"id"`
    Status string `json:"status"`
}

func login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    expectedPassword, ok := users[creds.Username]
    if !ok || expectedPassword != creds.Password {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Username: creds.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })
}

func getFlights(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    tokenStr := cookie.Value
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    flights := []Flight{
        {ID: 1, Status: "On Time"},
        {ID: 2, Status: "Delayed"},
    }
    json.NewEncoder(w).Encode(flights)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/login", login).Methods("POST")
    r.HandleFunc("/api/flights", getFlights).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}
