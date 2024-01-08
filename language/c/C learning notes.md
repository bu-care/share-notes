# 学习资源

Linux C编程一站式学习（https://akaedu.github.io/book/index.html）



# 数据类型

在 C 语言中，数据类型指的是用于声明不同类型的变量或函数的一个广泛的系统。变量的类型**决定了变量存储占用的空间**，以及如何解释存储的位模式。C 中的类型可分为以下几种：

<img src="images/C learning notes.assets/image-20221222155722925.png" alt="image-20221222155722925"  />

| 序号 | 类型与描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **基本类型：** 它们是算术类型，包括两种类型：整数类型和浮点类型。 |
| 2    | **枚举类型：** 它们也是算术类型，被用来定义在程序中只能赋予其一定的离散整数值的变量。 |
| 3    | **void 类型：** 类型说明符 *void* 表明没有可用的值。         |
| 4    | **派生类型：** 它们包括：指针类型、数组类型、结构类型、共用体类型和函数类型。 |

数组类型和结构类型统称为**聚合类型**。函数的类型指的是函数**返回值的类型**。在本章节接下来的部分我们将介绍基本类型，其他几种类型会在后边几个章节中进行讲解。

## 整数类型

下表列出了关于标准整数类型的存储大小和值范围的细节：

| 类型           | 存储大小    | 值范围                                               |
| :------------- | :---------- | :--------------------------------------------------- |
| char           | 1 字节      | -128 到 127 或 0 到 255                              |
| unsigned char  | 1 字节      | 0 到 255                                             |
| signed char    | 1 字节      | -128 到 127                                          |
| int            | 2 或 4 字节 | -32,768 到 32,767 或 -2,147,483,648 到 2,147,483,647 |
| unsigned int   | 2 或 4 字节 | 0 到 65,535 或 0 到 4,294,967,295                    |
| short          | 2 字节      | -32,768 到 32,767                                    |
| unsigned short | 2 字节      | 0 到 65,535                                          |
| long           | 4 字节      | -2,147,483,648 到 2,147,483,647                      |
| unsigned long  | 4 字节      | 0 到 4,294,967,295                                   |

注意，各种类型的存储大小与系统位数有关，但目前通用的以64位系统为主。

为了得到某个类型或某个变量在特定平台上的准确大小，您可以使用 **sizeof** 运算符。表达式 *sizeof(type)* 得到对象或类型的存储字节大小。

## void 类型

void 类型指定没有可用的值。它通常用于以下三种情况下：

| 类型              | 描述                                                         |
| :---------------- | :----------------------------------------------------------- |
| **函数返回为空**  | C 中有各种函数都不返回值，或者您可以说它们返回空。不返回值的函数的返回类型为空。例如 **void exit (int status);** |
| **函数参数为空**  | C 中有各种函数不接受任何参数。不带参数的函数可以接受一个 void。例如 **int rand(void);** |
| **指针指向 void** | 类型为 void * 的指针代表对象的地址，而不是类型。例如，内存分配函数 **void \*malloc( size_t size );** 返回指向 void 的指针，**可以转换为任何数据类型**。 |

## 占用空间（64位机器为例）

-  char ： 1个字节
-  int ：4个字节
-  float：4个字节
-  double：8个字节

## 基本类型书写

整数

-  a，默认为10进制 ，10 ，20。
-  b，以0开头为8进制，045，021。
-  c.，以0b开头为2进制，0b11101101。
-  d，以0x开头为16进制，0x21458adf。

**小数**

单精度常量：2.3f 。

双精度常量：2.3，默认为双精度。

**字符型常量**

用英文**单引号**括起来，只保存一个字符'a'、'b' 、'*' ，还有转义字符 '\n' 、'\t'。

**字符串常量**

用英文的**双引号**引起来 可以保存多个字符："abc"。

1、数据类型转换：C 语言中如果一个表达式中含有不同类型的常量和变量，在计算时，会将它们自动转换为同一种类型；在 C 语言中也可以对数据类型进行强制转换；

2、自动转换规则：

-  a）浮点数赋给整型，该浮点数小数被舍去；
-  b）整数赋给浮点型，数值不变，但是被存储到相应的浮点型变量中；

3、强制类型转换形式: **(类型说明符)(表达式)**

# C 变量

变量其实只不过是程序可操作的存储区的名称。C 中每个变量都有特定的类型，类型决定了变量存储的大小和布局，该范围内的值都可以存储在内存中，运算符可应用于变量上。

## C 中的变量定义

变量定义就是告诉编译器在何处创建变量的存储，以及如何创建变量的存储。变量定义指定一个数据类型，并包含了该类型的一个或多个变量的列表，如下所示：

```c
int    i, j, k;
char   c, ch;
float  f, salary;
double d;
```

变量可以在声明的时候被初始化（指定一个初始值）。初始化器由一个等号，后跟一个常量表达式组成，如下所示：

```c
type variable_name = value;

extern int d = 3, f = 5;    // d 和 f 的声明与初始化
int d = 3, f = 5;           // 定义并初始化 d 和 f
byte z = 22;                // 定义并初始化 z
char x = 'x';               // 变量 x 的值为 'x'
```

## C 中的变量声明

变量声明向编译器保证变量以指定的类型和名称存在，这样编译器在不需要知道变量完整细节的情况下也能继续进一步的编译。变量声明只在编译时有它的意义，在程序连接时编译器需要实际的变量声明。

变量的声明有两种情况：

- 1、一种是需要建立存储空间的。例如：int a 在声明的时候就已经建立了存储空间。

- 2、另一种是不需要建立存储空间的，通过使用extern关键字声明变量名而不定义它。 例如：extern int a 其中变量 a 可以在别的文件中定义的。

- 除非有extern关键字，否则都是变量的定义。

  ```c
  extern int i; //声明，不是定义
  int i; //声明，也是定义
  ```

如果需要在一个源文件中引用另外一个源文件中定义的变量，我们只需在引用的文件中将变量加上 extern 关键字的声明即可。

```c
//addtwonum.c 文件代码
#include <stdio.h>
/*外部变量声明*/
extern int x ;
extern int y ;
int addtwonum()
{
    return x+y;
}
```

```c
//test.c 文件代码
#include <stdio.h>
  
/*定义两个全局变量*/
int x=1;
int y=2;
int addtwonum();
int main(void)
{
    int result;
    result = addtwonum();
    printf("result 为: %d\n",result);
    return 0;
}
```

当上面的代码被编译和执行时，它会产生下列结果：

```shell
$ gcc addtwonum.c test.c -o main
$ ./main
result 为: 3
```

# C 常量

## 整数常量

整数常量可以是十进制、八进制或十六进制的常量。前缀指定基数：0x 或 0X 表示十六进制，0 表示八进制，不带前缀则默认表示十进制。

整数常量也可以带一个后缀，后缀是 U 和 L 的组合，U 表示无符号整数（unsigned），L 表示长整数（long）。后缀可以是大写，也可以是小写，U 和 L 的顺序任意。

```c
85         /* 十进制 */
0213       /* 八进制 */
0x4b       /* 十六进制 */
30         /* 整数 */
30u        /* 无符号整数 */
30l        /* 长整数 */
30ul       /* 无符号长整数 */
```

## 浮点常量

浮点常量由整数部分、小数点、小数部分和指数部分组成。您可以使用小数形式或者指数形式来表示浮点常量。

当使用小数形式表示时，必须包含整数部分、小数部分，或同时包含两者。当使用指数形式表示时， 必须包含小数点、指数，或同时包含两者。带符号的指数是用 e 或 E 引入的。

下面列举几个浮点常量的实例：

```c
3.14159       /* 合法的 */
314159E-5L    /* 合法的 */
510E          /* 非法的：不完整的指数 */
210f          /* 非法的：没有小数或指数 */
.e55          /* 非法的：缺少整数或分数 */
```



## 定义常量

在 C 中，有两种简单的定义常量的方式：

1. 使用 **#define** 预处理器。
2. 使用 **const** 关键字。

### #define 预处理器

下面是使用 #define 预处理器定义常量的形式：

```c
#define identifier value
```

```c
#include <stdio.h>
 
#define LENGTH 10   
#define WIDTH  5
#define NEWLINE '\n'
 
int main()
{
 
   int area;  
  
   area = LENGTH * WIDTH;
   printf("value of area : %d", area);
   printf("%c", NEWLINE);
 
   return 0;
}
```

### const 关键字

您可以使用 **const** 前缀声明指定类型的常量，如下所示：

```c
const type variable = value;
```

```c
#include <stdio.h>
 
int main()
{
   const int  LENGTH = 10;
   const int  WIDTH  = 5;
   const char NEWLINE = '\n';
   int area;  
   
   area = LENGTH * WIDTH;
   printf("value of area : %d", area);
   printf("%c", NEWLINE);
 
   return 0;
}
```

**请注意，把常量定义为大写字母形式，是一个很好的编程习惯。**



# C 存储类

存储类定义 C 程序中变量/函数的范围（可见性）和生命周期。这些说明符放置在它们所修饰的类型之前。下面列出 C 程序中可用的存储类：

- auto、register、static、extern

## auto 存储类

**auto** 存储类是所有局部变量默认的存储类。

```c
{
   int mount;
   auto int month;
}
```

上面的实例定义了两个带有相同存储类的变量，auto 只能用在函数内，即 **auto 只能修饰局部变量**。

## register 存储类

**register** 存储类用于定义存储在**寄存器**中而不是 RAM 中的局部变量。这意味着变量的最大尺寸等于寄存器的大小（通常是一个字），且不能对它应用一元的 '&' 运算符（因为它没有内存位置）。

```c
{
   register int  miles;
}
```

寄存器只用于需要快速访问的变量，比如计数器。还应注意的是，定义 'register' 并不意味着变量将被存储在寄存器中，它意味着**变量可能存储在寄存器中**，这取决于硬件和实现的限制。

## static 存储类

**static** 存储类指示编译器在程序的生命周期内保持局部变量的存在，而不需要在每次它进入和离开作用域时进行创建和销毁。因此，使用 static 修饰局部变量可以**在函数调用之间保持局部变量的值**。

static 修饰符也可以应用于全局变量。当 static 修饰全局变量时，会使变量的作用域限制在声明它的文件内。

全局声明的一个 static 变量或方法可以被任何函数或方法调用，只要这些方法出现在跟 static 变量或方法同一个文件中。

以下实例演示了 static 修饰全局变量和局部变量的应用：

```c
#include <stdio.h>
 
/* 函数声明 */
void func1(void);
 
static int count=10;        /* 全局变量 - static 是默认的 */
 
int main()
{
  while (count--) {
      func1();
  }
  return 0;
}
 
void func1(void)
{
/* 'thingy' 是 'func1' 的局部变量 - 只初始化一次
 * 每次调用函数 'func1' 'thingy' 值不会被重置。
 */                
  static int thingy=5;
  thingy++;
  printf(" thingy 为 %d ， count 为 %d\n", thingy, count);
}
//实例中 count 作为全局变量可以在函数内使用，thingy 使用 static 修饰后，不会在每次调用时重置。
```

## extern 存储类

**extern** 存储类用于提供一个全局变量的**引用**，全局变量对所有的程序文件都是可见的。当您使用 **extern** 时，对于无法初始化的变量，会把变量名**指向一个之前定义过的存储位置**。

当您有多个文件且定义了一个可以在其他文件中使用的全局变量或函数时，可以在其他文件中使用 *extern* 来得到已定义的变量或函数的引用。可以这么理解，*extern* 是用来在另一个文件中声明一个全局变量或函数。

extern 修饰符通常用于当有**两个或多个文件共享相同的全局变量或函数**的时候，如下所示

```c
//第一个文件：main.c
#include <stdio.h>
 
int count ;
extern void write_extern();//指向一个之前定义过的存储位置
 
int main()
{
   count = 5;
   write_extern();
}
```

```c
//第二个文件：support.c
#include <stdio.h>
 
extern int count;
 
void write_extern(void)
{
   printf("count is %d\n", count);
}
```

