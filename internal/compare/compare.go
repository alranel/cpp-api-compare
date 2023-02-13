package compare

import (
	"fmt"
	"sort"

	"github.com/alranel/cpp-api-compare/internal/library"
)

type ComparedMethod struct {
	library.Method
	ReturnTypes []string
	Presence    []bool
}

// IsCommon returns true if the method is present in all libraries,
// regardless of the return type.
func (method *ComparedMethod) IsCommon() bool {
	for _, presence := range method.Presence {
		if !presence {
			return false
		}
	}

	return true
}

func (method *ComparedMethod) HasDifferingReturnTypes() bool {
	returnTypes := make(map[string]bool)
	for _, rt := range method.ReturnTypes {
		returnTypes[rt] = true
	}
	return len(returnTypes) > 1
}

// CompareResult contains the result of a comparison between two singletons
// or two classes having the same name, implemented in different libraries.
type CompareResult struct {
	ClassName string
	Methods   []ComparedMethod
}

func (result *CompareResult) CommonMethods(sameRT bool) []ComparedMethod {
	var commonMethods []ComparedMethod
	for _, method := range result.Methods {
		if method.IsCommon() && method.HasDifferingReturnTypes() != sameRT {
			commonMethods = append(commonMethods, method)
		}
	}
	return commonMethods
}

func (result *CompareResult) NonCommonMethods() []ComparedMethod {
	var commonMethods []ComparedMethod
	for _, method := range result.Methods {
		if !method.IsCommon() {
			commonMethods = append(commonMethods, method)
		}
	}
	return commonMethods
}

func Compare(libs []library.Library, className string) CompareResult {
	// Find the correponding classs from each library
	classes := make(map[string]library.Class)
	for _, lib := range libs {
		class, foundClass := lib.Classes[className]

		if !foundClass {
			fmt.Printf("Class %s not found in library %s\n", className, lib.LibraryName)
			class = library.Class{
				ClassName:    className,
				DoesNotExist: true,
			}
		}
		classes[lib.LibraryName] = class
	}

	result := CompareResult{
		ClassName: className,
	}

	// Find the common methods
	methods := make(map[string]ComparedMethod)
	for _, lib := range libs {
		class := classes[lib.LibraryName]
		for _, method := range class.Methods {
			// In case of singletons, ignore constructors
			if class.IsSingleton && method.Ctor {
				continue
			}

			comparedMethod := ComparedMethod{Method: method}
			for _, otherLib := range libs {
				if lib.LibraryName == otherLib.LibraryName {
					comparedMethod.ReturnTypes = append(comparedMethod.ReturnTypes, method.ReturnType)
					comparedMethod.Presence = append(comparedMethod.Presence, true)
					continue
				}

				otherClass := classes[otherLib.LibraryName]
				compatibleMethod := otherClass.FindCompatibleMethod(method)
				if compatibleMethod == nil {
					comparedMethod.ReturnTypes = append(comparedMethod.ReturnTypes, "")
					comparedMethod.Presence = append(comparedMethod.Presence, false)
				} else {
					comparedMethod.ReturnTypes = append(comparedMethod.ReturnTypes, compatibleMethod.ReturnType)
					comparedMethod.Presence = append(comparedMethod.Presence, true)
				}
			}

			methods[method.Signature(true, false)] = comparedMethod
		}
	}

	for _, m := range methods {
		result.Methods = append(result.Methods, m)
	}

	// Sort methods in result
	sort.Slice(result.Methods, func(i, j int) bool {
		return library.CompareMethods(result.Methods[i].Method, result.Methods[j].Method)
	})

	return result
}
