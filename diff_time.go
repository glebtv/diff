/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package diff

import (
	"reflect"
	"time"
)

func (cl *Changelog) diffTime(path []string, a, b reflect.Value) error {
	if a.Kind() == reflect.Invalid {
		cl.add(CREATE, path, nil, b.Interface())
		return nil
	}

	if b.Kind() == reflect.Invalid {
		cl.add(DELETE, path, a.Interface(), nil)
		return nil
	}

	if a.Kind() != b.Kind() {
		return ErrTypeMismatch
	}
	ai := a.Interface()
	bi := b.Interface()
	at := ai.(time.Time)
	bt := bi.(time.Time)
	if duration := at.Sub(bt); duration != 0 {
		cl.add(UPDATE, path, a.Interface(), b.Interface())
	}
	return nil
}
