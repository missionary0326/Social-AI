package model

type Post struct {
    ID      string `json:"id"`
    User    string `json:"user"`
    Message string `json:"message"`
    URL     string `json:"url"`
    Type    string `json:"type"`
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Age      int64  `json:"age"`
    Gender   string `json:"gender"`
}