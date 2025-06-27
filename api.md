# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderInputParam">OrderInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#Order">Order</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{orderId}">client.Orders.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /orders/{orderId}">client.Orders.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderListParams">OrderListParams</a>) ([]<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk">ordersapi</a>.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /orders/{orderId}">client.Orders.<a href="https://pkg.go.dev/github.com/rameshsunkara/orders-sdk#OrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
