package utils

type Doctor struct {
	Name        string
	Email       string
	Address     string
	ZipCode     string
	City        string
	Country     string
	Specialty   string
	Description string
}

func GetDoctors() []Doctor {
	return []Doctor{
		{
			Name:        "Dr. John Doe",
			Email:       "john.doe@example.org",
			Address:     "1234 Main St",
			ZipCode:     "12345",
			City:        "Springfield",
			Country:     "USA",
			Specialty:   "General Practitioner",
			Description: "Dr. John Doe is a general practitioner with over 20 years of experience.",
		},
	}
}
