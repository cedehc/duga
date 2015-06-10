/*
   This file is part of DUGA.

   DUGA is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as
   published by the Free Software Foundation, either version 3 of the
   License, or (at your option) any later version.

   DUGA is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with DUGA.  If not, see <http://www.gnu.org/licenses/>.
*/
package api

import (
	"github.com/coocood/jas"
)

type Keeper interface {
	Count() int
	Fetch(uid string) (ComicItem, error)
}

type Comic struct {
	ComicKeeper Keeper
}

func (*Comic) Gap() string {
	return ":uid"
}

type ComicItem struct {
	Uid   string
	Title string
	Link  string
}

func (p *Comic) Get(ctx *jas.Context) {
	uid := ctx.GapSegment("")

	tmp, err := p.ComicKeeper.Fetch(uid)

	if err == nil {
		ctx.Data = tmp
	} else {
		ctx.Error = jas.NewRequestError(err.Error())
	}
}

type ComicError struct {
	Msg string
}

func (e *ComicError) Error() string {
	return e.Msg
}
