# Assignment
Sample Output
C:\Users\Welcome\go\Assignment>main.exe get
Rover positions
Rover 1 =  0   0   N
Rover 2 =  0   0   N
##Check
C:\Users\Welcome\go\Assignment>main.exe set -RoverNumber=0 -SetMaxXPosition1=5 -SetMaxYPosition1=5 -SetXPosition1=1 -SetYPosition1=2 -SetDirectionChar1="N" -SetMoveMessage1="LMLMLMLMM" -SetXPosition2=3 -SetYPosition2=3 -SetDirectionChar2="E" -SetMoveMessage2="MMRMMRMRRM"

1   3   N
5   1   E

C:\Users\Welcome\go\Assignment>main.exe update -UpdateMoveMessage1="R"
1   3   E
5   1   E

C:\Users\Welcome\go\Assignment>main.exe update -UpdateMoveMessage1="L"
1   3   N
5   1   E

C:\Users\Welcome\go\Assignment>main.exe get
Rover positions
Rover 1 =  1   3   N
Rover 2 =  5   1   E

---Problem statement---
Go Developer
INSTRUCTIONS
A squad of robotic rovers are to be landed by NASA on a plateau on Mars.
This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the
surrounding terrain to send back to Earth.
A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points.
The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left
corner and facing North.
In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90
degrees left or right respectively, without moving from its current spot.
'M' means move forward one grid point, and maintain the same heading.
Assume that the square directly North from (x, y) is (x, y+1).
Input (hard coded, input from keyboard, input from rest api):
The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are assumed to be 0,0. The rest of the input is
information pertaining to the rovers that have been deployed. Each rover has two lines of input. The first line gives the rover's position, and
the second line is a series of instructions telling the rover how to explore the plateau.
The position is made up of two integers and a letter separated by spaces, corresponding to the x and y co-ordinates and the rover's
orientation. Each rover will be finished sequentially, which means that the second rover won't start to move until the first one has finished
moving. Output:
The output for each rover should be its final co-ordinates and heading.
Plateau max X and Y, Starting coordinates, direction and path for two rovers:
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
Output and new coordinates:
1 3 N
5 1 E
DELIVERABLE:
Command line interface (CLI) with fully documented “--help” and other paths. The CLI should permit singular or compound input. The
rover’s position should persist across invocations of the CLI.
Code should be packaged (either as an ipk, deb or snap file) and can be written in Go
Local Webserver based interface (built using libhttpd or similar tool) for issuing the same commands as the CLI Standalone REST
API, built using GO , that uses http verbs (like GET, POST and PUT ) to provide the same features as the command line interface,
and which would potentially allow for a web front end to create a dashboard for controlling the Rovers
(Optional) Containerize the app and API: Create a Dockerfile and possibly a docker_compose.yml (in case more than
one container is to be created) with README.md clear instructions to build/run the app
If uploading the repo to github (instead of .zip) please:
make it private and share it with users: ks-iot
remove any references to Leader Group

IMPORTANT:
There is no time limit to deliver the result, and time will be taken into account
Tests are more than welcome
