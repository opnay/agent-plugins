// git-codex provides the bounded message-file and commit lifecycle for toolkit:git.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unicode/utf8"
)

const version = "0.1.0"
const prefix = "toolkit-git-message."

func fail(code, message string) { fmt.Fprintf(os.Stderr, "git-codex: %s: %s\n", code, message) }
func usage() { fmt.Println("Usage:\n  git codex --help\n  git codex --version\n  git codex install [--force|--check]\n  git codex message create\n  git codex message validate <file>\n  git codex commit <file>") }

func tempRoot() (string, error) { return filepath.EvalSymlinks(os.TempDir()) }
func safeFile(path string) (syscall.Stat_t, error) {
	var zero syscall.Stat_t
	if !filepath.IsAbs(path) || !strings.HasPrefix(filepath.Base(path), prefix) { return zero, fmt.Errorf("path_invalid") }
	root, err := tempRoot(); if err != nil { return zero, err }
	parent, err := filepath.EvalSymlinks(filepath.Dir(path)); if err != nil || parent != root { return zero, fmt.Errorf("path_invalid") }
	info, err := os.Lstat(path); if err != nil || !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 || info.Mode().Perm()&0400 == 0 { return zero, fmt.Errorf("file_unsafe") }
	stat, ok := info.Sys().(*syscall.Stat_t); if !ok || stat.Uid != uint32(os.Geteuid()) { return zero, fmt.Errorf("file_unsafe") }
	return *stat, nil
}
func validate(path string) (syscall.Stat_t, error) {
	stat, err := safeFile(path); if err != nil { return stat, err }
	b, err := os.ReadFile(path); if err != nil || !utf8.Valid(b) || bytes.IndexByte(b, 0) >= 0 || bytes.Contains(b, []byte(`\n`)) { return stat, fmt.Errorf("message_invalid") }
	for _, line := range bytes.Split(b, []byte("\n")) { line = bytes.TrimSuffix(line, []byte("\r")); if bytes.ContainsRune(line, '\r') { return stat, fmt.Errorf("message_invalid") } }
	lines := bytes.Split(b, []byte("\n")); if len(lines) == 0 || len(bytes.TrimSpace(bytes.TrimSuffix(lines[0], []byte("\r")))) == 0 { return stat, fmt.Errorf("message_invalid") }
	if len(lines) > 2 && len(bytes.TrimSpace(bytes.Join(lines[1:], nil))) > 0 && len(bytes.TrimSuffix(lines[1], []byte("\r"))) != 0 { return stat, fmt.Errorf("message_invalid") }
	return stat, nil
}
func same(a, b syscall.Stat_t) bool { return a.Dev == b.Dev && a.Ino == b.Ino }
func result(reason string, attempted bool) { fmt.Fprintf(os.Stderr, "git-codex: commit_result: reason=%s attempted=%t\n", reason, attempted) }
func run(name string, args ...string) error { c:=exec.Command(name,args...); c.Stdin=os.Stdin;c.Stdout=os.Stdout;c.Stderr=os.Stderr;return c.Run() }
func head() (string,error) { b,e:=exec.Command("git","rev-parse","--verify","HEAD").Output(); return strings.TrimSpace(string(b)),e }
func safeDir(path string) error { info,e:=os.Lstat(path);if e!=nil||!info.IsDir()||info.Mode()&os.ModeSymlink!=0||info.Mode().Perm()&0022!=0{return fmt.Errorf("unsafe directory")};s,ok:=info.Sys().(*syscall.Stat_t);if !ok||s.Uid!=uint32(os.Geteuid()){return fmt.Errorf("unsafe directory")};return nil }
func install(force, check bool) error { self,e:=os.Executable();if e!=nil{return e}; home:=os.Getenv("HOME");if home==""{return fmt.Errorf("HOME is required")}; local:=filepath.Join(home,".local"); bin:=filepath.Join(local,"bin"); target:=filepath.Join(bin,"git-codex"); if check { if e=safeDir(local);e!=nil{return e};if e=safeDir(bin);e!=nil{return e}; a,e:=os.ReadFile(self);if e!=nil{return e};b,e:=os.ReadFile(target);if e!=nil||!bytes.Equal(a,b){return fmt.Errorf("installed binary does not match")};fmt.Println(target);return nil }; for _,d:=range []string{local,bin}{if _,e=os.Lstat(d);os.IsNotExist(e){if e=os.Mkdir(d,0755);e!=nil{return e}};if e=safeDir(d);e!=nil{return e}}; if info,e:=os.Lstat(target);e==nil {if !info.Mode().IsRegular()||info.Mode()&os.ModeSymlink!=0{return fmt.Errorf("refusing to replace a symlink or directory")};a,_:=os.ReadFile(self);b,_:=os.ReadFile(target);if bytes.Equal(a,b)&&info.Mode().Perm()==0755{return nil};if !force{return fmt.Errorf("existing file is preserved; rerun with --force")}}; src,e:=os.Open(self);if e!=nil{return e};defer src.Close();tmp,e:=os.CreateTemp(bin,".git-codex.*");if e!=nil{return e};if _,e=io.Copy(tmp,src);e!=nil{tmp.Close();os.Remove(tmp.Name());return e};if e=tmp.Chmod(0755);e==nil{e=tmp.Close()};if e==nil{e=os.Rename(tmp.Name(),target)};return e }
func main() {
	a:=os.Args[1:]; if len(a)==1 && (a[0]=="--help"||a[0]=="-h") {usage();return}; if len(a)==1&&a[0]=="--version" {fmt.Printf("toolkit-git-codex %s\n",version);return}; if len(a)>=1&&a[0]=="install" {if len(a)>2{usage();os.Exit(1)};force:=len(a)==2&&a[1]=="--force";check:=len(a)==2&&a[1]=="--check";if len(a)==2&&!force&&!check{usage();os.Exit(1)};if e:=install(force,check);e!=nil{fail("install_failed",e.Error());os.Exit(1)};return}
	if len(a)==2&&a[0]=="message"&&a[1]=="create" { root,e:=tempRoot();if e!=nil{fail("path_invalid","OS temp directory is unavailable");os.Exit(1)}; f,e:=os.CreateTemp(root,prefix+"*");if e!=nil{fail("create_failed","could not allocate a message file");os.Exit(1)};f.Close();os.Chmod(f.Name(),0600);if _,e=safeFile(f.Name());e!=nil{os.Remove(f.Name());fail("create_failed","allocated file is unsafe");os.Exit(1)};fmt.Println(f.Name());return }
	if len(a)==3&&a[0]=="message"&&a[1]=="validate" { if _,e:=validate(a[2]);e!=nil{fail(e.Error(),"message file failed validation");os.Exit(1)};return }
	if len(a)==2&&a[0]=="commit" { stat,e:=validate(a[1]);if e!=nil{fail(e.Error(),"message file failed validation");result("validation_failed",false);os.Exit(1)}; before,e:=head(); unborn:=e!=nil; if unborn { if e:=run("git","rev-parse","--is-inside-work-tree");e!=nil{result("head_unavailable",false);os.Exit(1)} }; err:=run("git","commit","-F",a[1]); after,afterErr:=head(); if err==nil&&afterErr==nil&&(unborn||after!=before) { now,e:=safeFile(a[1]);if e!=nil||!same(stat,now)||os.Remove(a[1])!=nil {fmt.Printf("commit=%s cleanup=failed message_file=%s\n",after,a[1]);os.Exit(2)};fmt.Printf("commit=%s cleanup=deleted\n",after);return }; if err!=nil&&((!unborn&&afterErr==nil&&after==before)||(unborn&&afterErr!=nil)){result("commit_failed",true);os.Exit(1)};result("outcome_unknown",true);os.Exit(3) }
	_ = runtime.GOOS; usage(); os.Exit(1)
}
