mkdir -p bin
nasm -f elf 01.hello.asm
ld -m elf_i386 -s -o bin/01.hello 01.hello.o
./bin/01.hello