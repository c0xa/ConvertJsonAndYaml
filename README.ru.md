

# :construction: 

- [:ru: RUS](./README.ru.md) - Русская документация
- [:uk: ENG](./README.md) - English documentation

# О проекте
Это небольшая утилита для конвертации файлов,
написанная на Goland. Она позволяет конвертировать из JSON в YAML и из 
YAML  в JSON 
с помощью аргументов командной строки.

# Поддерживаемые типы

- [x] Yaml
- [x] Json

# Примеры


```
convert --input file.yaml --output file.json --from yaml --to json
```

```
convert --input file.json--output file.yaml   --from json --to yaml 
```