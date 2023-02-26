import socket
import sys

HOST = "127.0.0.1"
PORT = 3610

file = sys.argv[1]

# recipeFileName needs to end in .txt and collocated in the current directory
# if needed, will add functionality to allow for absolute paths
def SetUpClient(recipeFileName):
    with open("/Users/carly/cs361FinalProject/microservice/" + recipeFileName, "r", encoding= "utf8") as recipeText:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((HOST, PORT))
            # First send the txt file name so PDF is saved identically
            s.sendall(recipeFileName.encode())

            # get/print feedback from microservice (needed to interupt sendalls)
            feedbackMsg = s.recv(1024)
            print(feedbackMsg.decode())

            # read all contents of txt file to recipe
            recipe = recipeText.read()

            # send contents to microservice, END is appended to signal when all contents received
            s.sendall(recipe.encode() + b"END")

            # get/print feedback from microservice
            feedbackMsg = s.recv(1024)
            print(feedbackMsg.decode())

SetUpClient(file)