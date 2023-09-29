package fixtures

import (
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
)

var emailAdress, _ = value_objects.NewEmailAddress("matheus@email.com")
var phone, _ = value_objects.NewPhone("559999048223")
var studentType = entities.NewUserType("student", "Student")

var StudentUser, _ = entities.NewUser(
	"Matheus",
	emailAdress,
	phone,
	studentType,
	"12345678",
)
