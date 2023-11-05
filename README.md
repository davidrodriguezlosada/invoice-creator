# Invoice Creator

This small Go program is designed to read an Excel file and generate PDF files containing
invoices for companies. It simplifies the billing process by automating the creation of 
invoices in PDF format.

## Features
Reads Excel files with three specific tabs: Data, Companies, and Invoices.
Generates PDF files for invoices marked as "not created" in the Invoices tab.

## Getting Started
### Prerequisites
Before using this program, ensure you have the following prerequisites:

* Go programming language installed on your system.

You can install all dependencies with:

    go get

Check the example Excel file in resources folder.A valid Excel file with the specified structure (Data, Companies, and Invoices tabs).

### Installation & Usage
Clone the repository:

    git clone https://github.com/davidrodriguezlosada/invoice_creator.git

Build the program:

    go build

Run the program:

    go run .

## Contributing
Feel free to contribute to this project by submitting issues, feature requests, or pull requests. Your contributions are highly appreciated.

## License
This project is licensed under the MIT License - see the LICENSE file for details.