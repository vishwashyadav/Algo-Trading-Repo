package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type KiteCredentials struct {
    ApiKey    string `json:"apiKey"`
    ApiSecret string `json:"apiSecret"`
    UserId    string `json:"userId"`
    Password  string `json:"password"`
}

func kiteLoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds KiteCredentials
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // TODO: Implement actual Kite authentication logic here
    // For now, just echo credentials back (do NOT do in production!)
    log.Printf("Received credentials: %+v\n", creds)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "success",
        "message": "Kite login simulated (implement real auth)",
    })
}

func main() {
    http.HandleFunc("/api/kite-login", kiteLoginHandler)
    log.Println("Backend running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}