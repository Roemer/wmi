// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS3BC654AA_392A_4C67_BCCB_7765E6837951
//////////////////////////////////////////////
package ns3bc654aa_392a_4c67_bccb_7765e6837951

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __MethodProviderRegistration struct
type __MethodProviderRegistration struct {
	*__ProviderRegistration
}

func New__MethodProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__MethodProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__MethodProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__MethodProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__MethodProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__MethodProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}
