#include <stdio.h>
#include <string.h>


#define PW "piwei"

#define EGRETI "You know nothing, Jon Snow"

#define UP "update for develop ... test merge tool"


/*
 * only for test 
 * hello world :)
 */
int _main_()
{
    printf("hello world!\n");
    return 0;
}


#ifdef WIN
int WinMain(int argc, char *argv[])
#else
int main(int argc, char *argv[])
#endif
{
    return _main_();
}

