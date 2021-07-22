mkdir -p bin
nasm -f elf hello.asm
ld -m elf_i386 -s -o bin/hello hello.o
./bin/hello