#include <stdio.h>
#include <string.h>

/* time  20151108104027 */


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

#define EGRETI "You know nothing, Jon Snow"


