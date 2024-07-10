package class

type Employed struct {
	Id    int
	name  string
	phone string
	email string
}

func (e *Employed) SetEmployed(name, phone, email string) {
	e.SetName(name)
	e.SetPhone(phone)
	e.SetEmail(email)
}

func (e *Employed) GetId() int {
	return e.Id
}

func (e *Employed) GetName() string {
	return e.name
}

func (e *Employed) GetPhone() string {
	return e.phone
}

func (e *Employed) GetEmail() string {
	return e.email
}

func (e *Employed) SetId(id int) {
	e.Id = id
}

func (e *Employed) SetName(name string) {
	e.name = name
}

func (e *Employed) SetPhone(phone string) {
	e.phone = phone
}

func (e *Employed) SetEmail(email string) {
	e.email = email
}
