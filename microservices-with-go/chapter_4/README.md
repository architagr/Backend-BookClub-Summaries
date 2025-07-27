# 📘 Chapter 4 – Serialization

## Prerequisites

Before working with Protocol Buffers, ensure you have the following setup:

1. Install Protocol Buffer Compiler
   [Official installation guide](https://protobuf.dev/installation/)

   **For macOS:**

   ```bash
   brew install protobuf
   ```

2. Install required Go packages:

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Why Do We Need Serialization?

Serialization helps in encoding structured data so it can be easily stored, transmitted, and reconstructed. Typical use cases include:

- Transferring data between backend and frontend
- Communication between microservices
- Encoding/decoding data for persistence
- Storing configuration files
- Logging structured events

## Types of Serialization Formats

1. JSON (JavaScript Object Notation)
2. XML (Extensible Markup Language)
3. YAML (YAML Ain’t Markup Language)
4. Apache Thrift
5. Apache Avro
6. Protocol Buffers (Protobuf)

### 1. JSON

Example:

```json
{
  "id": 1,
  "name": "TV",
  "categoryID": 1
}
```

#### Benefits

- Widely supported across programming languages
- Natively supported in all browsers
- Easy to read and debug

#### Limitations

- Larger payload size compared to Protocol Buffers
- Slower encoding/decoding speed

### 2. XML

Example:

```xml
<subcategory>
    <id>1</id>
    <name>TV</name>
    <categoryID>1</categoryID>
</subcategory>
```

#### Limitation

- Produces larger output → slower serialization/deserialization

### 3. YAML

Example:

```yml
subcategory:
  id: 1
  name: TV
  categoryID: 1
```

#### Benefits of YAML

- Human-readable
- Supports comments → useful for configuration files

#### Limitations of YAML

- Extra indentation → larger size → slower performance
- More complex to parse than JSON

### 4. Protocol Buffers (Protobuf)

Introduced by Google in 2008, Protocol Buffers are a highly efficient binary serialization format.

#### Key Benefits

1. Much smaller payload size
1. Fastest serialization/deserialization speed
1. Ability to define services along with data structures
1. Supports code generation for multiple languages
1. Backed by Google → mature and widely adopted

#### Schema Advantages

1. Explicit schema definition: Decoupled from code → easy to read
1. Code generation: Generate boilerplate code automatically
1. Cross-language support: Use the same .proto schema for Go, Python, Java, etc.