在这里，第二个文件中的 *extern* 关键字用于声明已经在第一个文件 main.c 中定义的 *count*。现在 ，编译这两个文件，如下所示：

```shell
 $ gcc main.c support.c
 ##这会产生 a.out 可执行程序，当程序被执行时，它会产生下列结果：
 count is 5
```

# C 运算符

运算符是一种告诉编译器执行特定的数学或逻辑操作的符号。C 语言内置了丰富的运算符，并提供了以下类型的运算符：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 杂项运算符

## 位运算符

位运算符作用于位，并逐位执行操作。&、 | 和 ^ 的真值表如下所示：

```shell
# ^ 异或运算符，按二进制位进行"异或"运算。运算规则：不同为1，相同为0
0^0=0;   
0^1=1;   
1^0=1;  
1^1=0;
```

假设如果 A = 60，A = 0011 1100

```shell
# << 二进制左移运算符。将一个运算对象的各二进制位全部左移若干位（左边的二进制位丢弃，右边补0）。
A << 2 将得到 240，即为 1111 0000

# >> 二进制右移运算符。将一个数的各二进制位全部右移若干位，正数左补0，负数左补1，右边丢弃。
A >> 2 将得到 15，即为 0000 1111
```

## 杂项运算符 ↦ sizeof & 三元

下表列出了 C 语言支持的其他一些重要的运算符，包括 **sizeof** 和 **?  :**。

| 运算符   | 描述             | 实例                                 |
| -------- | ---------------- | ------------------------------------ |
| sizeof() | 返回变量的大小。 | sizeof(a) 将返回 4，其中 a 是整数。  |
| &        | 返回变量的地址。 | &a; 将给出变量的实际地址。           |
| *        | 指向一个变量。   | *a; 将指向一个变量。                 |
| ? :      | 条件表达式       | 如果条件为真 ? 则值为 X : 否则值为 Y |

# C 函数

函数**声明**告诉编译器函数的**名称、返回类型和参数**。函数定义提供了函数的实际主体。

C 标准库提供了大量的程序可以调用的内置函数。例如，函数 **strcat()** 用来连接两个字符串，函数 **memcpy()** 用来复制内存到另一个位置。

## 定义函数

C 语言中的函数定义的一般形式如下：

```c
return_type function_name( parameter list )
{
   body of the function
}
```

在 C 语言中，函数由一个函数头和一个函数主体组成。下面列出一个函数的所有组成部分：

- **返回类型：**一个函数可以返回一个值。**return_type** 是函数返回的值的数据类型。有些函数执行所需的操作而**不返回值**，在这种情况下，return_type 是关键字 **void**。
- **函数名称：**这是函数的实际名称。函数名和参数列表一起构成了函数签名。
- **参数：**参数就像是占位符。当函数被调用时，您向参数传递一个值，这个值被称为实际参数。参数列表包括函数参数的类型、顺序、数量。参数是可选的，也就是说，函数可能不包含参数。
- **函数主体：**函数主体包含一组定义函数执行任务的语句。

## 函数声明

函数**声明**会告诉编译器函数名称及如何调用函数。函数的**实际主体可以单独定义**。

函数声明包括以下几个部分：

```c
return_type function_name( parameter list );
```

在函数声明中，参数的名称并不重要，只有**参数的类型是必需的**，因此下面也是有效的声明：

```c
int max(int, int);
```

当您在一个源文件中定义函数且在另一个文件中调用函数时，函数声明是必需的。在这种情况下，您应该在调用函数的文件顶部声明函数。

## 调用函数

当程序调用函数时，程序控制权会转移给被调用的函数。被调用的函数执行已定义的任务，当函数的返回语句被执行时，或到达函数的结束括号时，会把程序控制权交还给主程序。

调用函数时，传递所需参数，如果函数返回一个值，则可以**存储返回值**。例如：

```c
#include <stdio.h>
 
/* 函数声明 */
int max(int num1, int num2);
 
int main ()
{
   /* 局部变量定义 */
   int a = 100;
   int b = 200;
   int ret;
 
   /* 调用函数来获取最大值 */
   ret = max(a, b);
 
   printf( "Max value is : %d\n", ret );
 
   return 0;
}
 
/* 函数返回两个数中较大的那个数 */
int max(int num1, int num2) 
{
   /* 局部变量声明 */
   int result;
 
   if (num1 > num2)
      result = num1;
   else
      result = num2;
 
   return result; 
}
```

## 函数参数

如果函数要使用参数，则必须声明接受参数值的变量。这些变量称为函数的**形式参数**。形式参数就像函数内的其他局部变量，在进入函数时被创建，退出函数时被销毁。当调用函数时，有两种向函数传递参数的方式：

