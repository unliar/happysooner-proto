package grpc_client

import (
	"time"

	"github.com/unliar/happysooner-proto/v2/account"
	"github.com/unliar/happysooner-proto/v2/pay"
	"github.com/unliar/happysooner-proto/v2/push"
	"github.com/unliar/happysooner-proto/v2/search"
	"github.com/unliar/happysooner-proto/v2/writing"
	"go-micro.dev/v4/client"
)

var c = client.NewClient(func(o *client.Options) {
	o.CallOptions.RequestTimeout = 15 * time.Second
})

var AccountSVService = account.NewAccountSVService("happysooner-account-rpc", c)

var PushService = push.NewPushSVService("happysooner-push-rpc", c)

var WritingService = writing.NewWritingSVService("happysooner-writing-rpc", c)

var SearchService = search.NewSearchSVService("happysooner-search-rpc", c)

var PayService = pay.NewPaySVService("happysooner-pay-rpc", c)
