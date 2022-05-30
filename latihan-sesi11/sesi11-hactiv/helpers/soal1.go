package helpers

func IsAllowedToWork(age int, gender string) bool {
	allowed := false
	if (gender == "pria") && (age > 17 && age < 60) {
		allowed = true
	} else if (gender == "wanita") && (age > 19 && age < 60) {
		allowed = true
	}
	return allowed
}
