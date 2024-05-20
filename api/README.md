# Generating API schemas

This document describes the schema generation workflow from APIs defined using protobuf, and customizations that can be applied to the generated schemas.


## Worklow

### Steps

Update necessary dependencies:
```bash
make install-go-tools
```

Add or update proto files.

Generate the schemas:
```bash
make generate-apis
```

### How it works

APIs defined in proto files are transpiled to CRDs using a combination of tools and plugins. At a high level, it works as follows:

1. The `make generate-apis` command invokes the code generator tooling provided by [skv2](https://github.com/solo-io/skv2/).

1. `skv2` processes proto files using `protoc`, and uses the [protoc-gen-openapi](https://github.com/solo-io/protoc-gen-openapi) plugin to generate OpenAPI v3 schemas from the protobuf descriptors parsed by the plugin. `skv2` invokes the `protoc` command with the `--openapi_out` argument to invoke the `protoc-gen-openapi` plugin for schema generation.

1. `protoc-gen-openapi` writes the schemas to output files, that are then processed by `skv2`.

1. `skv2` post-processes the generated schemas to render them as CRDs in the expected shape and form.

## Schema customizations

Schema customizations are provided using `+kubebuilder` style comment markers in the leading comments of proto messages and fields. The comment should be attached to the leading comment block of the message or field. For readability and consistency, schema markers should be grouped together and separated from the message or field description comments by an empty comment line `//`.

### Style

Good:
```proto
// This is a message
// that does foo.
//
// +kubebuilder:<validation rule>
message Foo {
  // This is a string that
  // does nothing.
  //
  // +kubebuilder:<validation rule>
  string s = 1;
}
```

Bad:
```proto
// This is a message
// that does foo.

// +kubebuilder:<validation rule>
message Foo {
  // This is a string that
  // does nothing.
  // +kubebuilder:<validation rule>
  //
  // +kubebuilder:<validation rule>
  string s = 1;
}
```
In the "Bad" example:
1. The schema marker for message `Foo` is detatched from its leading comment, so the comment will not be a part of the generated message description. It should be linked by an empty comment line.
1. The schema markers for field `s` are not grouped together and are not separated from the field comments by an empty comment line.


### Type overrides

By default, the schemas generated from the proto files uses schema types that are inferred from the proto type, i.e., if a field is a `string`, the OpenAPI schema for that field will be of type `string`. This is applicable to all the data types.

In some cases, we require to be able to override the default type for two reasons:
1. The type in the generated schema needs to be represented differently than the proto field. For e.g., we may want the schema to be an opaque object type or value.
1. Messages in proto files could be recursive, in which case schema generation can experience an infinite recursion while trying to transpile protobuf to OpenAPI schemas; we want to avoid that.

The type in the schema for proto field can be overridden by specifying the `+kubebuilder:validation:Type=` marker in the leading comments attached to the field.

Consider the following example proto:
```proto
message Foo {
  // This is a recursive object
  //
  // +kubebuilder:validation:Type=object
  repeated Recursive obj = 1;
}

message Recursive {
  Recursive r = 1;
}
```

Without the `+kubebuilder:validation:Type=object` or `+kubebuilder:validation:Type=value` comment marker, the schema generation for this proto fill crash due to stack overflow. These markers terminate the schema generation of the corresponding field.

Structurally, both `Type=object` and `Type=value` produce a similar schema, except that when using `Type=object`, the schema is not annotated with `type: object`. `Type=object` should be used with the `google.protobuf.Struct` proto type, whereas `Type=value` should be used with the `google.protobuf.Value` proto type.

Schema corresponding to `Type=object`:
```yaml
obj:
  type: object
  x-kubernetes-preserve-unknown-fields: true
```

Schema corresponding to `Type=value`:
```yaml
obj:
  x-kubernetes-preserve-unknown-fields: true
```

Resursive fields should always use the `+kubebuilder:validation:Type=object` or `+kubebuilder:validation:Type=value` comment marker.

### Using Kubernetes CEL validation rules

Validation rules using the [Common Expression Language](https://github.com/google/cel-spec) can be used to custom resource values. The rules can be specified in the CRD using the `x-kubernetes-validations` extension header as follows:
```yaml
openAPIV3Schema:
    type: object
    properties:
      spec:
        type: object
        x-kubernetes-validations:
          - rule: "self.minReplicas <= self.replicas"
            message: "replicas should be greater than or equal to minReplicas."
```

In order to generate CEL validation rules in CRDs, the `+kubebuilder:validation:XValidation:` comment marker can be used with proto messages and fields.

#### Anatomy of CEL expressions

An expression specified by the `+kubebuilder:validation:XValidation:` marker can contain the following fields, where the value must be enclosed in double quotes `".."`:

1. `rule` [required]: e.g., `// +kubebuilder:validation:XValidation:rule="self.x > 0"`

1. `message` [optional]: e.g., `// +kubebuilder:validation:XValidation:rule="self.x > 0",message="x must be greater than 0"`

1. `messageExpression` [optional]: e.g., `// +kubebuilder:validation:XValidation:rule="self.x > 0",messageExpression="'x must be greater than 0, got: ' + string(self.x)"`
  > Note: single quotes must be used within the double quoted value of a `messageExpression`.


Multiple expressions for a message or field may be grouped together:
```proto
// This is a top-level message.
//
// +kubebuilder:validation:XValidation:rule="self.x > 5",message="x must be greater than 5"
// +kubebuilder:validation:XValidation:rule="has(self.foo) && self.foo.baz != 'invalid'",messageExpression="'foo.baz is invalid:' + self.foo.baz"
message Msg {
  ...
}
```

#### Examples

Consider the following example:
```proto
syntax = "proto3";

package test;

// This is a top-level message.
//
// +kubebuilder:validation:XValidation:rule="self.x > 5",message="x must be greater than 5"
message Msg {

  // x is an integer
  // that must be greater than 5.
  int32 x = 1;

  // a is a nested
  // field.
  //
  // +kubebuilder:validation:XValidation:rule="self.a != ''",message="a cannot be empty"
  Nested w = 2;

  // z is an integer that must be less than 100
  //
  // +kubebuilder:validation:XValidation:rule="self < 100",messageExpression="'z got invalid value: ' + self"
  uint32 z = 3;

  // Nested message
  message Nested {
    string a = 1;
  }
}
```

This will transpile to the following OpenAPI schema:
```yaml
schemas:
  test.Msg:
    description: This is a top-level message.
    properties:
      w:
        description: |-
          a is a nested
          field.
        properties:
          a:
            type: string
        type: object
        x-kubernetes-validations:
        - message: a cannot be empty
          rule: self.a != ''
      x:
        description: |-
          x is an integer
          that must be greater than 5.
        format: int32
        type: integer
      z:
        description: z is an integer that must be less than 100
        maximum: 4294967295
        minimum: 0
        type: integer
        x-kubernetes-validations:
        - messageExpression: '''z got invalid value: '' + self'
          rule: self < 100
    type: object
    x-kubernetes-validations:
    - message: x must be greater than 5
      rule: self.x > 5
```

As seen above, CEL expressions can be specified for proto messages and fields using the `+kubebuilder:validation:XValidation:` command marker.

#### Common Complex Rules

It's common for several Gloo CRDs to share a field type and related constraints.
When the CEL rule(s) to enforce these constraints are complex, this section will be used as a reference
to document the rules themselves as well as any explanatory comments to help developers understand them.

##### `ValidateDuration()` : google.protobuf.Duration

Any field that is validated by Istio's [`ValidateDuration()` function](https://github.com/istio/istio/blob/1.21.1/pkg/config/validation/validation.go#L1606)
must only represent durations greater than or equal to one millisecond
and must only represent durations with granularity greater than or equal to one millisecond.
This can be largely enforced in CEL with the following rules:

```
// +kubebuilder:validation:XValidation:rule="duration(self) >= duration('1ms')",message="Must be greater than or equal to 1ms"
// +kubebuilder:validation:XValidation:rule="!self.contains('ns') && !self.contains('us')",message="Cannot have granularity smaller than 1 millisecond"
// +kubebuilder:validation:XValidation:rule="(duration(self)-duration('1ns')).getMilliseconds() == duration(self).getMilliseconds()-1",message="Cannot have granularity smaller than 1 millisecond"
```
- The first rule is straight forward; it checks that the duration is greater than or equal to 1 millisecond.

- The second rule checks that units smaller than a millisecond are not used, which handles basic granularity cases.

- The third rule checks that the field value is less than 1 nanosecond above an integer number of milliseconds. Valid and invalid examples can be seen below.

  ![image](https://github.com/solo-io/gloo-mesh-enterprise/assets/11544864/d67975ae-bbd6-4c7d-aaf9-ddb6c838bfd3)

  `getMilliseconds()` returns an integer count of the total milliseconds in the given value. In other words, it rounds down to the nearest millisecond.
  <br>If the field value is 2 milliseconds, `duration(self).getMilliseconds()` will return 2. Subtracting 1ns from this value and then using `getMilliseconds()` will return 1.
  <br>If the field value is 2ms0.5ns, `duration(self).getMilliseconds()` will return 2. Subtracting 1ns from this value and then using `getMilliseconds()` will result in 1, just as the previous example, and this is valid.
  <br>However, if the field value is 2.5ms0.5ns, `duration(self).getMilliseconds()` will still return 2. But subtracting 1ns from this value will not move it across a whole millisecond line, and using `getMilliseconds()` will still return 2. From this we know that we have an invalid value.

  This is a good approximation of millisecond granularity and will avoid breaking Istio,
as 1 nanosecond is the smallest value `time.Duration` recognizes.
However, "1s0.5ns" will still be accepted by k8s and will show up as 1 second in our generated Istio resource.

  Note: if we make the '1ns' in this rule smaller, it's ignored (again since 1 nanosecond is the smallest value `time.Duration` recognizes), and this rule becomes `duration(self).getMilliseconds() == duration(self).getMilliseconds()-1` which blocks all values.

With the combination of these rules, "1s0.0000001ms" will still be accepted and will result in a 1-second delay.

### Kubebuilder markers

[Kubebuilder comment markers](https://book.kubebuilder.io/reference/markers/crd-validation#crd-validation) can be used to specify field and message level schema constraints. A marker may be applicable either to fields, messages, or both, depending on the marker.

Consider the following example:
```proto
syntax = "proto3";

package test;

// This is a top-level message.
//
// +kubebuilder:pruning:PreserveUnknownFields
message Msg {
  // +kubebuilder:pruning:PreserveUnknownFields
  Nested nested = 1;

  // +kubebuilder:validation:Maximum=100
  // +kubebuilder:validation:Minimum=5
  // +kubebuilder:validation:ExclusiveMaximum=true
  // +kubebuilder:validation:ExclusiveMinimum=true
  // +kubebuilder:validation:MultipleOf=2
  int32 a = 2;

  // +kubebuilder:validation:MinItems=1
  // +kubebuilder:validation:MaxItems=5
  // +kubebuilder:validation:UniqueItems=true
  repeated string blist = 3;

  // This is a nested message.
  //
  // +kubebuilder:validation:MinProperties=1
  // +kubebuilder:validation:MaxProperties=2
  message Nested {
    // +kubebuilder:validation:Pattern="^[a-zA-Z0-9_]*$"
    // +kubebuilder:validation:Required
    string a = 1;

    // +kubebuilder:validation:Enum=Allow;Forbid;Replace
    // +kubebuilder:validation:Required
    string b = 2;

    // +kubebuilder:validation:MaxLength=100
    // +kubebuilder:validation:MinLength=1
    string c = 3;

    // +kubebuilder:validation:Format=date-time
    string d = 4;

    // +kubebuilder:validation:XIntOrString
    string int_or_string = 5;

    // +kubebuilder:default=forty-two
	  // +kubebuilder:example=forty-two
    string default_value = 6;

    // Schemaless field
    //
    // +kubebuilder:validation:Schemaless
    string schemaless = 7;

    // +kubebuilder:validation:EmbeddedResource
    // +kubebuilder:validation:Nullable
    string embedded = 8;
  }
}
```

This will transpile to the following OpenAPI schema:
```yaml
schemas:
  test6.Msg:
    description: This is a top-level message.
    properties:
      a:
        exclusiveMaximum: true
        exclusiveMinimum: true
        format: int32
        maximum: 100
        minimum: 5
        multipleOf: 2
        type: integer
      blist:
        items:
          type: string
        maxItems: 5
        minItems: 1
        type: array
        uniqueItems: true
      nested:
        maxProperties: 2
        minProperties: 1
        properties:
          a:
            pattern: ^[a-zA-Z0-9_]*$
            type: string
          b:
            enum:
            - Allow
            - Forbid
            - Replace
            type: string
          c:
            maxLength: 100
            minLength: 1
            type: string
          d:
            format: date-time
            type: string
          defaultValue:
            default: forty-two
            example: forty-two
            type: string
          embedded:
            nullable: true
            type: string
            x-kubernetes-embedded-resource: true
          intOrString:
            type: string
            x-kubernetes-int-or-string: true
          schemaless:
            description: Schemaless field
        required:
        - a
        - b
        type: object
        x-kubernetes-preserve-unknown-fields: true
    type: object
    x-kubernetes-preserve-unknown-fields: true
```
