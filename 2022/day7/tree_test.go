package day7

import (
	"reflect"
	"testing"
)

func TestCDINTO(t *testing.T) {
	fs := NewFS()

	fs.CDInto("dir1")
	fs.CDInto("dir2")

	expectedResult := &Stack{path: []string{"/", "dir1", "dir2"}, length: 3}
	result := fs.wd

	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Error - CDInto(): expected: %v, got: %v", expectedResult, result)
	}
}

func TestCDTOROOT(t *testing.T) {
	fs := NewFS()

	fs.CDInto("dir1")
	fs.CDInto("dir2")
	fs.CDInto("dir3")
	fs.CDInto("dir4")
	fs.CDInto("/")

	expectedResult := &Stack{path: []string{"/"}, length: 1}
	result := fs.wd

	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Error - CDInto(\"/\"): expected: %v, got: %v", expectedResult, result)
	}

	if _, exists := fs.root.dirs["dir1"].dirs["dir2"].dirs["dir3"].dirs["dir4"]; !exists {
		t.Fatalf("Error - CDInto(\"/\"): file system state was not preserved")
	}
}

func TestCDBACK(t *testing.T) {
	fs := NewFS()

	fs.CDInto("dir1")
	fs.CDInto("dir2")
	fs.CDInto("dir3")
	fs.CDInto("dir4")
	fs.CDBack()

	expectedResult := &Stack{path: []string{"/", "dir1", "dir2", "dir3"}, length: 4}
	result := fs.wd

	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Error - CDBack(): expected: %v, got: %v", expectedResult, result)
	}

	if _, exists := fs.root.dirs["dir1"].dirs["dir2"].dirs["dir3"].dirs["dir4"]; !exists {
		t.Fatalf("Error - CDBack(): file system state was not preserved")
	}
}

func TestState(t *testing.T) {
	fs := NewFS()

	fs.CDInto("dir1")
	fs.CDInto("dir2")
	fs.CDInto("dir3")
	fs.CDInto("dir4")
	fs.CDBack()
	fs.CDBack()
	fs.CDInto("dir5")
	fs.CDInto("dir6")
	fs.CDInto("/")
	fs.CDInto("dir7")
	fs.CDBack()

	expectedResult := &Stack{path: []string{"/"}, length: 1}
	result := fs.wd

	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Error: expected: %v, got: %v", expectedResult, result)
	}

	if _, exists := fs.root.dirs["dir1"].dirs["dir2"].dirs["dir3"].dirs["dir4"]; !exists {
		t.Fatalf("Error: file system state was not preserved")
	}

	if _, exists := fs.root.dirs["dir1"].dirs["dir2"].dirs["dir5"].dirs["dir6"]; !exists {
		t.Fatalf("Error: file system state was not preserved")
	}

	if _, exists := fs.root.dirs["dir7"]; !exists {
		t.Fatalf("Error: file system state was not preserved")
	}
}

func TestGETFILE(t *testing.T) {
	fs := NewFS()

	file1 := file{
		name: "file1",
		size: 10,
	}

	file2 := file{
		name: "file2",
		size: 20,
	}

	file3 := file{
		name: "file3",
		size: 30,
	}

	file4 := file{
		name: "file4",
		size: 40,
	}

	fs.CDInto("dir1")
	fs.GetFile(file1.name, file1.size)
	fs.GetFile(file2.name, file2.size)
	fs.CDBack()
	fs.CDInto("dir2")
	fs.GetFile(file3.name, file3.size)
	fs.GetFile(file4.name, file4.size)
	fs.CDInto("/")

	expectedResult := 30
	result := fs.currDir.dirs["dir1"].size

	expectedResult2 := 70
	result2 := fs.currDir.dirs["dir2"].size

	if expectedResult != result {
		t.Fatalf("Error - GetFile(\"dir1\"): expected: %v, got: %v", expectedResult, result)
	}

	if expectedResult2 != result2 {
		t.Fatalf("Error - GetFile(\"dir1\"): expected: %v, got: %v", expectedResult2, result2)
	}
}
