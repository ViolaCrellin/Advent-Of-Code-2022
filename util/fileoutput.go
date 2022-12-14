package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"
)

// File contains information about a file that enables us to access its contents.
type File struct {
	path    string
	scanner *bufio.Scanner
	absPath string
	handle  *os.File
}

var (
	// ErrInvalidPathArgs to group all errors relating to cmd argument not pointing us to valid directories.
	ErrInvalidPathArgs = errors.New("path args invalid")

	// ErrInputFileNotLexicographicallySorted describes an error when a supplied input file is not lexicographically sorted.
	ErrInputFileNotLexicographicallySorted = errors.New("input file is not lexicographically sorted")
)

/////////////////////////////
//         BUILDER
/////////////////////////////

// New creates a new file struct
func NewFile(path string) *File {
	return &File{
		path: path,
	}
}

// WithLineScanner opens up a handle to read the file with a bufio scanner.
// Note that this will need to be closed with f.Close() when done.
// I separated this out here so that I can error check/handle that we can open the file before kicking off lots of goroutines.
func (f *File) WithLineScanner() (*File, error) {
	readFile, err := os.Open(f.path)
	if err != nil {
		return f, errors.Wrap(ErrInvalidPathArgs, fmt.Sprintf("could not open input file to scan: %s, error: %s", f.path, err))
	}
	f.scanner = bufio.NewScanner(readFile)
	f.handle = readFile
	return f, nil
}

// WithWriteableFile checks that the file is writeable by writing 0 bytes to it and returns error if not.
// Note that the handle is kept open to the file and upon finishing with it you should call f.Close()
func (f *File) WithWriteableFile() (*File, error) {
	// We write an empty file to check it is possible to write to.
	err := f.Write([]byte{})
	if err != nil {
		return f, errors.Wrap(ErrInvalidPathArgs, fmt.Sprintf("could not open file to write: %s, error: %s", f.path, err))
	}
	return f, nil
}

// Closes any handle to a file.
func (f *File) Close() error {
	return f.handle.Close()
}

/////////////////////////////
//         READ
/////////////////////////////

// Read reads the entire file contents. Not advisable on large files.
func (f *File) Read() ([]byte, error) {
	return os.ReadFile(f.path)
}

/////////////////////////////
//         WRITE
/////////////////////////////

// Append writes appending to a file. If the file does not already exist it will not create it.
func (f *File) Append(content string) error {
	absPath, err := filepath.Abs(f.path)
	if err != nil {
		return err
	}

	if f.handle == nil {
		// If we don't have it open then open it.
		file, err := os.OpenFile(absPath, os.O_WRONLY|os.O_APPEND|os.O_SYNC, 0644)
		if err != nil {
			return err
		}
		f.handle = file
	}

	// Write. Job's a good'un
	_, err = f.handle.WriteString(content)
	return err
}

// Write writes all the given content to a file. Note that it will close the file once done as it considers this a "make the file thus" action.
// If the file does not exist (or path to it) it will create it (and parent directories) with writeonly permissions.
// If it does exist it will overwrite any contents by first truncating the file.
// This function can be used with an empty byte array to guarantee or error that an empty file exists associated with receiver *File that can be written to.
// TODO this needs a better name.
func (f *File) Write(content []byte) error {
	absPath := f.absPath
	var err error
	if len(absPath) == 0 {
		absPath, err = filepath.Abs(f.path)
		if err != nil {
			return err
		}
	}

	err = os.MkdirAll(filepath.Dir(absPath), os.ModePerm)
	if err != nil {
		return err
	}

	writefile, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		return err
	}
	defer writefile.Close()

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return err
	}

	// There is something already here. Burn contents with flame.
	if fileInfo.Size() > 0 {
		// Might be polite to warn somebody somewhere about emptying contents of files? For those who don't like reading docs?
		if err := writefile.Truncate(0); err != nil {
			return err
		}
	}

	// Write. Job's a good'un
	_, err = writefile.Write(content)
	f.absPath = absPath
	return err
}

// Delete deletes the file
func (f *File) Delete() error {
	absPath := f.absPath
	var err error
	if len(absPath) == 0 {
		absPath, err = filepath.Abs(f.path)
		if err != nil {
			return err
		}
	}

	err = os.Remove(absPath)
	if perr, ok := err.(*os.PathError); ok {
		// Meh. File didn't exist anyway.
		if perr.Err.(syscall.Errno) == syscall.ENOENT {
			return nil
		}
	}

	return err
}
