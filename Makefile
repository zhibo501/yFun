
.PHONY:all clean

all: yfun

yfun: CFLAGS+=-Werror
yfun: main.o
	gcc -o $@ $<


clean:
	rm -f *.so *.o *.obj
	rm -f a.out a.so a.exe
	rm -f yfun yfun.exe yfun.out
