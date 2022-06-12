# golang

1. Types
    1.1. The bool keyword is for Boolean type which represents a True or False state.
    1.2. Many numeric types being the most common:
        - The int type represents a number from 0 to 4294967295 in 32 bits machines and from 0 to 18446744073709551615 in 64 bits.
        - The byte type represents a number from 0 to 255.
        - The float32 and float64 types are the set of all IEEE-754 64/-bit floating-point numbers respectively.
        - You also have signed int type like rune which is an alias of int32 type, a number that goes from -2147483648 to 2147483647 and complex64 and complex128 which are the set of all complex numbers with float32/ float64 real and imaginary parts like 2.0i.
    1.3. The string keyword for string type represents an array of characters enclosed in quotes like "golang" or "computer".
    1.4. An array that is a numbered sequence of elements of a single type and a fixed size. A list of numbers or lists of words with a fixed size is considered arrays.
    1.5. The slice type is a segment of an underlying array. This type is a bit confusing at the beginning because it seems like an array but we will see that actually, they are more powerful.
    1.6. The structures that are the objects that are composed of another objects or types.
    1.7. The pointers are like directions in the memory of our program.
    1.8. The functions are interesting. You can also define functions as variables and pass them to other functions.
    1.9. The interface is incredibly important for the language as they provide many encapsulation and abstraction functionalities that we'll need often. 

2. Flow Control
    Go does not have ternary conditions like condition ? true : false.
    The len method is used to know the length of a collection.

3. Functions
    func [function_name] (param1 type, param2 type...) (returned type1, returned type2...) {
        //Function body
    }
    - Naming returned types
    Have you realized that we have given a name to the returned type? Usually, our declaration would be written as func sum(args int) int but you can also name the variable that you'll use within the function as a return value. Naming the variable in the return type would also zero-value it (in this case, an int will be initialized as zero). At the end, you just need to return the function (without value) and it will take the respective variable from the scope as returned value. This also makes easier to follow the mutation that the returning variable is suffering as well as to ensure that you aren't returning a mutated argument.

4. Arrays
    Go will initialize every value in a bool array to false . We will look deeper to zero-initialization later in this chapter.

5. Slices
    Slices are similar to arrays, but their size can be altered on runtime. This is achieved, thanks to the underlying structure of a slice that is an array. So, like arrays, you have to specify the type of the slice and its size.

6. Visibility
    Visibility is the attribute of a function or a variable to be visible to different parts of the program. So a variable can be used only in the function that is declared, in the entire package or in the entire program. How can I set the visibility of a variable or function? Well, it can be confusing at the beginning but it cannot be simpler:
        Uppercase definitions are public (visible in the entire program).
        Lowercase are private (not seen at the package level) and function definitions (variables within functions) are visible just in the scope of the function.

7. TDD
    TDD is the acronym for Test Driven Development. It consists of writing the tests first before writing the function. 
    TDD changes the way to write code and structure code so that it can be tested (a lot of code you can find in GitHub, even code that you have probably written in the past is probably very difficult, if not impossible, to test).
    So, how does it work? Let's explain this with a real life example–imagine that you are in summer and you want to be refreshed somehow. You can build a pool, fill it with cold water, and jump into it. But in TDD terms, the steps will be:
        1. You jump into a place where the pool will be built (you write a test that you know it will fail).
        2. It hurts… and you aren't cool either (yes… the test failed, as we predicted).
        3. You build a pool and fill it with cold water (you code the functionality).
        4. You jump into the pool (you repeat the point 1 test again).
        5. You're cold now. Awesome! Object completed (test passed).
        6. Go to the fridge and take a beer to the pool. Drink. Double awesomeness (refactor the code).

8. Go tools
    8-1. golint
    8-2. gofmt
    8-3. godoc