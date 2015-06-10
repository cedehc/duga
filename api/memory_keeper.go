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

type FakeDb map[string]ComicItem

func NewFakeDb() *FakeDb {
	var tmp FakeDb = make(map[string]ComicItem)

	return &tmp
}

func (p *FakeDb) FakeData() {
	(*p)["xkcd"] = ComicItem{"xkcd", "XKCD", "xkcd.com"}
	(*p)["abstrusegoose"] = ComicItem{"abstrusegoose", "Abstruse Goose", "abstrusegoose.com"}
}

func (p *FakeDb) Fetch(uid string) (ComicItem, error) {
	tmp, ok := (*p)[uid]
	var err error

	if !ok {
		tmp = ComicItem{}
		err = &ComicError{"invalid id"}
	}

	return tmp, err
}

func (p *FakeDb) Count() int {
	return len(*p)
}
