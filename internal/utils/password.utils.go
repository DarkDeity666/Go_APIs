package utils

import "golang.org/x/crypto/bcrypt"


//hashing password
func HashPassword(password string)(string, error){
	hashedPassword,err := bcrypt.GenerateFromPassword(
		//we convert string into byte bcz it perform its operation on byte
		[]byte(password),
		//DefaulCost is a salt value and by default its value is 10
		bcrypt.DefaultCost,
	)
	
	if err != nil{
		return "",err
	}
	return string(hashedPassword),nil
}


//for verifing the password
func CheckPassword(password, hash string) bool {
	err:= bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
	//if password matches it will return true and if not then it will return false
	return err == nil
}
