package activitypub

import (
	as "github.com/go-ap/activitystreams"
	"reflect"
	"testing"
)

func TestLikedNew(t *testing.T) {
	l := LikedNew()

	id := as.ObjectID("liked")
	if l.ID != id {
		t.Errorf("%T should be initialized with %q as %T", l, id, id)
	}
	if len(l.Name) != 0 {
		t.Errorf("%T should be initialized with 0 length Name", l)
	}
	if len(l.Content) != 0 {
		t.Errorf("%T should be initialized with 0 length Content", l)
	}
	if len(l.Summary) != 0 {
		t.Errorf("%T should be initialized with 0 length Summary", l)
	}
	if l.TotalItems != 0 {
		t.Errorf("%T should be initialized with 0 TotalItems", l)
	}
}

func TestLikedCollection_GetID(t *testing.T) {
	l := LikedCollection{}
	if *l.GetID() != "" {
		t.Errorf("%T should be initialized with empty %T", l, l.GetID())
	}
	id := as.ObjectID("test_out_stream")
	l.ID = id
	if *l.GetID() != id {
		t.Errorf("%T should have %T as %q", l, id, id)
	}
}

func TestLikedCollection_GetType(t *testing.T) {
	l := LikedCollection{}

	if l.GetType() != "" {
		t.Errorf("%T should be initialized with empty %T", l, l.GetType())
	}

	l.Type = as.OrderedCollectionType
	if l.GetType() != as.OrderedCollectionType {
		t.Errorf("%T should have %T as %q", l, l.GetType(), as.OrderedCollectionType)
	}
}

func TestLikedCollection_Append(t *testing.T) {
	l := LikedCollection{}

	val := as.Object{ID: as.ObjectID("grrr")}

	l.Append(val)
	if l.TotalItems != 1 {
		t.Errorf("%T should have exactly an element, found %d", l, l.TotalItems)
	}
	if !reflect.DeepEqual(l.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", l, l.OrderedItems, val.ID)
	}
}

func TestLiked_Append(t *testing.T) {
	l := LikedNew()

	val := as.Object{ID: as.ObjectID("grrr")}

	l.Append(val)
	if l.TotalItems != 0 {
		t.Errorf("%T should have exactly an element, found %d", l, l.TotalItems)
	}
	if !reflect.DeepEqual(l.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", l, l.OrderedItems, val.ID)
	}
}

func TestLiked_Collection(t *testing.T) {
	t.Skipf("TODO")
}

func TestLiked_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestLiked_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLiked_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestLiked_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLiked_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestLikedCollection_Collection(t *testing.T) {
	t.Skipf("TODO")
}

func TestLikedCollection_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLikedCollection_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLikedCollection_IsObject(t *testing.T) {
	t.Skipf("TODO")
}
