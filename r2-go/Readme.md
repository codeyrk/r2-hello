### Scripting R2 using Golang

Reference: https://github.com/radareorg/radare2-r2pipe/blob/master/go/example/example.go

Usage: 

First Create a folder .../binaries with following structure
```
➜  Data ll -R binaries
total 0
drwxr-xr-x  2 yogesh  staff    64B Sep 15 18:11 input
drwxr-xr-x  5 yogesh  staff   160B Sep 15 18:10 output
drwxr-xr-x  6 yogesh  staff   192B Sep 15 18:11 processed
```
- Put your binary files in input folder.

- Configure the r2 commands in config.yml file.
    ```
    r2commands:
    - cmd: "fs strings; fj"
      idx: "idx_strings"

    - cmd: "ij"
      idx: idx_binary_info
    ```
    Here, all r2comands specified in cmd will be run. 
    For example output of "fs strings; fj" will be written to 
    ```
    binaries/output/idx_strings.json
    ```
    and output of "ij" will be written to
    ```
    binaries/output/idx_binary_info.json
    ```
- Build
    ```
    docker build -f Dockerfile -t local/r2go .
    ```
- Run
    ```
    docker run -ti -v /Users/yogesh/Work/Data/binaries:/binaries local/r2go
    ```

Functionality: Done

Dockerfile: Done

#### Currently acheived scripting on OSX, using Radare2 and Golang libs. 
- The configured commands can be run on any binary placed in input folder.
- The idea is to generate json records for several analysis flows
- Each analysis flow will have a different set of commands mostly ending with a list returning json output. 
Example:
    ```
    fs strings; fj
    ```

    ```
    ij
    ```
- Configured set of command will be run on all input binaries. Output from different flows will be loaded in seperate elastic indices.
- Dashboard can be created on top of it.

