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

// SourceListenerSettings struct
type SourceListenerSettings struct {
	*EmbeddedObject

	//
	Listeners []ListenerElement
}

func NewSourceListenerSettingsEx1(instance *cim.WmiInstance) (newInstance *SourceListenerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SourceListenerSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewSourceListenerSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SourceListenerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SourceListenerSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetListeners sets the value of Listeners for the instance
func (instance *SourceListenerSettings) SetPropertyListeners(value []ListenerElement) (err error) {
	return instance.SetProperty("Listeners", (value))
}

// GetListeners gets the value of Listeners for the instance
func (instance *SourceListenerSettings) GetPropertyListeners() (value []ListenerElement, err error) {
	retValue, err := instance.GetProperty("Listeners")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ListenerElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ListenerElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ListenerElement(valuetmp))
	}

	return
}
