# ccwc
 This project is a hands-on challenge to create a custom version of the Unix wc (word count) command-line tool, named ccwc (Coding Challenges word count). It guides you through incrementally building features like counting bytes, lines, words, and characters, as well as handling default options and input from standard input (stdin). This dcope of the exercise is to emphasizes Unix philosophies of creating simple, interconnected tools and provides practical experience in command-line interface (CLI) development.

# Installation & Setup

1. Clone the project from GitHub:
   ```bash
   git clone https://github.com/PLCodingStuff/ccwc.git
   cd ccwc.git
   ```
2. Ensure that dependencies are synced:
   ```bash
   go mod tidy
   ```
3. Run the aplication:
   ```bash
   go run main.go [OPTION]... [FILE]...
   ```
4. (Optional) Test the application with the tests in `test` folder:
   ```bash
   go test
   ```
# Usage
### Synopsis

```bash
ccwc [OPTION]... [FILE]...
```
### Description
Print newline, word, character, and byte counts for each FILE, and a total line if more than one FILE is specified. A word is a non-zero-length sequence of characters delimited by white space.

### Options

| Flag| Description|
| --- | ------------------------------ |
| -c | Print the byte counts.|
| -l | Print the newline counts. |
| -w | Print the word counts.|
| -m | Print the character counts.|

### Example

```bash
go run ccwc.go -c tests/test.txt -l tests/test2.txt -wml tests/test3.txt tests/test4.txt
```
