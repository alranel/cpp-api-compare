# C++ API Compare Tool 

This tool performs an assessment of feature parity and API interoperability between two or more libraries. It generates a report with the list of classes, methods, global functions and singletons, in which their signatures and return values are compared across the supplied libraries.

Features include:

* parsing with Doxygen for robustness (code is not compiled so you can supply even incomplete codebases)
* expansion of preprocessor directives
* parsing of functions defined as macros
* parsing of class singletons
* custom namespaces
* custom defines

For an example, see the [Arduino API comparison](https://github.com/alranel/arduino-api-compare) repository which is the primary use case for which this project was developed.

## Getting started

1. Install Doxygen with your preferred package manager:

    ```
    apt-get install doxygen
    ```

2. Clone this repository and compile its contents:

    ```
    git clone https://github.com/alranel/cpp-api-compare
    cd cpp-api-compare
    go build
    ```

3. Duplicate [config.yml.template](config.yml.template) and fill it: the paths to the libraries you want to compare as well as the list of classes/singletons/functions to consider, :

    ```
    cp config.yml.template config.yml
    vi config.yml
    ```

    You'll need to enter the paths to the libraries you want to compare, the list of classes/singletons/functions to consider in the comparison, and any additional define or namespace specification needed to compile the libraries.

4. Run the tool:

    ```
    ./cpp-api-compare run --config config.yml --output report.md
    ```

## Credits and license

This tool was written by [Alessandro Ranellucci](https://github.com/alranel) and is licensed under the terms of the Affero GNU General Public License v3.
