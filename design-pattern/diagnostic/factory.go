package diagnostic

// DiagnosisFactory 추상 팩토리 인터페이스
type DiagnosisFactory interface {
	CreateDiagnosis() Diagnosis
}

// TopographyFactory 생성 팩토리 선언
type TopographyFactory struct {
	Curvature string
	Surface   string
}

func (f TopographyFactory) CreateDiagnosis() Diagnosis {
	return Topography{
		Curvature: f.Curvature,
		Surface:   f.Surface,
	}
}

// AxialLengthFactory 생성 팩토리 선언
type AxialLengthFactory struct {
	Length float64
}

func (f AxialLengthFactory) CreateDiagnosis() Diagnosis {
	return AxialLength{
		Length: f.Length,
	}
}

// VisualFieldFactory 생성 팩토리 선언
type VisualFieldFactory struct {
	Count   int
	Success int
	Fail    int
}

func (f VisualFieldFactory) CreateDiagnosis() Diagnosis {
	return VisualField{
		Count:   f.Count,
		Success: f.Success,
		Fail:    f.Fail,
	}
}
