package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/solo-apis/codegen"
)

func main() {
	var resources []string
	for _, group := range codegen.GlooGroups() {
		for _, resource := range group.Resources {
			resources = append(resources, resource.Kind)
		}
	}
	moduleRoot := util.GetModuleRoot()
	err := filepath.Walk(
		filepath.Join(moduleRoot, "api/gloo"),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(info.Name(), ".proto") {
				return nil
			}
			file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
			if err != nil {
				return err
			}
			defer file.Close()
			wholeFileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			// If the proto file is not from a relevant package skip it.
			wholeFileBytes, ok := isRelevantFile(wholeFileBytes)
			if !ok {
				return nil
			}

			var fileByLine []string
			scan := bufio.NewScanner(bytes.NewBuffer(wholeFileBytes))
			for scan.Scan() {
				line := scan.Text()

				if strings.Contains(line, "option (core.solo.io.resource)") ||
					strings.Contains(line, "solo-kit/api/v1/status.proto") ||
					strings.Contains(line, "solo-kit/api/v1/metadata.proto") {
					continue
				} else if strings.Contains(line, "core.solo.io.Metadata") ||
					strings.Contains(line, "core.solo.io.Status") {
					continue
				}

				fileByLine = append(fileByLine, line)
			}
			err = file.Truncate(0)
			if err != nil {
				return err
			}
			_, err = file.Seek(0, 0)
			if err != nil {
				return err
			}
			_, err = file.WriteString(strings.Join(fileByLine, "\n"))
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}

func isRelevantFile(file []byte) ([]byte, bool) {
	var relevantMessage bool
	glooGroups := codegen.GlooGroups()
	relevantTypes := getRelevantTypes(file, glooGroups)
	// If there are no relevant types in this file, return
	if len(relevantTypes) == 0 {
		return nil, false
	}
	for _, relevantType := range relevantTypes {
		// Make sure to add extra space after to check specifically for only Message
		oldMessageBytes := []byte(fmt.Sprintf("message %s ", relevantType))
		if bytes.Contains(file, oldMessageBytes) {
			oldMessageBytes = oldMessageBytes[:len(oldMessageBytes)-1]
			file = bytes.ReplaceAll(file, oldMessageBytes, append(oldMessageBytes, []byte("Spec")...))
			statusBytes := &bytes.Buffer{}
			if err := statusTemplate.Execute(statusBytes, relevantType); err != nil {
				panic(err)
			}
			file = append(file, statusBytes.Bytes()...)
			file = addImport(file)
			relevantMessage = true
		}
	}

	return file, relevantMessage
}

func getRelevantTypes(file []byte, glooGroups []model.Group) []string {
	var resources []string
	for _, group := range glooGroups {
		for _, resource := range group.Resources {
			if resource.Spec.Type.ProtoPackage != "" {
				if bytes.Contains(file, []byte(fmt.Sprintf("package %s", resource.Spec.Type.ProtoPackage))) {
					resources = append(resources, resource.Kind)
				}
			} else if resource.Status != nil && resource.Status.Type.ProtoPackage != "" {
				if bytes.Contains(file, []byte(fmt.Sprintf("package %s", resource.Spec.Type.ProtoPackage))) {
					resources = append(resources, resource.Kind)
				}
			} else if bytes.Contains(file, []byte(fmt.Sprintf("package %s", group.Group))) {
				resources = append(resources, resource.Kind)
			}
		}
	}
	return resources
}

func addImport(file []byte) []byte {
	// already present, no need to add
	if bytes.Contains(file, []byte(`import "google/protobuf/struct.proto"`)) {
		return file
	}
	idx := bytes.Index(file, []byte("import"))
	builder := &bytes.Buffer{}
	builder.Write(file[:idx])
	builder.Write([]byte("import \"google/protobuf/struct.proto\";\n"))
	builder.Write(file[idx:])
	return builder.Bytes()
}

var statusTemplate = template.Must(template.New("status").Parse(`

message {{ . }}Status {
	enum State {
			// Pending status indicates the resource has not yet been validated
			Pending = 0;
			// Accepted indicates the resource has been validated
			Accepted = 1;
			// Rejected indicates an invalid configuration by the user
			// Rejected resources may be propagated to the xDS server depending on their severity
			Rejected = 2;
			// Warning indicates a partially invalid configuration by the user
			// Resources with Warnings may be partially accepted by a controller, depending on the implementation
			Warning = 3;
	}
	// State is the enum indicating the state of the resource
	State state = 1;
	// Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty
	string reason = 2;
	// Reference to the reporter who wrote this status
	string reported_by = 3;
	// Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource
	map<string, {{ . }}Status> subresource_statuses = 4;

	// Opaque details about status results
	google.protobuf.Struct details = 5;
}

`))