| 调用类型                                                     | 描述                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [传值调用](https://www.runoob.com/cprogramming/c-function-call-by-value.html) | 该方法把参数的实际值复制给函数的形式参数。在这种情况下，修改函数内的形式参数不会影响实际参数。（将原本的实际值拷贝一份到新开辟的一块内存中） |
| [引用调用](https://www.runoob.com/cprogramming/c-function-call-by-pointer.html) | 通过指针传递方式，形参为指向实参地址的指针，当对形参的指向操作时，就相当于对实参本身进行的操作。（对指针指向的原本地址的内存中的值进行操作） |

# C 作用域规则

任何一种编程中，作用域是程序中定义的变量所存在的区域，超过该区域变量就不能被访问。C 语言中有三个地方可以声明变量：

1. 在函数或块内部的**局部**变量
2. 在所有函数外部的**全局**变量
3. 在**形式**参数的函数参数定义中

## 形式参数

函数的参数，形式参数，被当作该函数内的局部变量，如果与全局变量同名它们会优先使用。

**全局变量与局部变量在内存中的区别**：

- 全局变量保存在内存的**全局存储区**中，占用**静态的存储单元**；
- 局部变量保存在**栈**中，只有在所在函数**被调用时才动态地为变量分配存储单元**。

## 初始化局部变量和全局变量

当局部变量被定义时，系统不会对其初始化，您必须自行对其初始化。**定义全局变量时，系统会自动对其初始化**，如下所示：

```c
int	      0
char	    '\0'
float	    0
double	  0
pointer	  NULL
```

# C 数组

C 语言支持**数组**数据结构，它可以存储一个**固定大小的相同类型元素**的顺序集合。

所有的数组都是由**连续的内存位置**组成。最低的地址对应第一个元素，最高的地址对应最后一个元素。

## 声明数组

在 C 中要声明一个数组，需要指定元素的类型和元素的数量，如下所示：

```c
type arrayName [ arraySize ];
```

**arraySize** 必须是一个大于零的整数常量，**type** 可以是任意有效的 C 数据类型。

## 初始化数组

在 C 中，您可以逐个初始化数组，也可以使用一个初始化语句，如下所示：

```c
double balance[5] = {1000.0, 2.0, 3.4, 7.0, 50.0};
//如果您省略掉了数组的大小，数组的大小则为初始化时元素的个数。
double balance[] = {1000.0, 2.0, 3.4, 7.0, 50.0};
```

## 访问数组元素

数组元素可以通过数组名称加索引进行访问。例如：

```c
double salary = balance[9];
```

声明数组、数组赋值、访问数组：

```c
#include <stdio.h>
 
int main ()
{
   int n[ 10 ]; /* n 是一个包含 10 个整数的数组 */
   int i,j;
 
   /* 初始化数组元素 */         
   for ( i = 0; i < 10; i++ )
   {
      n[ i ] = i + 100; /* 设置元素 i 为 i + 100 */
   }
   
   /* 输出数组中每个元素的值 */
   for (j = 0; j < 10; j++ )
   {
      printf("Element[%d] = %d\n", j, n[j] );
   }
 
   return 0;
}
```

## C 中数组详解

下面列出了 C 程序员必须清楚的一些与数组相关的重要概念：

| 概念                                                         | 描述                                                         |
| :----------------------------------------------------------- | :----------------------------------------------------------- |
| [多维数组](https://www.runoob.com/cprogramming/c-multi-dimensional-arrays.html) | C 支持多维数组。多维数组最简单的形式是二维数组。             |
| [传递数组给函数](https://www.runoob.com/cprogramming/c-passing-arrays-to-functions.html) | 您可以通过指定不带索引的数组名称来给函数传递一个**指向数组的指针**。（数组名是一个指向数组中第一个元素的常量指针，**balance** 是一个指向 &balance[0] 的指针） |
| [从函数返回数组](https://www.runoob.com/cprogramming/c-return-arrays-from-function.html) | C 语言不允许返回一个完整的数组作为函数的参数。但是，您可以通过指定不带索引的数组名来**返回一个指向数组的指针**。 |
| [指向数组的指针](https://www.runoob.com/cprogramming/c-pointer-to-an-array.html) | 您可以通过指定不带索引的数组名称来生成一个指向数组中**第一个元素的指针**。 |

```
使用数组名作为常量指针是合法的，反之亦然。因此，*(balance + 4) 是一种访问 balance[4] 数据的合法方式。一旦您把第一个元素的地址存储在 p 中，您就可以使用 *p、*(p+1)、*(p+2) 等来访问数组元素。
```

```c
#include <stdio.h>
 
int main ()
{
   /* 带有 5 个元素的整型数组 */
   double balance[5] = {1000.0, 2.0, 3.4, 17.0, 50.0};
   double *p;
   int i;
 
   p = balance;//将 balance 数组的首地址赋值给 p
 
   /* 输出数组中每个元素的值 */
   printf( "使用指针的数组值\n");
   for ( i = 0; i < 5; i++ )
   {
       printf("*(p + %d) : %f\n",  i, *(p + i) ); //首地址的值加上 i
   }
 
   printf( "使用 balance 作为地址的数组值\n");
   for ( i = 0; i < 5; i++ )
   {
       printf("*(balance + %d) : %f\n",  i, *(balance + i) ); //首地址的值加上 i
   }
 
   return 0;
}
```

p 是一个指向 double 型的指针，这意味着它可以存储一个 double 类型的变量。一旦我们有了 p 中的地址，***p** 将给出存储在 p 中相应地址的值。

### C 从函数返回数组

C 语言不允许返回一个完整的数组作为函数的参数。但是，您可以通过指定不带索引的数组名来返回一个指向数组的指针。

如果您想要从函数返回一个一维数组，您必须声明一个**返回指针的函数**，如下：

```c
int * myFunction()
{
}
```

另外，C 不支持在函数外返回**局部变量的地址**，除非定义局部变量为 **static** 变量。

```c
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
 
/* 要生成和返回随机数的函数，返回的是r数组的首地址；static才可以返回 */
int * getRandom( )
{
  static int  r[10];
  int i;
 
  /* 设置种子 */
  srand( (unsigned)time( NULL ) );
  for ( i = 0; i < 10; ++i)
  {
     r[i] = rand();
     printf( "r[%d] = %d\n", i, r[i]);
 
  }
 
  return r;
}
 
/* 要调用上面定义函数的主函数 */
int main ()
{
   /* 一个指向整数的指针 */
   int *p;
   int i;
 
   p = getRandom();
   for ( i = 0; i < 10; i++ )
   {
       printf( "*(p + %d) : %d\n", i, *(p + i));
   }
 
   return 0;
}
```

srand((unsigned)time(NULL))是初始化随机函数种子：

-  1、是拿当前系统时间作为种子，由于时间是变化的，种子变化，可以产生不相同的随机数。计算机中的随机数实际上都不是真正的随机数，如果两次给的种子一样，是会生成同样的随机序列的。 所以，一般都会以当前的时间作为种子来生成随机数，这样更加的随机。
-  2、使用时，参数可以是unsigned型的任意数据，比如srand(10)；
-  3、如果不使用srand，用rand()产生的随机数，在多次运行，结果是一样的。

# C enum(枚举)

### 枚举类型

枚举是 C 语言中的一种基本数据类型，它可以让数据更简洁，更易读。枚举语法定义格式为：

```C
enum　枚举名　{枚举元素1,枚举元素2,……};
```

比如：一星期有 7 天，如果不用枚举，我们需要使用 #define 来为每个整数定义一个别名：

```C
#define MON  1
#define TUE  2
#define WED  3
#define THU  4
#define FRI  5
#define SAT  6
#define SUN  7
```

这个看起来代码量就比较多，接下来我们看看使用枚举的方式：

```C
enum DAY
{
      MON=1, TUE, WED, THU, FRI, SAT, SUN
};
```

**注意：**第一个枚举成员的默认值为整型的 0，后续枚举成员的值在前一个成员上加 1。我们在这个实例中把第一个枚举成员的值定义为 1，第二个就为 2，以此类推。

可以在定义枚举类型时改变枚举元素的值：

```c
enum season {spring, summer=3, autumn, winter};
//没有指定值的枚举元素，其值为前一元素加 1。也就说 spring 的值为 0，summer 的值为 3，autumn 的值为 4，winter 的值为 5
```

### 定义枚举变量

前面我们只是声明了枚举类型，接下来我们看看如何**定义枚举变量**。我们可以通过以下三种方式来定义枚举变量

1. 先定义枚举类型，再定义枚举变量

   ```c
   enum DAY
   {
         MON=1, TUE, WED, THU, FRI, SAT, SUN
   };
   enum DAY day;
   ```

2. **定义枚举类型的同时定义枚举变量**

   ```c
   enum DAY
   {
         MON=1, TUE, WED, THU, FRI, SAT, SUN
   } day;
   ```

3. **省略枚举名称，直接定义枚举变量**

   ```c
   enum
   {
         MON=1, TUE, WED, THU, FRI, SAT, SUN
   } day;
   ```

```c
#include <stdio.h>
 
enum DAY
{
      MON=1, TUE, WED, THU, FRI, SAT, SUN
};
 
int main()
{
    enum DAY day;
    day = WED;
    printf("%d",day);
    return 0;
}//输出为3
```

在C 语言中，枚举类型是被当做 int 或者 unsigned int 类型来处理的，所以按照 C 语言规范是没有办法遍历枚举类型的。不过在一些特殊的情况下，枚举类型必须连续是可以实现有条件的遍历。以下实例使用 for 来遍历枚举的元素：

```c
#include <stdio.h>
 
enum DAY
{
      MON=1, TUE, WED, THU, FRI, SAT, SUN
} day;
int main()
{
    // 遍历枚举元素
    for (day = MON; day <= SUN; day++) {
        printf("枚举元素：%d \n", day);
    }
}
```

# C 指针

每一个变量都有一个内存位置，每一个内存位置都定义了可使用 **&** 运算符访问的地址，它表示了在内存中的一个地址。

```c
#include <stdio.h>
 
int main ()
{
    int var_runoob = 10;
    int *p;              // 定义指针变量
    p = &var_runoob;
 
   printf("var_runoob 变量的地址： %p\n", p);
   return 0;
}
//var_runoob 变量的地址： 0x7ffeeaae08d8
```

![img](images/C learning notes.assets/c-pointer.png)

## 什么是指针？

指针也就是内存地址，指针变量是用来存放内存地址的变量。就像其他变量或常量一样，您必须在使用指针存储其他变量地址之前，对其进行声明。指针变量声明的一般形式为：

```
type *var_name;
```

**type** 是指针的基类型，它必须是一个有效的 C 数据类型，**var_name** 是指针变量的名称。用来声明指针的星号 ***** 与乘法中使用的星号是相同的。但是，在这个语句中，星号是用来指定一个变量是指针。以下是有效的指针声明：

```c
int    *ip;    /* 一个整型的指针 */
double *dp;    /* 一个 double 型的指针 */
float  *fp;    /* 一个浮点型的指针 */
char   *ch;    /* 一个字符型的指针 */
```

所有实际数据类型，不管是整型、浮点型、字符型，还是其他的数据类型，指针的值的类型都是一样的，都是一个十六进制数的**内存地址**。唯一的不同是，指针所指向的变量或常量的数据类型不同。

## C 中的 NULL 指针

在变量声明的时候，如果没有确切的地址可以赋值，为指针变量赋一个 NULL 值是一个良好的编程习惯。赋为 NULL 值的指针被称为**空指针**。NULL 指针是一个定义在标准库中的值为零的常量。请看下面的程序：

```c
#include <stdio.h>
 
int main ()
{
   int  *ptr = NULL; 
   printf("ptr 的地址是 %p\n", ptr  );
   return 0;
}//ptr 的地址是 0x0
```

在大多数的操作系统上，程序不允许访问地址为 0 的内存，因为该内存是操作系统保留的。然而，内存地址 0 有特别重要的意义，它表明该指针不指向一个可访问的内存位置。但按照惯例，如果指针包含空值（零值），则**假定它不指向任何东西**。

如需检查一个空指针，您可以使用 if 语句，如下所示：

```c
if(ptr)     /* 如果 p 非空，则完成 */
if(!ptr)    /* 如果 p 为空，则完成 */
```

## C 指针详解

下面列出了 C 程序员必须清楚的一些与指针相关的重要概念：

| 概念                                                         | 描述                                                     |
| :----------------------------------------------------------- | :------------------------------------------------------- |
| [指针的算术运算](https://www.runoob.com/cprogramming/c-pointer-arithmetic.html) | 可以对指针进行四种算术运算：++、--、+、-                 |
| [指针数组](https://www.runoob.com/cprogramming/c-array-of-pointers.html) | 可以定义用来存储指针的数组。                             |
| [指向指针的指针](https://www.runoob.com/cprogramming/c-pointer-to-pointer.html) | C 允许指向指针的指针。                                   |
| [传递指针给函数](https://www.runoob.com/cprogramming/c-passing-pointers-to-functions.html) | 通过引用或地址传递参数，使传递的参数在调用函数中被改变。 |
| [从函数返回指针](https://www.runoob.com/cprogramming/c-return-pointer-from-functions.html) | C 允许函数返回指针到局部变量、静态变量和动态内存分配。   |

### C 指针的算术运算

C 指针是一个用**数值表示**的地址。因此，您可以对指针**执行算术运算**。可以对指针进行四种算术运算：++、--、+、-。

假设 **ptr** 是一个指向地址 1000 的整型指针，是一个 32 位的整数，让我们对该指针执行下列的算术运算：

```C
ptr++
```

在执行完上述的运算之后，**ptr** 将指向位置 1004，因为 ptr 每增加一次，它都将指向下一个整数位置（整型是4个字节），即当前位置往后移 4 字节。如果 **ptr** 指向一个地址为 1000 的字符（字符是1个字节），上面的运算会导致指针指向位置 1001，因为下一个字符位置是在 1001。

我们概括一下：

- 指针的每一次递增，它其实会指向下一个元素的存储单元。
- 指针的每一次递减，它都会指向前一个元素的存储单元。
- 指针在递增和递减时跳跃的字节数取决于指针所指向**变量数据类型长度**，比如 int 就是 4 个字节。

#### 递增一个指针

我们喜欢在程序中使用指针代替数组，因为变量指针可以递增，而数组不能递增，数组可以看成一个指针常量。下面的程序递增变量指针，以便顺序访问数组中的每一个元素：

```c
#include <stdio.h>
 
const int MAX = 3;
 
int main ()
{
   int  var[] = {10, 100, 200};
   int  i, *ptr;
 
   /* 指针中的数组地址 */
   ptr = var;//把数组的首地址给指针变量
   for ( i = 0; i < MAX; i++)
   {
 
      printf("存储地址：var[%d] = %p\n", i, ptr );
      printf("存储值：var[%d] = %d\n", i, *ptr );
 
      /* 指向下一个位置 */
      ptr++;//使用指针的自加去遍历数组的值
   }
   return 0;
}
/* 整型是4个字节，地址的值在增加4
存储地址：var[0] = e4a298cc
存储值：var[0] = 10
存储地址：var[1] = e4a298d0
存储值：var[1] = 100
存储地址：var[2] = e4a298d4
存储值：var[2] = 200
*/
```

#### 指针的比较

指针可以用关系运算符进行比较，如 ==、< 和 >。如果 p1 和 p2 指向两个相关的变量，比如同一个数组中的不同元素，则可对 p1 和 p2 进行大小比较。

```c
#include <stdio.h>
 
const int MAX = 3;
 
int main ()
{
   int  var[] = {10, 100, 200};
   int  i, *ptr;
 
   /* 指针中第一个元素的地址 */
   ptr = var;
   i = 0;
   while ( ptr <= &var[MAX - 1] )//地址比较，本质是两个十六进制数比较
   {
 
      printf("存储地址：var[%d] = %p\n", i, ptr );
      printf("存储值：var[%d] = %d\n", i, *ptr );
 
      /* 指向上一个位置 */
      ptr++;
      i++;
   }
   return 0;
}
/*
存储地址：var[0] = 0x7ffeee2368cc
存储值：var[0] = 10
存储地址：var[1] = 0x7ffeee2368d0
存储值：var[1] = 100
存储地址：var[2] = 0x7ffeee2368d4
存储值：var[2] = 200
*/
```

### C 指针数组

可能有一种情况，我们想要让数组存储指针(指针指向 int 或 char 或其他数据类型)。下面是一个指向整数的指针数组的声明：

```c
int *ptr[MAX];
```

把 **ptr** 声明为一个数组，由 MAX 个整数指针组成。因此，ptr 中的每个元素，都是一个指向 int 值的指针。

```c
#include <stdio.h>
 
const int MAX = 3;
 
int main ()
{
   int  var[] = {10, 100, 200};
   int i, *ptr[MAX];
 
   for ( i = 0; i < MAX; i++)
   {
      ptr[i] = &var[i]; /* 赋值为整数的地址，将数组的地址放进指针数组中 */
   }
   for ( i = 0; i < MAX; i++)
   {
      printf("Value of var[%d] = %d\n", i, *ptr[i] );
   }
   return 0;
}
```

也可以用一个**指向字符**的指针数组来存储一个字符串列表，如下：

```c
#include <stdio.h>
 
const int MAX = 4;
 
int main ()
{
   const char *names[] = {
                   "Zara Ali",
                   "Hina Ali",
                   "Nuha Ali",
                   "Sara Ali",
   };
   int i = 0;
 
   for ( i = 0; i < MAX; i++)
   {
      printf("Value of names[%d] = %s\n", i, names[i] );//这里不太明白，取值不用使用指针么，这样取出来不是16进制的地址么？
   }
   return 0;
}
/*
Value of names[0] = Zara Ali
Value of names[1] = Hina Ali
Value of names[2] = Nuha Ali
Value of names[3] = Sara Ali*/
```

### C 指向指针的指针

指向指针的指针是一种多级间接寻址的形式，或者说是一个指针链。通常，一个指针包含一个变量的地址。当我们定义一个指向指针的指针时，第一个指针包含了第二个指针的地址，第二个指针指向包含实际值的位置。

![C 中指向指针的指针](images/C learning notes.assets/pointer_to_pointer.jpg)

一个指向指针的指针变量必须如下声明，即在变量名前放置两个星号。例如，下面声明了一个指向 int 类型指针的指针：

```c
int **var;
```

```c
#include <stdio.h>
 
int main ()
{
   int  V;
   int  *Pt1;
   int  **Pt2;
 
   V = 100;
 
   /* 获取 V 的地址 */
   Pt1 = &V;
 
   /* 使用运算符 & 获取 Pt1 的地址 */
   Pt2 = &Pt1;
 
   /* 使用 pptr 获取值 */
   printf("var = %d\n", V );
   printf("Pt1 = %p\n", Pt1 );
   printf("*Pt1 = %d\n", *Pt1 );
    printf("Pt2 = %p\n", Pt2 );
   printf("**Pt2 = %d\n", **Pt2);
 
   return 0;
}
/*
var = 100
Pt1 = 0x7ffee2d5e8d8
*Pt1 = 100
Pt2 = 0x7ffee2d5e8d0
**Pt2 = 100
*/
```

### 传递指针给函数

C 语言允许您传递指针给函数，只需要简单地声明函数参数为指针类型即可。

下面的实例中，我们传递一个无符号的 long 型指针给函数，并在函数内改变这个值：

```c
#include <stdio.h>
#include <time.h>
 
void getSeconds(unsigned long *par);

int main ()
{
   unsigned long sec;


   getSeconds( &sec );

   /* 输出实际值 */
   printf("Number of seconds: %ld\n", sec );

   return 0;
}

void getSeconds(unsigned long *par)
{
   /* 获取当前的秒数 */
   *par = time( NULL );
   return;
}
//Number of seconds :1294450468
```

能接受指针作为参数的函数，也能接受数组作为参数，如下所示：

```c
#include <stdio.h>
 
/* 函数声明 */
double getAverage(int *arr, int size);
 
int main ()
{
   /* 带有 5 个元素的整型数组  */
   int balance[5] = {1000, 2, 3, 17, 50};
   double avg;
 
   /* 传递一个指向数组的指针作为参数 */
   avg = getAverage( balance, 5 ) ;
 
   /* 输出返回值  */
   printf("Average value is: %f\n", avg );
   
   return 0;
}

double getAverage(int *arr, int size)
{
  int    i, sum = 0;      
  double avg;          
 
  for (i = 0; i < size; ++i)
  {
    sum += arr[i];
  }
 
  avg = (double)sum / size;
 
  return avg;
}
//Average value is: 214.40000
```

### 从函数返回指针

C 允许您从函数返回指针。为了做到这点，您必须声明一个返回指针的函数，如下所示：

```c
int * myFunction()
{
}
```

另外，C 语言不支持在调用函数时返回局部变量的地址，除非定义局部变量为 **static** 变量。

下面的函数，它会生成 10 个随机数，并使用表示指针的数组名（即第一个数组元素的地址）来返回它们，具体如下：

```c
#include <stdio.h>
#include <time.h>
#include <stdlib.h> 
 
/* 要生成和返回随机数的函数 */
int * getRandom( )
{
   static int  r[10];//局部变量想返回地址，必须声明为 static 类型
   int i;
 
   /* 设置种子 */
   srand( (unsigned)time( NULL ) );
   for ( i = 0; i < 10; ++i)
   {
      r[i] = rand();
      printf("%d\n", r[i] );
   }
 
   return r;
}
 
/* 要调用上面定义函数的主函数 */
int main ()
{
   /* 一个指向整数的指针 */
   int *p;
   int i;
 
   p = getRandom();
   for ( i = 0; i < 10; i++ )
   {
       printf("*(p + [%d]) : %d\n", i, *(p + i) );
   }
 
   return 0;
}
```

### 总结注意

>C 不支持在调用函数时返回局部变量的地址，除非定义局部变量为 **static 变量**。
>
>因为局部变量是存储在内存的**栈区**内，当函数调用结束后，局部变量所占的内存地址便被释放了，因此当其函数执行完毕后，函数内的变量便不再拥有那个内存地址，所以**不能返回其指针**。
>
>除非将其变量定义为 static 变量，static 变量的值存放在内存中的**静态数据区**，不会随着函数执行的结束而被清除，故能返回其地址。



## 函数指针与回调函数

### 函数指针

函数指针是**指向函数**的指针变量。

通常我们说的指针变量是指向一个整型、字符型或数组等变量，而函数指针是指向函数。函数指针可以像一般函数一样，用于调用函数、传递参数。函数指针变量的声明：

```c
typedef int (*fun_ptr)(int,int); // 声明一个指向同样参数、返回值的函数指针类型
```

以下实例声明了函数指针变量 p，指向函数 max：

```c
#include <stdio.h>
 
int max(int x, int y)
{
    return x > y ? x : y;
}
 
int main(void)
{
    /* p 是函数指针 */
    int (* p)(int, int) = & max; // &可以省略
    int a, b, c, d;
 
    printf("请输入三个数字:");
    scanf("%d %d %d", & a, & b, & c);
 
    /* 与直接调用函数等价，d = max(max(a, b), c) */
    d = p(p(a, b), c); 
 
    printf("最大的数字是: %d\n", d);
 
    return 0;
}
```

### 回调函数

函数指针变量可以作为某个函数的参数来使用的，回调函数就是一个通过函数指针调用的函数。

简单讲：回调函数是由别人的函数执行时调用你实现的函数。(就是函数的参数是一个指针，它指向另一个函数名)

```c
#include <stdlib.h>  
#include <stdio.h>
 
void populate_array(int *array, size_t arraySize, int (*getNextValue)(void))
{
    for (size_t i=0; i<arraySize; i++)
        array[i] = getNextValue();
}
 
// 获取随机值
int getNextRandomValue(void)
{
    return rand();
}
 
int main(void)
{
    int myarray[10];
    /* getNextRandomValue 不能加括号，否则无法编译，因为加上括号之后相当于传入此参数时传入了 int , 而不是函数指针*/
    //加上括号是相当于传入函数执行的返回值么？？
    populate_array(myarray, 10, getNextRandomValue);
    for(int i = 0; i < 10; i++) {
        printf("%d ", myarray[i]);
    }
    printf("\n");
    return 0;
}
//16807 282475249 1622650073 984943658 1144108930 470211272 101027544 1457850878 1458777923 2007237709 
```

# 字符串

在 C 语言中，字符串实际上是使用空字符 **\0 结尾的一维字符数组**。因此，**\0** 是用于标记字符串的结束。

**空字符（Null character**）又称结束符，缩写 **NUL**，是一个数值为 **0** 的控制字符，**\0** 是转义字符，意思是告诉编译器，这不是字符 **0**，而是空字符。

下面的声明和初始化创建了一个 **RUNOOB** 字符串。由于在数组的末尾存储了空字符 **\0**，所以字符数组的大小比单词 **RUNOOB** 的字符数**多一个**。

```c
char site[7] = {'R', 'U', 'N', 'O', 'O', 'B', '\0'};
```

依据数组初始化规则，您可以把上面的语句写成以下语句：

```c
char site[] = "RUNOOB";
```

以下是 C/C++ 中定义的字符串的内存表示：

![C/C++ 中的字符串表示](images/C learning notes.assets/c-strings-2020-12-21.png)

不需要把 **null** 字符放在字符串常量的末尾。C 编译器会在初始化数组时，自动把 **\0** 放在字符串的末尾。

C 中有大量操作字符串的函数：

| 函数                | 目的                                                         |
| :------------------ | :----------------------------------------------------------- |
| **strcpy(s1, s2);** | 复制字符串 s2 到字符串 s1。                                  |
| **strcat(s1, s2);** | 连接字符串 s2 到字符串 s1 的末尾。                           |
| **strlen(s1);**     | 返回字符串 s1 的长度。                                       |
| **strcmp(s1, s2);** | 如果 s1 和 s2 是相同的，则返回 0；如果 s1<s2 则返回小于 0；如果 s1>s2 则返回大于 0。 |
| **strchr(s1, ch);** | 返回一个指针，指向字符串 s1 中**字符** ch 的第一次出现的位置。 |
| **strstr(s1, s2);** | 返回一个指针，指向字符串 s1 中**字符串** s2 的第一次出现的位置。 |

```c
#include <stdio.h>
#include <string.h>
 
int main ()
{
   char str1[14] = "runoob";
   char str2[14] = "google";
   char str3[14];
   int  len ;
 
   /* 复制 str1 到 str3 */
   strcpy(str3, str1);
   printf("strcpy( str3, str1) :  %s\n", str3 );
 
   /* 连接 str1 和 str2 */
   strcat( str1, str2);
   printf("strcat( str1, str2):   %s\n", str1 );
 
   /* 连接后，str1 的总长度 */
   len = strlen(str1);
   printf("strlen(str1) :  %d\n", len );
 
   return 0;
}
/*
strcpy( str3, str1) :  runoob
strcat( str1, str2):   runoobgoogle
strlen(str1) :  12
*/
```

# 结构体

C 语言中数组允许存储相同类型数据项，而**结构**是 C 编程中另一种用户自定义的可用的数据类型，它允许您**存储不同类型**的数据项。

结构用于表示一条记录，假设您想要跟踪图书馆中书本的动态，您可能需要跟踪每本书的下列属性：Title、Author、Subject、Book ID

## 定义结构

为了定义结构，您必须使用 **struct** 语句。struct 语句定义了一个包含多个成员的新的数据类型，struct 语句的格式如下：

```c
struct tag { 
    member-list
    member-list 
    member-list  
    ...
} variable-list ;
```

**tag** 是结构体标签。

**member-list** 是标准的**变量定义**，比如 int i; 或者 float f，或者其他有效的变量定义。

**variable-list** 结构变量，定义在结构的末尾，最后一个分号之前，您可以指定**一个或多个结构变量**。下面是声明 Book 结构的方式：

```c
struct Books
{
   char  title[50]; //数组变量，其实就是表示字符串
   char  author[50];
   char  subject[100];
   int   book_id;
} book;
```

1. 在一般情况下，**tag、member-list、variable-list** 这 3 部分至少要出现 2 个。以下为实例：

```c
//此声明声明了拥有3个成员的结构体，分别为整型的a，字符型的b和双精度的c
//同时又声明了结构体变量s1
//这个结构体并没有   标明其标签
struct 
{
    int a;
    char b;
    double c;
} s1;
 
//此声明声明了拥有3个成员的结构体，分别为整型的a，字符型的b和双精度的c
//结构体的标签被命名为SIMPLE,没有   声明变量
struct SIMPLE
{
    int a;
    char b;
    double c;
};

//用SIMPLE标签的结构体，另外声明了 变量t1、t2、t3（没有属性的空结构体）
struct SIMPLE t1, t2[20], *t3;
 
//也可以用typedef   创建新类型
typedef struct
{
    int a;
    char b;
    double c; 
} Simple2;

//现在可以用Simple2作为类型声明新的结构体变量
Simple2 u1, u2[20], *u3;
```

在上面的声明中，第一个和第二声明被编译器当作两个完全不同的类型，即使他们的成员列表是一样的，如果令 t3=&s1，则是非法的。

2. 结构体的成员可以包含其他结构体，也可以包含**指向自己结构体类型的指针**，而通常这种指针的应用是为了实现一些更高级的数据结构如**链表和树**等。

```c
//此结构体的声明包含了其他的结构体
struct COMPLEX
{
    char string[100];
    struct SIMPLE a;//struct SIMPLE 是类型，a是变量
};
 
//此结构体的声明包含了指向自己类型的指针
struct NODE
{
    char string[100];
    struct NODE *next_node;//指向自己结构体类型的指针变量，这个指针指向的地址存放的是NODE类型的结构体数据
};
```

3. 如果两个结构体互相包含，则需要对其中一个结构体进行**不完整声明**，如下所示：

```c
struct B;    //对结构体B进行不完整声明
 
//结构体A中包含指向结构体B的指针
struct A
{
    struct B *partner;
    //other members;
};
 
//结构体B中包含指向结构体A的指针，在A声明完后，B也随之进行声明
struct B
{
    struct A *partner;
    //other members;
};
```

## 结构体变量的初始化

和其它类型变量一样，对结构体变量可以在定义时指定初始值。

```c
#include <stdio.h>
 
struct Books
{
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
} book = {"C 语言", "RUNOOB", "编程语言", 123456};
 
int main()
{
    printf("title : %s\nauthor: %s\nsubject: %s\nbook_id: %d\n", book.title, book.author, book.subject, book.book_id);
}
/*
title : C 语言
author: RUNOOB
subject: 编程语言
book_id: 123456
*/
```

## 访问结构成员

为了访问结构的成员，我们使用**成员访问运算符（.）**。成员访问运算符是结构变量名称和我们要访问的结构成员之间的一个句号。您可以使用 **struct** 关键字来定义结构类型的变量。下面的实例演示了结构的用法：

```c
#include <stdio.h>
#include <string.h>
 
struct Books
{
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
};
 
int main( )
{
   struct Books Book1;        /* 声明 Book1，类型为 Books */
   struct Books Book2;        /* 声明 Book2，类型为 Books */
 
   /* Book1 详述 */
   strcpy( Book1.title, "C Programming");
   strcpy( Book1.author, "Nuha Ali"); 
   strcpy( Book1.subject, "C Programming Tutorial");
   Book1.book_id = 6495407;
 
   /* Book2 详述 */
   strcpy( Book2.title, "Telecom Billing");
   strcpy( Book2.author, "Zara Ali");
   strcpy( Book2.subject, "Telecom Billing Tutorial");
   Book2.book_id = 6495700;
 
   /* 输出 Book1 信息 */
   printf( "Book 1 title : %s\n", Book1.title);
   printf( "Book 1 author : %s\n", Book1.author);
   printf( "Book 1 subject : %s\n", Book1.subject);
   printf( "Book 1 book_id : %d\n", Book1.book_id);
 
   /* 输出 Book2 信息 */
   printf( "Book 2 title : %s\n", Book2.title);
   printf( "Book 2 author : %s\n", Book2.author);
   printf( "Book 2 subject : %s\n", Book2.subject);
   printf( "Book 2 book_id : %d\n", Book2.book_id);
 
   return 0;
}
```

## 结构作为函数参数

您可以把结构作为函数参数，传参方式与其他类型的变量或指针类似。您可以使用上面实例中的方式来访问结构变量：

这种方式是**值传递**，在函数内部从新开辟内存将传来的结构体拷贝一份，函数内部修改结构体不会改变外部的结构体属性

```c
#include <stdio.h>
#include <string.h>
 
struct Books
{
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
};
 
/* 函数声明 */
void printBook( struct Books book );
int main( )
{
   struct Books Book1;        /* 声明 Book1，类型为 Books */
   struct Books Book2;        /* 声明 Book2，类型为 Books */
 
   /* Book1 详述 */
   strcpy( Book1.title, "C Programming");
   strcpy( Book1.author, "Nuha Ali"); 
   strcpy( Book1.subject, "C Programming Tutorial");
   Book1.book_id = 6495407;
 
   /* Book2 详述 */
   strcpy( Book2.title, "Telecom Billing");
   strcpy( Book2.author, "Zara Ali");
   strcpy( Book2.subject, "Telecom Billing Tutorial");
   Book2.book_id = 6495700;
 
   /* 输出 Book1 信息 */
   printBook( Book1 );
   printf( "Book1 title after modify: %s\n", Book1.title);
   /* 输出 Book2 信息 */
   printBook( Book2 );
 
   return 0;
}
void printBook( struct Books book )
{
   printf( "Book title : %s\n", book.title);
   strcpy( book.title, "modify title");
   printf( "Book title in printBook: %s\n", book.title);
   printf( "Book author : %s\n", book.author);
   printf( "Book subject : %s\n", book.subject);
   printf( "Book book_id : %d\n", book.book_id);
}
```

## 指向结构的指针

您可以定义指向结构的指针，方式与定义指向其他类型变量的指针相似，如下所示：

```c
struct Books *struct_pointer;//指针指向的地址存放的是 Books 类型的结构体数据
```

现在，您可以在上述定义的指针变量中存储结构变量的地址。为了查找结构变量的地址，请把 & 运算符放在结构名称的前面，如下所示：

```c
struct_pointer = &Book1;
```

为了使用指向该结构的指针访问结构的成员，您必须使用 -> 运算符，如下所示：

```c
struct_pointer->title;
```

这种方式是**引用传递**，函数参数是指针，传入函数的是内存地址，函数内部修改结构体属性实际上是在修改外部的结构体属性。

```c
#include <stdio.h>
#include <string.h>
 
struct Books
{
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
};
 
/* 函数声明 */
void printBook( struct Books *book );
int main( )
{
   struct Books Book1;        /* 声明 Book1，类型为 Books */
   struct Books Book2;        /* 声明 Book2，类型为 Books */
 
   /* Book1 详述 */
   strcpy( Book1.title, "C Programming");
   strcpy( Book1.author, "Nuha Ali"); 
   strcpy( Book1.subject, "C Programming Tutorial");
   Book1.book_id = 6495407;
 
   /* Book2 详述 */
   strcpy( Book2.title, "Telecom Billing");
   strcpy( Book2.author, "Zara Ali");
   strcpy( Book2.subject, "Telecom Billing Tutorial");
   Book2.book_id = 6495700;
 
   /* 通过传 Book1 的地址来输出 Book1 信息 */
   printBook( &Book1 );
   printf( "Book1 title after modify: %s\n", Book1.title);
 
   /* 通过传 Book2 的地址来输出 Book2 信息 */
   printBook( &Book2 );
   printf( "Book2 title after modify: %s\n", Book1.title);
 
   return 0;
}
void printBook( struct Books *book )
{
   printf( "Book title : %s\n", book->title);
   strcpy( book->title, "modify title");
   printf( "Book title in printBook: %s\n", book->title);
}
/*
Book title : C Programming
Book title in printBook: modify title
Book1 title after modify: modify title
Book title : Telecom Billing
Book title in printBook: modify title
Book2 title after modify: modify title
*/
```

# 共用体

共用体是一种特殊的数据类型，允许您在**相同的内存位置**存储不同的数据类型。您可以定义一个带有多成员的共用体，但是任何时候**只能有一个成员带有值**。共用体提供了一种使用相同的内存位置的有效方式。

## 定义共用体

为了定义共用体，您必须使用 **union** 语句，方式与定义结构类似。union 语句定义了一个新的数据类型，带有多个成员。union 语句的格式如下：

```c
union [union tag]
{
   member definition;
   member definition;
   ...
   member definition;
} [one or more union variables];
```

**union tag** 是可选的，每个 member definition 是标准的变量定义，比如 int i; 或者 float f; 或者其他有效的变量定义。在共用体定义的末尾，最后一个分号之前，您可以指定一个或多个共用体变量，这是可选的。下面定义一个名为 Data 的共用体类型，有三个成员 i、f 和 str：

```c
union Data
{
   int i;
   float f;
   char  str[20];
} data;
```

现在，**Data** 类型的变量可以存储一个整数、一个浮点数，或者一个字符串。这意味着一个变量（相同的内存位置）可以存储多个多种类型的数据。您可以根据需要在一个共用体内使用任何内置的或者用户自定义的数据类型。

**共用体占用的内存应足够存储共用体中最大的成员**。例如，在上面的实例中，Data 将占用 20 个字节的内存空间，因为在各个成员中，字符串所占用的空间是最大的。下面的实例将显示上面的共用体占用的总内存大小：

```c
#include <stdio.h>
#include <string.h>
 
union Data
{
   int i;
   float f;
   char  str[20];
};
 
int main( )
{
   union Data data;        
   printf( "Memory size occupied by data : %d\n", sizeof(data));
   return 0;
}
//Memory size occupied by data : 20
```

## 访问共用体成员

为了访问共用体的成员，我们使用**成员访问运算符（.）**。下面的实例演示了共用体的用法：

```c
#include <stdio.h>
#include <string.h>
 
union Data
{
   int i;
   float f;
   char  str[20];
};
 
int main( )
{
   union Data data;        
 
   data.i = 10;
   data.f = 220.5;
   strcpy( data.str, "C Programming");
 
   printf( "data.i : %d\n", data.i);
   printf( "data.f : %f\n", data.f);
   printf( "data.str : %s\n", data.str);
 
   return 0;
}
/*
data.i : 1917853763
data.f : 4122360580327794860452759994368.000000
data.str : C Programming
*/
```

在这里，我们可以看到共用体的 **i** 和 **f** 成员的**值有损坏**，因为最后赋给变量的值**占用了内存位置**，这也是 **str** 成员能够完好输出的原因。现在让我们再来看一个相同的实例，这次我们在同一时间只使用一个变量，这也演示了使用共用体的主要目的：

```c
#include <stdio.h>
#include <string.h>
 
union Data
{
   int i;
   float f;
   char  str[20];
};
 
int main( )
{
   union Data data;        
 
   data.i = 10;
   printf( "data.i : %d\n", data.i);
   
   data.f = 220.5;
   printf( "data.i after f: %d\n", data.i); //损毁
   printf( "data.f : %f\n", data.f);
   
   strcpy( data.str, "C Programming");
   printf( "data.i after str: %d\n", data.i); //损毁
   printf( "data.f after str: %f\n", data.f); //损毁
   printf( "data.str : %s\n", data.str);
 
   return 0;
}
/*data.i : 10
data.i after f: 1130135552  //损毁
data.f : 220.500000
data.i after str: 1917853763  //损毁
data.f after str: 4122360580327794860452759994368.000000  //损毁
data.str : C Programming*/
```

# 位域

如果程序的结构中包含多个开关量，只有 **TRUE/FALSE** 变量，如下：

```c
struct
{
  unsigned int widthValidated;
  unsigned int heightValidated;
} status;
```

**这种结构需要 8 字节的内存空间**（为什么是8字节？），但在实际上，**在每个变量中，我们只存储 0 或 1**。在这种情况下，C 语言提供了一种**更好的利用内存空间的方式**。如果您在结构内使用这样的变量，您可以定义变量的宽度来告诉编译器，您将只使用这些字节。例如，上面的结构可以重写成：

```c
struct
{
  unsigned int widthValidated : 1;
  unsigned int heightValidated : 1;
} status;
```

现在，上面的结构中，status 变量将占用 4 个字节的内存空间（每段内存是4个字节），但是只有 2 位被用来存储值。如果您用了 32 个变量，每一个变量宽度为 1 位，那么 status 结构将使用 4 个字节，但只要您再多用一个变量，如果使用了 33 个变量，那么它将分配**内存的下一段**来存储第 33 个变量，这个时候就开始使用 8 个字节。让我们看看下面的实例来理解这个概念：

```c
#include <stdio.h>
#include <string.h>
 
/* 定义简单的结构 */
struct
{
  unsigned int widthValidated;
  unsigned int heightValidated;
} status1;
 
/* 定义位域结构 */
struct
{
  unsigned int widthValidated : 1;
  unsigned int heightValidated : 1;
} status2;
 
int main( )
{
   printf( "Memory size occupied by status1 : %d\n", sizeof(status1));
   printf( "Memory size occupied by status2 : %d\n", sizeof(status2));
 
   return 0;
}
/*Memory size occupied by status1 : 8
Memory size occupied by status2 : 4*/
```

## 位域声明

有些信息在存储时，并**不需要占用一个完整的字节**，而只需占几个或一个二进制位。例如在存放一个开关量时，只有 0 和 1 两种状态，用 1 位二进位即可。为了**节省存储空间**，并使处理简便，C 语言又提供了一种数据结构，称为"位域"或"位段"。

所谓"位域"是把**一个字节中的二进位划分为几个不同的区域**，并说明每个区域的位数。每个域有一个域名，允许在程序中按域名进行操作。这样就可以**把几个不同的对象用一个字节的二进制位域来表示**。

典型的实例：

- 用 1 位二进位存放一个**开关量**时，只有 0 和 1 两种状态。
- 读取外部文件格式——可以读取非标准的文件格式。例如：9 位的整数。

### 位域的定义和位域变量的说明

位域定义与结构定义相仿，其形式为：

```c
struct 位域结构名 
{
 type [member_name] : width;
};
```

下面是有关位域中变量元素的描述：

| 元素        | 描述                                                         |
| :---------- | :----------------------------------------------------------- |
| type        | 只能为 int(整型)，unsigned int(无符号整型)，signed int(有符号整型) 三种类型，决定了如何解释位域的值。 |
| member_name | 位域的名称。                                                 |
| width       | 位域中位的数量。宽度必须小于或等于指定类型的位宽度。         |

带有预定义宽度的变量被称为**位域**。位域可以存储多于 1 位的数，例如，需要一个变量来存储从 0 到 7 的值，您可以定义一个宽度为 3 位的位域，如下：

```c
struct
{
  unsigned int age : 3;
} Age;
```

上面的结构定义指示 C 编译器，age 变量将只使用 3 位来存储这个值，如果您试图使用超过 3 位，则无法完成。

```c
struct bs{
    int a:8;
    int b:2;
    int c:6;
}data;
```

data 为 bs 变量，共占两个字节。其中位域 a 占 8 位，位域 b 占 2 位，位域 c 占 6 位。

```C
struct packed_struct {
  unsigned int f1:1;
  unsigned int f2:1;
  unsigned int f3:1;
  unsigned int f4:1;
  unsigned int type:4;
  unsigned int my_int:9;
} pack;
//在这里，packed_struct 包含了 6 个成员：四个 1 位的标识符 f1..f4、一个 4 位的 type 和一个 9 位的 my_int。
```

```C
#include <stdio.h>
#include <string.h>
 
struct
{
  unsigned int age : 3;
} Age;
 
int main( )
{
   Age.age = 4;
   printf( "Sizeof( Age ) : %d\n", sizeof(Age) );
   printf( "Age.age : %d\n", Age.age );
 
   Age.age = 7;
   printf( "Age.age : %d\n", Age.age );
 
   Age.age = 8; // 二进制表示为 1000 有四位，超出
   printf( "Age.age : %d\n", Age.age );
 
   return 0;
}
/*Sizeof( Age ) : 4
Age.age : 4
Age.age : 7
Age.age : 0
*/
```

**对于位域的定义尚有以下几点说明：**

- 一个位域存储在同一个字节中，如一个字节所剩空间不够存放另一位域时，则会从下一单元起存放该位域。也可以**有意使某位域从下一单元开始**。例如：

  ```c
  struct bs{
      unsigned a:4;
      unsigned  :4;    /* 空域 */
      unsigned b:4;    /* 从下一单元开始存放 */
      unsigned c:4
  }
  ```

  在这个位域定义中，a 占第一字节的 4 位，后 4 位填 0 表示不使用，b 从第二字节开始，占用 4 位，c 占用 4 位。

- 位域的宽度**不能超过它所依附的数据类型的长度**，成员变量都是有类型的，这个类型限制了成员变量的最大长度，**:** 后面的数字不能超过这个长度。

- 位域可以是**无名位域**，这时它只用来作**填充或调整位置**。无名的位域**是不能使用的**。

从以上分析可以看出，位域在本质上就是一种结构类型，不过其成员是按二进位分配的。

### 位域的使用

位域的使用和结构成员的使用相同，其一般形式为：

```c
位域变量名.位域名
位域变量名->位域名
```

位域允许用各种格式输出。

```c
#include <stdio.h>
 
int main(){
    struct bs{
        unsigned a:1;
        unsigned b:3;
        unsigned c:4;
    } bit,*pbit;
    bit.a=1;    /* 给位域赋值（应注意赋值不能超过该位域的允许范围） */
    bit.b=7;    /* 给位域赋值（应注意赋值不能超过该位域的允许范围） */
    bit.c=15;    /* 给位域赋值（应注意赋值不能超过该位域的允许范围） */
    printf("%d,%d,%d\n",bit.a,bit.b,bit.c);    /* 以整型量格式输出三个域的内容 */
    pbit=&bit;    /* 把位域变量 bit 的地址送给指针变量 pbit */
    pbit->a=0;    /* 用指针方式给位域 a 重新赋值，赋为 0 */
    pbit->b&=3;    /* 使用了复合的位运算符 "&="，相当于：pbit->b=pbit->b&3，位域 b 中原有值为 7，与 3 作按位与运算的结果为 3（111&011=011，十进制值为 3） */
    pbit->c|=1;    /* 使用了复合位运算符"|="，相当于：pbit->c=pbit->c|1，其结果为 15 */
    printf("%d,%d,%d\n",pbit->a,pbit->b,pbit->c);    /* 用指针方式输出了这三个域的值 */
}
```

# typedef

C 语言提供了 **typedef** 关键字，您可以使用它来**为类型取一个新的名字**。下面的实例为单字节数字定义了一个术语 **BYTE**：

```c
typedef unsigned char BYTE;
```

在这个类型定义之后，标识符 BYTE 可作为类型 **unsigned char 的缩写**，例如：

```c
BYTE  b1, b2;
```

按照惯例，定义时会**大写字母**，以便提醒用户类型名称是一个象征性的缩写，但您**也可以使用小写字母**。

您也可以使用 **typedef** 来为用户**自定义的数据类型取一个新的名字**。例如，您可以对结构体使用 typedef 来定义一个**新的数据类型名字**，然后**使用这个新的数据类型来直接定义结构变量**，如下：

```c
#include <stdio.h>
#include <string.h>
 
typedef struct Books
{
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
} Book;
 
int main( )
{
   Book book;//使用这个新的数据类型来直接定义 结构变量book
 
   strcpy( book.title, "C 教程");
   strcpy( book.author, "Runoob"); 
   strcpy( book.subject, "编程语言");
   book.book_id = 12345;
 
   printf( "书标题 : %s\n", book.title);
 
   return 0;
}
```

## typedef vs #define

**#define** 是 C **指令**，用于为各种数据类型**定义别名**，与 **typedef** 类似，但是它们有以下几点不同：

- **typedef** 仅限于为类型定义符号名称，**#define** 不仅可以为类型定义别名，也能为数值定义别名，比如您可以定义 1 为 ONE。
- **typedef** 是由**编译器执行解释**的，**#define** 语句是由**预编译器**进行处理的。

下面是 #define 的最简单的用法：

```c
#include <stdio.h>
 
#define TRUE  1
#define FALSE 0
 
int main( )
{
   printf( "TRUE 的值: %d\n", TRUE);
   printf( "FALSE 的值: %d\n", FALSE);
 
   return 0;
}
/*TRUE 的值: 1
FALSE 的值: 0*/
```

# 输入 & 输出

## 标准文件

C 语言把所有的设备**都当作文件**。所以设备（比如显示器）被处理的方式与文件相同。以下三个文件会在程序执行时自动打开，以便访问键盘和屏幕。

| 标准文件 | 文件指针 | 设备     |
| :------- | :------- | :------- |
| 标准输入 | stdin    | 键盘     |
| 标准输出 | stdout   | 屏幕     |
| 标准错误 | stderr   | 您的屏幕 |

文件指针是访问文件的方式，C 语言中的 I/O (输入/输出) 通常使用 printf() 和 scanf() 两个函数。

**实例解析：**

- 所有的 C 语言程序都需要包含 **main()** 函数。 代码从 **main()** 函数开始执行。
- **printf()** 用于格式化输出到屏幕。**printf()** 函数在 **"stdio.h"** 头文件中声明。
- **stdio.h** 是一个头文件 (标准输入输出头文件) and **#include** 是一个预处理命令，用来引入头文件。 当编译器遇到 **printf()** 函数时，如果没有找到 **stdio.h** 头文件，会发生编译错误。
- **return 0;** 语句用于表示退出程序。

## getchar() & putchar() 函数

**int getchar(void)** 函数从屏幕读取下一个可用的字符，并把它返回为一个整数。这个函数在同一个时间内只会读取一个单一的字符。您可以在循环内使用这个方法，以便从屏幕上读取多个字符。

**int putchar(int c)** 函数把字符输出到屏幕上，并返回相同的字符。这个函数在同一个时间内只会输出一个单一的字符。您可以在循环内使用这个方法，以便在屏幕上输出多个字符。

```c
#include <stdio.h>
 
int main( )
{
   int c;
 
   printf( "Enter a value :");
   c = getchar( );
 
   printf( "\nYou entered: ");
   putchar( c );
   printf( "\n");
   return 0;
}
```

它会等待您输入一些文本，当您输入一个文本并按下回车键时，程序会继续并只会读取一个单一的字符，显示如下：

```shell
$./a.out
Enter a value :runoob

You entered: r
```

## gets() & puts() 函数

**char \*gets(char \*s)** 函数从 **stdin** 读取一行到 **s** 所指向的缓冲区，直到一个终止符或 EOF。

**int puts(const char \*s)** 函数把字符串 s 和一个尾随的换行符写入到 **stdout**。

```c
#include <stdio.h>
 
int main( )
{
   char str[100];
 
   printf( "Enter a value :");
   gets( str );
 
   printf( "\nYou entered: ");
   puts( str );
   return 0;
}
```

它会等待您输入一些文本，当您输入一个文本并按下回车键时，程序会继续并读取一整行直到该行结束，显示如下：

```
$./a.out
Enter a value :runoob

You entered: runoob
```

# 文件读写

## 打开文件

您可以使用 **fopen( )** 函数来创建一个新的文件或者打开一个已有的文件，这个调用会初始化类型 **FILE** 的一个对象，类型 **FILE** 包含了所有用来控制流的必要的信息。下面是这个函数调用的原型：

```c
FILE *fopen( const char *filename, const char *mode );
```

在这里，**filename** 是字符串，用来命名文件，访问模式 **mode** 的值可以是下列值中的一个：

| 模式 | 描述                                                         |
| :--- | :----------------------------------------------------------- |
| r    | 打开一个已有的文本文件，允许读取文件。                       |
| w    | 打开一个文本文件，允许写入文件。如果文件不存在，则会创建一个新文件。在这里，您的程序会从文件的开头写入内容。如果文件存在，则该会被截断为零长度，重新写入。 |
| a    | 打开一个文本文件，以追加模式写入文件。如果文件不存在，则会创建一个新文件。在这里，您的程序会在已有的文件内容中追加内容。 |
| r+   | 打开一个文本文件，允许读写文件。                             |
| w+   | 打开一个文本文件，允许读写文件。如果文件已存在，则文件会被截断为零长度，如果文件不存在，则会创建一个新文件。 |
| a+   | 打开一个文本文件，允许读写文件。如果文件不存在，则会创建一个新文件。读取会从文件的开头开始，写入则只能是追加模式。 |

如果处理的是二进制文件，则需使用下面的访问模式来取代上面的访问模式：

```shell
"rb", "wb", "ab", "rb+", "r+b", "wb+", "w+b", "ab+", "a+b"
```

## 关闭文件

为了关闭文件，请使用 fclose( ) 函数。函数的原型如下：

```
 int fclose( FILE *fp );
```

如果成功关闭文件，**fclose( )** 函数返回零，如果关闭文件时发生错误，函数返回 **EOF**。这个函数实际上，会清空缓冲区中的数据，关闭文件，并释放用于该文件的所有内存。EOF 是一个定义在头文件 **stdio.h** 中的常量。

## 写入文件

下面是把字符写入到流中的最简单的函数：

```c
int fputc( int c, FILE *fp );
```

函数 **fputc()** 把参数 c 的字符值写入到 **fp 所指向的输出流**中。如果写入成功，它会返回写入的字符，如果发生错误，则会返回 **EOF**。您可以使用下面的函数来把一个以 null 结尾的字符串写入到流中：

```c
int fputs( const char *s, FILE *fp );
```

函数 **fputs()** 把字符串 **s** 写入到 **fp 所指向的输出流**中。如果写入成功，它会返回一个非负值，如果发生错误，则会返回 **EOF**。您也可以使用 **int fprintf(FILE \*fp,const char \*format, ...)** 函数把一个字符串写入到文件中。尝试下面的实例：

> **注意：**请确保您有可用的 **tmp** 目录，如果不存在该目录，则需要在您的计算机上先创建该目录。
>
> **/tmp** 一般是 Linux 系统上的临时目录，如果你在 Windows 系统上运行，则需要修改为本地环境中已存在的目录，例如: **C:\tmp**、**D:\tmp**等。

```c
#include <stdio.h>
 
int main()
{
   FILE *fp = NULL;
 
   fp = fopen("/tmp/test.txt", "w+");
   fprintf(fp, "This is testing for fprintf...\n");
   fputs("This is testing for fputs...\n", fp);
   fclose(fp);
}
```

## 读取文件

下面是从文件读取单个字符的最简单的函数：

```c
int fgetc( FILE * fp );
```

**fgetc()** 函数从 fp 所指向的输入文件中读取一个字符。返回值是读取的字符，如果发生错误则返回 **EOF**。下面的函数允许您从流中读取一个字符串：

```c
char *fgets( char *buf, int n, FILE *fp );
```

函数 **fgets()** 从 fp 所指向的输入流中读取 n - 1 个字符。它会把读取的字符串复制到缓冲区 **buf**，并在最后追加一个 **null** 字符来终止字符串。如果这个函数在读取最后一个字符之前就遇到一个换行符 '\n' 或文件的末尾 EOF，则**结束返回读取到的字符**，包括换行符。

您也可以使用 **int fscanf(FILE \*fp, const char \*format, ...)** 函数来从文件中读取字符串，但是在遇到第一个空格和换行符时，它**会停止读取**。

```c
#include <stdio.h>
 
int main()
{
   FILE *fp = NULL;
   char buff[255];
 
   fp = fopen("/tmp/test.txt", "r");
   fscanf(fp, "%s", buff);//遇到第一个空格和换行符时，它会停止读取
   printf("1: %s\n", buff );
 
   fgets(buff, 255, (FILE*)fp);//继续读取buff，遇到一个换行符 '\n' 或文件的末尾 EOF，则结束，返回读取到的字符包括换行符
   printf("2: %s\n", buff );
   
   fgets(buff, 255, (FILE*)fp);//继续读取buff
   printf("3: %s\n", buff );
   fclose(fp);
 
}
/*
1: This
2: is testing for fprintf...

3: This is testing for fputs...
*/
```

## 二进制 I/O 函数

下面两个函数用于二进制输入和输出：

```c
size_t fread(void *ptr, size_t size_of_elements, size_t number_of_elements, FILE *a_file);
              
size_t fwrite(const void *ptr, size_t size_of_elements, size_t number_of_elements, FILE *a_file);
```

这两个函数都是用于**存储块**的读写，通常是数组或结构体。

# 预处理器

**C 预处理器**不是编译器的组成部分，但是它是**编译过程中一个单独的步骤**。简言之，C 预处理器只不过是一个**文本替换工具**而已，它们会指示编译器在实际编译之前完成所需的预处理。我们将把 C 预处理器（C Preprocessor）简写为 **CPP**。

所有的**预处理器命令**都是以井号（#）开头。它必须是第一个非空字符，为了增强可读性，预处理器指令应从第一列开始。下面列出了所有重要的预处理器指令：

| 指令     | 描述                                                        |
| :------- | :---------------------------------------------------------- |
| #define  | 定义 **宏**                                                 |
| #include | 包含一个源代码文件                                          |
| #undef   | 取消已定义的宏                                              |
| #ifdef   | 如果宏已经定义，则返回真                                    |
| #ifndef  | 如果宏没有定义，则返回真                                    |
| #if      | 如果给定条件为真，则**编译下面代码**                        |
| #else    | #if 的替代方案                                              |
| #elif    | 如果前面的 #if 给定条件不为真，当前条件为真，则编译下面代码 |
| #endif   | **结束一个** #if……#else 条件**编译块**                      |
| #error   | 当遇到标准错误时，输出错误消息                              |
| #pragma  | 使用标准化方法，向编译器发布特殊的命令到编译器中            |

## 预处理器实例

分析下面的实例来理解不同的指令。

```c
//这个指令告诉 CPP 把所有的 MAX_ARRAY_LENGTH 定义为 20。使用 *#define* 定义常量来增强可读性。
#define MAX_ARRAY_LENGTH 20

//这些指令告诉 CPP 从系统库中获取 stdio.h，并添加文本到当前的源文件中。
#include <stdio.h>
//告诉 CPP 从本地目录中获取 myheader.h，并添加内容到当前的源文件中。
#include "myheader.h"

//取消已定义的 FILE_SIZE，并定义它为 42。
#undef  FILE_SIZE
#define FILE_SIZE 42

//只有当 MESSAGE 未定义时，才定义 MESSAGE。
#ifndef MESSAGE
   #define MESSAGE "You wish!"
#endif

//如果定义了 DEBUG，则执行处理语句。
#ifdef DEBUG
   /* Your debugging statements here */
#endif
```

## 预定义宏

define宏定义

以#号开头的都是编译预处理指令，而它们不是C语言的成分。
#define用来定义一个宏，程序在预处理阶段将会使用define定义的宏去**将内容进行了替换**。

```
不带参数的宏：#define <名字> <值>
带参数的宏(像函数的宏)：#define 宏名(形参表）字符串
```

>宏定义不是说明或语句，在行末不必加分号，如加上分号则连分号也一起置换。宏定义其作用域为宏定义命令起到源程序结束。如要终止其作用域可使用#undef命令。

ANSI C 定义了许多宏。在编程中您可以**使用这些宏**，但是**不能直接修改这些预定义的宏**。

| 宏       | 描述                                                |
| :------- | :-------------------------------------------------- |
| __DATE__ | 当前日期，一个以 "MMM DD YYYY" 格式表示的字符常量。 |
| __TIME__ | 当前时间，一个以 "HH:MM:SS" 格式表示的字符常量。    |
| __FILE__ | 当前文件名，一个字符串常量。                        |
| __LINE__ | 当前行号，一个十进制常量。                          |
| __STDC__ | 当编译器以 ANSI 标准编译时，则定义为 1。            |

```C
//代码在文件 test.c 中
#include <stdio.h>
 
main()
{
   printf("File :%s\n", __FILE__ );
   printf("Date :%s\n", __DATE__ );
   printf("Time :%s\n", __TIME__ );
   printf("Line :%d\n", __LINE__ );
   printf("ANSI :%d\n", __STDC__ );
 
}
/*File :test.c
Date :Jun 2 2012
Time :03:36:24
Line :8
ANSI :1*/
```

## 预处理器运算符

C 预处理器提供了下列的运算符来帮助您创建宏：

### 宏延续运算符（\）

一个宏通常写在一个单行上。但是如果宏太长，一个**单行容纳不下**，则使用宏延续运算符（\）。例如：

```C
#define  message_for(a, b)  \
    printf(#a " and " #b ": We love you!\n")
```

### 字符串常量化运算符（#）

在宏定义中，当需要把一个**宏的参数转换为字符串常量**时，则使用字符串常量化运算符（#）。在宏中使用的该运算符有一个特定的参数或参数列表。例如：

```c
#include <stdio.h>
 
#define  message_for(a, b)  \
    printf(#a " and " #b ": We love you!\n")
 
int main(void)
{
   message_for(Carole, Debra);
   return 0;
}
//Carole and Debra: We love you!
```

### 标记粘贴运算符（##）

宏定义内的标记粘贴运算符（##）会合并两个参数。它允许在宏定义中两个独立的标记被合并为一个标记。例如：

```c
#include <stdio.h>
 
#define tokenpaster(n) printf ("token" #n " = %d", token##n)
 
int main(void)
{
   int token34 = 40; 
   tokenpaster(34);
   return 0;
}
```

当上面的代码被编译和执行时，它会产生下列结果：

```c
token34 = 40
```

这是怎么发生的，因为这个实例会从编译器产生下列的实际输出：

```c
printf ("token34 = %d", token34);
```

这个实例演示了 token##n 会连接到 token34 中，在这里，我们使用了**字符串常量化运算符（#）**和**标记粘贴运算符（##）**。

### defined() 运算符

预处理器 **defined** 运算符是用在常量表达式中的，用来确定一个**标识符是否已经使用 #define 定义过**。如果指定的标识符已定义，则值为真（非零）。如果指定的标识符未定义，则值为假（零）。下面的实例演示了 defined() 运算符的用法：

```c
#include <stdio.h>
 
#if !defined (MESSAGE)
   #define MESSAGE "You wish!"
#endif
 
int main(void)
{
   printf("Here is the message: %s\n", MESSAGE);  
   return 0;
}
//Here is the message: You wish!
```

## 参数化的宏

CPP 一个强大的功能是可以使用**参数化的宏来模拟函数**。例如，下面的代码是计算一个数的平方：

```C
int square(int x) {
   return x * x;
}
```

我们可以使用宏重写上面的代码，如下：

```C
#define square(x) ((x) * (x))
```

在使用带有参数的宏之前，必须使用 **#define** 指令定义。参数列表是括在圆括号内，且必须紧跟在宏名称的后边。宏名称和左圆括号之间不允许有空格。例如：

```C
#include <stdio.h>
 
#define MAX(x,y) ((x) > (y) ? (x) : (y))
 
int main(void)
{
   printf("Max between 20 and 10 is %d\n", MAX(10, 20));  
   return 0;
}
//Max between 20 and 10 is 20
```

```c
#include <stdio.h>
 
#define ADD(x,y) ((x) + (y))
 
int main(void)
{
   printf("%d\n", 4*ADD(2, 3));//用((2) + (3))替换ADD(2, 3)
   return 0;
}
//20
```

# 头文件

头文件是扩展名为 **.h** 的文件，包含了 C 函数声明和宏定义，被多个源文件中引用共享。有两种类型的头文件：程序员编写的头文件和编译器自带的头文件。

在程序中要使用头文件，需要使用 C 预处理指令 **#include** 来引用它。引用头文件相当于复制头文件的内容。

## 引用头文件的语法

使用预处理指令 **#include** 可以引用用户和系统头文件。它的形式有以下两种：

引用系统头文件。它在系统目录的标准列表中搜索名为 file 的文件。在编译源代码时，您可以通过 -I 选项把目录前置在该列表前。

```c
#include <file>
```

引用用户头文件。它在包含当前文件的目录中搜索名为 file 的文件。在编译源代码时，您可以通过 -I 选项把目录前置在该列表前。

```c
#include "file"
```

## 引用头文件的操作

**#include** 指令会指示 C 预处理器浏览指定的文件作为输入。预处理器的输出包含了已经生成的输出，被引用文件生成的输出以及 **#include** 指令之后的文本输出。例如，如果您有一个头文件 header.h，如下：

```c
char *test (void);
```

和一个使用了头文件的主程序 *program.c*，如下：

```c
int x;
#include "header.h"

int main (void)
{
   puts (test ());
}
```

编译器会看到如下的代码信息：

```c
int x;
char *test (void);

int main (void)
{
   puts (test ());
}
```

## 只引用一次头文件

如果一个头文件被引用两次，编译器会处理两次头文件的内容，这将产生错误。为了防止这种情况，标准的做法是把文件的整个内容放在条件编译语句中，如下：

```c
#ifndef HEADER_FILE
#define HEADER_FILE

the entire header file file

#endif
```

这种结构就是通常所说的包装器 **#ifndef**。当再次引用头文件时，条件为假，因为 HEADER_FILE 已定义。此时，预处理器会跳过文件的整个内容，编译器会忽略它。

## 有条件引用

有时需要从多个不同的头文件中选择一个引用到程序中。例如，需要指定在不同的操作系统上使用的配置参数。您可以通过一系列条件来实现这点，如下：

```c
#if SYSTEM_1
   # include "system_1.h"
#elif SYSTEM_2
   # include "system_2.h"
#elif SYSTEM_3
   ...
#endif
```

但是如果头文件比较多的时候，这么做是很不妥当的，预处理器**使用宏来定义头文件的名称**。这就是所谓的**有条件引用**。它不是用头文件的名称作为 **#include** 的直接参数，您只需要使用宏名称代替即可：

```c
 #define SYSTEM_H "system_1.h"
 ...
 #include SYSTEM_H
```

SYSTEM_H 会扩展，预处理器会查找 system_1.h，就像 **#include** 最初编写的那样。SYSTEM_H 可通过 -D 选项**被您的 Makefile 定义**。

## 笔记

在有多个 **.h** 文件和多个 **.c** 文件的时候，往往我们会用一个 **global.h** 的头文件来包括所有的 **.h** 文件，然后在除 **global.h** 文件外的头文件中 包含 **global.h** 就可以实现所有头文件的包含，同时不会乱。方便在各个文件里面调用其他文件的函数或者变量。

```c
#ifndef _GLOBAL_H
#define _GLOBAL_H
#include <fstream>
#include <iostream>
#include <math.h>
#include <Config.h>
```

**C 语言中 include <> 与include "" 的区别?**

**#include < >** 引用的是编译器的类库路径里面的头文件。

**#include " "** 引用的是你程序目录的相对路径中的头文件，如果在程序目录没有找到引用的头文件则到编译器的类库路径的目录下找该头文件。

# 强制类型转换

强制类型转换是把变量从一种类型转换为另一种数据类型。例如，如果您想存储一个 long 类型的值到一个简单的整型中，您需要把 long 类型强制转换为 int 类型。您可以使用**强制类型转换运算符**来把值显式地从一种类型转换为另一种类型，如下所示：

```c
(type_name) expression
```

```c
#include <stdio.h>
 
int main()
{
   int sum = 17, count = 5;
   double mean;
 
   mean = (double) sum / count;//sum 的值首先被转换为 double 型，然后除以 count
   printf("Value of mean : %f\n", mean );
 
}
//Value of mean : 3.400000
```

强制类型转换运算符的优先级大于除法，因此 **sum** 的值首先被转换为 **double** 型，然后除以 count，得到一个类型为 double 的值。

## 整数提升

整数提升是指把小于 **int** 或 **unsigned int** 的整数类型转换为 **int** 或 **unsigned int** 的过程。请看下面的实例，在 int 中添加一个字符：

```c
#include <stdio.h>
 
int main()
{
   int  i = 17;
   char c = 'c'; /* ascii 值是 99 */
   int sum;
 
   sum = i + c;
   printf("Value of sum : %d\n", sum );//Value of sum : 116
 
}
```

在这里，sum 的值为 116，因为编译器进行了整数提升，在执行实际加法运算时，**把 'c' 的值转换为对应的 ascii 值**。

## 常用的算术转换

**常用的算术转换**是隐式地把值强制转换为相同的类型。编译器首先执行**整数提升**，如果操作数类型不同，则它们会被转换为下列层次中出现的最高层次的类型：

![Usual Arithmetic Conversion](images/C learning notes.assets/usual_arithmetic_conversion.png)

常用的算术转换不适用于赋值运算符、逻辑运算符 && 和 ||。让我们看看下面的实例来理解这个概念：

```c
#include <stdio.h>
 
int main()
{
   int  i = 17;
   char c = 'c'; /* ascii 值是 99 */
   float sum;
 
   sum = i + c; //c 首先被转换为整数，然后把 i 和 c 转换为浮点型，然后相加得到一个浮点数
   printf("Value of sum : %f\n", sum );
 
}
```

在这里，c 首先被转换为整数，但是由于最后的值是 float 型的，所以会应用常用的算术转换，编译器会把 i 和 c 转换为浮点型，并把它们相加得到一个浮点数。

# 错误处理

C 语言不提供对错误处理的直接支持，但是作为一种系统编程语言，它以返回值的形式允许您访问底层数据。在发生错误时，大多数的 C 或 UNIX 函数调用返回 1 或 NULL，同时会设置一个错误代码 **errno**，该错误代码是全局变量，表示在函数调用期间发生了错误。您可以在 errno.h 头文件中找到各种各样的错误代码。

所以，C 程序员可以通过检查返回值，然后根据返回值决定采取哪种适当的动作。开发人员应该在程序初始化时，把 errno 设置为 0，这是一种良好的编程习惯。0 值表示程序中没有错误。

## errno、perror() 和 strerror()

C 语言提供了 **perror()** 和 **strerror()** 函数来显示与 **errno** 相关的文本消息。

- **perror()** 函数显示您传给它的字符串，后跟一个冒号、一个空格和当前 errno 值的文本表示形式。
- **strerror()** 函数，返回一个指针，指针指向当前 errno 值的文本表示形式。

让我们来模拟一种错误情况，尝试打开一个不存在的文件。您可以使用多种方式来输出错误消息，在这里我们使用函数来演示用法。另外有一点需要注意，您应该使用 **stderr** 文件流来输出所有的错误。

```c
#include <stdio.h>
#include <errno.h>
#include <string.h>
 
extern int errno ;
 
int main ()
{
   FILE * pf;
   int errnum;
   pf = fopen ("unexist.txt", "rb");
   if (pf == NULL)
   {
      errnum = errno;
      fprintf(stderr, "错误号: %d\n", errno);
      perror("通过 perror 输出错误");
      fprintf(stderr, "打开文件错误: %s\n", strerror( errnum ));
   }
   else
   {
      fclose (pf);
   }
   return 0;
}
/*
错误号: 2
通过 perror 输出错误: No such file or directory
打开文件错误: No such file or directory
*/
```

## 程序退出状态

通常情况下，程序成功执行完一个操作正常退出的时候会带有值 EXIT_SUCCESS。在这里，EXIT_SUCCESS 是宏，它被定义为 0。

如果程序中存在一种错误情况，当您退出程序时，会带有状态值 EXIT_FAILURE，被定义为 -1。

```c
#include <stdio.h>
#include <stdlib.h>
 
main()
{
   int dividend = 20;
   int divisor = 5;
   int quotient;
 
   if( divisor == 0){
      fprintf(stderr, "除数为 0 退出运行...\n");
      exit(EXIT_FAILURE);
   }
   quotient = dividend / divisor;
   fprintf(stderr, "quotient 变量的值为: %d\n", quotient );
 
   exit(EXIT_SUCCESS);
}
//quotient 变量的值为 : 4
```

# 可变参数

希望函数带有可变数量的参数，而不是预定义数量的参数。为了使用这个功能，您需要使用 **stdarg.h** 头文件，该文件提供了实现可变参数功能的函数和宏。具体步骤如下：

1. 函数参数为省略号

   定义一个函数，最后一个参数为**省略号**，省略号前面可以设置自定义参数。

2. 定义参数列表

   在函数定义中创建一个 **va_list** 类型变量，该类型是在 stdarg.h 头文件中定义的。

3. 初始化参数列表

   使用 **int** 参数和 **va_start** 宏来初始化 **va_list** 变量为一个参数列表。宏 va_start 是在 stdarg.h 头文件中定义的。

4. 访问参数列表

   使用 **va_arg** 宏和 **va_list** 变量来访问参数列表中的每个项。

5. 清理内存

   使用宏 **va_end** 来清理赋予 **va_list** 变量的内存。

```c
#include <stdio.h>
#include <stdarg.h>
 
double average(int num,...)
{
 
    va_list valist;
    double sum = 0.0;
    int i;
 
    /* 为 num 个参数初始化 valist */
    va_start(valist, num);
 
    /* 访问所有赋给 valist 的参数 */
    for (i = 0; i < num; i++)
    {
       sum += va_arg(valist, int);
    }
    /* 清理为 valist 保留的内存 */
    va_end(valist);
 
    return sum/num;
}
 
int main()
{
   printf("Average of 2, 3, 4, 5 = %f\n", average(4, 2,3,4,5));
   printf("Average of 5, 10, 15 = %f\n", average(3, 5,10,15));
}
```

# 内存管理

本章将讲解 C 中的动态内存管理。C 语言为内存的分配和管理提供了几个函数。这些函数可以在 **<stdlib.h>** 头文件中找到。

| 序号 | 函数和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **void \*calloc(int num, int size);**     在内存中  **动态地分配**  num 个长度为 size 的连续空间，并将每一个字节都初始化为 0。所以它的结果是分配了 num*size 个字节长度的内存空间，并且每个字节的值都是0。 |
| 2    | **void free(void \*address);**              该函数**释放 address 所指向的内存块**，释放的是动态分配的内存空间。 |
| 3    | **void \*malloc(int num);**                   在**堆区分配一块指定大小**的内存空间，用来存放数据。这块内存空间在函数**执行完成后不会被初始化**，它们的值是未知的。 |
| 4    | **void \*realloc(void \*address, int newsize);**                      该函数**重新分配内存**，把内存扩展到 **newsize**。 |

**注意：**void * 类型表示未确定类型的指针。C、C++ 规定 void * 类型可以通过类型转换强制转换为任何其它类型的指针。

## 动态分配内存

编程时，如果预先不知道需要存储的文本长度，例如您想存储有关一个主题的详细描述。在这里，我们需要定义一个指针，该指针指向未定义所需内存大小的字符，后续再根据需求来分配内存，如下所示：

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
 
int main()
{
   char name[100];
   char *description;
 
   strcpy(name, "Zara Ali");
 
   /* 动态分配内存 */
   description = (char *)malloc( 200 * sizeof(char) );
   if( description == NULL )
   {
      fprintf(stderr, "Error - unable to allocate required memory\n");
   }
   else
   {
      strcpy( description, "Zara ali a DPS student in class 10th");
   }
   printf("Name = %s\n", name );
   printf("Description: %s\n", description );
}
/*
Name = Zara Ali
Description: Zara ali a DPS student in class 10th
*/
```

上面的程序也可以使用 **calloc()** 来编写，只需要把 malloc 替换为 calloc 即可，如下所示：

```c
calloc(200, sizeof(char));//分配200个长度为 size 的连续空间
```

当动态分配内存时，您有完全控制权，可以传递任何大小的值。而那些预先定义了大小的数组，一旦定义则无法改变大小。

## 重新调整内存的大小和释放内存

当程序退出时，操作系统会自动释放所有分配给程序的内存，但是，建议您在不需要内存时，都应该调用函数 **free()** 来释放内存。

或者，您可以通过调用函数 **realloc()** 来增加或减少已分配的内存块的大小。

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
 
int main()
{
   char name[100];
   char *description;
 
   strcpy(name, "Zara Ali");
 
   /* 动态分配内存 */
   description = (char *)malloc( 30 * sizeof(char) );
   if( description == NULL )
   {
      fprintf(stderr, "Error - unable to allocate required memory\n");
   }
   else
   {
      strcpy( description, "Zara ali a DPS student.");
   }
   /* 假设您想要存储更大的描述信息 */
   description = (char *) realloc( description, 100 * sizeof(char) );
   if( description == NULL )
   {
      fprintf(stderr, "Error - unable to allocate required memory\n");
   }
   else
   {
      strcat( description, "She is in class 10th");//拼接字符串
   }
   
   printf("Name = %s\n", name );
   printf("Description: %s\n", description );
 
   /* 使用 free() 函数释放内存 */
   free(description);
}
/*
Name = Zara Ali
Description: Zara ali a DPS student.She is in class 10th
*/
```

# 命令行参数

命令行参数是使用 main() 函数参数来处理的，其中，**argc** 是指传入参数的个数，**argv[]** 是一个指针数组，指向传递给程序的每个参数。

Linux 下我们可使用 getopt 和 getopt_long 来对命令行参数进行解析，例如：

```c
int main(int argc, char *argv[]){
    char *optstr = "p:n:m:c:";
    struct option opts[] = {
        {"path", 1, NULL, 'p'},
        {"name", 1, NULL, 'n'},
        {"mtime", 1, NULL, 'm'},
        {"ctime", 1, NULL, 'c'},
        {0, 0, 0, 0},
    };
    int opt;
    while((opt = getopt_long(argc, argv, optstr, opts, NULL)) != -1){
        switch(opt) {
            case 'p':
                strcpy(path, optarg);
                break;
            case 'n':
                strcpy(targetname, optarg);
                break;
            case 'm':
                modifiedtime = atoi(optarg);
                break;
            case 'c':
                changetime = atoi(optarg);
                break;
            case '?':
                if(strchr(optstr, optopt) == NULL){
                    fprintf(stderr, "unknown option '-%c'\n", optopt);
                }else{
                    fprintf(stderr, "option requires an argument '-%c'\n", optopt);
                }
                return 1;
        }
    }
    findInDir(path);
    return 0;
}
```

