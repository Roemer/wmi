// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SecurityPolicySection struct
type SecurityPolicySection struct {
	*ConfigurationSectionWithCollection

	//
	SecurityPolicy []TrustLevel
}

func NewSecurityPolicySectionEx1(instance *cim.WmiInstance) (newInstance *SecurityPolicySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SecurityPolicySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewSecurityPolicySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SecurityPolicySection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SecurityPolicySection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetSecurityPolicy sets the value of SecurityPolicy for the instance
func (instance *SecurityPolicySection) SetPropertySecurityPolicy(value []TrustLevel) (err error) {
	return instance.SetProperty("SecurityPolicy", (value))
}

// GetSecurityPolicy gets the value of SecurityPolicy for the instance
func (instance *SecurityPolicySection) GetPropertySecurityPolicy() (value []TrustLevel, err error) {
	retValue, err := instance.GetProperty("SecurityPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TrustLevel)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TrustLevel is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TrustLevel(valuetmp))
	}

	return
}
