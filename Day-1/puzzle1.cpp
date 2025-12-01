#include<iostream>
#include<vector>

using namespace std;

int getPassword(vector<vector<int>>& rotations){
    int numZeroes = 0;
    short int pos = 50;

    for(vector<int> rotation : rotations){
        pos = (pos + (rotation[0] * rotation[1]) + 100) % 100;

        if (pos == 0) numZeroes++;
    }

    return numZeroes;
}


int main(){
    // read input file from puzzle.in
    vector<vector<int>> rotations;

    freopen("puzzle1.in", "r", stdin);

    char dirChar;
    int steps;
    while(cin >> dirChar >> steps){
        int direction = (dirChar == 'R') ? 1 : -1;
        rotations.push_back({direction, steps});
    }
    
    cout << getPassword(rotations) << endl;
}