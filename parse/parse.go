package parse

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	//"golang.org/x/tools/go/ast/astutil"
	"bytes"
)

// If the type represents a struct (not an alias or primitive), it will have a package and file name identifier
type Type struct {
	Literal string `json:"literal"`
	// Package string `json:"package"`
	Structs []string `json:"structs"`
}

type ClientStruct struct {
	Packages []Package `json:"packages"`
	Edges    []Edge    `json:"edges"`
}

type Package struct {
	Name  string `json:"name"`
	Files []File `json:"files"`
}

type File struct {
	Name    string   `json:"name"`
	Structs []Struct `json:"structs"`
}

type Struct struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
	//Methods []Method `json:"methods"`
}

// TODO multiple names per field
type Field struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}

type Method struct {
	Name       string `json:"name"`
	ReturnType []Type `json:"returnType"`
}

type Node struct {
	FieldTypeName string `json:"fieldTypeName"`
	StructName    string `json:"structName"`
	PackageName   string `json:"packageName"`
	FileName      string `json:"fileName"`
}

type Edge struct {
	To   *Node `json:"to"`
	From *Node `json:"from"`
}

func GetStructsFile(fset *token.FileSet, f *ast.File, fname string, packageName string) (File, []Edge) {
	structs := []Struct{}
	edges := []Edge{}
	//ast.Print(fset, f)
	// For all declarations
	for _, d := range f.Decls {
		if g, ok := d.(*ast.GenDecl); ok && g.Tok == token.TYPE {
			// For all type declarations
			for _, s := range g.Specs {
				if ts, ok := s.(*ast.TypeSpec); ok {
					if st, ok := ts.Type.(*ast.StructType); ok {
						fields := []Field{}
						for _, field := range st.Fields.List {
							// TODO: why can a field have multiple names?
							for _, name := range field.Names {
								// TODO: can the type be an expression?
								//fmt.Println(astutil.NodeDescription(field.Type))
								var buf bytes.Buffer
								if err := format.Node(&buf, fset, field.Type); err != nil {
									panic(err)
								}
								// stpackage, stname = GetType(field.Type)
								// fieldtype := Type{Literal: string(buf.Bytes()), Package: stpackage, Struct: stname}
								stname, toNodes := GetTypes(field.Type, packageName)
								fieldtype := Type{Literal: string(buf.Bytes()), Structs: stname}
								fi := Field{Name: name.Name, Type: fieldtype}
								fields = append(fields, fi)

								// Add edges
								for _, toNode := range toNodes {
									edges = append(edges, Edge{From: &Node{FieldTypeName: name.Name, StructName: ts.Name.Name, FileName: fname, PackageName: packageName}, To: toNode})
								}
							}
						}
						structs = append(structs, Struct{Name: ts.Name.Name, Fields: fields})
					}
				}
			}
		}
	}
	return File{Name: fname, Structs: structs}, edges
}

// TODO: don't deeply nest
// https://golang.org/ref/spec#Struct_types
// func GetStructsFileName(filename string) File {
// 	fset := token.NewFileSet()

// 	f, err := parser.ParseFile(fset, filename, nil, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return GetStructsFile(fset, f, filename)
// }

// TODO use a map instead of this craziness
func GetFileName(toNode *Node, pkgs []Package) string {
	for _, pkg := range pkgs {
		if pkg.Name == toNode.PackageName {
			for _, file := range pkg.Files {
				for _, st := range file.Structs {
					if st.Name == toNode.StructName {
						return file.Name
					}
				}
			}
		}
	}
	// Because we don't index types like funcs yet, those won't be found
	// and we can't find the filename. Just leave it as is for now. TODO
	fmt.Println("Matching file not found for struct", toNode.StructName, "(probably a library package)")
	return ""
}

func getPackagesEdgesDirName(path string, fset *token.FileSet) ([]Package, []Edge, map[string]*ast.Package) {
	var packages []Package
	var edges []Edge
	packagemap, err := parser.ParseDir(fset, path, nil, 0)
	if err != nil {
		panic(err)
	}
	for packagename, packageval := range packagemap {
		// TODO ignore main for now because of conflicts
		// Assumes that packagenames are unique
		if packagename != "main" {
			files := []File{}
			for fname, f := range packageval.Files {
				newfile, newedges := GetStructsFile(fset, f, fname, packagename)
				files = append(files, newfile)
				edges = append(edges, newedges...)
			}
			packages = append(packages, Package{Name: packagename, Files: files})
		}
	}
	return packages, edges, packagemap
}

func pathInDirectories(path string, directories []string) bool {
	for _, directory := range directories {
		if path == directory {
			return true
		}
	}
	return false
}

// Get Structs from a dirname
// including vendor if needed
func GetStructsDirName(path string) (*ClientStruct, map[string]*ast.Package) {
	directories := []string{}
	packages := []Package{}
	edges := []Edge{}
	pkgmap := map[string]*ast.Package{}
	fset := token.NewFileSet()

	directories = []string{path}
	// Get all the directories, but not current app folder
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			if !strings.Contains(path, "app") && !pathInDirectories(path, directories) {
				directories = append(directories, path)
			}
		}
		return nil
	})
	fmt.Printf("Process %d directories, %s\n", len(directories), directories)

	for _, directory := range directories {
		newpackages, newedges, newpkgmap := getPackagesEdgesDirName(directory, fset)
		// Merge the packages, edges, and pkgmap from this directory with our other results
		packages = append(packages, newpackages...)
		edges = append(edges, newedges...)
		for k, v := range newpkgmap {
			pkgmap[k] = v
		}
	}

	// The To Nodes in Edges are currently missing filename because that's unknown when we are going through the AST
	// Here we fill in what the filename is
	validedges := []Edge{}
	for _, edge := range edges {
		if name := GetFileName(edge.To, packages); name != "" {
			edge.To.FileName = name
			validedges = append(validedges, edge)
		}
	}
	return &ClientStruct{Packages: packages, Edges: validedges}, pkgmap
}

