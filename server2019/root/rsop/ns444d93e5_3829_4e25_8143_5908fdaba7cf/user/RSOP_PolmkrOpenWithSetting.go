// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS444D93E5_3829_4E25_8143_5908FDABA7CF.User
//
// ////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrOpenWithSetting struct
type RSOP_PolmkrOpenWithSetting struct {
	*RSOP_PolmkrFolderOptionSetting
}

func NewRSOP_PolmkrOpenWithSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrOpenWithSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrOpenWithSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrOpenWithSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrOpenWithSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrOpenWithSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}
