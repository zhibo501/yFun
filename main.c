#include <stdio.h>


/*
 * only for test 
 * hello world :)
 */
int _m_mian()
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
    return _m_mian();
}
