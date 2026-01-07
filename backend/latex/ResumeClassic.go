package latex

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
