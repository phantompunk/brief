package command

import (
	"io"
	"os"
)

type Editor struct {
	Args []string
	IDE  string
}

func NewEditor() Editor {
	return Editor{Args: []string{"vim"}}
}

func saveFile(path string) (f *os.File, err error) {
	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	return
}

func copyContents(f *os.File, r io.Reader) error {
	if _, err := io.Copy(f, r); err != nil {
		os.Remove(f.Name())
		return err
	}
	return nil
}

func (e *Editor) SaveTemplateFile(path string, r io.Reader) (f *os.File, err error) {
	f, err = saveFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f, copyContents(f, r)
}

// func (e *Editor) LaunchTempFile(prefix, suffix string, r io.Reader) ([]byte, string, error) {
// 	f, err := tempFile(prefix, suffix)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer f.Close()
// 	path := f.Name()
// 	if _, err := io.Copy(f, r); err != nil {
// 		os.Remove(path)
// 		return nil, path, err
// 	}
// 	f.Close()
// 	if err := e.Launch(path); err != nil {
// 		return nil, path, err
// 	}
// 	bytes, err := ioutil.ReadFile(path)
// 	return bytes, path, err
// }

// func (e *Editor) Launch(path string) (err error) {
// 	fmt.Println("launching")
// 	abs, err := filepath.Abs(path)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("path:", abs, "IDE", e.IDE)
// 	command := exec.Command(e.IDE, abs)
// 	command.Stderr = os.Stderr
// 	command.Stdout = os.Stdout
// 	command.Stdin = os.Stdin
// 	err = command.Run()
// 	fmt.Print(err)
// 	return
// }

// func tempFile(prefix, suffix string) (f *os.File, err error) {
// 	// dir := os.TempDir()
// 	dir := ""
// 	// name := filepath.Join(dir, prefix, suffix)
// 	name := fmt.Sprintf("%s%s%s\n", dir, prefix, suffix)
// 	fmt.Printf("prefix: %s, suffix: %s, filename: %s", prefix, suffix, name)
// 	// for i := 0; i < 10000; i++ {
// 	// name := filepath.Join(dir, prefix+randSeq(5)+suffix)
// 	f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	// if os.IsExist(err) {
// 	// continue
// 	// }
// 	// break
// 	// }
// 	return
// }

// var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// func randSeq(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(b)
// }
