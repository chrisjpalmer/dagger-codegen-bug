# non-deterministic codegen repro

This is a simple repro for an issue in dagger's codegen for go, which shows that in certain cases the codegen is non-deterministic. 
I believe the issue is that the token position used to order objects is actually non-deterministic.
This seems to be something with the `golang.org/x/tools/go/packages` library.

To reproduce the issue, I have copied two functions from the dagger codegen source:
- [a bit of the `getTypes` function](https://github.com/dagger/dagger/blob/main/cmd/codegen/generator/go/templates/modules.go#L73)
- [the `loadPackage` function](https://github.com/dagger/dagger/blob/main/cmd/codegen/generator/go/loader.go#L21)
 
Additionally I have checked in the function stubs from a module that was causing us trouble (the actual function implementations aren't here but it shouldn't matter since function implementations are not used for this process).

The repro code prints two lists. The first is the name sorted objects that come from a call to `pkgScope.Names()`. This list
always seems to be in the right order. The second is the list after resorting the objects by their position:

```go
sort.Slice(objs, func(i, j int) bool {
		return objs[i].Pos() < objs[j].Pos()
	})
```

This list seems to have an unstable order.

In the resulting printout of two executive runs, you can see that the order of the second list changes.
The object positions are different each time.


```
cd ./test-codegen

❯ go run .

GitRepo pos: 13437305
New pos: 13437434
RegistryAuth pos: 13437218
Remote pos: 13438737
ghHost pos: 13437109
httpsMatch pos: 13438944
parseRemote pos: 13439021
repository pos: 13437927
sshMatch pos: 13438878
---------
---------
---------
dagger/git-repo.RegistryAuth
dagger/git-repo.GitRepo
func(defaultBranch string, src invalid type, ssh invalid type) *dagger/git-repo.GitRepo
dagger/git-repo.Remote

❯ go run .

GitRepo pos: 13440911
New pos: 13441040
RegistryAuth pos: 13440824
Remote pos: 13437652
ghHost pos: 13440715
httpsMatch pos: 13437859
parseRemote pos: 13437936
repository pos: 13441533
sshMatch pos: 13437793
---------
---------
---------
dagger/git-repo.Remote
dagger/git-repo.RegistryAuth
dagger/git-repo.GitRepo
func(defaultBranch string, src invalid type, ssh invalid type) *dagger/git-repo.GitRepo
```