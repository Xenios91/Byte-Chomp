package bytechomp

type sectionOfInterest string

const (
	roData    sectionOfInterest = ".rodata"
	gopclntab sectionOfInterest = ".gopclntab"
)

var sectionsOfInterest = []sectionOfInterest{
	roData,
	gopclntab,
}
