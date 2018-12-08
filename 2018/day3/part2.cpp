#include <iostream>
#include <stdio.h>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
using namespace std;

typedef struct Input {
    string id;
    int posX;
    int posY;
    int dimensionX;
    int dimensionY;
} Input;


int main() {
    int n = 1200;

    int fabric[n][n];
    for (int i = 0; i < n;i++) {
        for (int j = 0; j < n; j++) {
            fabric[i][j] = 0;
        }
    }

    ifstream file("input.txt");
    string line;

    string id;
    int posX;
    int posY;
    int dimensionX;
    int dimensionY;
    vector<Input> inputs;
    while (std::getline(file, line)) {
        replace(line.begin(), line.end(), '@', ' ');
        replace(line.begin(), line.end(), ':', ' ');
        replace(line.begin(), line.end(), ',', ' ');
        replace(line.begin(), line.end(), 'x', ' ');

        istringstream iss(line);
        if (!(iss >> id >> posX >> posY >> dimensionX >> dimensionY)) {
            break;
        }
        Input input = {
            id,
            posX,
            posY,
            dimensionX,
            dimensionY
        };

        inputs.push_back(input);

        for (int i = posX; i < posX + dimensionX; i++) {
            for (int j = posY; j < posY + dimensionY; j++) {
                fabric[i][j]++;
            }
        }
    }


    for (int i = 0; i < inputs.size(); i++) {
        Input input = inputs[i];
        bool failed = false;
        for (int j = input.posX; (j < input.posX + input.dimensionX) && !failed; j++) {
            for (int k = input.posY; k < input.posY + input.dimensionY; k++) {
                if (fabric[j][k] > 1) {
                    failed = true;
                    break;
                }
            }
        }

        if (!failed) {
            cout<<input.id<<endl;
            return 0;
        }
    }

}
