// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ads_dfsconfiguration struct
type ads_dfsconfiguration struct {
	*ds_top
}

func Newads_dfsconfigurationEx1(instance *cim.WmiInstance) (newInstance *ads_dfsconfiguration, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dfsconfiguration{
		ds_top: tmp,
	}
	return
}

func Newads_dfsconfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dfsconfiguration, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dfsconfiguration{
		ds_top: tmp,
	}
	return
}
