import csv

with open ("C:\Users\dougd\Dziuba&Associates Dropbox\Doug\SourceCode\Python\testcsv.csv", newline="") as file:
    reader = csv.reader(file)
    for row in reader:
        print(f"{row[0]} ({row[1]})")