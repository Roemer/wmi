// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NSB8ABF180_FFEF_497C_BC82_26B24617DE91
//////////////////////////////////////////////
package nsb8abf180_ffef_497c_bc82_26b24617de91

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __SecurityRelatedClass struct
type __SecurityRelatedClass struct {
	*cim.WmiInstance
}

func New__SecurityRelatedClassEx1(instance *cim.WmiInstance) (newInstance *__SecurityRelatedClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &__SecurityRelatedClass{
		WmiInstance: tmp,
	}
	return
}

func New__SecurityRelatedClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__SecurityRelatedClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__SecurityRelatedClass{
		WmiInstance: tmp,
	}
	return
}
