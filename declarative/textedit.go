// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package declarative

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type TextEdit struct {
	// Window

	Background         Brush
	ContextMenuItems   []MenuItem
	Enabled            Property
	Font               Font
	MaxSize            Size
	MinSize            Size
	Name               string
	OnKeyDown          walk.KeyEventHandler
	OnKeyPress         walk.KeyEventHandler
	OnKeyUp            walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Persistent         bool
	RightToLeftReading bool
	ToolTipText        Property
	Visible            Property

	// Widget

	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	Row                int
	RowSpan            int
	StretchFactor      int

	// TextEdit

	Alignment     Alignment1D
	AssignTo      **walk.TextEdit
	HScroll       bool
	MaxLength     int
	OnTextChanged walk.EventHandler
	ReadOnly      Property
	Text          Property
	VScroll       bool
}

func (te TextEdit) Create(builder *Builder) error {
	var style uint32
	if te.HScroll {
		style |= win.WS_HSCROLL
	}
	if te.VScroll {
		style |= win.WS_VSCROLL
	}

	w, err := walk.NewTextEditWithStyle(builder.Parent(), style)
	if err != nil {
		return err
	}

	return builder.InitWidget(te, w, func() error {
		if err := w.SetAlignment(walk.Alignment1D(te.Alignment)); err != nil {
			return err
		}

		if te.MaxLength > 0 {
			w.SetMaxLength(te.MaxLength)
		}

		if te.OnTextChanged != nil {
			w.TextChanged().Attach(te.OnTextChanged)
		}

		if te.AssignTo != nil {
			*te.AssignTo = w
		}

		return nil
	})
}
