// Copyright 2018 The go-saloon Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

func SendWelcome() error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Welcome"
	m.From = ""
	m.To = []string{}
	err := m.AddBody(r.HTML("welcome.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
