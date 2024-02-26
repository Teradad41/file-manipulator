import sys


class FileManipulator:
    def __init__(self, input_file=None, output_file=None):
        self.input_file = input_file
        self.output_file = output_file

    def reverse(self, input_file, output_file):
        with open(input_file, 'r') as in_file:
            original_text = in_file.read()

        with open(output_file, 'w') as out_file:
            reversed_text = ''.join(reversed(original_text))
            out_file.write(reversed_text)

    def copy(self, input_file, output_file):
        with open(input_file, 'r') as in_file:
            original_text = in_file.read()

        with open(output_file, 'w') as out_file:
            out_file.write(original_text)

    def duplicate(self, input_file, n):
        with open(input_file, 'r+') as in_file:
            original_text = in_file.read()
            duplicated_text = original_text * int(n)

        with open(input_file, 'w') as in_file:
            in_file.write(duplicated_text)

    def replace_string(self, input_file, needle, new_string):
        with open(input_file, 'r') as in_file:
            original_text = in_file.read()
            replaced_string = original_text.replace(needle, new_string)

        with open(input_file, 'w') as in_file:
            in_file.write(replaced_string)


def main():
    if len(sys.argv) < 4:
        print("Usage: program <command> <inputFile> <arguments...>")
        sys.exit(1)

    command = sys.argv[1]
    args = sys.argv[2:]

    file_manipulator = FileManipulator()

    if command == 'reverse':
        file_manipulator.reverse(*args)
    elif command == 'copy':
        file_manipulator.copy(*args)
    elif command == 'duplicate':
        file_manipulator.duplicate(*args)
    elif command == 'replace-string':
        file_manipulator.replace_string(*args)
    else:
        print("Invalid command")
        sys.exit(1)

    print("Operation completed successfully.")

if __name__ == '__main__':
    main()