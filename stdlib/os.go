package stdlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/zeaphoo/nanojs/v2"
)

var osModule = map[string]nanojs.Object{
	"o_rdonly":            &nanojs.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &nanojs.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &nanojs.Int{Value: int64(os.O_RDWR)},
	"o_append":            &nanojs.Int{Value: int64(os.O_APPEND)},
	"o_create":            &nanojs.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &nanojs.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &nanojs.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &nanojs.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &nanojs.Int{Value: int64(os.ModeDir)},
	"mode_append":         &nanojs.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &nanojs.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &nanojs.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &nanojs.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &nanojs.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &nanojs.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &nanojs.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &nanojs.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &nanojs.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &nanojs.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &nanojs.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &nanojs.Int{Value: int64(os.ModeType)},
	"mode_perm":           &nanojs.Int{Value: int64(os.ModePerm)},
	"path_separator":      &nanojs.Char{Value: os.PathSeparator},
	"path_list_separator": &nanojs.Char{Value: os.PathListSeparator},
	"dev_null":            &nanojs.String{Value: os.DevNull},
	"seek_set":            &nanojs.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &nanojs.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &nanojs.Int{Value: int64(io.SeekEnd)},
	"args": &nanojs.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &nanojs.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &nanojs.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &nanojs.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &nanojs.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &nanojs.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &nanojs.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &nanojs.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &nanojs.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &nanojs.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &nanojs.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &nanojs.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &nanojs.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &nanojs.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &nanojs.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &nanojs.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &nanojs.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &nanojs.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &nanojs.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &nanojs.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &nanojs.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &nanojs.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &nanojs.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &nanojs.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &nanojs.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &nanojs.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &nanojs.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &nanojs.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &nanojs.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &nanojs.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &nanojs.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &nanojs.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &nanojs.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &nanojs.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &nanojs.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &nanojs.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &nanojs.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &nanojs.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &nanojs.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	fname, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > nanojs.MaxBytesLen {
		return nil, nanojs.ErrBytesLimit
	}
	return &nanojs.Bytes{Value: bytes}, nil
}

func osStat(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	fname, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"name":  &nanojs.String{Value: stat.Name()},
			"mtime": &nanojs.Time{Value: stat.ModTime()},
			"size":  &nanojs.Int{Value: stat.Size()},
			"mode":  &nanojs.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = nanojs.TrueValue
	} else {
		fstat.Value["directory"] = nanojs.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	s1, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	s1, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 3 {
		return nil, nanojs.ErrWrongNumArguments
	}
	s1, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	i2, ok := nanojs.ToInt(args[1])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := nanojs.ToInt(args[2])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 0 {
		return nil, nanojs.ErrWrongNumArguments
	}
	arr := &nanojs.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > nanojs.MaxStringLen {
			return nil, nanojs.ErrStringLimit
		}
		arr.Value = append(arr.Value, &nanojs.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *nanojs.UserFunction {
	return &nanojs.UserFunction{
		Name: name,
		Value: func(args ...nanojs.Object) (nanojs.Object, error) {
			if len(args) != 2 {
				return nil, nanojs.ErrWrongNumArguments
			}
			s1, ok := nanojs.ToString(args[0])
			if !ok {
				return nil, nanojs.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "string(compatible)",
					Found:    args[0].TypeName(),
				}
			}
			i2, ok := nanojs.ToInt64(args[1])
			if !ok {
				return nil, nanojs.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	s1, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, ok := os.LookupEnv(s1)
	if !ok {
		return nanojs.FalseValue, nil
	}
	if len(res) > nanojs.MaxStringLen {
		return nil, nanojs.ErrStringLimit
	}
	return &nanojs.String{Value: res}, nil
}

func osExpandEnv(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	s1, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > nanojs.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > nanojs.MaxStringLen {
		return nil, nanojs.ErrStringLimit
	}
	return &nanojs.String{Value: s}, nil
}

func osExec(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) == 0 {
		return nil, nanojs.ErrWrongNumArguments
	}
	name, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := nanojs.ToString(arg)
		if !ok {
			return nil, nanojs.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}
	i1, ok := nanojs.ToInt(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...nanojs.Object) (nanojs.Object, error) {
	if len(args) != 4 {
		return nil, nanojs.ErrWrongNumArguments
	}
	name, ok := nanojs.ToString(args[0])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *nanojs.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *nanojs.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := nanojs.ToString(args[2])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *nanojs.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *nanojs.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []nanojs.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*nanojs.String)
		if !ok {
			return nil, nanojs.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
