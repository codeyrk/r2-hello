### Scripting R2 using Golang

Reference: https://github.com/radareorg/radare2-r2pipe/blob/master/go/example/example.go

Usage: TODO

Functionality: TODO

Dockerfile: TODO

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

