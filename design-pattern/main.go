package main

import (
	"design-pattern.com/diagnostic"
	"fmt"
)

type Patient struct {
	Name      string
	Diagnoses []diagnostic.Diagnosis
}

func (p *Patient) AddDiagnosis(d diagnostic.Diagnosis) {
	p.Diagnoses = append(p.Diagnoses, d)
}

func (p Patient) PrintDiagnoses() {
	fmt.Println("Patient Name:", p.Name)
	for _, d := range p.Diagnoses {
		fmt.Printf("- [%s] %s\n", d.GetType(), d.GetDetails())
	}
}

func main() {
	// 환자 생성
	patient := Patient{Name: "Seung"}

	topography := diagnostic.TopographyFactory{
		Curvature: "Steep",
		Surface:   "Smooth",
	}.CreateDiagnosis()
	axialLength := diagnostic.AxialLengthFactory{
		Length: 23.45,
	}.CreateDiagnosis()
	visualField := diagnostic.VisualFieldFactory{
		Count:   12,
		Success: 8,
		Fail:    4,
	}.CreateDiagnosis()

	// 환자에 진단 정보 추가

	// 환자에 진단을 넣을 때와
	patient.AddDiagnosis(topography)
	patient.AddDiagnosis(axialLength)
	patient.AddDiagnosis(visualField)

	// 데이터를 출력할 때 기존 내용을 전혀 몰라도 됨 => 추상 팩토리 사용 이유
	patient.PrintDiagnoses()
}
