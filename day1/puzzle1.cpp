#include<iostream>
#include<vector>

using namespace std;

/*Implements the logic to solve the first puzzle*/
int getPassword(vector<vector<int>>& rotations){
    int numZeroes = 0;
    short int pos = 50;

    for(vector<int> rotation : rotations){
        pos = (pos + (rotation[0] * rotation[1]) + 100) % 100;

        if (pos == 0) numZeroes++;
    }

    return numZeroes;
}

/*Parses Input from the puzzle.in file to read into a 2-D vector*/
vector<vector<int>> parseInput(){
    
    vector<vector<int>> rotations;

    freopen("puzzle1.in", "r", stdin);

    char dir;
    int rotation;
    while(cin >> dir >> rotation){
        int direction = (dir == 'R')? 1 : -1;
        rotations.push_back({direction, rotation});
    }

    return rotations;
}

int main(){
    vector<vector<int>> rotations = parseInput();
    
    cout << getPassword(rotations) << endl;
}