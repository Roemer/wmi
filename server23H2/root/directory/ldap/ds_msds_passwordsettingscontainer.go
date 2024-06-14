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

// ds_msds_passwordsettingscontainer struct
type ds_msds_passwordsettingscontainer struct {
	*ads_msds_passwordsettingscontainer
}

func Newds_msds_passwordsettingscontainerEx1(instance *cim.WmiInstance) (newInstance *ds_msds_passwordsettingscontainer, err error) {
	tmp, err := Newads_msds_passwordsettingscontainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_passwordsettingscontainer{
		ads_msds_passwordsettingscontainer: tmp,
	}
	return
}

func Newds_msds_passwordsettingscontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_passwordsettingscontainer, err error) {
	tmp, err := Newads_msds_passwordsettingscontainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_passwordsettingscontainer{
		ads_msds_passwordsettingscontainer: tmp,
	}
	return
}
