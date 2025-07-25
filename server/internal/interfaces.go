/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package internal

import (
	"io"
	"net/url"

	"github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/common"
	"github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/model/message"
)

// IWebSocketClient - interface that manages a weboscket connection.
// Maps handlers by actions and/or responses by requestID.
//
//go:generate mockgen -destination ./mock/client.go -package=mock . IWebSocketClient
type IWebSocketClient interface {
	io.Closer
	Connect(url *url.URL) error
	SendRequest(req MessageGetter, resp chan<- common.Outcome) error
	AddHandler(action message.MessageAction, handler func([]byte))
	CancelRequest(requestID string)
}

// MessageGetter - interface representing the data type that contains request.Request.
type MessageGetter interface {
	GetMessage() message.Message
}
