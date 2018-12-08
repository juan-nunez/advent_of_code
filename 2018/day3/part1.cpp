#include <iostream>
#include <stdio.h>
#include <fstream>
#include <sstream>
#include <string>
using namespace std;


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

    while (std::getline(file, line)) {
        replace(line.begin(), line.end(), '@', ' ');
        replace(line.begin(), line.end(), ':', ' ');
        replace(line.begin(), line.end(), ',', ' ');
        replace(line.begin(), line.end(), 'x', ' ');

        istringstream iss(line);
        if (!(iss >> id >> posX >> posY >> dimensionX >> dimensionY)) {
            break;
        }

        for (int i = posX; i < posX + dimensionX; i++) {
            for (int j = posY; j < posY + dimensionY; j++) {
                fabric[i][j]++;
            }
        }
    }

    int overlap = 0;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            if (fabric[i][j] > 1 ) {
                overlap++;
            }
        }
    }
    cout<<overlap<<endl;
}
