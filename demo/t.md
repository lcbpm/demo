好的，这里是将我之前的回答翻译成中文：

要在 GoLand 中自动填充模板文件的值，你通常会利用 Go 提供的 `text/template` 或 `html/template` 模板引擎。下面，我将分步骤说明如何创建一个模板、定义占位符值，然后在 Go 代码中动态填充它们：

### 第一步：定义你的模板

创建一个模板文件，例如 `template.tmpl`：

```gotemplate
type {{.StructName}} struct {
    {{- range .Fields }}
    {{ .Name }} {{ .Type }} `{{ .Tag }}`
    {{- end }}
}

func New{{ .StructName }}({{ range .Fields }}{{ .Name }} {{ .Type }}, {{ end }}) *{{ .StructName }} {
    return &{{ .StructName }}{
        {{- range .Fields }}
        {{ .Name }}: {{ .Name }},
        {{- end }}
    }
}

{{- range .Fields }}
func (s *{{ $.StructName }}) Set{{ .Name }}(val {{ .Type }}) {
    s.{{ .Name }} = val
}

func (s *{{ $.StructName }}) Get{{ .Name }}() {{ .Type }} {
    return s.{{ .Name }}
}
{{- end }}
```

### 第二步：在 Go 代码中加载和解析模板

创建你的 Go 文件，这个文件将加载模板并用值填充它。以下是 `main.go` 的示例：

```go
package main

import (
    "os"
    "text/template"
)

type Field struct {
    Name string
    Type string
    Tag  string
}

type StructTemplate struct {
    StructName string
    Fields     []Field
}

func main() {
    tmpl, err := template.ParseFiles("template.tmpl")
    if err != nil {
        panic(err)
    }

    data := StructTemplate{
        StructName: "TalkerGiftConfig",
        Fields: []Field{
            {"ID", "int64", "`json:\"id\"`"},
            {"GiftType", "int32", "`json:\"gift_type\"`"},
            {"GiftName", "string", "`json:\"gift_name\"`"},
            // 根据需要添加更多字段...
        },
    }

    file, err := os.Create("generated_struct.go")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    err = tmpl.Execute(file, data)
    if err != nil {
        panic(err)
    }

    println("Struct and methods generated successfully.")
}
```

### 第三步：运行你的 Go 程序

- 保存你的 `template.tmpl` 和 `main.go` 文件。
- 运行你的 Go 程序以生成新的结构体和方法：

```sh
go run main.go
```

这将创建一个名为 `generated_struct.go` 的文件，其中包含你填充值后的结构体和方法。

通过使用模板引擎并在代码中定义你的值，你可以很容易地在一个地方管理和更新你的代码，使其更易于维护。

如果你遇到任何问题或有更多问题，欢迎随时问我！ 😊