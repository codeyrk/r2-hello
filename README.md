# r2-hello

## Notes

1.  Installation (Windows)

    - Download static binaries from here:
      https://github.com/radareorg/radare2/releases/download/4.5.1/radare2-windows-static-4.5.1.zip

    Contents:

    ```
    D:\Work\Tools\radare2-install\bin (master -> origin)
    λ ll
    total 145500
    -rwxr-xr-x 1 yogesh 197121 12375552 Sep 14 09:25 r2agent.exe
    -rwxr-xr-x 1 yogesh 197121 17101 Sep 14 09:25 r2pm
    -rwxr-xr-x 1 yogesh 197121 445440 Sep 14 09:25 r2r.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 rabin2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 radare2.exe
    -rwxr-xr-x 1 yogesh 197121 12371968 Sep 14 09:25 radiff2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 rafind2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 ragg2.exe
    -rwxr-xr-x 1 yogesh 197121 12383744 Sep 14 09:25 rahash2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 rarun2.exe
    -rwxr-xr-x 1 yogesh 197121 12374016 Sep 14 09:25 rasign2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 rasm2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:25 rax2.exe
    -rwxr-xr-x 1 yogesh 197121 12372992 Sep 14 09:27 r2.exe
    -rwxr-xr-x 1 yogesh 197121 17101 Sep 14 10:19 r2pm.exe
    ```

2.  Getting Started

        Download ELF file from here: https://github.com/ITAYC0HEN/A-journey-into-Radare2/blob/master/Part%201%20-%20Simple%20crackme/megabeets_0x1

        Check help

              ```
              r2 -h
              ```

        Check binary info

              ```
              $ rabin2 -I megabeets_0x1
                arch x86
                baddr 0x8048000
                binsz 6220
                bintype elf
                bits 32
                canary false
                class ELF32
                compiler GCC: (Ubuntu 5.4.0-6ubuntu1~16.04.4) 5.4.0 20160609
                crypto false
                endian little
                havecode true
                intrp /lib/ld-linux.so.2
                laddr 0x0
                lang c
                linenum true
                lsyms true
                machine Intel 80386
                maxopsz 16
                minopsz 1
                nx false
                os linux
                pcalign 0
                pic false
                relocs true
                relro partial
                rpath NONE
                sanitiz false
                static false
                stripped false
                subsys linux
                va true
            ```

Load the binary in r2

        ```
        D:\Work\Tools\cmder (master -> origin)
        λ cd D:\Work\Tools\radare2-install\bin

        D:\Work\Tools\radare2-install\bin (master -> origin)
        λ r2 ..\..\..\Data\megabeets_0x1
        -- Setup dbg.fpregs to true to visualize the fpu registers in the debugger view.
        [0x08048370]>
        ```

Check entrypoint information

        D:\Work\Tools\radare2-install\bin (master -> origin)
        λ r2 ..\..\..\Data\megabeets_0x1
        -- Setup dbg.fpregs to true to visualize the fpu registers in the debugger view.
        [0x08048370]> ie
        [Entrypoints]
        vaddr=0x08048370 paddr=0x00000370 haddr=0x00000018 hvaddr=0x08048018 type=program

        1 entrypoints

        [0x08048370]>

Analyse the binary using aa or aaa

```
[0x08048370]> aaa
[x] Analyze all flags starting with sym. and entry0 (aa)
[x] Analyze function calls (aac)
[x] Analyze len bytes of instructions for references (aar)
[x] Check for objc references
Warning: aao experimental on 32bit binaries
[x] Check for vtables
[x] Type matching analysis for all functions (aaft)
[x] Propagate noreturn information
[x] Use -AA or aaaa to perform additional experimental analysis.
[0x08048370]>
```

### Analyze command identifies interesting offsets in binary as Flags.

Check Flags with fs.

```
[0x08048370]> fs
    0 * classes
    0 * functions
    5 * imports
    5 * relocs
   31 * sections
   10 * segments
    4 * strings
   37 * symbols
   27 * symbols.sections
[0x08048370]>
```

Select imports and have a look at it. Pipe commands using ; sign

```
[0x08048370]> fs imports; f
0x00000360 16 loc.imp.__gmon_start
0x08048320 6 sym.imp.strcmp
0x08048330 6 sym.imp.strcpy
0x08048340 6 sym.imp.puts
0x08048350 6 sym.imp.__libc_start_main
[0x08048370]>
```

Select strings flagspace, list using f

