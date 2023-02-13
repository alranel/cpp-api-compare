package library

import (
	"strings"
)

type Parameter struct {
	ParameterName string
	Type          string
	DefaultValue  string
}

type Method struct {
	MethodName   string
	ReturnType   string
	Ctor         bool
	Parameters   []Parameter
	DoesNotExist bool
}

type Class struct {
	ClassName    string
	Methods      []Method
	IsSingleton  bool
	DoesNotExist bool
}

type Library struct {
	LibraryName string
	Classes     map[string]Class
	Defines     map[string]string
}

// FindCompatibleMethod finds a compatible method in the current class
// by comparing the method name and the parameter types, but ignoring
// the return types.
func (class *Class) FindCompatibleMethod(method Method) *Method {
	for _, myMethod := range class.Methods {
		if myMethod.MethodName == method.MethodName {
			// If this method has fewer parameters, it is not compatible
			if len(myMethod.Parameters) < len(method.Parameters) {
				continue
			}

			// Check parameter types
			for i, parameter := range method.Parameters {
				if parameter.Type != myMethod.Parameters[i].Type {
					continue
				}
			}

			// If this method has more parameters but they are optional,
			// it is still compatible
			if len(myMethod.Parameters) > len(method.Parameters) {
				if myMethod.Parameters[len(method.Parameters)].DefaultValue == "" {
					continue
				}
			}

			return &myMethod
		}
	}
	return nil
}

func (method *Method) Overrides(other Method) bool {
	if method.MethodName != other.MethodName {
		return false
	}
	if len(method.Parameters) != len(other.Parameters) {
		return false
	}
	for i, parameter := range method.Parameters {
		if parameter.Type != other.Parameters[i].Type {
			return false
		}
	}
	return true
}

// Signature returns a string representing the method signature
func (method *Method) Signature(withParams bool, withReturnType bool) string {
	signature := method.MethodName + "("
	for i, parameter := range method.Parameters {
		signature += parameter.Type
		if withParams {
			signature += " " + parameter.ParameterName
		}
		if parameter.DefaultValue != "" {
			signature += " = " + parameter.DefaultValue
		}
		if i < len(method.Parameters)-1 {
			signature += ", "
		}
	}
	signature += ")"

	if withReturnType && !method.Ctor {
		signature = method.ReturnType + " " + signature
	}

	return signature
}

func CompareMethods(a, b Method) bool {
	if a.Ctor != b.Ctor {
		return a.Ctor
	}
	cmp := strings.Compare(a.MethodName, b.MethodName)
	if cmp == 0 {
		return len(a.Parameters) < len(b.Parameters)
	}
	return cmp == -1
}
