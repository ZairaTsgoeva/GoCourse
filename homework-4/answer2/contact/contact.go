package contact

import "strings"

type Book []contact

type contact struct {
	name string
	phones []int
	emails []string
}

func CreateBook() *Book {
	return &Book{}
}

func (b *Book) AddContact(c *contact) {
	*b = append(*b, *c)
}

func (b *Book) Len() int {
	return len(*b)
}
func (b *Book) Less(i ,j int) bool {
	return strings.Compare((*b)[i].name, (*b)[j].name) < 0
}
func (b *Book) Swap(i ,j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func CreateContact(name string) *contact {
	return &contact{name: name}
}

func (c *contact) GetName() string {
	return c.name
}

func (c *contact) AddPhone(phone int) {
	c.phones = append(c.phones, phone)
}

func (c *contact) addEmail(email string) {
	c.emails = append(c.emails, email)
}