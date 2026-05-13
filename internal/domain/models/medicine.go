package models

type Medicine struct {
	ID                int64  `json:"id"`
	Article           int64  `json:"article"`
	Name              string `json:"name"`
	ActiveSubstance   string `json:"active_substance"`
	Form              string `json:"form"`
	Contraindications string `json:"contraindications"`
	Prescription      bool   `json:"prescription"`
}
