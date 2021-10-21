## Сustomization ATC config file(.atc.yaml)

Config file use syntax Yaml

### USAGE:
## Action Inputs
- **Path**: Path to your package manager. Path shouldn't have preffix "/".
- **Behavior**: What is commits will be used for create tag: `After` or `Before`.
- **Template**: Template your tag. Template should conteins {{.Version}}.
- **Branch**: .
- **RegexStr**: .[Contribution guidelines for this project](https://github.com/DmitriyLewen/atcDemo/blob/master/README.md#action-inputs)

## Example
This example will create a release when a tag is pushed:

```yaml
path: "contents/pom.xml"
behavior: "after"
template: "v{{.Version}}"
```
