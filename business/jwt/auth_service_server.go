package jwt

import (
	"change_password_service/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elgs/gojq"
)

type JwtService interface {
	ValidateToken(token string) bool
}
type jwtService struct {
}

func NewJwtService() JwtService {
	return &jwtService{}
}

func (c *jwtService) ValidateToken(reqToken string) bool {
	server := config.GetJwtServer()
	baseUrl1 := fmt.Sprintf("http://%s:%s/ValidateToken", server.Host, server.Port)
	req, err := http.NewRequest("GET", baseUrl1, nil)
	req.Header.Set("Authorization", reqToken)
	if err != nil {
		fmt.Println("request error", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	sb := string(body)
	parser, _ := gojq.NewStringQuery(sb)

	status, _ := parser.QueryToBool("status")

	return status
}