func isPrimitive(name string) bool {
	for _, basicType := range types.Typ {
		if name == basicType.Name() {
			return true
		}
	}
	if name == "error" || name == "byte" {
		return true
	}
	return false
}

// This will fill in the packageName, and filename is figured out later after the whole
// AST has been traversed. TODO clean this up to not separate those steps
func GetType(node *ast.Ident, packageName string) (string, *Node) {
	var toNode *Node
	name := node.Name
	if !isPrimitive(name) {
		toNode = &Node{StructName: name, PackageName: packageName}
	}
	return name, toNode
}

func GetTypes(node ast.Expr, packageName string) ([]string, []*Node) {
	switch node.(type) {
	case *ast.Ident:
		toNodes := []*Node{}
		name, toNode := GetType(node.(*ast.Ident), packageName)
		if toNode != nil {
			toNodes = append(toNodes, toNode)
		}
		return []string{name}, toNodes
	case *ast.SelectorExpr:
		// TODO: This assumes the selector expression is of type Ident. Use scope.lookup instead.
		xPackageName := node.(*ast.SelectorExpr).X.(*ast.Ident).Name
		toNodes := []*Node{}
		name, toNode := GetType(node.(*ast.SelectorExpr).Sel, xPackageName)
		if toNode != nil {
			toNodes = append(toNodes, toNode)
		}
		return []string{name}, toNodes
	case *ast.ArrayType:
		return GetTypes(node.(*ast.ArrayType).Elt, packageName)
	case *ast.MapType:
		// TODO: also handle the value
		return GetTypes(node.(*ast.MapType).Key, packageName)
	case *ast.StarExpr:
		return GetTypes(node.(*ast.StarExpr).X, packageName)
	case *ast.FuncType:
		return []string{"TODO"}, nil
	case *ast.InterfaceType:
		return []string{"TODO"}, nil
	case *ast.ChanType:
		return []string{"TODO"}, nil
	default:
		fmt.Println(reflect.TypeOf(node))
		panic("Need to cover all Type Exprs")
	}
}

// (path, packages) -> (Write to directory)
// Given a client side struct data structure, deserialize it into an AST
// and write to the given directory
func WriteClientPackages(dirpath string, pkgs map[string]*ast.Package, clientpackages []Package) error {
	var err error
	for _, clientpackage := range clientpackages {
		for _, clientfile := range clientpackage.Files {
			packagename := clientpackage.Name
			packageast := pkgs[packagename]
			// Get the AST with the matching file name
			f := packageast.Files[clientfile.Name]

			if f == nil {
				fmt.Println("Couldn't find", packagename, packageast.Files, clientfile.Name)
			}
			// Update the AST with the values from the client
			f, err = clientFileToAST(clientfile, f)
			if err != nil {
				return err
			}
			writeFileAST(clientfile.Name, f)
		}
	}
	return nil
}

// Write the given AST to the filepath
func writeFileAST(filepath string, f *ast.File) {
	fset := token.NewFileSet()
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	err := ioutil.WriteFile(filepath, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

// ClientFileToAST convert
func clientFileToAST(clientfile File, f *ast.File) (*ast.File, error) {
	var newdecls []ast.Decl
	if len(f.Decls) > 0 {
		newdecls = []ast.Decl{f.Decls[0]}
	} else {
		newdecls = []ast.Decl{}
	}

	f.Decls = removeStructDecls(f.Decls)
	newstructs, err := clientFileToDecls(clientfile)
	if err != nil {
		return nil, err
	}
	// Assume import is the first decl. Add typedefs after that
	newdecls = append(newdecls, newstructs...)
	if len(f.Decls) > 0 {
		newdecls = append(newdecls, f.Decls[1:]...)
	} else {
		newdecls = append(newdecls, f.Decls...)
	}
	f.Decls = newdecls
	return f, nil
}

// Given a client-formatted File struct, return a list of AST declarations
func clientFileToDecls(clientfile File) ([]ast.Decl, error) {
	decls := []ast.Decl{}
	for _, clientstruct := range clientfile.Structs {
		decl := &ast.GenDecl{Tok: token.TYPE}
		fieldList := []*ast.Field{}
		for _, clientfield := range clientstruct.Fields {
			// TODO assuming literal == struct. Change to support more than ident
			// TODO tags
			// TODO support maps
			parsedtype, err := parser.ParseExpr(clientfield.Type.Literal)
			if err != nil {
				return nil, err
			}
			field := ast.Field{
				Names: []*ast.Ident{{Name: clientfield.Name}},
				Type:  parsedtype}
			fieldList = append(fieldList, &field)
		}
		fields := &ast.FieldList{List: fieldList}
		structExpr := &ast.StructType{Struct: token.NoPos, Fields: fields}
		spec := ast.TypeSpec{
			Name: ast.NewIdent(clientstruct.Name),
			Type: structExpr}
		decl.Specs = append(decl.Specs, &spec)
		decls = append(decls, decl)
	}
	return decls, nil
}

func removeStructDecls(decls []ast.Decl) []ast.Decl {
	// TODO type definitions that aren't structs
	newdecls := []ast.Decl{}
	for _, decl := range decls {
		if g, ok := decl.(*ast.GenDecl); !ok || g.Tok != token.TYPE {
			newdecls = append(newdecls, decl)
		}
	}
	return newdecls
}
