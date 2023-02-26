import os
import socket
from fpdf import FPDF

HOST = "127.0.0.1"
PORT = 3610

while True:
    print("Creating socket connection...")
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((HOST, PORT))
        s.listen()
        print("Awaiting txt file data...")
        conn, addr = s.accept()
        with conn:
            print(f"Connected by {addr}")

            # First get the txt file name, which we'll carry over to the PDF file
            fileName = ""

            # Continue to wait for data until we see the .txt ending
            while True:
                data = conn.recv(1024)
                fileName += data.decode()
                if len(fileName) > 4:
                    ending = fileName[-4:]
                    if ending == '.txt':
                        break

            # Print feedback and send same feedback to main
            print("Saving " + fileName + " as PDF...")
            conn.sendall(b"Saving '" + fileName.encode() + b"' as PDF...")

            # Then we get all of the recipe text and save it in a string
            recipeText = ""
            # Continue to capture the data until we reach END
            while True:
                data = conn.recv(2048)
                recipeText += data.decode()
                if len(recipeText) > 3:
                    ending = recipeText[-3:]
                    if ending == 'END':
                        break
            
            # Create PDF from recipeText
            print("Setting up PDF ...")
            print(recipeText)
            pdf = FPDF()
            pdf.set_right_margin(margin=10)
            pdf.set_font("Arial", size = 12)
            pdf.add_page()
            pdf.multi_cell(w=0, h=10, txt = recipeText[:-3], align = 'J')

            # We get the user's unique path to their downloads folder
            userpath = os.path.expanduser("~/Downloads") 

            print("Line 56 ...")
            # then save the PDF in the downloads folder
            pdf.output(f"{userpath}\\{fileName[:-4]}.pdf")

            # Print complete and send final feedback to main
            print("Saved!")
            conn.sendall(b"Saved!")