package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"sort"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkg, _, err := loadPackage(context.TODO(), "../git-repo", false)
	if err != nil {
		log.Fatal(err)
	}

	tps, err := getTypes(pkg)
	if err != nil {
		log.Fatal(err)
	}

	for _, tp := range tps {
		fmt.Println(tp.String())
	}
}

func getTypes(modulePkg *packages.Package) ([]types.Type, error) {
	var objs []types.Object

	pkgScope := modulePkg.Types.Scope()

	for _, name := range pkgScope.Names() {
		obj := pkgScope.Lookup(name)
		if obj == nil {
			return nil, fmt.Errorf("%q should exist in scope, but doesn't", name)
		}

		fmt.Println(name, "pos:", obj.Pos())
		objs = append(objs, obj)
	}

	fmt.Println("---------")
	fmt.Println("---------")
	fmt.Println("---------")

	// preserve definition order, so developer can keep more important /
	// entrypoint types higher up
	sort.Slice(objs, func(i, j int) bool {
		return objs[i].Pos() < objs[j].Pos()
	})

	tps := []types.Type{}
	for _, obj := range objs {
		// ignore any private definitions, they may be part of the runtime itself
		// e.g. marshalCtx
		if !obj.Exported() {
			continue
		}

		// // check if this is the constructor func, save it for later if so
		// if ok := ps.checkConstructor(obj); ok {
		// 	continue
		// }

		// // check if this is the DaggerObject interface
		// if ok, err := ps.checkDaggerObjectIface(obj); err != nil {
		// 	return nil, err
		// } else if ok {
		// 	continue
		// }

		// if ps.checkMainModuleObject(obj) || ps.isDaggerGenerated(obj) {
		// }
		tps = append(tps, obj.Type())
	}

	// if ps.daggerObjectIfaceType == nil && strict {
	// 	return nil, fmt.Errorf("cannot find default codegen %s interface in:\n[%s]", daggerObjectIfaceName, strings.Join(pkgScope.Names(), ", "))
	// }

	return tps, nil
}

func loadPackage(ctx context.Context, dir string, allowEmpty bool) (_ *packages.Package, _ *token.FileSet, rerr error) {

	fset := token.NewFileSet()
	pkgs, err := packages.Load(&packages.Config{
		Context: ctx,
		Dir:     dir,
		Tests:   false,
		Fset:    fset,
		Mode: packages.NeedName |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedModule,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			astFile, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
			if err != nil {
				return nil, err
			}
			// strip function bodies since we don't need them and don't need to waste time in packages.Load with type checking them
			for _, decl := range astFile.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {
					fn.Body = nil
				}
			}
			return astFile, nil
		},
		// Print some debug logs with timing information to stdout
		Logf: func(format string, args ...any) {
			fmt.Printf(format+"\n", args...)
		},
	}, ".")
	if err != nil {
		return nil, nil, err
	}
	switch len(pkgs) {
	case 0:
		return nil, nil, fmt.Errorf("no packages found in %s", dir)
	case 1:
		if pkgs[0].Name == "" && !allowEmpty {
			// this can happen when:
			// - loading an empty dir within an existing Go module
			// - loading a dir that is not included in a parent go.work
			return nil, nil, fmt.Errorf("package name is empty")
		}
		return pkgs[0], fset, nil
	default:
		// this would mean I don't understand how loading '.' works
		return nil, nil, fmt.Errorf("expected 1 package, got %d", len(pkgs))
	}
}
