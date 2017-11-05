# Simple template command using go templates

## Build and install
    make
    make install
    
## Usage
```
tmpl \
    -tmpl=./path/mytemple.tpl \
    -data=./path/some.yaml
```
## Supported functions
| Function    | Example                                    |
|-------------|--------------------------------------------|
| Env         | {{ Env "MYVAR" }}                          | 
| Format      | {{ Format "%s - %s" "Test1" "Test2" }}     |
| ToUpper     | {{ ToUpper "Text" }}                       |
| ToLower     | {{ ToLower "Text" }}                       |
| Repeat      | {{ Repeat "TEXT" 3 }}                      | 
| Replace     | {{ Replace "TextText" "ex" "aa" 1 }}       |
| Trim        | {{ Trim " ! Text ! " "! " }}               |
| TrimSpace   | {{ TrimSpace " ! Text ! " }}               |
| File        | {{ File "./test/a_file.txt" }}             |
| Date        | {{ Date }}                                 |
