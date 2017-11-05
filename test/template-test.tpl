======= START OF TEMPLATE ======
Env           : {{ Env "MYVAR" }}
ToUpper       : {{ ToUpper "Text" }}
ToLower       : {{ ToLower "Text" }}
Title         : {{ Title "TEXT" }}
Repeat        : {{ Repeat "TEXT" 3 }}
Replace       : {{ Replace "TextText" "ex" "aa" 1 }}
Trim          : {{ Trim " ! Text ! " "! " }}
TrimSpace     : {{ TrimSpace " ! Text ! " }}
File          : {{ File "./test/a_file.txt" }}
File->ToUpper : {{ File "./test/a_file.txt" | ToUpper }}
Date          : {{ Date }}
Data          : {{ .root.var1 }}
Format        : {{ Format "%s - %s" .root.var1 .root.var2 }}

======= END OF TEMPLATE ========
