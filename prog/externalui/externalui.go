// Code generated by "esc -o prog/externalui/externalui.go -pkg externalui -prefix client/build-external -include \.html$ client/build-external"; DO NOT EDIT.

package externalui

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/index.html": {
		name:    "index.html",
		local:   "client/build-external/index.html",
		size:    1447,
		modtime: 1548712236,
		compressed: `
H4sIAAAAAAAC/7xUf4vbRhD9P59ishTSQqS1fcbnXiWVcFVJmzQ259CjhCBGu2NrL9LusruS4pZ+9yLL
Cg4lpaQ0f2n1ZufHezOzyWNpRDhagio0dfYoGT4gavQ+ZdpED55ljwCSilAOB4CkoYAgKnSeQsrasI/W
7GwKKtSU3RN2BDthLCV8hC48NTaUMkleOGWDMpqBMDqQDiljf7/YKeqtceHiVq9kqFJJnRIUnX6egtIq
KKwjL7CmdD4FGpNAjfrQ4oFS9oAdjiDLeqWl6eOiuM+f/Zrfb+5e7Irb3d2PxevNi/wVpMC+Ki6AYvvy
2W3+fPPyh/yuKNh3CR8DnfQ5Hz8OWexuN9u8eP08/yXfQQo/7zavYjsI9/WTP9jAx6EP7IZVIVh/w7m/
irHB343G3sfCNNwHDErE/aBo3Bv3znM/6Bq1ivtwrCmaokShooYiXEuxKFHOFrS4Xq7l1Xoxj4X3388E
ytVsPyvn11fruViUy9UVe8q0cQ3W/60EtDZaoZTL/Wq1KpfyWySJ8/L6n/LatqyV2GKoPjc3+/PJNx+1
oFb6HVSO9umXJQOO6pSdvH1FFFiW8GldktLI43kWH0fRG7WHOsBPOcxnb0cYILHTvpXO9J4qtPbIst9M
C+gIWq/0AVBD4oMz+pCZNkgMJBN+BmD0czFsa0JPkOCFDjecX8Q9qcCy1h4cSoKjad3knnDMIBhQjXWm
O9vovSWnSAuKE24nJm9IS7V/G0VnQKpu4tA7tOwDtcGgZMrQ2kEWqbqzx3ScNnR4glIW6H3gFysK3onP
7GZHWhrn44dPdC37MDv/WwmfHKV/URMfByfh47P8VwAAAP//oBvEbacFAAA=
`,
	},

	"/terminal.html": {
		name:    "terminal.html",
		local:   "client/build-external/terminal.html",
		size:    1078,
		modtime: 1548712236,
		compressed: `
H4sIAAAAAAAC/7RUX2vbPhR976e4Fb/Hn62kf9Kukz1Kl7HRspZ2rIxSzI10E6u1JSHJdrNPPxzHJTDG
xmBPls/VPdY951hiX1kZ146gjHWV74n+AbLCEDJmbPIUWL4HIEpC1S8ARE0RQZboA8WMNXGZnLJtKepY
UX5P2BLcSetI8AHa6TRYU8YUBem1i9oaBtKaSCZmjP28sdXUOevjzq5Oq1hmilotKdm8/A/a6KixSoLE
irLpSDR8BCo0qwZXlLEnbHEAWd5po2yXFsX9/Pzr/P769vKuuLi7/VB8ub6cf4YM2H/FDlDcXJ1fzD9e
X72f3xYFeyv4QLTRp9LmGUpPy4yVMbpwxnk4TLHG79ZgF1Jpax4iRi3Trpcn7ax/Djz0IiWN5iGuK0oi
+VobrBJ0LpnS0eJ4djKjxdHBG0Unx3JymsoQ3k0kqtlkOVlMTw5Pp/JgcTQ7ZOCpytiGJpREkeWCj6aJ
hVXrrSL7SfKgl1BF+DSH6eRxgAGEG11feNsFKtG5Ncu/2QbQEzRBmxWgARGit2aV2yYqjKQE3wIw9PkU
birCQCBwR5Azznd4N3KwvHErj4pgbRs/tguOOUQLunbettsavTjymoykVHA3TvJARunlY5JsAaXbcYbO
o2Ovo/UFrTKGzvWyKN1uO8blmJP+R8hYpJfId4ICwcu/tLUlo6wP6dMvXMtfQ/TPjvD7TP3B4fiQIMGH
W+JHAAAA///9vU16NgQAAA==
`,
	},

	"/": {
		name:  "/",
		local: `client/build-external`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"client/build-external": {
		_escData["/index.html"],
		_escData["/terminal.html"],
	},
}
