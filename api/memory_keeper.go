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

type FakeDb []Comic

func NewFakeDb() *FakeDb {
	var tmp FakeDb = make([]Comic, 0, 10)

	return &tmp
}

func (p *FakeDb) FakeData() {
	*p = append(*p, Comic{"xkcd", "xkcd.com"})
	*p = append(*p, Comic{"abstrusegoose", "abstrusegoose.com"})
}

func (p *FakeDb) Fetch(id int) (Comic, error) {
	if id < p.Count() {
		return (*p)[id], nil
	}

	return Comic{}, &ComicError{"invalid id"}
}

func (p *FakeDb) Count() int {
	return len(*p)
}
