package db

import "ecommerce/model"

func GetPass(email string) string {
	query := "SELECT PASSWORD from users where email ='" + email + "';"

	var password string
	Db.QueryRow(query).Scan(&password)

	return password

}
func INSERT(name, email, pass string) error {

	query := "INSERT into users (name,email,password) VALUES ('" + name + "','" + email + "', '" + pass + "');"
	_, err := Db.Exec(query)
	return err
}
func GetUser(email string, usrchan chan model.User) {

	query := "SELECT id, name, email FROM users WHERE email = '" + email + "';"
	var user model.User

	err = Db.QueryRow(query).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		usrchan <- model.User{}
		close(usrchan)
		return
	}

	usrchan <- user
	close(usrchan)
}
