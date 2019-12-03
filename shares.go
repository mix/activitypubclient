package activitypub

type (
	// SharesCollection is a list of all Announce activities with this object as the object property,
	// added as a side effect. The shares collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication
	// is given.
	SharesCollection = Shares

	// Shares is a type alias for an Ordered Collection
	Shares OrderedCollection
)

// SharesNew initializes a new Shares
func SharesNew() *Shares {
	id := ObjectID("Shares")

	i := Shares{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an Shares
func (o *Shares) Append(ob Item) error {
	o.OrderedItems = append(o.OrderedItems, ob)
	o.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to Shares
func (o Shares) GetID() ObjectID {
	return o.Collection().GetID()
}

// GetLink returns the IRI corresponding to the current Shares object
func (o Shares) GetLink() IRI {
	return IRI(o.ID)
}

// GetType returns the Shares's type
func (o Shares) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink returns false for an Shares object
func (o Shares) IsLink() bool {
	return false
}

// IsObject returns true for a Shares object
func (o Shares) IsObject() bool {
	return true
}

// UnmarshalJSON
func (o *Shares) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	c := OrderedCollection(*o)
	err := c.UnmarshalJSON(data)

	*o = Shares(c)

	return err
}

// Collection returns the underlying Collection type
func (o Shares) Collection() CollectionInterface {
	c := OrderedCollection(o)
	return &c
}
