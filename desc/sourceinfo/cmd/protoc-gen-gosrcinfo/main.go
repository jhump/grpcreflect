// Command protoc-gen-gosrcinfo is a protoc plugin. It emits Go code, into files
// named "<file>.pb.srcinfo.go". These source files include source code info for
// processed proto files and register that info with the srcinfo package.
package main

import (
	"fmt"
	"path"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/gopoet"
	"github.com/jhump/goprotoc/plugins"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/jhump/protoreflect/desc"
)

func main() {
	plugins.PluginMain(genSourceInfo)
}

func genSourceInfo(req *plugins.CodeGenRequest, resp *plugins.CodeGenResponse) error {
	args, err := parseArgs(req.Args)
	if err != nil {
		return err
	}
	names := plugins.GoNames{
		ImportMap:      args.importMap,
		ModuleRoot:     args.moduleRoot,
		SourceRelative: args.sourceRelative,
	}
	if args.importPath != "" {
		// if we're overriding import path, go ahead and query
		// package for each file, which will cache the override name
		// so all subsequent queries are consistent
		for _, fd := range req.Files {
			// Only use the override for files that don't otherwise have an
			// entry in the specified import map
			if _, ok := args.importMap[fd.GetName()]; !ok {
				names.GoPackageForFileWithOverride(fd, args.importPath)
			}
		}
	}
	for _, fd := range req.Files {
		if err := generateSourceInfo(fd, &names, resp, args); err != nil {
			if fe, ok := err.(*gopoet.FormatError); ok {
				if args.debug {
					return fmt.Errorf("%s: error in generated Go code: %v:\n%s", fd.GetName(), err, fe.Unformatted)
				} else {
					return fmt.Errorf("%s: error in generated Go code: %v (use debug=true arg to show full source)", fd.GetName(), err)
				}
			} else {
				return fmt.Errorf("%s: %v", fd.GetName(), err)
			}
		}
	}
	return nil
}

var typeOfSourceInfo = reflect.TypeOf((*descriptorpb.SourceCodeInfo)(nil)).Elem()

func generateSourceInfo(fd *desc.FileDescriptor, names *plugins.GoNames, resp *plugins.CodeGenResponse, args codeGenArgs) error {
	si := fd.AsFileDescriptorProto().GetSourceCodeInfo()
	if len(si.GetLocation()) == 0 {
		return nil
	}
	pkg := names.GoPackageForFile(fd)
	filename := names.OutputFilenameFor(fd, ".pb.srcinfo.go")
	f := gopoet.NewGoFile(path.Base(filename), pkg.ImportPath, pkg.Name)

	f.FileComment = "Code generated by protoc-gen-gosrcinfo. DO NOT EDIT.\n" +
		"source: " + fd.GetName()

	siBytes, err := proto.Marshal(si)
	if err != nil {
		return fmt.Errorf("failed to serialize source code info: %w", err)
	}

	srcCodeInfo := f.EnsureTypeImported(gopoet.TypeNameForReflectType(typeOfSourceInfo))
	srcInfoPkg := f.RegisterImport("github.com/jhump/protoreflect/desc/sourceinfo", "sourceinfo")
	protoPkg := f.RegisterImport("google.golang.org/protobuf/proto", "proto")

	varName := "srcInfo_" + clean(fd.GetName())
	var initBlock gopoet.CodeBlock
	initBlock.Println("[]byte{")
	for len(siBytes) > 0 {
		var chunk []byte
		if len(siBytes) < 16 {
			chunk = siBytes
			siBytes = nil
		} else {
			chunk = siBytes[:16]
			siBytes = siBytes[16:]
		}
		for _, b := range chunk {
			initBlock.Printf(" 0x%02x,", b)
		}
		initBlock.Println("")
	}
	initBlock.Println("}")
	f.AddVar(gopoet.NewVar(varName).SetInitializer(&initBlock))
	f.AddElement(gopoet.NewFunc("init").
		Printlnf("var si %s", srcCodeInfo).
		Printlnf("if err := %sUnmarshal(%s, &si); err != nil {", protoPkg, varName).
		Println("    panic(err)").
		Println("}").
		Printlnf("%sRegisterSourceInfo(%q, &si)", srcInfoPkg, fd.GetName()))

	out := resp.OutputFile(filename)
	return gopoet.WriteGoFile(out, f)
}

func clean(name string) string {
	name = strings.TrimSuffix(name, ".proto")
	data := ([]byte)(name)
	for i, b := range data {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
			continue
		}
		data[i] = '_'
	}
	return string(data)
}

type codeGenArgs struct {
	debug          bool
	importPath     string
	importMap      map[string]string
	moduleRoot     string
	sourceRelative bool
}

func parseArgs(args []string) (codeGenArgs, error) {
	var result codeGenArgs
	for _, arg := range args {
		vals := strings.SplitN(arg, "=", 2)
		switch vals[0] {
		case "debug":
			val, err := boolVal(vals)
			if err != nil {
				return result, err
			}
			result.debug = val

		case "import_path":
			if len(vals) == 1 {
				return result, fmt.Errorf("plugin option 'import_path' requires an argument")
			}
			result.importPath = vals[1]

		case "module":
			if len(vals) == 1 {
				return result, fmt.Errorf("plugin option 'module' requires an argument")
			}
			result.moduleRoot = vals[1]

		case "paths":
			if len(vals) == 1 {
				return result, fmt.Errorf("plugin option 'paths' requires an argument")
			}
			switch vals[1] {
			case "import":
				result.sourceRelative = false
			case "source_relative":
				result.sourceRelative = true
			default:
				return result, fmt.Errorf("plugin option 'paths' accepts 'import' or 'source_relative' as value, got %q", vals[1])
			}

		default:
			if len(vals[0]) > 1 && vals[0][0] == 'M' {
				if len(vals) == 1 {
					return result, fmt.Errorf("plugin 'M' options require an argument: %s", vals[0])
				}
				if result.importMap == nil {
					result.importMap = map[string]string{}
				}
				result.importMap[vals[0][1:]] = vals[1]
				break
			}

			return result, fmt.Errorf("unknown plugin option: %s", vals[0])
		}
	}

	if result.sourceRelative && result.moduleRoot != "" {
		return result, fmt.Errorf("plugin option 'module' cannot be used with 'paths=source_relative'")
	}

	return result, nil
}

func boolVal(vals []string) (bool, error) {
	if len(vals) == 1 {
		// if no value, assume "true"
		return true, nil
	}
	switch strings.ToLower(vals[1]) {
	case "true", "on", "yes", "1":
		return true, nil
	case "false", "off", "no", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean arg for option '%s': %s", vals[0], vals[1])
	}
}
