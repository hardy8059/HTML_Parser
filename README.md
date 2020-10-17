# HTML Parser

This repo implements a html parser package, which extracts the links and texts from "a"(hyperlink) tag and returns as a array of json objects.

## Steps to run

* Clone or download the github repo from [here](https://github.com/hardymunjal/HTML_Parser.git).
* [Data](https://github.com/hardymunjal/HTML_Parser/tree/master/data/templates) contains dummy templates for testing. You can also create your own file in the same directory which you would want to test.

* Go to the file [main.go](https://github.com/hardymunjal/HTML_Parser/tree/master/main) and assign *filename* variable with the name of your file.
    > filename := "test2.html" #replace test2.html

* Run the file using the below command:

    ```shell
    go run main.go
    ```

## Folder Structure

* [Images](https://github.com/hardymunjal/HTML_Parser/tree/master/Images) - Contains the screenshot of output of test2.html file.

* [data](https://github.com/hardymunjal/HTML_Parser/tree/master/data/templates) - Contains the template files for testing.

* [main](https://github.com/hardymunjal/HTML_Parser/tree/master/main) - Contains the *main* file for execution.

* [src](https://github.com/hardymunjal/HTML_Parser/tree/master/src) - Contains the source code.

## Example Output

![alt text](https://github.com/hardymunjal/HTML_Parser/blob/master/Images/Test2.PNG?raw=true)

## Packages Explored

* [HTTP](https://golang.org/pkg/net/http/)

* [Filepath](https://golang.org/pkg/path/filepath/)

* [HTML](https://godoc.org/golang.org/x/net/html)

* [Strings](https://golang.org/pkg/strings/)

* [Template](https://golang.org/pkg/text/template/)


*Credits - Project idea by [Jon Calhoun](https://courses.calhoun.io/)*
