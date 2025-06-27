// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ordersapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rameshsunkara/orders-sdk/internal/apijson"
	"github.com/rameshsunkara/orders-sdk/internal/apiquery"
	"github.com/rameshsunkara/orders-sdk/internal/requestconfig"
	"github.com/rameshsunkara/orders-sdk/option"
	"github.com/rameshsunkara/orders-sdk/packages/param"
	"github.com/rameshsunkara/orders-sdk/packages/respjson"
)

// OrderService contains methods and other services that help with interacting with
// the ordersapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderService] method instead.
type OrderService struct {
	Options []option.RequestOption
}

// NewOrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrderService(opts ...option.RequestOption) (r OrderService) {
	r = OrderService{}
	r.Options = opts
	return
}

// Create a new order in the system
func (r *OrderService) New(ctx context.Context, body OrderNewParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific order by ID
func (r *OrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update details of a specific order by ID
func (r *OrderService) Update(ctx context.Context, orderID string, body OrderUpdateParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of orders in the system with pagination support
func (r *OrderService) List(ctx context.Context, query OrderListParams, opts ...option.RequestOption) (res *[]Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific order by ID
func (r *OrderService) Delete(ctx context.Context, orderID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Object representing an order in the system.
type Order struct {
	// The date and time when the order was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The unique identifier for the order
	OrderID string `json:"orderId,required" format:"uuid"`
	// The list of products in the order
	Products []OrderProduct `json:"products,required"`
	// The status of the order
	//
	// Any of "OrderPending", "OrderProcessing", "OrderShipped", "OrderDelivered",
	// "OrderCancelled".
	Status OrderStatus `json:"status,required"`
	// The total amount of the order
	TotalAmount float64 `json:"totalAmount,required"`
	// The user who placed the order
	User string `json:"user,required" format:"string"`
	// The version of the order
	Version int64 `json:"version,required"`
	// The date and time when the order was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// The list of updates made to the order
	Updates []OrderUpdate `json:"updates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		OrderID     respjson.Field
		Products    respjson.Field
		Status      respjson.Field
		TotalAmount respjson.Field
		User        respjson.Field
		Version     respjson.Field
		UpdatedAt   respjson.Field
		Updates     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object representing a product in the system.
type OrderProduct struct {
	// The name of the product
	Name string `json:"name" format:"string"`
	// The price of the product
	Price float64 `json:"price"`
	// Additional remarks or description about the product
	Remarks string `json:"remarks" format:"string"`
	// The status of the product
	Status string `json:"status" format:"string"`
	// The date and time when the product information was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Price       respjson.Field
		Remarks     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderProduct) RawJSON() string { return r.JSON.raw }
func (r *OrderProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the order
type OrderStatus string

const (
	OrderStatusOrderPending    OrderStatus = "OrderPending"
	OrderStatusOrderProcessing OrderStatus = "OrderProcessing"
	OrderStatusOrderShipped    OrderStatus = "OrderShipped"
	OrderStatusOrderDelivered  OrderStatus = "OrderDelivered"
	OrderStatusOrderCancelled  OrderStatus = "OrderCancelled"
)

// Object representing an update to an order.
type OrderUpdate struct {
	// The person or system who handled the update
	HandledBy string `json:"handledBy" format:"string"`
	// Additional notes or comments about the update
	Notes string `json:"notes" format:"string"`
	// The date and time when the order was updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HandledBy   respjson.Field
		Notes       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderUpdate) RawJSON() string { return r.JSON.raw }
func (r *OrderUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object representing the input data for creating or updating an order.
//
// The property Products is required.
type OrderInputParam struct {
	// The list of products in the order
	Products []OrderInputProductParam `json:"products,omitzero,required"`
	paramObj
}

func (r OrderInputParam) MarshalJSON() (data []byte, err error) {
	type shadow OrderInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object representing the input data for creating or updating a product belonging
// to an order.
type OrderInputProductParam struct {
	// The name of the product
	Name param.Opt[string] `json:"name,omitzero" format:"string"`
	// The price of the product
	Price param.Opt[float64] `json:"price,omitzero"`
	// No.of items
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	paramObj
}

func (r OrderInputProductParam) MarshalJSON() (data []byte, err error) {
	type shadow OrderInputProductParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderInputProductParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewParams struct {
	// Object representing the input data for creating or updating an order.
	OrderInput OrderInputParam
	paramObj
}

func (r OrderNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.OrderInput)
}
func (r *OrderNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.OrderInput)
}

type OrderUpdateParams struct {
	// Object representing the input data for creating or updating an order.
	OrderInput OrderInputParam
	paramObj
}

func (r OrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.OrderInput)
}
func (r *OrderUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.OrderInput)
}

type OrderListParams struct {
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrderListParams]'s query parameters as `url.Values`.
func (r OrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
