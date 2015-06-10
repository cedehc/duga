/*
   DUGA is the Dynamic Universal Gomic API
   Copyright (C) 2015 Igor Almeida

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as
   published by the Free Software Foundation, either version 3 of the
   License, or (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/coocood/jas"

	"duca/api"
)

func main() {
	db := api.NewFakeDb()
	db.FakeData()

	comics := new(api.Comic)
	comics.ComicKeeper = db

	router := jas.NewRouter(comics)
	router.BasePath = "/"

	// show paths
	fmt.Println(router.HandledPaths(true))

	http.Handle(router.BasePath, router)
	http.ListenAndServe(":8080", nil)
}
