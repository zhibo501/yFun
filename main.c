#include <stdio.h>
#include <string.h>


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
int WinMain()
#else
int main()
#endif
{
    return _main_();
}


#define PW "piwei"
#define UP "update for develop"

