# General template command using go templates


## Build and install
    make
    make install
    
## Usage
```
tmpl \
    -tmpl=./path/mytemple.tpl \
    -dest=./path/result.txt \
```
## Supported functions
| Function    | Example                                    |
|-------------|--------------------------------------------|
| Env         | {{ Env "MYVAR" }}                          | 
| ToUpper     | {{ ToUpper "Text" }}                       |
| ToLower     | {{ ToLower "Text" }}                       |
| Repeat      | {{ Repeat "TEXT" 3 }}                      | 
| Replace     | {{ Replace "TextText" "ex" "aa" 1 }}       |
| Trim        | {{ Trim " ! Text ! " "! " }}               |
| TrimSpace   | {{ TrimSpace " ! Text ! " }}               |
| File        | {{ File "./test/a_file.txt" }}             |
| Date        | {{ Date }}                                 | 
