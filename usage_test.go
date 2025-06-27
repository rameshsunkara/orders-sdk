// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ordersapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/rameshsunkara/orders-sdk"
	"github.com/rameshsunkara/orders-sdk/internal/testutil"
	"github.com/rameshsunkara/orders-sdk/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ordersapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	orders, err := client.Orders.List(context.TODO(), ordersapi.OrderListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", orders)
}
