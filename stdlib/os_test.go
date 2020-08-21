package stdlib_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/zeaphoo/nanojs/v2"
	"github.com/zeaphoo/nanojs/v2/require"
)

func TestReadFile(t *testing.T) {
	content := []byte("the quick brown fox jumps over the lazy dog")
	tf, err := ioutil.TempFile("", "test")
	require.NoError(t, err)
	defer func() { _ = os.Remove(tf.Name()) }()

	_, err = tf.Write(content)
	require.NoError(t, err)
	_ = tf.Close()

	module(t, "os").call("read_file", tf.Name()).
		expect(&nanojs.Bytes{Value: content})
}

func TestReadFileArgs(t *testing.T) {
	module(t, "os").call("read_file").expectError()
}
func TestFileStatArgs(t *testing.T) {
	module(t, "os").call("stat").expectError()
}

func TestFileStatFile(t *testing.T) {
	content := []byte("the quick brown fox jumps over the lazy dog")
	tf, err := ioutil.TempFile("", "test")
	require.NoError(t, err)
	defer func() { _ = os.Remove(tf.Name()) }()

	_, err = tf.Write(content)
	require.NoError(t, err)
	_ = tf.Close()

	stat, err := os.Stat(tf.Name())
	if err != nil {
		t.Logf("could not get tmp file stat: %s", err)
		return
	}

	module(t, "os").call("stat", tf.Name()).expect(&nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"name":      &nanojs.String{Value: stat.Name()},
			"mtime":     &nanojs.Time{Value: stat.ModTime()},
			"size":      &nanojs.Int{Value: stat.Size()},
			"mode":      &nanojs.Int{Value: int64(stat.Mode())},
			"directory": nanojs.FalseValue,
		},
	})
}

func TestFileStatDir(t *testing.T) {
	td, err := ioutil.TempDir("", "test")
	require.NoError(t, err)
	defer func() { _ = os.RemoveAll(td) }()

	stat, err := os.Stat(td)
	require.NoError(t, err)

	module(t, "os").call("stat", td).expect(&nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"name":      &nanojs.String{Value: stat.Name()},
			"mtime":     &nanojs.Time{Value: stat.ModTime()},
			"size":      &nanojs.Int{Value: stat.Size()},
			"mode":      &nanojs.Int{Value: int64(stat.Mode())},
			"directory": nanojs.TrueValue,
		},
	})
}

func TestOSExpandEnv(t *testing.T) {
	curMaxStringLen := nanojs.MaxStringLen
	defer func() { nanojs.MaxStringLen = curMaxStringLen }()
	nanojs.MaxStringLen = 12

	_ = os.Setenv("NANOJS", "FOO BAR")
	module(t, "os").call("expand_env", "$NANOJS").expect("FOO BAR")

	_ = os.Setenv("NANOJS", "FOO")
	module(t, "os").call("expand_env", "$NANOJS $NANOJS").expect("FOO FOO")

	_ = os.Setenv("NANOJS", "123456789012")
	module(t, "os").call("expand_env", "$NANOJS").expect("123456789012")

	_ = os.Setenv("NANOJS", "1234567890123")
	module(t, "os").call("expand_env", "$NANOJS").expectError()

	_ = os.Setenv("NANOJS", "123456")
	module(t, "os").call("expand_env", "$NANOJS$NANOJS").expect("123456123456")

	_ = os.Setenv("NANOJS", "123456")
	module(t, "os").call("expand_env", "${NANOJS}${NANOJS}").
		expect("123456123456")

	_ = os.Setenv("NANOJS", "123456")
	module(t, "os").call("expand_env", "$NANOJS $NANOJS").expectError()

	_ = os.Setenv("NANOJS", "123456")
	module(t, "os").call("expand_env", "${NANOJS} ${NANOJS}").expectError()
}
