package latex

type ResumeClassic struct {
	templateName string
	Person       PersonClassic
	Education    EducationClassic
	Positions    []PositionClassic
}

type PersonClassic struct {
	Name     string
	Position string
	Contacts []ContactClassic
}

type ContactClassic struct {
	Title string
	Ref   string
}

type EducationClassic struct {
	Instituion string
	Specialty  string
	StartDate  string
	EndDate    string
}

type ExperienceClassic struct {
}

type PositionClassic struct {
	Position    string
	Company     string
	Location    string
	StartDate   string
	EndDate     string
	Description []string
}
