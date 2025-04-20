package fd

import (
	"os"
	"runtime/pprof"
)

// `pprof.NewProfile("fd.inuse")`
// - 이 프로파일은 현재 사용 중인 파일 디스크립터(FD) 개수를 추적하는 용도로 사용됨
// - `pprof.Lookup("fd.inuse")`로 다른 패키지에서도 조회 가능
// - 전역 변수로 설정하면 특정 파일에서만 사용하기 어려우므로 일반적으로는 지양함
var fdProfile = pprof.NewProfile("fd.inuse")

// File 구조체는 os.File을 래핑하여 pprof와 연동되도록 설계됨
type File struct {
	*os.File
}

// 파일을 열고 프로파일링에 추가하는 함수
func Open(name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	// `Add(f, 2)`에서 2는 스택 프레임을 몇 개 건너뛸지를 결정함
	// 1: Open() 자신
	// 2: Open()을 호출한 상위 함수
	// 프로파일러에 Open 상위 함수를 기록하기 위한 용도
	fdProfile.Add(f, 2)
	return &File{File: f}, nil
}

// 파일을 닫을 때 pprof 프로파일에서 제거하는 함수
func (f *File) Close() error {
	// 추적하던 파일이 닫히는 순간, 프로파일에서 제거
	defer fdProfile.Remove(f.File)
	return f.File.Close()
}

// 현재 열려 있는 파일 디스크립터 목록을 pprof 파일로 저장하는 함수
func Write(profileOutPath string) error {
	out, err := os.Create(profileOutPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 현재 pprof 프로파일 기록 (0: 기본 출력 레벨)
	if err := fdProfile.WriteTo(out, 0); err != nil {
		return err
	}
	return nil
}
