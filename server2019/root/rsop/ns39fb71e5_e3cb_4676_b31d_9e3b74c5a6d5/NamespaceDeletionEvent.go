// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS39FB71E5_E3CB_4676_B31D_9E3B74C5A6D5
//
// ////////////////////////////////////////////
package ns39fb71e5_e3cb_4676_b31d_9e3b74c5a6d5

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __NamespaceDeletionEvent struct
type __NamespaceDeletionEvent struct {
	*__NamespaceOperationEvent
}

func New__NamespaceDeletionEventEx1(instance *cim.WmiInstance) (newInstance *__NamespaceDeletionEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__NamespaceDeletionEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}

func New__NamespaceDeletionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__NamespaceDeletionEvent, err error) {
	tmp, err := New__NamespaceOperationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__NamespaceDeletionEvent{
		__NamespaceOperationEvent: tmp,
	}
	return
}
