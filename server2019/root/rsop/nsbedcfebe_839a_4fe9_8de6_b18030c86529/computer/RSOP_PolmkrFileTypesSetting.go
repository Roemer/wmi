// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NSBEDCFEBE_839A_4FE9_8DE6_B18030C86529.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrFileTypesSetting struct
type RSOP_PolmkrFileTypesSetting struct {
	*RSOP_PolmkrFolderOptionSetting
}

func NewRSOP_PolmkrFileTypesSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrFileTypesSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFileTypesSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrFileTypesSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrFileTypesSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFileTypesSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}
