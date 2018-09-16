package main


type user struct{
    sub string  `json:"sub"`
    name string `json:"name"`
    given_name string `json:"given_name"`
    family_name string `json:"family_name"`
    email string `json:"email"`
    picture string `json:"picture"`
}

