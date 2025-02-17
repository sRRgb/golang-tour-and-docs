## Formatting with `fmt.Printf`

The `fmt.Printf` function allows you to format and print values. Here's a breakdown of the common format specifiers:

| Specifier | Description                                | Example                                  | Output                              |
|-----------|--------------------------------------------|------------------------------------------|-------------------------------------|
| `%v`      | Default format for any value               | `fmt.Printf("%v", 123)`                  | `123`                               |
| `%+v`     | Default format with struct field names     | `fmt.Printf("%+v", Person{"Alice", 30})` | `{Name:Alice Age:30}`               |
| `%#v`     | Go-syntax representation                   | `fmt.Printf("%#v", Person{"Alice", 30})` | `main.Person{Name:"Alice", Age:30}` |
| `%T`      | Prints the type                            | `fmt.Printf("%T", 3.14)`                 | `float64`                           |
| `%t`      | Boolean (`true` or `false`)                | `fmt.Printf("%t", true)`                 | `true`                              |
| `%b`      | Binary representation                      | `fmt.Printf("%b", 10)`                   | `1010`                              |
| `%c`      | Unicode character                          | `fmt.Printf("%c", 65)`                   | `A`                                 |
| `%d`      | Decimal integer                            | `fmt.Printf("%d", 123)`                  | `123`                               |
| `%o`      | Octal integer                              | `fmt.Printf("%o", 64)`                   | `100`                               |
| `%q`      | Quoted character or string                 | `fmt.Printf("%q", 65)`                   | `'A'`                               |
|           |                                            | `fmt.Printf("%q", "Hello\nGo")`          | `"Hello\nGo"`                       |
| `%x`      | Hexadecimal (lowercase)                    | `fmt.Printf("%x", 255)`                  | `ff`                                |
| `%X`      | Hexadecimal (uppercase)                    | `fmt.Printf("%X", 255)`                  | `FF`                                |
| `%U`      | Unicode format (U+1234 'char')             | `fmt.Printf("%U", 9731)`                 | `U+2603 '☃'`                        |
| `%e`      | Scientific notation (lowercase) for floats | `fmt.Printf("%e", 123.456)`              | `1.234560e+02`                      |
| `%E`      | Scientific notation (uppercase) for floats | `fmt.Printf("%E", 123.456)`              | `1.234560E+02`                      |
| `%f`      | Decimal floating-point                     | `fmt.Printf("%f", 123.456)`              | `123.456000`                        |
| `%F`      | Same as `%f`                               | `fmt.Printf("%F", 123.456)`              | `123.456000`                        |
| `%g`      | Compact float format (chooses %e or %f)    | `fmt.Printf("%g", 123.456)`              | `123.456`                           |
| `%G`      | Similar to `%g`, using %E if needed        | `fmt.Printf("%G", 123.456)`              | `123.456`                           |
| `%s`      | Plain string                               | `fmt.Printf("%s", "Hello")`              | `Hello`                             |
| `%p`      | Pointer address in hexadecimal             | `fmt.Printf("%p", &x)`                   | `0x...` (an address)                |
| `%%`      | Literal percent sign                       | `fmt.Printf("%%")`                       | `%`                                 |

**Note:**  The example using `%p` requires a variable `x` to have been declared.  The output will be the memory address of that variable.  The example for `%+v` and `%#v` assumes a struct named `Person` with fields `Name` and `Age`.

This table provides a concise overview of the most commonly used format specifiers. Refer to the official Go documentation for a complete list and more details.
