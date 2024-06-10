// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_msauthz_centralaccesspolicies struct
type ds_msauthz_centralaccesspolicies struct {
	*ads_msauthz_centralaccesspolicies
}

func Newds_msauthz_centralaccesspoliciesEx1(instance *cim.WmiInstance) (newInstance *ds_msauthz_centralaccesspolicies, err error) {
	tmp, err := Newads_msauthz_centralaccesspoliciesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccesspolicies{
		ads_msauthz_centralaccesspolicies: tmp,
	}
	return
}

func Newds_msauthz_centralaccesspoliciesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msauthz_centralaccesspolicies, err error) {
	tmp, err := Newads_msauthz_centralaccesspoliciesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msauthz_centralaccesspolicies{
		ads_msauthz_centralaccesspolicies: tmp,
	}
	return
}
