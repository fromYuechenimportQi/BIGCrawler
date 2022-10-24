package src

import "gorm.io/gorm"

type NeedInfo struct {
	gorm.Model
	Primayid             string `gorm:"primaryKey" gorm:"unique" gorm:"uniqueIndex"`
	Accession            string
	ScientificName       string
	Bioproject           string
	Biosample            string
	ReleaseDate          string
	AssemblyLevel        string
	GenomeRepresentation string
	DNAFile              string
	GFFFile              string
	RNAFile              string
	ProteinFile          string
}
