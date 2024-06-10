// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors struct {
	*Win32_PerfRawData

	//
	ErrorsPersec uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrorsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrorsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetErrorsPersec sets the value of ErrorsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors) SetPropertyErrorsPersec(value uint64) (err error) {
	return instance.SetProperty("ErrorsPersec", (value))
}

// GetErrorsPersec gets the value of ErrorsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLErrors) GetPropertyErrorsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
