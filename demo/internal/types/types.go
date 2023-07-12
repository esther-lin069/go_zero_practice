// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile,optional"`
	Gender   string `json:"gender,options=M|F"`
	Nickname string `json:"nickname"`
}

type SignUpResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
