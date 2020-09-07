package auth

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/leegeobuk/Bible-User/auth/kakao"
)

// Login authenticates user and decide whether to login or not
func Login(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.QueryStringParameters["type"] == "kakao" {
		resp := kakao.Login(&request)
		addHeaders(resp.Headers, headers)
		return resp, nil
	}

	resp := events.APIGatewayProxyResponse{Headers: headers}

	return resp, nil
}
