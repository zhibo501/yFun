#include <stdio.h>
#include <string.h>

/* time  20151108104027 */

#define PW "piwei"

#define EGRETI "You know nothing, Jon Snow"

#define UP "update for develop ... test merge tool"

/* 可执行程序参数字符串最大长度 */
#define M_YFUN_ARGS_STRING_MAX    (10)



/*
 * only for test 
 * hello world :)
 */
int real_fun(char * str)
{
    if (NULL == str)
        return printf("hello world!\n");
    else
        return printf("hello world!\n    %s\n", str);
}


int pack_arg2str(int argc, char *argv[], char *pbuf, int len)
{
    int i, dest_len, tmp;

    for (dest_len = 0, i = 0; i < argc; i++)
    {
        tmp = snprintf(pbuf + dest_len, len - dest_len, "%s ", argv[i]);
        if (tmp > (len - dest_len)) break;

        dest_len = strlen(pbuf);
    }

    if (i < argc)
    {
        printf("%s : buf(0x%x/%d) is less than args(%d/%d/%s)\n",
            __FUNCTION__, pbuf, len, argc, i, argv[i]);
        return 1;
    }

    return 0;
}


char g_args_str[M_YFUN_ARGS_STRING_MAX];

int _main_(int argc, char *argv[])
{
    int ret;

    if (argc > 1)
    {
        pack_arg2str(argc - 1, &(argv[1]), g_args_str, M_YFUN_ARGS_STRING_MAX);
    }

    return real_fun(g_args_str);
}


#ifdef WIN
int WinMain(int argc, char *argv[])
#else
int main(int argc, char *argv[])
#endif
{
    return _main_(argc, argv);
}