```
[0x08048370]> fs strings
[0x08048370]> f
0x08048700 21 str..::_Megabeets_::.
0x08048715 23 str.Think_you_can_make_it
0x0804872c 10 str.Success
0x08048736 22 str.Nop__Wrong_argument.
[0x08048370]>
```

List string in data sections

- iz - list strings in data section
- izz - list strings in whole binary

```
[0x08048370]> iz
[Strings]
nth paddr      vaddr      len size section type  string
-------------------------------------------------------
0   0x00000700 0x08048700 20  21   .rodata ascii \n  .:: Megabeets ::.
1   0x00000715 0x08048715 22  23   .rodata ascii Think you can make it?
2   0x0000072c 0x0804872c 9   10   .rodata ascii Success!\n
3   0x00000736 0x08048736 21  22   .rodata ascii Nop, Wrong argument.\n

[0x08048370]>
```

### Analyze XRefs To command

axt @@ str.\*
This command helps find references to this address. for help ax?
You can see in the output where in main function these are referenced

```
[0x08048370]> axt @@ str.*
main 0x8048609 [DATA] push str..::_Megabeets_::.
main 0x8048619 [DATA] push str.Think_you_can_make_it
main 0x8048646 [DATA] push str.Success
main 0x8048658 [DATA] push str.Nop__Wrong_argument.
[0x08048370]>
```

### Seeking

To navigate from offsets to offsets we need seek command - s
for help: s?

Analyze function list: afl

```
[0x080485f5]> afl
0x08048370    1 33           entry0
0x08048350    1 6            sym.imp.__libc_start_main
0x080483b0    4 43           sym.deregister_tm_clones
0x080483e0    4 53           sym.register_tm_clones
0x08048420    3 30           sym.__do_global_dtors_aux
0x08048440    4 43   -> 40   entry.init0
0x080486e0    1 2            sym.__libc_csu_fini
0x080483a0    1 4            sym.__x86.get_pc_thunk.bx
0x0804846b   19 282          sym.rot13
0x080486e4    1 20           sym._fini
0x08048585    1 112          sym.beet
0x08048330    1 6            sym.imp.strcpy
0x08048320    1 6            sym.imp.strcmp
0x08048680    4 93           sym.__libc_csu_init
0x080485f5    5 127          main
0x080482ec    3 35           sym._init
0x08048340    1 6            sym.imp.puts
[0x080485f5]>
```

Seek to main function

```
[0x08048370]> s main
[0x080485f5]>
```

### Disassembling

Commanf to ask r2 to print disassembled function: pdf

