# Struct Notes - Go Programming

## Basic Topics Covered

### 1. Bufio and Strings Package
- **`bufio` Package**: Provides buffered I/O operations, allowing for efficient reading and writing of data.
    - **`NewReader`**: Creates a new buffered reader.
    - **`ReadString`**: Reads input until a specified delimiter.
- **`strings` Package**: Useful for manipulating string data.
    - **`ReplaceAll`**: Replaces all occurrences of a substring within a string.
    - **`Trim`**: Removes specified characters from the beginning and end of a string.
    - **`ToLower`**: Converts a string to lowercase.

### 2. Encoding/JSON Package
- **`encoding/json`**: This package helps to handle JSON data in Go.
    - **`Marshal`**: Converts Go data structures (like structs) into JSON format.

### 3. Struct Tags in Go
Struct tags are special annotations attached to struct fields to provide additional metadata that can be used by libraries, frameworks, or functions.
- Example:
    ```go
    type Person struct {
        FirstName string `json:"first_name"`  // JSON field name: first_name
        LastName  string `json:"last_name"`   // JSON field name: last_name
        Age       int    `json:"age"`         // JSON field name: age
    }
    ```
- In the above example:
    - **`json:"first_name"`** specifies the JSON field name when the struct is marshaled into JSON. It helps map Go struct fields to specific JSON keys during serialization and deserialization.
