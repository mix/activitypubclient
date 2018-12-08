package activitystreams

// ItemCollection represents an array of items
type ItemCollection []Item

// Item struct
type Item ObjectOrLink

// GetID returns the ObjectID corresponding to ItemCollection
func (i ItemCollection) GetID() *ObjectID {
	return nil
}

// GetLink returns the empty IRI
func (i ItemCollection) GetLink() IRI {
	return IRI("")
}

// GetType returns the ItemCollection's type
func (i ItemCollection) GetType() ActivityVocabularyType {
	return i.First().GetType()
}

// IsLink returns false for an ItemCollection object
func (i ItemCollection) IsLink() bool {
	return false
}

// IsObject returns true for a ItemCollection object
func (i ItemCollection) IsObject() bool {
	return false
}

// Append facilitates adding elements to Item arrays
// and ensures ItemCollection implements the Collection interface
func (i *ItemCollection) Append(o Item) error {
	oldLen := len(*i)
	d := make(ItemCollection, oldLen+1)
	for k, it := range *i {
		d[k] = it
	}
	d[oldLen] = o
	*i = d
	return nil
}

// First returns the ObjectID corresponding to ItemCollection
func (i ItemCollection) First() Item {
	if len(i) == 0 {
		return nil
	}
	return i[0]
}

// Collection returns the current object as collection interface
func (i *ItemCollection) Collection() CollectionInterface {
	return i
}