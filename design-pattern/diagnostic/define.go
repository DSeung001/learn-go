package diagnostic

import "fmt"

type Diagnosis interface {
	GetType() string
	GetDetails() string
}

// Diagnosis 를 구현하는 Topography
type Topography struct {
	Curvature string
	Surface   string
}

func (t Topography) GetType() string {
	return "Topography"
}

func (t Topography) GetDetails() string {
	return "Curvature: " + t.Curvature + ", Surface: " + t.Surface
}

// Diagnosis 를 구현하는 AxialLength
type AxialLength struct {
	Length float64
}

func (a AxialLength) GetType() string {
	return "AxialLength"
}

func (a AxialLength) GetDetails() string {
	return "Length: " + fmt.Sprintf("%.2f mm", a.Length)
}

// Diagnosis 를 구현하는 AxialLength
type VisualField struct {
	Count   int
	Success int
	Fail    int
}

func (v VisualField) GetType() string {
	return "VisualField"
}

func (v VisualField) GetDetails() string {
	return fmt.Sprintf("Count: %d", v.Count)
}
