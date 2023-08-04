package go_ernie

import (
	"context"
	"net/http"
)

const ernireBotTurboURL = "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant"

type ErnieBotTurboRequest struct {
	Messages     []ChatCompletionMessage `json:"messages"`
	Temperature  float64                 `json:"temperature"`
	TopP         float64                 `json:"top_p"`
	PenaltyScore float64                 `json:"penalty_score"`
	Stream       bool                    `json:"stream"`
	UserId       string                  `json:"user_id"`
}

type ErnieBotTurboResponse struct {
	ErnieBotResponse
}

func (c *Client) CreateErnieBotTurboChatCompletion(
	ctx context.Context,
	request ErnieBotTurboRequest,
) (response ErnieBotTurboResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(ernireBotTurboURL), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	if response.ErrorCode != 0 {
		err = &response.APIError
	}
	return
}
