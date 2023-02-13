package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/alranel/cpp-api-compare/internal/compare"
	"github.com/alranel/cpp-api-compare/internal/config"
	"github.com/alranel/cpp-api-compare/internal/library"
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the API comparison",
	Long:  `This command performs the API comparison between the supplied libraries.`,
	Run:   runRun,
}

func init() {
	runCmd.PersistentFlags().StringP("config", "c", "config.yml", "The YAML configuration file.")
	runCmd.PersistentFlags().StringP("output", "o", "report.md", "The file to write the Markdown report to.")
	runCmd.PersistentFlags().Bool("keep-build-dir", false, "Do not delete build directories (for debugging purposes).")
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, cliArguments []string) {
	// Parse the supplied configuration file
	var config config.Config
	{
		path, _ := cmd.Flags().GetString("config")
		config.ReadFromFile(path)
	}

	// Parse other command line options
	keepBuildDir, _ := cmd.Flags().GetBool("keep-build-dir")

	// Read the libraries
	var libs []library.Library
	for _, lib := range config.Libraries {
		fmt.Printf("Processing %s...\n", lib.Name)
		parsedLib := library.ReadLibrary(lib, keepBuildDir)
		if parsedLib == nil {
			fmt.Printf("  Error reading library %s from paths %v\n", lib.Name, lib.Include)
			continue
		}
		libs = append(libs, *parsedLib)
	}

	// Comparison contains the result of a comparison between two libraries.
	fmt.Printf("Performing comparison...\n")
	type Comparison struct {
		GlobalFunctions compare.CompareResult
		Classes         map[string]compare.CompareResult
	}
	result := Comparison{
		GlobalFunctions: compare.CompareResult{ClassName: "_"},
		Classes:         make(map[string]compare.CompareResult),
	}

	// Perform comparison
	fmt.Printf("  * classes\n")
	for _, className := range config.Classes {
		result.Classes[className] = compare.Compare(libs, className)
	}
	fmt.Printf("  * singletons\n")
	for _, className := range config.Singletons {
		result.Classes[className] = compare.Compare(libs, className)
	}

	// Compare global functions (but only those that are present in config)
	fmt.Printf("  * global functions\n")
	{
		// Make a map of all functions that are present in the config for quick lookup
		functionMap := make(map[string]bool)
		for _, functionName := range config.Functions {
			functionMap[functionName] = true
		}

		// Compare all detected functions
		compResult := compare.Compare(libs, "_")
		for _, method := range compResult.Methods {
			if functionMap[method.MethodName] {
				result.GlobalFunctions.Methods = append(result.GlobalFunctions.Methods, method)
				delete(functionMap, method.MethodName)
			}
		}

		// Warn about functions that were not found
		for functionName := range functionMap {
			fmt.Printf("    Warning: function %s not found in any libs\n", functionName)
		}
	}

	// Generate a sorted list of all classes
	var classNames []string
	for _, classResult := range result.Classes {
		if classResult.ClassName != "_" {
			classNames = append(classNames, classResult.ClassName)
		}
	}
	sort.Strings(classNames)

	// Write output file
	outputPath, _ := cmd.Flags().GetString("output")
	fmt.Printf("Writing report to %s\n", outputPath)
	f, _ := os.Create(outputPath)
	f.WriteString("# API comparison\n\n")
	{
		f.WriteString("This document compares the API of the following libraries:\n\n")
		var rows [][]string
		for i, lib := range libs {
			row := []string{lib.LibraryName, config.Libraries[i].Description}
			if config.Libraries[i].Link != "" {
				row[0] = fmt.Sprintf("[%s](%s)", lib.LibraryName, config.Libraries[i].Link)
			}
			rows = append(rows, row)
		}
		table, err := markdown.NewTableFormatterBuilder().
			WithPrettyPrint().
			Build("Library", "Description").
			Format(rows)
		if err != nil {
			fmt.Printf("Error generating table: %+v\n", err)
		}
		f.WriteString(table + "\n")
	}
	{
		f.WriteString("## Table of Contents\n\n")
		i := 1
		if len(result.GlobalFunctions.Methods) > 0 {
			f.WriteString("1. [Global Functions](#global-functions)\n")
			i++
		}
		for _, className := range classNames {
			f.WriteString(fmt.Sprintf("%d. [%s](#%s)\n", i, className, strings.ToLower(className)))
			i++
		}
		f.WriteString("\n")
	}
	if len(result.GlobalFunctions.Methods) > 0 {
		f.WriteString(printClass(libs, result.GlobalFunctions))
	}
	for _, className := range classNames {
		f.WriteString(printClass(libs, result.Classes[className]))
	}
	{
		f.WriteString("_Generated on ")
		t := time.Now().UTC()
		f.WriteString(t.Format(time.RFC1123))
		f.WriteString(" with [cpp-api-compare](https://github.com/alranel/cpp-api-compare)._\n\n")
	}
	f.Close()
}

func printClass(libs []library.Library, classResult compare.CompareResult) string {
	var out string
	if classResult.ClassName == "_" {
		out += "## Global Functions\n\n"
	} else {
		out += fmt.Sprintf("## %s\n\n", classResult.ClassName)
	}

	// Prepare tabular data
	columnNames := []string{"", ""}
	for _, lib := range libs {
		columnNames = append(columnNames, string(lib.LibraryName))
	}
	rows := [][]string{}

	commonMethodsSameRT := classResult.CommonMethods(true)
	if len(commonMethodsSameRT) > 0 {
		rows = append(rows, []string{"", "**Fully compatible**"})
		for range libs {
			rows[0] = append(rows[0], "")
		}
		for _, method := range commonMethodsSameRT {
			row := []string{
				"`" + method.ReturnType + "`",
				"`" + method.Signature(true, false) + "`",
			}
			if method.ReturnType == "" {
				row[0] = ""
			}
			for range libs {
				row = append(row, "✔️")
			}
			rows = append(rows, row)
		}
	}

	commonMethodsDifferingRT := classResult.CommonMethods(false)
	if len(commonMethodsDifferingRT) > 0 {
		rows = append(rows, []string{"", "**Differing return values**"})
		for range libs {
			rows[len(rows)-1] = append(rows[len(rows)-1], "")
		}
		for _, method := range commonMethodsDifferingRT {
			row := []string{
				"",
				"`" + method.Signature(true, false) + "`",
			}
			for _, rt := range method.ReturnTypes {
				row = append(row, "`"+rt+"`")
			}
			rows = append(rows, row)
		}
	}

	nonCommonMethods := classResult.NonCommonMethods()
	if len(nonCommonMethods) > 0 {
		rows = append(rows, []string{"", "**Limited compatibility**"})
		for range libs {
			rows[len(rows)-1] = append(rows[len(rows)-1], "")
		}
		for _, method := range nonCommonMethods {
			row := []string{
				"`" + method.ReturnType + "`",
				"`" + method.Signature(true, false) + "`",
			}
			if method.ReturnType == "" {
				row[0] = ""
			}
			for _, p := range method.Presence {
				if p {
					row = append(row, "✔️")
				} else {
					row = append(row, "")
				}
			}
			rows = append(rows, row)
		}
	}

	// Generate table
	table, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build(columnNames...).
		Format(rows)
	if err != nil {
		fmt.Printf("Error generating table: %+v\n", err)
	}
	out += table + "\n"

	return out
}
