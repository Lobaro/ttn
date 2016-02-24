// Copyright © 2015 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package handlers

import (
	core "github.com/TheThingsNetwork/ttn/refactor"
	. "github.com/TheThingsNetwork/ttn/refactor/adapters/http"
	"github.com/TheThingsNetwork/ttn/utils/errors"
	"github.com/brocaar/lorawan"
)

// type pubSubRegistration implements the core.Registration interface
type pubSubRegistration struct {
	recipient HttpRecipient
	appEUI    lorawan.EUI64
	nwkSKey   lorawan.AES128Key
	devEUI    lorawan.EUI64
}

// Recipient implements` the core.Registration interface
func (r pubSubRegistration) Recipient() core.Recipient {
	return r.recipient
}

// AppEUI implements the core.Registration interface
func (r pubSubRegistration) AppEUI() (lorawan.EUI64, error) {
	return r.appEUI, nil
}

// AppSKey implements the core.Registration interface
func (r pubSubRegistration) AppSKey() (lorawan.AES128Key, error) {
	return lorawan.AES128Key{}, errors.New(errors.Implementation, "AppSKey noy supported on pubsub registration")
}

// DevEUI implements the core.Registration interface
func (r pubSubRegistration) DevEUI() (lorawan.EUI64, error) {
	return r.devEUI, nil
}

// NwkSKey implement the core.Registration interface
func (r pubSubRegistration) NwkSKey() (lorawan.AES128Key, error) {
	return r.nwkSKey, nil
}
