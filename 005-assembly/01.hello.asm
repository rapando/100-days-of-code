; An assembly language program has three sections
; data section [section.data]   - constants are declared here
; bss section  [section.bss]    - variables are declared here
; text section [section.text]   - the actual code
; stack - contains data values passed to functions and procedures
; You can replace section with segment
;
; syntax : [label] mnemonic [operands] [;comment]

section .text
    global _start   ; must be declared for linker (ld)

_start:             ; tells the linker the starting point
    mov edx, len    ; message length
    mov ecx, msg    ; message to write
    mov ebx, 1      ; file descriptor (stdout)
    mov eax, 4      ; system call number (sys_write)
    int 0x80        ; call kernel

    mov eax, 1      ; system call number (sys_exit)
    int 0x80        ; call kernel

section .data
msg db 'Hello, world', 0xa  ; string to be printed
len equ $ - msg             ; length of the string