package googleads

import (
	"context"
	"fmt"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	*grpc.ClientConn
	services.CampaignServiceClient
	services.ConversionActionServiceClient
}

func NewGoogleApisClient() *Client {
	conn, err := grpc.Dial("googleads.googleapis.com:443",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Google Ads API: %v", err)
	}

	return &Client{
		ClientConn:                    conn,
		CampaignServiceClient:         services.NewCampaignServiceClient(conn),
		ConversionActionServiceClient: services.NewConversionActionServiceClient(conn),
	}
}

func (c *Client) AddConversionAction(ctx context.Context, customerID string,
	conversionAction *resources.ConversionAction) (string, error) {
	// Tạo yêu cầu MutationRequest
	req := &services.MutateConversionActionsRequest{
		CustomerId: customerID,
		Operations: []*services.ConversionActionOperation{
			{
				Operation: &services.ConversionActionOperation_Create{
					Create: conversionAction,
				},
			},
		},
	}

	// Gửi yêu cầu
	resp, err := c.MutateConversionActions(ctx, req)
	if err != nil {
		log.Printf("Failed to create conversion action: %v", err)
		return "", err
	}

	// Hiển thị kết quả
	fmt.Printf("Created conversion action with resource name: %s\n", resp.Results[0].ResourceName)

	return "", nil
}
