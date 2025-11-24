// Нет классов, но можно описывать типы объектов на основе struct:
package user

type User struct { // struct embedding
    login, email string
    UserRole			// composition via struct embedding: описать anonymous field как имя типа
}

// К такому типу можно присоединить поведение с помощью методов:
func (u User) Login() string            { return u.login }
func (u User) Email() string            { return u.email }
func (u *User) SetEmail(mailbox string) { u.email = mailbox }

// Это не конструктор, а обычная функция, которую можно назвать New или NewUser
func New(l, e string) (u User) {
    u = User{login: l, email: e, UserRole: UserRole{"Гость", 9} }
    return u
}

type UserRole struct {
    name string
    id   int
}
func (r UserRole) Role() string { return r.name }

