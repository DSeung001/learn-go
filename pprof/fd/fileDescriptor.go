package fd

import (
	"os"
	"runtime/pprof"
)

// 파일 디스크립터 : 리눅스, 유닉스 계열에서 프로세스가 파일을 다룰 때 사용하는 개념으로 프로세스에서 특정 파일에 접근할 때 사용하는 추상적인 값

// profile의 이름 설정, fd.inuse인 이유는 현재 사용중인 파일 디스크립터를 추적하겠다는 의미를 내포
// 전역으로 생성해서, 타 패키지서 pprof.Lookup("fd.inuse")로 사용 가능
// 이 예제 서만 전역으로 지정 했을 뿐 일반 적으로는 권하지 않는 방법
var fdProfile = pprof.NewProfile("fd.inuse")

// File is a wrapper on os.file that tracks file descriptor lifetime
type File struct {
	*os.File
}

// Open opens a file and tracks it in the `fd` profile.
func Open(name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	// Add로 객체 정보 저장, Open 함수를 샘플 생성에 대한 참조로 사용하기 위해 2개의 스택 프레임을 건너 뜀
	fdProfile.Add(f, 2)
	return &File{File: f}, err
}

// Close closes files and updates profile.
func (f *File) Close() error {
	// 추적하던 파일이 닫히는 순간, 객체를 제거
	defer fdProfile.Remove(f.File)
	return f.File.Close()
}

// Write saves the profiles of the currently open file
// descriptors into a file in pprof format
func Write(profileOutPath string) error {
	out, err := os.Create(profileOutPath)
	if err != nil {
		return err
	}
	// GO의 프로파일러는 WriteTo를 통해 전체 pprof 파일 내용을 지정한 대상 위치에 바이트 형식을 저장할 수 있다
	// 그러나 여기서는 파일 형태를 기록하기 위해 벼롣로 Write를 추가로 구현
	if err := fdProfile.WriteTo(out, 0); err != nil {
		_ = out.Close()
		return err
	}
	return out.Close()
}
