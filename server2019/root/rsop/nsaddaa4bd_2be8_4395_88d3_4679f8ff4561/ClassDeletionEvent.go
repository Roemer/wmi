// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NSADDAA4BD_2BE8_4395_88D3_4679F8FF4561
//
// ////////////////////////////////////////////
package nsaddaa4bd_2be8_4395_88d3_4679f8ff4561

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ClassDeletionEvent struct
type __ClassDeletionEvent struct {
	*__ClassOperationEvent
}

func New__ClassDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__ClassDeletionEvent, err error) {
	tmp, err := New__ClassOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassDeletionEvent{
		__ClassOperationEvent: tmp,
	}
	return
}

func New__ClassDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassDeletionEvent, err error) {
	tmp, err := New__ClassOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassDeletionEvent{
		__ClassOperationEvent: tmp,
	}
	return
}
