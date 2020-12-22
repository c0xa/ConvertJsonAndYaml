
# :construction: 

- [:ru: RUS](./README.ru.md) - Русская документация
- [:uk: ENG](./README.md) - English documentation

# About
This is small application for conversion files written on Goland.
This application allows you to convert from json to yaml or  
from yaml to json with command-line arguments.

# Supported types

- [x] Yaml
- [x] Json

# Examples

```
go run . --input data.yaml --output data.json --from yaml --to json
```

```
go run . --input data.json--output data.yml   --from json --to yaml 
```
