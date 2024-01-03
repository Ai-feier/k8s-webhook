package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	authentication "k8s.io/api/authentication/v1beta1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/authenticate", func(writer http.ResponseWriter, request *http.Request) {
		// 获取请求中的 TokonReview
		decoder := json.NewDecoder(request.Body)
		var tokenReview authentication.TokenReview
		if err := decoder.Decode(&tokenReview); err != nil {
			log.Println("decode error: ", err)
			// 响应返回错误
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(map[string]interface{}{
				"apiVersion": "authentication.k8s.io/v1beta1",
				"kind":       "TokenReview",
				"status": authentication.TokenReviewStatus{
					Authenticated: false,
				},
			})
			return
		}

		log.Println("receive request")

		// 检查用户
		// 创建 oauth2 对象, 发送给 github
		oauthToken := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: tokenReview.Spec.Token, // tokenReview 的 token
		})
		oauthClient := oauth2.NewClient(context.Background(), oauthToken)
		client := github.NewClient(oauthClient)
		user, _, err := client.Users.Get(context.Background(), "")
		if err != nil {
			log.Println("github auth error: ", err)
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(map[string]any{
				"apiVersion": "authentication.k8s.io/v1beta1",
				"kind":       "TokenReview",
				"status": authentication.TokenReviewStatus{
					Authenticated: false, // 返回认证失败
				},
			})
			return
		}

		// 认证成功
		log.Println("github login success, as: ", *user.Login)
		status := authentication.TokenReviewStatus{
			Authenticated: true,
			User: authentication.UserInfo{
				Username: *user.Login,
				UID: *user.Login,
			},
		}
		json.NewEncoder(writer).Encode(map[string]any{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind":       "TokenReview",
			"status": status,
		})
	})
	log.Println(http.ListenAndServe(":3000", nil))
}
