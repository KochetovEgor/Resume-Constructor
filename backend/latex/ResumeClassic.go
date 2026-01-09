package latex

import "net/url"

const resumeClassicName = "resume_classic.tex"

type ResumeClassic struct {
	Person     *PersonClassic    `json:"person"`
	Education  *EducationClassic `json:"education"`
	Experience []PositionClassic `json:"experience"`
	Courses    []CourseClassic   `json:"courses"`
	Projects   []ProjectClassic  `json:"projects"`
	Skills     []SkillClassic    `json:"skills"`
	AboutMe    string            `json:"about_me"`
}

type PersonClassic struct {
	Name     string           `json:"name"`
	Position string           `json:"position"`
	Contacts []ContactClassic `json:"contacts"`
}

type ContactClassic struct {
	Title string `json:"title"`
	Ref   string `json:"ref"`
}

type EducationClassic struct {
	Institution string `json:"institution"`
	Specialty   string `json:"specialty"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type PositionClassic struct {
	Position    string   `json:"position"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Description []string `json:"description"`
}

type CourseClassic struct {
	Title  string `json:"title"`
	Period string `json:"period"`
	Author string `json:"author"`
}

type ProjectClassic struct {
	Title       string   `json:"title"`
	Stack       string   `json:"stack"`
	Ref         string   `json:"ref"`
	Description []string `json:"description"`
}

type SkillClassic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *ResumeClassic) TemplateName() string {
	return resumeClassicName
}

func (r *ResumeClassic) Escape() *ResumeClassic {
	var escPerson *PersonClassic
	if r.Person != nil {
		escPerson = &PersonClassic{
			Name:     EscapeLaTeX(r.Person.Name),
			Position: EscapeLaTeX(r.Person.Position),
		}
		escPersonContacts := make([]ContactClassic, len(r.Person.Contacts))
		for i, c := range r.Person.Contacts {
			escPersonContacts[i] = ContactClassic{
				Title: EscapeLaTeX(c.Title),
				Ref:   url.PathEscape(c.Ref),
			}
		}
		escPerson.Contacts = escPersonContacts
	}

	var escEducation *EducationClassic
	if r.Education != nil {
		escEducation = &EducationClassic{
			Institution: EscapeLaTeX(r.Education.Institution),
			Specialty:   EscapeLaTeX(r.Education.Specialty),
			Location:    EscapeLaTeX(r.Education.Location),
			StartDate:   EscapeLaTeX(r.Education.StartDate),
			EndDate:     EscapeLaTeX(r.Education.EndDate),
		}
	}

	escExperience := make([]PositionClassic, len(r.Experience))
	for i, ex := range r.Experience {
		escPosition := PositionClassic{
			Position:  EscapeLaTeX(ex.Position),
			Company:   EscapeLaTeX(ex.Company),
			Location:  EscapeLaTeX(ex.Location),
			StartDate: EscapeLaTeX(ex.StartDate),
			EndDate:   EscapeLaTeX(ex.EndDate),
		}
		escPositionDescription := make([]string, len(ex.Description))
		for j, desc := range ex.Description {
			escPositionDescription[j] = EscapeLaTeX(desc)
		}
		escPosition.Description = escPositionDescription
		escExperience[i] = escPosition
	}

	escCourses := make([]CourseClassic, len(r.Courses))
	for i, c := range r.Courses {
		escCourse := CourseClassic{
			Title:  EscapeLaTeX(c.Title),
			Period: EscapeLaTeX(c.Period),
			Author: EscapeLaTeX(c.Author),
		}
		escCourses[i] = escCourse
	}

	escProjects := make([]ProjectClassic, len(r.Projects))
	for i, p := range r.Projects {
		escProject := ProjectClassic{
			Title: EscapeLaTeX(p.Title),
			Stack: EscapeLaTeX(p.Stack),
			Ref:   url.PathEscape(p.Ref),
		}
		escProjectDescription := make([]string, len(p.Description))
		for j, d := range p.Description {
			escProjectDescription[j] = EscapeLaTeX(d)
		}
		escProject.Description = escProjectDescription
		escProjects[i] = escProject
	}

	escSkills := make([]SkillClassic, len(r.Skills))
	for i, s := range r.Skills {
		escSkill := SkillClassic{
			Name:        EscapeLaTeX(s.Name),
			Description: EscapeLaTeX(s.Description),
		}
		escSkills[i] = escSkill
	}

	escAboutMe := EscapeLaTeX(r.AboutMe)

	return &ResumeClassic{
		Person:     escPerson,
		Education:  escEducation,
		Experience: escExperience,
		Courses:    escCourses,
		Projects:   escProjects,
		Skills:     escSkills,
		AboutMe:    escAboutMe,
	}
}