```
[0x080485f5]> s main
[0x080485f5]> pdf
            ; DATA XREF from entry0 @ 0x8048387
/ 127: int main (char **argv);
|           ; var int32_t var_8h @ ebp-0x8
|           ; arg char **argv @ esp+0x24
|           0x080485f5      8d4c2404       lea ecx, [argv]
|           0x080485f9      83e4f0         and esp, 0xfffffff0
|           0x080485fc      ff71fc         push dword [ecx - 4]
|           0x080485ff      55             push ebp
|           0x08048600      89e5           mov ebp, esp
|           0x08048602      53             push ebx
|           0x08048603      51             push ecx
|           0x08048604      89cb           mov ebx, ecx
|           0x08048606      83ec0c         sub esp, 0xc
|           0x08048609      6800870408     push str..::_Megabeets_::.  ; 0x8048700 ; "\n  .:: Megabeets ::." ; const char *s
|           0x0804860e      e82dfdffff     call sym.imp.puts           ; int puts(const char *s)
|           0x08048613      83c410         add esp, 0x10
|           0x08048616      83ec0c         sub esp, 0xc
|           0x08048619      6815870408     push str.Think_you_can_make_it ; 0x8048715 ; "Think you can make it?" ; const char *s
|           0x0804861e      e81dfdffff     call sym.imp.puts           ; int puts(const char *s)
|           0x08048623      83c410         add esp, 0x10
|           0x08048626      833b01         cmp dword [ebx], 1
|       ,=< 0x08048629      7e2a           jle 0x8048655
|       |   0x0804862b      8b4304         mov eax, dword [ebx + 4]
|       |   0x0804862e      83c004         add eax, 4
|       |   0x08048631      8b00           mov eax, dword [eax]
|       |   0x08048633      83ec0c         sub esp, 0xc
|       |   0x08048636      50             push eax
|       |   0x08048637      e849ffffff     call sym.beet
|       |   0x0804863c      83c410         add esp, 0x10
|       |   0x0804863f      85c0           test eax, eax
|      ,==< 0x08048641      7412           je 0x8048655
|      ||   0x08048643      83ec0c         sub esp, 0xc
|      ||   0x08048646      682c870408     push str.Success            ; 0x804872c ; "Success!\n" ; const char *s
|      ||   0x0804864b      e8f0fcffff     call sym.imp.puts           ; int puts(const char *s)
|      ||   0x08048650      83c410         add esp, 0x10
|     ,===< 0x08048653      eb10           jmp 0x8048665
|     |||   ; CODE XREFS from main @ 0x8048629, 0x8048641
|     |``-> 0x08048655      83ec0c         sub esp, 0xc
|     |     0x08048658      6836870408     push str.Nop__Wrong_argument. ; 0x8048736 ; "Nop, Wrong argument.\n" ; const char *s
|     |     0x0804865d      e8defcffff     call sym.imp.puts           ; int puts(const char *s)
|     |     0x08048662      83c410         add esp, 0x10
|     |     ; CODE XREF from main @ 0x8048653
|     `---> 0x08048665      b800000000     mov eax, 0
|           0x0804866a      8d65f8         lea esp, [var_8h]
|           0x0804866d      59             pop ecx
|           0x0804866e      5b             pop ebx
|           0x0804866f      5d             pop ebp
|           0x08048670      8d61fc         lea esp, [ecx - 4]
\           0x08048673      c3             ret
[0x080485f5]>
```

Visual And Graph Mode
V - Visual Mode
VV -Graph Mode

Graph

```
[0x080485f5]> 0x80485f5 # int main (char **argv);


                                                                                        .----------------------------------.
                                                                                        | [0x80485f5]                      |
                                                                                        |   ; DATA XREF from entry0 @ 0x80 |
                                                                                        | 127: int main (char **argv);     |
                                                                                        | ; var int32_t var_8h @ ebp-0x8   |
                                                                                        | ; arg char **argv @ esp+0x24     |
                                                                                        | lea ecx, [argv]                  |
                                                                                        | and esp, 0xfffffff0              |
                                                                                        | push dword [ecx - 4]             |
                                                                                        | push ebp                         |
                                                                                        | mov ebp, esp                     |
                                                                                        | push ebx                         |
                                                                                        | push ecx                         |
                                                                                        | mov ebx, ecx                     |
                                                                                        | sub esp, 0xc                     |
                                                                                        | ; const char *s                  |
                                                                                        | ; 0x8048700                      |
                                                                                        | ; "\n  .:: Megabeets ::."        |
                                                                                        | push str..::_Megabeets_::.       |
                                                                                        | ; int puts(const char *s)        |
                                                                                        | call sym.imp.puts;[oa]           |
                                                                                        | add esp, 0x10                    |
                                                                                        | sub esp, 0xc                     |
                                                                                        | ; const char *s                  |
                                                                                        | ; 0x8048715                      |
                                                                                        | ; "Think you can make it?"       |
                                                                                        | push str.Think_you_can_make_it   |
                                                                                        | ...                              |
                                                                                        `----------------------------------'
                                                                                                f t
                                                                                                | |
                                                                                                | '----------.
                                                                                   .------------'            |
                                                                                   |                         |
                                                                               .-------------------------.   |
                                                                               |  0x804862b [od]         |   |
                                                                               | mov eax, dword [ebx + 4 |   |
                                                                               | add eax, 4              |   |
                                                                               | mov eax, dword [eax]    |   |
                                                                               | sub esp, 0xc            |   |
                                                                               | push eax                |   |
                                                                               | call sym.beet;[oc]      |   |
                                                                               | ...                     |   |
                                                                               `-------------------------'   |
                                                                                       f t                   |
                                                                                       | |                   |
                                                                                       | '----------.        |
                                                                    .------------------'            |        |
                                                                    |                               | .------'
                                                                    |                               | |
                                                                .------------------------.    .------------------------------------------.
                                                                |  0x8048643 [oe]        |    |  0x8048655 [of]                          |
                                                                | sub esp, 0xc           |    | ; CODE XREFS from main @ 0x8048629, 0x80 |
                                                                | ; const char *s        |    | sub esp, 0xc                             |
                                                                | ; 0x804872c            |    | ; const char *s                          |
                                                                | ; "Success!\n"         |    | ; 0x8048736                              |
                                                                | push str.Success       |    | ; "Nop, Wrong argument.\n"               |
                                                                | ; int puts(const char  |    | push str.Nop__Wrong_argument.            |
                                                                | ...                    |    | ...                                      |
                                                                `------------------------'    `------------------------------------------'
                                                                    v                             v
                                                                    |                             |
                                                                    '-----------------.           |
                                                                                      | .---------'
                                                                                      | |
                                                                                .-------------------------------.
                                                                                |  0x8048665 [og]               |
                                                                                | ; CODE XREF from main @ 0x804 |
                                                                                | mov eax, 0                    |
                                                                                | lea esp, [var_8h]             |
                                                                                | pop ecx                       |
                                                                                | pop ebx                       |
                                                                                | ...                           |
                                                                                `-------------------------------'

```

Jump to beet function
type oc in graph mode. Letters next to sym.beet;

```
[0x08048585]>  # sym.beet (char *src);




                                                                                     .-----------------------------------------.
                                                                                     | [0x8048585]                             |
                                                                                     |   ; CALL XREF from main @ 0x8048637     |
                                                                                     | 112: sym.beet (char *src);              |
                                                                                     | ; var char *s2 @ ebp-0x92               |
                                                                                     | ; var int32_t var_8eh @ ebp-0x8e        |
                                                                                     | ; var int32_t var_8ah @ ebp-0x8a        |
                                                                                     | ; var char *dest @ ebp-0x88             |
                                                                                     | ; arg char *src @ ebp+0x8               |
                                                                                     | push ebp                                |
                                                                                     | mov ebp, esp                            |
                                                                                     | sub esp, 0x98                           |
                                                                                     | sub esp, 8                              |
                                                                                     | ; const char *src                       |
                                                                                     | push dword [src]                        |
                                                                                     | lea eax, [dest]                         |
                                                                                     | ; char *dest                            |
                                                                                     | push eax                                |
                                                                                     | ; char *strcpy(char *dest, const char * |
                                                                                     | call sym.imp.strcpy;[oa]                |
                                                                                     | add esp, 0x10                           |
                                                                                     | ; 'Mega'                                |
                                                                                     | mov dword [s2], 0x6167654d              |
                                                                                     | ; 'beet'                                |
                                                                                     | mov dword [var_8eh], 0x74656562         |
                                                                                     | ; 's'                                   |
                                                                                     | ; 115                                   |
                                                                                     | mov word [var_8ah], 0x73                |
                                                                                     | sub esp, 0xc                            |
                                                                                     | lea eax, [s2]                           |
                                                                                     | push eax                                |
                                                                                     | call sym.rot13;[ob]                     |
                                                                                     | add esp, 0x10                           |
                                                                                     | sub esp, 8                              |
                                                                                     | lea eax, [s2]                           |
                                                                                     | ; const char *s2                        |
                                                                                     | push eax                                |
                                                                                     | lea eax, [dest]                         |
                                                                                     | ; const char *s1                        |
                                                                                     | push eax                                |
                                                                                     | ; int strcmp(const char *s1, const char |
                                                                                     | call sym.imp.strcmp;[oc]                |
                                                                                     | ...                                     |
                                                                                     `-----------------------------------------'
```

We can see that the given argument is copied to a buffer. The buffer is located at ebp - local_88h. ‘local_88h’ is actually 0x88 which is 136 in decimal. We can see this by executing ? 0x88. To execute r2 command from inside Visual Graph mode use : and then write the command.

```
:> ? 0x88
int32   136
uint32  136
hex     0x88
octal   0210
unit    136
segment 0000:0088
string  "\x88"
fvalue: 136.0
float:  0.000000f
double: 0.000000
binary  0b1000100
trits   0t12001
:>
```

It seems that the strcmp function is comparing the input string with out put of a function sym.rot13
rot13 is a hashing function (need to read about this)

The binary performs rot13 hashing on a string "Megabeets" and compares the result with our input argument.

let calculate the rot13 hash straight from the command line. r2 provides functionality to do so.

```
:> !rahash2 -E rot -S s:13 -s "Megabeets\n"
Zrtnorrgf
:>
```

Zrtnorrgf is the string being compared to, lets pass this argument to the program in debug mode.
repoen program in debug mode using ood. For help use ood?

```
:> ood Zrtnorrgf
fork_and_ptraceme/CreateProcess: %1 is not a valid Win32 application.

fork_and_ptraceme/CreateProcess: %1 is not a valid Win32 application.

r_core_file_reopen: Cannot reopen file: dbg://D:\\Work\\Data\\megabeets_0x1  Zrtnorrgf with perms 0x7, attempting to open read-only.
:>
```

Here i encountered roadbloack as i am running the ELF file on windows. Need to switch to Linux OS.
Will come back later......
