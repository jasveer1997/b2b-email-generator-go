package storage

var domainData = []StorageDomain{
	{
		DomainName: "babbel.com",
		EmailPref:  "FIRST_NAME_INITIAL_LAST_NAME",
	},
	{
		DomainName: "linkedin.com",
		EmailPref:  "FIRST_NAME_LAST_NAME",
	},
	{
		DomainName: "google.com",
		EmailPref:  "FIRST_NAME_LAST_NAME",
	},
}

var userData = []StorageUser{
	{
		FirstName: "Jane",
		LastName:  "Doe",
		Domain:    "babbel.com",
		Email:     "jdoe@babbel.com",
	},
	{
		FirstName: "Jay",
		LastName:  "Arun",
		Domain:    "linkedin.com",
		Email:     "jayarun@linkedin.com",
	},
	{
		FirstName: "David",
		LastName:  "Stein",
		Domain:    "google.com",
		Email:     "davidstein@google.com",
	},
	{
		FirstName: "Mat",
		LastName:  "Lee",
		Domain:    "google.com",
		Email:     "matlee@google.com",
	},
	{
		FirstName: "Marta",
		LastName:  "Dahl",
		Domain:    "babbel.com",
		Email:     "mdahl@babbel.com",
	},
	{
		FirstName: "Vanessa",
		LastName:  "Boom",
		Domain:    "babbel.com",
		Email:     "vboom@babbel.com",
	},
}
