package api

import (
	"github.com/coocood/jas"
)

type ComicId struct {
	ComicKeeper Keeper
}

type Comic struct {
	Title string
	Link  string
}

func (p *ComicId) Get(ctx *jas.Context) {
	id := int(ctx.Id)

	tmp, err := p.ComicKeeper.Fetch(id)

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

type FakeDb []Comic

type Keeper interface {
	Count() int
	Fetch(id int) (Comic, error)
}

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
